1. nginx
    - web server
        - serve static files ( HTML, CSS, image)
    - reverse proxy
        - sit infront of backend services
    - l7 load balancer
        - route HTTP traffic
    - TLS termination
        - handle HTTPS
    - ingress control

2. why NGINX
    - event-driven, non-blocking I/O
    - few worker processes handles multiple connections
    - for high traffic, concurrent, reverse proxy
    - rate limiting