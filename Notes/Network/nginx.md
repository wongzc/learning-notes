1. nginx
    - like apache web server
        - also server as reverse proxy ( accept https,terminate ssl, forward request to django)
        - django gunicorn dont have built in TLS management, protection etc
        - gunicorn not production web server, not optimized for internet exposure
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