1. firewall
    - network security control based on IP. port, protocol, connection state
    - operate on
        - layer 3: ip filtering
        - layer 4: tcp/udp port filtering ( no fixed port. follow rules like allow TCP on port 443 etc)
    - type
        1. packet filtering
            - stateless, no session awareness
            - based on IP +port
        2. stateful
            - track TCP connection state ( SYN, ESTABLISHED, FIN)
2. web application firewall
    - inspect & filter HTTP/HTTPS traffic
    - layer 7, terminate TLS ( need decrypt to check)
    - check
        - http headers
        - query param
        - JSON payload
        - Cookies
        - POST bodies
    - block sql injection, XSS, CSRF etc
    - compare against
        - signature rules
        - regex patterns
        - heuristic
        - behavioral models