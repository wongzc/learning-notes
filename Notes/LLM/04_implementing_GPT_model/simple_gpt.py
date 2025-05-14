import torch
import torch.nn as nn

class MultiHeadAttention(nn.Module):
    def __init__(self, d_in, d_out, context_length, dropout, num_heads, qkv_bias=False):
        super().__init__()
        assert (d_out % num_heads == 0), "d_out must be divisible by num_heads"
        
        self.d_out = d_out
        self.num_heads = num_heads
        self.head_dim = d_out // num_heads # reduce projection dim to match desired output dim

        self.W_query = nn.Linear(d_in, d_out, bias=qkv_bias)
        self.W_key = nn.Linear(d_in, d_out, bias=qkv_bias)
        self.W_value = nn.Linear(d_in, d_out, bias=qkv_bias)

        self.out_proj = nn.Linear(d_out, d_out) # use linear layer to combine head outputs
        self.dropout = nn.Dropout(dropout)
        self.register_buffer(
            "mask",
            torch.triu(torch.ones(context_length, context_length),
                diagonal=1)
        )

    def forward(self, x):
        b, num_tokens, d_in = x.shape

        keys = self.W_key(x)
        queries = self.W_query(x)
        values = self.W_value(x) 

        keys = keys.view(b, num_tokens, self.num_heads, self.head_dim)
        values = values.view(b, num_tokens, self.num_heads, self.head_dim)
        queries = queries.view(
            b, num_tokens, self.num_heads, self.head_dim
        )
        # reshape the matrix with self.head_dim & self.num_head. 
        # it was (b, num_tokens, d_out) previously


        # transpose the 3 from (b, num_tokens, num_heads, head_dim) to (b, num_heads, num_tokens, head_dim)
        keys = keys.transpose(1, 2) 
        queries = queries.transpose(1, 2)
        values = values.transpose(1, 2)

        attn_scores = queries @ keys.transpose(2, 3) # get dot product of each head
        # (b, num_heads, num_tokens, head_dim) @ (b, num_heads, head_dim, num_tokens) => (b, num_heads, num_tokens, num_tokens)
        # (...,m,k)@(...,k,n)=> (...m,n)
        mask_bool = self.mask.bool()[:num_tokens, :num_tokens] # mask truncated to number of tokens

        attn_scores.masked_fill_(mask_bool, -torch.inf) # use mask to fill attention scores

        attn_weights = torch.softmax(
            attn_scores / keys.shape[-1]**0.5, dim=-1)
        attn_weights = self.dropout(attn_weights)

        context_vec = (attn_weights @ values).transpose(1, 2) # transpose back to shape (b, num_tokens, num_heads, head_dim)
        context_vec = context_vec.contiguous().view(
            b, num_tokens, self.d_out
        ) # combine heads. self.d_out = self.num_heads * self.head_dim

        context_vec = self.out_proj(context_vec) # optional linear projection
        return context_vec
    
# configuration
GPT_CONFIG_124M = {
 "vocab_size": 50257, # Vocabulary size, as used by BPE tokenizer
 "context_length": 1024, # Context length, max number of input tokens the model can hanlde via positonal embedding
 "emb_dim": 768, # Embedding dimension, each token will become 768 dimension
 "n_heads": 12, # Number of attention heads
 "n_layers": 12, # Number of layers (transformer block)
 "drop_rate": 0.1, # Dropout rate
 "qkv_bias": False # Query-Key-Value bias, will be discussed again in ch06
}

class GELU(nn.Module): 
    def __init__(self):
        super().__init__()
    def forward(self, x):
        return 0.5 * x * (1 + torch.tanh(
        torch.sqrt(torch.tensor(2.0 / torch.pi)) *
        (x + 0.044715 * torch.pow(x, 3))
        ))
    
class FeedForward(nn.Module): # neuron module for LLM transformer block
    def __init__(self, cfg):
        super().__init__()
        # 2 linear layer+ GeLU activation
        self.layers = nn.Sequential(
        nn.Linear(cfg["emb_dim"], 4 * cfg["emb_dim"]), # output size of 4*768
        GELU(),
        nn.Linear(4 * cfg["emb_dim"], cfg["emb_dim"]), # output size of 768
        )
    def forward(self, x):
        return self.layers(x)

