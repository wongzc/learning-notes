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