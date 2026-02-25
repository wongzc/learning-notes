1. to prevent cascading failure
    - circuit breakers
        - error rate > threshold, open circuit to fail fast
    - timeouts
        - limit time wait for response
    - retry with exponentially backoff
        - retry with increase delay
    - load shedding
        - drop excess traffic when overload
    - rate limiting
        - control by client or globally

2. rate limiting
    - token bucket: request use token, token add at rate
    - leaky bucket: queue drain at fixed rate, if full, request drop
    - fixed window: n request per time window

3. detect partial failure
    - latency monitoring p99, p999
    - error rate spike
    - health check

4. user report high latency
    - Check monitoring dashboards (CPU, memory, network)
    - Check error rates
    - Check backend health
    - Inspect connection counts
    - Check for packet drops
    - Verify DNS resolution
    - Trace request path

5. Backend is healthy but traffic not reaching
    - routing rule misconfigure
    - SNAT failed
    - firewall rule
    - healthcheck mismatch
    - security group restriction
    - connection tracking limit reached

6. what is connection tracking
    - maintain state of active TCP
    - for load balancer, NAT, firewall

7. handle Tbps traffic scale
    - horizontal scaling
    - anycast
    - Kernel-level load balancing
    - Edge POPs

8. metrics are critical for traffic platforms
    - SYN backlog: half open
    - p99 latency
        - 99% of requests complete within this time
    - QPS
    - packet drop rate
    - connection count: full open
    - CPU softirq usage
    - Error rate

9. softirq
    - network packets in linux processed via softirq
    - if saturated:
        - packet drop
        - latency increase

10. DNS failover vs BGP failover
    - DNS: slower, depends on TTL
    - BGP: faster, auto route

