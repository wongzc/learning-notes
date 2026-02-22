1. ELK stack
    - component:
        - elasticsearch
            - store, index log, full text seacrh
        - logstash
            - collect, parse, transform log
            - send to elasticsearch
        - kibana
            - search UI
    - distributed logging:
        - centralized log collection
        - full text search
        - infra/app monitoring

2. architecture
    - app -> log shipper -> logstash -> elastic search -> kibana
