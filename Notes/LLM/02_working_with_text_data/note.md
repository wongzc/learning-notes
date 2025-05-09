1. understanding word embeddings
    - LLM cannot process raw text directly
    - text is categorical, need to translate to continuous valued vector
        - known as embedding
    - different data format need distinct embedding models.
        - embedding model for text not suitable for audio/ video
    - embedding map discrete objects like words to points in a continuous vector space
        - can also embed sentence, paragraph, document, for RAG
        - RAG combined generation with retrieval to pull relevant information
    - Word2Vec
        - train neural network to generate word embedding by:
            - generate word given context or vice versa
            - similar context word have similar meaning
            - can have many dimension for context
    - LLM usually produce own embeddings that are part of input layer, and updated during training.
        - so that embedding are optimized to specific task and data at hand
    - LLM use higher dimension to capture context
    - GPT embedding size (also known as dimensionality)
        - GPT2: 768 dimensions (~117M param)
        - GPT3: 12288 dimensions (~175B param)

2. tokenizing text
    - split input text into individual token
    - token can be word or special character
    - when tokenizing, keep or remove white space depends on our need
        - remove can save memory
        - keep can make it sensitive to text structure maybe like for python

3. converting tokens into token ID
    - after tokenize, need to assign ID before convert to embedding vectors
    - and we also need to decode it to human readable way

4. special context tokens
    - <|unk|>: to address new and unknown words that were not part of training
    - <|endoftext|>: to separate 2 unrelated text sources. usually put before a document, so that GPT understand that it is unrelated to the previous document
    - other special token that may use:
        1. [BOS]: beginning of a sequence (text)
        2. [EOS]: end of text
        3. [PAD]: padding to ensure all text same length when training LLM with batch>1
    - GPT tokenizer only use <|endoftext|>
    - GPT also dont use <|unk|> for unknown word, but use byte pair encoding tokenizer

5. Byte Pair Encoding (BPE)
    - can use python library 'tiktoken'
    - initiate with `tokenizer = tiktoken.get_encoding("gpt2")`
    - encode with `tokenizer.encode('...')`
    - decode with `tokenizer.decode('...')`
    - BPE tokenizer: 50527 vocab when training GPT2,3
    - BPE break word into subword units
    - BPE start from single character, then merge character to form common subword

6. data sampling with sliding window
    - train model by predicting word 1 by 1, the other behind are masked
    - parameter:
        - `max_length`: recommend for LLM input size of 256 token
        - `stride`: number of position shift across sample
        - `batch`: how many sample at a time when iterate dataloader, smaller less memory but more noise. a trade off and parameter to experiment when training LLMs

7. creating token embeddings
    - convert token ID into embedding vector
    - need to initialize embedding weights with random value
    - `torch.nn.Embedding(vocab_size, dimension)` to initialize data

8. encoding word positions
    - self attention mechanism of LLM treat token without knowing position, so need to inject positional information into model!
    - 2 type of position-aware embeddings: 
        1. relative positional embedding: like xx is 3 word before yy
        2. absolute positional embedding
    - positional embeddings+ token embeddings to become input embeddings


