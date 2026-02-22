1. layer 4 acceleration
    - decide by IP+port
    - accelerate:
        - TCP connection
            - offload TCP to L4 ( from backend)
        - SYN flood protection
        - NAT/SNAT optimization
        - Anycast routing
    - technique
        - Hardware offload (ASIC-based LB)
        - Kernel bypass (DPDK)
        - TCP Fast Open
        - Global Anycast routing
    - accelerate:
        - not just faster, but offload expensive work from backend server

2. layer 7 acceleration
    - decide by URL, host header, cookies, headers, JWT (HTTP meta data)
    - accelerates
        - HTTP routing
        - TLS termination
            - offload TLS termination
            - connection pooling
        - caching
            - if response cachable, cache it and return
        - WAF filtering
        - rate limiting

3. ingress
    - is a entry point
    - usually l7, backed by LB, not alone

4. global ingress
    - globally distributed entry point, route user traffic to nearest/healthies region
    - global ingress route traffic to regional LB
    - mechanism:
        - anycast
        - geo DNS
        - health-aware routing
        - cross region failover
    - helps:
        - latency improve
        - regional outage
        - global scaling

5. CMAF
    - media container standard defined by MPEG
    - enable
        - low-latency streaming
        - unified format for HLS + DASH
        - chunked video delivery
    - standardized the format that used by apple ( HLS) and others (MPEG-DASH)
    - standardized fragmented MP4 & common chunk format
    - for live streaming
    - nor protocol, is packaging format