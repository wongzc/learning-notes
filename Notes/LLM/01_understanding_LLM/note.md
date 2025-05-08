1. Before LLM, traditional method excel at categorization
    - LLM "understand" mean able to generate text in way that appear coherent and contextually relevant
    - LLM from deep learning
    - previous NLP for specific task, but LLM with broader proficiency 
    - transformer architecture + vast amounts of data

2. what is LLM
    - neural network design to understand, generate, respond to human-like text
    - deep neural network trained on massive amount of text data
    - large: model parameter size and dataset size
        - 10~100 billion parameter
    - transformer model: pay selective attention to input when predict, able to handle complexities of human language
    - generative AI
    - machine learning: development of algorithms that can learn from and make predictions or decisions based on data, without explicitly programmed
    - deep learning: no need manual feature extraction
3. stage of building and using LLM
    - most LLM today implement using PyTorch deep learning library
    - advantage of building own LLM
        - custom built LLM, which tailored for specific task/ domain, can outperform general purpose LLMs
        - data privacy
        - smaller custom LLM enable deployment on customer device
            - lower latency
            - lower server cost
        - developer more control
    - stage: pretraining + finetuning
        - pretraining: 
            - LLM train on diverse dataset to understand language
            - large corpus of text ( raw text), without labeling
            - LLM use self-supervised learning to generate own labels from input
            - after this stage
                - model capable of text completion 
                - model can learn new task based on few example
        - finetuning
            - specifically trained on narrower dataset, more specific
            - train on labeled data
            - 2 type:
                1. instruction finetuning
                    - with instruction and answer pair
                2. classification finetuning
                    - with text and class label

4. transformer architecture
    - deep neural network architecture
    - original transformer: to translate English text to German and French
    - 2 submodule in transformer
        1. encoder
            - process input, encode into numerical representation/ vector that capture the context
        2. decoder
            - take encoded vector and generate output
    - both encoder and decoder consist of many layer , connected by "self-attention mechanism"
    - self attention mechanism:
        - weigh importance of words/ token in sequence relative to each other
        - capture long range dependency and contextual relationship within input
    - BERT: bidirectional encoder representation from transformer
        - specialized in masked word prediction at encoder ( GPT is next word prediction at decoder)
        - better at classification like sentiment prediction
            - X use BERT to detect toxic content 
    - GPT: generative pretrained transformer 
        - focus on decoder portion
        - for task that require generating text
            - translation, text summarization, writing
        - zero shot learning task
            - ability to generalize to new tasks without any prior specific example
        - few shot learning 
            - learn from minimal example
    - transformer vs LLM
        - LLM based on transformer
        - not all transformer are LLM, transformer can use for computer vision
        - not all LLM are transformer, some LLM based on recurrent and convolutional architectures

5. Large data set
    - vast array of topics and natural and computer languages to train 
    - web-crawl data, books, Wikipedia
    - large scale and diversity of training dataset allow model to perform well on diverse tasks
    - pretrained nature allow later finetuning
    - pretraining required many resources. GPT3 cost 4.6 million
    - but many pretrained LLM available as open source

6. GPT architecture
    - GPT trained on next word prediction
    - help model understand how words and phrases fit together in language
    - next word prediction:
        - self supervised learning, self labeling
        - don't need to collect labels for training data, but can use the structure of data itself
        -  use next word in sentence as label that model need to predict
        - create label on the fly, can use massive unlabeled data to train LLM
    - general GPT architecture: just decoder, no encoder
    - predict text 1 word at a time, GPT consider as autoregressive
        - autoregressive incorporate previous output as input for future prediction
    - original transformer repeat encoder and decoder block 6 times.
        - GPT3 has 69 transformer layers and 175 billion param
    - original transformer with encoder+ decoder, but GPT only decoder, and also can translate
        - emergent behavior: capability that isn't explicitly taught during training but emerge due to exposure to vast quantities of data

7. build a LLM, 3 stage
    1. build LLM
    2. pretrain LLM
    3. finetune LLM

