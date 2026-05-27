0. general
    - view pod running in `kube-system`: `kubectl get pods -n kube-system`
    - get nodes `kubectl get nodes`
    - get pods in namespace `kubectl get pods -n kube-system`
1. etcd
    - set key `./etcdctl put key1 value1`
    - get key `./etcdctl get key1`
    - view key stored in etcd (under namespace key-system): `kubectl exec etcd-master -n kube-system -- etcdctl get / --prefix --keys-only`