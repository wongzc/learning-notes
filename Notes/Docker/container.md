1. how container work under the hood
    - namespace
        - isolation, each container get its own view of system
        - have PID/net/UTS/MNT namespace
    - cgroups
        - resource control, on CPU, disk, memory , network etc
    - filesystem
        - image: layered filesystem (read only)
        - container: writable layer on top
        - use COW for efficient image usage
    - container runtime stack
        - docker CLI -> dockerd -> containerd -> runc -> linux kernel

2. Scheduling, orchestration, and cluster-level troubleshooting