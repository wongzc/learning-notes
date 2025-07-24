# L18-04

Let's deploy an Nginx container using both methods.

## Imperative

    kubectl create deployment mynginx1 --image=nginx

##  List the deployments in current Kubernetes namespace

    kubectl get deploy

## Declarative

    kubectl create -f deploy-example.yaml

## Cleanup

    kubectl delete deployment mynginx1
    kubectl delete deploy mynginx2

