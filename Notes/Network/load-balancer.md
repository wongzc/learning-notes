1. load balancer
    - traffic distribution
    - between client and backend
    - L4 & L7
        - can be used together
        - internet -> cloud L4 LB -> L7 LB/ Ingress -> Service
    - target
        - HA
        - Horizontal scaling
        - better performance ( latency)
    - steps:
        - accept client connections
        - select backend target
        - forward to target
        - return response to client
    - algorithm
        - round-robin: sequential, go to 1, next will be 2, then next 3,...
        - weighted round-robin: 75% to 1, ...
        - least connections: to server with fewest connection
        - IP hashing: use client ip to determine backend, for stickiness
            - but most time actually use cookie-based sticknicess (L7)
            - to achieve stateless ( no need to store state locally)
                - store in redis, so server no local session state
    - health check:
        - passive: check after connect fail
        - active: periodical check
        - l4: `TCP connect probe`
        - l7: `GET /health`, expect 200 OK
    - SSL termination
        - SSL termination 
            - client -> HTTPS, LB decrypt, LB -> backend via HTTP ( plain text)
            - LB hold TLS cert and private key, takeover TLS from CPU
            - used for trusted network ( same VPC, subnet)
        - SSL re-encryption
            - client -> HTTPS, LB decrypt and re-encrypt, LB -> Backend via HTTPS
            - double TLS
            - for compliance reason
            - can do L7 routing ( due to decrypted TLS to get full HTTP info like url, method, header, body)
                - SNI ( server name indication) no need TLS decrypt, can only see the SN ( domain)
        - TCP passthrough ( TLS passthrough )
            - client -> HTTPS, LB -> backend via raw TCP stream
            - SSL not terminate, TCP is end to end between client and backend
            - L4 LB, (only see IP, port)

2. layer 4 (transport) load balancer (NLB)
    - have information on: `source/dest ip+port, protocol ( UCP/TCP)`
        - route based on info above
        - not inspect app payload (HTTP headers, URL paths, cookies)
    - characteristic
        - fast
        - no content based routing
    - usage:
        - DB load balancing
        - high through put LB

3. layer 7 (application) load balancer (ALB)
    - have information on: `everything L4 see`+`HTTP header, cookies, URL path, query param`
        - path based routing ( like /api -> backend)
        - host based (api.exmaple.com)
        - cookie based
        - header based
        - rate limiting
            - strategy:
                - token bucket: request use token, token add at rate
                - leaky bucket: queue drain at fixed rate, if full, request drop
                - fixed window: n request per time window
        - WAF rules
    - characteristic
        - flexible
        - higher CPU overhead
        - terminate TLS
    - all L7 load balancers are reverse proxies ( but not vice versa)
        - normal proxy: client -> proxy -> internet
        - reverse proxy: client -> proxy -> backend servers
        - l7 LB terminate HTTP, generate new HTTP request to backend
            - it modify headers
        - L4 does not terminate app protocol, so not reverse proxy

4. deployment pattern for load balancers
    - single load balancers: lb down, system down
    - HA load balancers:
        - Active-passive: 1 active, another passive, standby and to take over
        - Active-active: 2 active, use DNS round robin
        - Floating IP failover: ip not permanently tied to a machine
        - DNS-based failover

5. LB vs API gateway
    - both: distribute traffic
    - LB: basic authentication, limited rate limiting, transformation, protocol translate
    - API gateway:  advance authentication , and can do all above

6. firewall, DNS, proxy, micro-seg, load balancer
    1. DNS: domain to LB IP
    2. firewall: L3/L4, block ports, IP range
    3. Load balancer
    4. WAF: L7 inspection, block SQL injection, XSS
    5. Reverse proxy
    6. Micro-seg
    7. service

