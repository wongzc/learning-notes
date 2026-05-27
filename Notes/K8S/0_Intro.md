1. kodekloud notes
    - https://notes.kodekloud.com/docs/Certified-Kubernetes-Administrator-CKA/Introduction/Course-Introduction/page

2. k8s node and cluster and pod
    - cluster:
        - entire k8s system
        - with:
            - control plane ( master node)
            - work node
            - network/storage/API server/ scheduler/ etcd
    - node: 
        - usually separate physical machine/VM
        - each with own OS, kubelet, container runtime, network stack, storage
        - can be master/worker
    - pod:
        - smallest deployment unit in k8s
        - pod wraps 1 or more container
        - container inside share same ip, localhost, volumes
    - container:
        - actual running application process
            - nginx/ python/ postgres/ redis etc
            - created by containerd ( k8s ask runtime containerd to run container)
        - isolated namespace, cgroups, image filesystem

3. separate containers in same pod vs separate pods vs separate nodes?
    - same pod:
        - process that must live together
        - must scale together (app+log shipper)
    - different pod
        - logically separate service (frontend, backend, redis)
        - can scale independently
    - same node:
        - scheduler picks node (to fullfill deployment replica number)
        - by hardware like CPU etc
    - separate node
        - resource isolation: db that heavy CPU/ memory
        - hardware requirement: ML that need GPU
        - security, availability reason
4. others
    - deployment
        - dont create pod directly, create deployment
    - service
        - pod may change, ip may change
        - service give stable endpoints
    - container runtime
        - run container
        - containerd/ CRI-O
    - namespace
        - logical grouping within a cluster
        - for:
            - access control
            - resource limits
            - organization
            - network policy
        - use 1 namespace within cluster rather than 2 cluster when:
            - want simple control ( 1 cluster means 1 control plane, monitoring stack, etcd)
            - same trust boundary (for sharing infra)
            - resource share, ok with all share same cluster failure
        - use differnt cluster when:
            - security isolation reason
            - blast radius isolation ( failure)
            - diff k8s version/ infra cluster requirement/ regional separation

5. kubeadm
    - `kubeadm init`
    - bootstrap cluster, setup control plane
    - main phase it do
        1. preflight checks: check machine readiness
            - swap disabled
                - swap space at disk, to swap in active ram
            - required port free
            - kubelet installed
            - container runtime reachable
            - required kernel settings
            - hostname sanity
                - check if hostname duplicate/valid, hostname important for communication
            - CPU/ memory enough
        2. Generate cert
            - generate TLS cert for secure communication, between component
            - `/etc/kubernetes/pki/`
        3. generate kubeconfig files
            - create credentials/ config for components in `/etc/kubernetes`
        4. configure kubelet
        5. create static pod manifest
            - `/etc/kubernetes/manifests/` created
        6. kubelet start controlplane based on `/etc/kubernetes/manifests/`
            - look for `etcd.yaml` or `apiserver.yaml` etc
            - use container runtime to start container
        7. start etcd based on `etcd.yaml`
        8. start API server based on  manifest
        9. start controller manager
        10. start scheduler
        11. check control plane health
        12. initialize cluster state
        13. upload cluster config
        14. mark node as control plane
        15. create bootstrap token
        16. generate join command
        17. install core addons (DNS, kube-proxy)
        18. setup RBAC
        19. Configure node bootstrap auth
        20. Pull required images
