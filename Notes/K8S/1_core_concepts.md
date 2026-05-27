1. cluster architecture
    - master node
        - etcd cluster: store cluster config/state
        - kube scheduler: determine best node for new container deploy
        - kube-apiserver: orchestrating operation within cluster
        - coontroller
    - worker nodes: run containerized app
        - kubelet: listen for instruction from kube-apiserver, manage node lifecycle, report node status
        - kube-proxy: communication within cluster

2. docker vs containerD
    - container runtime interface (CRI)
        - other container runtime need to adhere to Open container initiative (OCI) to run on k8s
        - image spec
            - how image should built
        - runtime spec
            - how container runtime should be developed
        - but docker is not built for CRI!!
            - but k8s have dockershim to support it
    - docker:
        - not just a container runtime
        - has:
            - docker CLI
            - docker API
            - build tool
            - containerd
                - CRI compatible, can work with k8s
                - can be used as runtime as its own, separated from docker
        - k8s v1.24: remove dockershim!!!
            - those docker image that built before remove continue to work
            - because they follow OCI standard, continue to work with containerd
    - containerd
        - part of docker, but now is separated project, member of CNCF
        - can install containerd without docker
        - ctr:
            - cli tool for debugging, come with containerd
            - limited features
        - nerdctl
            - better cli tool
            - support almost all options
        - crictl
            - cli for CRI compatible container runtime

3. docker deprecation
    - containerd support CRI
    - k8s deprecate support of full docker
        - not require docker as the runtime

4. etcd
    - reliable, distributed, key-value store
    - how to use
        - for windows: download from https://github.com/etcd-io/etcd/releases?utm_source=chatgpt.com
        - double click on `etcd.exe` to start
        - to set: `.\etcdctl.exe put mykey hello` (old version is 'set')
        - to get: `.\etcdctl.exe get mykey`

    - etcd in k8s
        - store info for cluster
        - add nodes/ deploy pods/ replica sets updated in etcd server
    - 2 types of k8s deployment method
        1. from scratch
            - deploy etcd by download its binary, and configuring etcd in masternode ourself
            - `advertise-client-urls`: address that etcd listen
            - `--initial-cluster`: allow etcd to know peer, for HA
            `--initial-cluster controller-0=https://${CONTROLLER0_IP}:2380,controller-1=https://${CONTROLLER1_IP}:2380 \`
        2. using kubeadm tool
            - etcd will be deploy as pod in servers
                - called static pod: 
                    - no deployment/replicaSet, direct created by kubelet from local file `/etc/kubernetes/manifests/etcd.yaml`
                        - normal pod source of truth in etcd, static pod in local file
                    - normal k8s pod created with k8s api `kubectl apply -f nginx.yaml`
                    - normal pod need API server, static no need
                    - kubelet start etcd, kube-apiserver, scheduler, controller-manager from local manifests, as statis pod
                    - kubelet restart static pod directly, and no scaling by scheduler, need to have manifest on nodex
                - usually in `/etc/kubernetes/manifests/`, if see etcd.yaml, kubelet launch container
            - for manual way: can run as a linux service or containerize as a pod

5. Kube-apiserver
    - primary management component
    - kubectl command reach here, kube api server validate it, then get data from etcd and response
    - the only component that interact directly with etcd
    - if using kubeadm tool, will be simple, deployed as a pod in kube-system namespace

6. kube -controller-manager
    - manage various controller in k8s
    - controller: continuos monitor state of component and update the status when `kubectl get nodes`
    - can be run as a pod ( setup with kubeadm)
    - can be run as a system service

7. kube- scheduler
    - decide which pod go which node ( only decide)
        - depends on criteria
            - CPU/ memory requirement
            - calculate amount of recourse that free after placing pod on it, higher better
            - criteria can be customized
    - kubelet: create and move pod

8. kubelet
    - lead all activity
    - kubelet in worker: register node with cluster, request container runtime when need to load pod
    - kubelet need to manual install, eventhough use kubeadm

9. kube-proxy
    - within cluster, every pod can reach another pod
    - ip of each node is not guarantee, need to use a service to expose and connect
        - virtual component in k8s memory
    - process that runs on each node in k8s cluster
        - look for new service
        - create rule on each node to forward traffic
        