class LayerNorm(nn.Module):
    def __init__(self, emb_dim):
        super().__init__()
        self.eps = 1e-5 # small constant added to variance to avoid divide by 0

        # 2 trainable param that LLM will adjust if doing so improve model performance
        self.scale = nn.Parameter(torch.ones(emb_dim))
        self.shift = nn.Parameter(torch.zeros(emb_dim))

    def forward(self, x):
        mean = x.mean(dim=-1, keepdim=True) # normalization implement on the -1 dimension of input tensor
        var = x.var(dim=-1, keepdim=True, unbiased=False)
        # unbiased=False means didnt apply bessel's correction ( which use n-1, not n for denominator)
        # result in biased estimation, but LLM embedding dimension is large, so -1 or not impact is small
        norm_x = (x - mean) / torch.sqrt(var + self.eps)
        return self.scale * norm_x + self.shift
    
class TransformerBlock(nn.Module):
    def __init__(self, cfg):
        super().__init__()
        self.att = MultiHeadAttention(
            d_in=cfg["emb_dim"],
            d_out=cfg["emb_dim"],
            context_length=cfg["context_length"],
            num_heads=cfg["n_heads"],
            dropout=cfg["drop_rate"],
            qkv_bias=cfg["qkv_bias"])
        
        self.ff = FeedForward(cfg)

        self.norm1 = LayerNorm(cfg["emb_dim"])
        self.norm2 = LayerNorm(cfg["emb_dim"])
        self.drop_shortcut = nn.Dropout(cfg["drop_rate"])

    def forward(self, x):

        shortcut = x # shortcut connection for attention block
        x = self.norm1(x) # pre-layer norm, layer norm before attention/feedfoward, and dropout after
        # older architecture use post-layer norm, which leads to worse training dynamics
        x = self.att(x)
        x = self.drop_shortcut(x)
        x = x + shortcut # add the orginal input back as shortcut
        
        shortcut = x # shortcut connection for feedforward
        x = self.norm2(x)
        x = self.ff(x)
        x = self.drop_shortcut(x)
        x = x + shortcut # shortcut
        return x


def generate_text_simple(model, idx, max_new_tokens, context_size):
    # idx shape: (batch, n_token)
    for _ in range(max_new_tokens): #iterates for a specified number of new tokens to be generated
        idx_cond = idx[:, -context_size:] # crop current context if it exceed the length that model support
        with torch.no_grad():
            logits = model(idx_cond) # compute prediction

        logits = logits[:, -1, :] # focus on last step (batch, n_token, vocab_size) => (batch, vocab_size)
        probas = torch.softmax(logits, dim=-1) # shape in (batch, vocab_size)
        # use softmax to convert logits into probability distribution
        # the softmax here can be skip actually, we still can get the argmax
        idx_next = torch.argmax(probas, dim=-1, keepdim=True) # shape (batch, 1)
        # use argmax to get position with highest value
        idx = torch.cat((idx, idx_next), dim=1) # add sample index into running sequence, (batch, n_token +1)
    return idx

class GPTModel(nn.Module):
    def __init__(self, cfg):
        super().__init__()
        self.tok_emb = nn.Embedding(cfg["vocab_size"], cfg["emb_dim"])
        self.pos_emb = nn.Embedding(cfg["context_length"], cfg["emb_dim"])
        self.drop_emb = nn.Dropout(cfg["drop_rate"])

        self.trf_blocks = nn.Sequential(*[TransformerBlock(cfg) for _ in range(cfg["n_layers"])])
        # * to unpack list

        self.final_norm = LayerNorm(cfg["emb_dim"])
        self.out_head = nn.Linear(
            cfg["emb_dim"], cfg["vocab_size"], bias=False
        )
    def forward(self, in_idx):
        batch_size, seq_len = in_idx.shape
        tok_embeds = self.tok_emb(in_idx)

        pos_embeds = self.pos_emb(torch.arange(seq_len, device=in_idx.device))
        #allow us to train model on CPU or GPU

        x = tok_embeds + pos_embeds
        x = self.drop_emb(x)
        x = self.trf_blocks(x)
        x = self.final_norm(x)
        logits = self.out_head(x)
        return logits

import tiktoken
tokenizer = tiktoken.get_encoding("gpt2")
model = GPTModel(GPT_CONFIG_124M)

start_context = "Hello, I am"
encoded = tokenizer.encode(start_context)
print("encoded:", encoded)
encoded_tensor = torch.tensor(encoded).unsqueeze(0)
print("encoded_tensor.shape:", encoded_tensor.shape)
model.eval()
# in eval mode, so random components like dropout is disabled
out = generate_text_simple(
 model=model,
 idx=encoded_tensor,
 max_new_tokens=6,
 context_size=GPT_CONFIG_124M["context_length"]
)
print("Output:", out)
print("Output length:", len(out[0]))
decoded_text = tokenizer.decode(out.squeeze(0).tolist())
print(decoded_text)
# generate random stuff as it is not trained yet!!