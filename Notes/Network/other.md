1. layer 4 acceleration
    - decide by IP+port
    - accelerate:
        - TCP connection
            - offload TCP to L4 ( from backend)
        - SYN flood protection
            - SYN flood: 
                - Denial-of-Service (DOS) attack, attacker send many SYN to server
                - use spoofed IP, server ACK never arrive
                - cause millions half open connections
                - TCP is stateful, server allocate memory for it before handshake done
            - to solve:
                - rate limiting per IP
                - only allocate memory after final ACK ( what LB doing)
                    - use stateless method
                        1. when client sync, server create ISN by hashing client+server ip, port, secret, time, etc
                        2. use ISN as "seq" field to return to client
                        3. if real client, return with iSN+1
                        4. server recompute with infor and compare with client replied seq
        - NAT/SNAT optimization
            - SNAT= source NAT
            - mapping table at LB, from client IP+port to LB IP + port
            - to makesure backend response send to LB, and LB know which client to return
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
    - usually k8s cluster entry point

4. global ingress
    - globally distributed entry point, route user traffic to nearest/healthies region
    - also k8s concept, but similar to global LB
    - use BGP anycast + regional load balancing
    - mechanism:
        - anycast
            - determine which to route to
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

6. Anycast routing
    - multi geographically distributed server share same IP
    - routing (BGP) direct client to topologically nearest
    - L3 ( IP level routing) 
    - advantage
        - low latency as it select by nearest
        - DDos resistance: traffic spread by PoP
        - high availability: 1-down, go to others

7. high availability infrastructure for global
    - DNS -> Global traffic management -> Anycast/ Global LB -> Regional L7 LB -> service cluster
        - DNS -> GTM
            - geo routing ( by region)
            - latency-based
            - weighted routing
        - GTM
            - to route for compliance reason
            - weighted/ latency based etc
        - Anycast
            - BGP
            - optimize within region
    - active-active
        - multi region serve prod traffic
    - health check
        - global LB health
            - check if region reachable
            - reflect dependency readiness ( DB, cache)
        - regional LB health
            - check backend target instance with some endpoint like `/heath/ready`

8. Backlog (TCP)
    - SYN backlog: for half open (received SYN, sent SYN-ACK, waiting  ACK)
        - SYN flood attack this
    - Accepted Queue
        - for state=Established

9. TIME_WAIT (TCP)
    -  a state after send final 'ACK' for close connection
    - why need:
        - ensure 'ACK' sent, to resend if opp send 'FIN' again
        - to have enough time for old packet expire before close, to avoid confuse next connection
        
10. ipvs
    - IP Virtual Server
    - linux kernel feature, implementing l4 LB
    - used by linux virtual server (LVS)
    - faster than Nginx, due to in kernel space, avoid user-space context switching and HTTP parsing

11. questions
    - l4 LB vs l7 LB
        - L4 at transport, based on port+ip+protocol, 5 tuple
            - faster due to at kernel space, no parsing content
        - l7 at application
            - smarter, but need terminate TCP, decrypt TLS, parse HTTP
            - CPU overhead + user-space context switching
    - global LB
        - distribute traffic cross region using
            - GeoDNS
            - Anycast ( BGP)
            - health check
            