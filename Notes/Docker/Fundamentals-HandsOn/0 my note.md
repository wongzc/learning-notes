1. in docker file, most instruction run only in build time, except
    - `CMD` (`CMD ["npm", "start"]`)
    - `ENTRYPOINT` (`ENTRYPOINT ["nginx", "-g", "daemon off;"]`)
    - `HEALTHCHECK` ( run multiple time at runtime, not once, `HEALTHCHECK --interval=30s CMD node healthcheck.js`)
    - `STOPSIGNAL` ( run only when stop the container, `STOPSIGNAL SIGTERM`)

2. build vs run
    - build: `docker compose build` or `docker build -t my-app .`
    - run: `docker compose up` or `docker run -p 3000:3000 my-app`
        - restart: `docker restart <container>`

3. docker basic command
    - `docker info`: display system info
    - `docker version`
    - `docker login` or `docker login -u <username>`

4. docker run [FLAGS]
    - `-d`: detached mode, no logs print from container
    - `-p`: port mapping, `docker run -p 8080:80 nginx`, map localhost 8080 to container 80
    - `--name`: custom name for container `docker run --name myweb nginx`
    - `-v`: mount volume, `docker run -v /host/data:/container/data mysql`
    - `-e`: set env variable, `docker run -e MYSQL_ROOT_PASSWORD=123456 mysql`
    - `--rm`: auto remove when container exit
    - `-it`: i, interactive mode, t, give shell like terminal
    - `-u`: run as specify user, not root (UID 0, GID 0)
        - `docker run -u 1000:1000 app` UID 1000, GID 1000

5. docker ps/ docker images [FLAGS]
    - `-a`: show running ( default)+stopped, exited container
    - `-q`: quiet mode, show only docker id
    - `-f`: filter container
        - `docker ps -f "status=exited"`
        - `docker ps -f "name=myapp"`
        - `docker ps -f "ancestor=nginx"`
    - `-n`: limit number, `docker ps -n 5`

6. execute with interactive shell
    `docker exec -it  <container_name or id> bash`

7. docker stop
    - `-t`: seconds to wait before kill, `docker stop -t 5 myweb` (5 seconds)
    - stop multiple docker
        - `docker stop $(docker ps -q) `
        - `docker stop myweb redis`

8. docker rm
    - `f`: force remove, kill container first
    - `v`: remove anonymous volume attached

9. docker rmi
    - `-f`: force
    - `--no-prune`: do not remove untagged parent image

10. compose file:
    - top level:
        - `services`: containers
        - `volumes`: defined name volume that used by services
        - `networks`: custom network for services to join
        - `configs`: for Swarm related config data
        - `secrets`
    - `services`:
        - `build`: build from docker file
            - context: path to docker file folder
            - dockerfile: custom docker file name
            - args: build argument that pass to docker file
            - target: build specific target stage (AS xxx)
        - `image`: use prebuilt image
        - `command`: override CMD in docker file
        - `environment`: environment variable
        - `ports`: port mapping, local: container
        - `secrets`: what secret the service can use
        - `networks`: which networks the services can use
        - `depends_on`: which container need to be up first
        - `volumes`: 
            - `- ./db:/etc/data`: unnamed volume
            - `- db-data:/var/lib/mysql` ( with `db-data` declared at `volumes`): named volume
        - `restart`:
            - restart policy for service, default no restart
            - `always`
            - `on-failure`: restart if exit code indicate error
            - `unless-stopped`: restart, unless service stopped/ removed

11. k8s
    - `kubectl cluster-info`: to get k8s info
    - context: (cluster + user + namespace of cluster)
        - `kubectl config get-contexts` or `kubectx`: list all context
        - `kubectl config current-context`: get current context
        - `kubectl config use-context [contextName]` or `kubectx [contextName]`: set current context
        - `kubectl config delete-context [contextName]` or `kubectx -d [contextName]`: delete context from config
        - `kubectx -`: switch back to prev context
        - `kubectl config rename-context [old-name] [new-name]` or `kubectx [old-name]=[new-name]`: rename context
    - install kubectx: `choco install kubectx-ps`
    - cluster context at: C:\Users\wzcla\.kube\config
    - create resource
        - imperative way
            - use kubectl + series of command to create resource
            - like code, good for learning, testing, troubleshoot
        - declarative way
            - use kubectl + YAML file
            - reproducible, repeatable, can save in source control
            - like data that can be parsed and modified
        - create file: `kubectl create -f <file.yaml>`
        - for YAML, go to https://kubernetes.io/docs to search and copy
            - or use template by vscode
            - or use kubectl to generate
                - `kubectl create deploy mynginx --image=nginx --port=80 --replicas=3 --dry-run=client -o yaml`
                    - use `>deploy.yaml` to send it to file
                    - then can apply using `kubectl apply -f deploy.yaml`
    - namespace: (separation of resource in the cluster)
        - allow to group resource, like Dev, Test, Prod
        - object can access object in different namespace
        - delete namespace will delete all child objects
        - to use namespace
            - create namespace ( with YAML)
            - specify namespace when defining objects
        - can assign network policy and limit resource for a namespace
        - commands:
            - `kubectl get namespace` - List all namespaces in current context cluster
            - `kubectl get ns` - Shortcut
            - `kubectl config set-context --current --namespace=[namespaceName]` - Set the current context to use a namespace
            - `kubectl create ns [namespaceName]` - Create a namespace
            - `kubectl delete ns [namespaceName]` - Delete a namespace
            - `kubectl get pods --all-namespaces` - List all pods in all namespaces
    - Master node
        - dont run container in master node
        - also known as control panel
        - kube-control-manager:
            - controller of controller
            - run node, replication endpoint, service account & token controller
        - cloud-control-manager: interact with cloud provider controller
            - node: check cloud provider to determine if node deleted after it stop respond
            - route: set route in underlying cloud infra
            - service: create, update, delete cloud provider load balancer
            - volume: create, attach, mount, interact with volume
        - kube-scheduler:
            - assign node for newly create pods
            - consider requirement like hardware, software, data locality, affinity, anti-affinity, resource requirement
        - kube-apiserver:
            - expose REST interface, save state to etcd
            - client interact with it
        - etcd: key value datastore for cluster state data
            - only kube-apiserver communicate with it
            - single source of truth
        - Addons:
            - DNS, Web UI, Cluster-level logging, resource monitoring
    - Worker node
        - container runtime
            - Multiple type of k8s container runtime like
                - Moby, Containerd, Cri-0, Rkt, Kata, Virtlet
        - kubelet
            - manage pod lifecycle
            - ensure container described in pods specs are running well
        - kube-proxy
            - network proxy that manage network rules on nodes
            - all network traffic go through kube-proxy
    - Nodes pool
        - group virtual machine with same size
        - cluster can have multiple node pools
            - each pool different size of VM
            - each pool autoscaled independently
    - Pods
        - smallest unit of work of k8s
        - 1 or more container
            - inside pod, container share same IP, same volume
            - containers communicate with Localhost, IPC
            - if multiple container, usually 1 is main worker, another is helper
        - unit of deployment
        - if pods fail/ need to update, just delete it and replace with new one, new IP
        - scale by more pods, not more container in a pod
        - node can have multiple pods
    - Pod lifecycle
        - CREATION:
            1. user send request to create pod:
                - `API server` receive, validate and store the desired state in `etcd`
            2. `scheduler`:
                - check from `API server` (which get from `etcd`) for pending pod
                - choose best node for the pod, based on:
                    1. Node request limit ( cpu, memory)
                    2. Node selector (custom label like disktype, strict)
                    3. Affinity & Anti-Affinity: prefer/avoid for node, not strict
                    4. Taint & Toleration: 
                        - taint, set on node to only accept certain pod
                        - toleration, set on pod so that can run on tainted node
                - update Pod object in `API server`, then save on `etcd`
            3. `kubelet` on worker node:
                - watch `API server` & detect new pod assigned
                - pull pod spec and instruct `Runtime` to create container
            4. `Runtime`:
                - create container: pull image, setup namespaces, cgroups
                - expose container status back to `kubelet`
            5. `kubelet`:
                - periodically send container status to `API server`
                - `API server` update `etcd` with current state
        - DELETION:
            1. user send request to delete pod:
                - `API server` receive, mark pod with deletion timestamp, set pod terminating
                - state save in `etcd`
                - `API server` notify `Endpoint controller` and `kubelet`
                - `Endpoint controller` remove pod IP from service endpoint list
            2. Grace period
                - `kubelet` detect pod terminating, send SIGTERM to containers in pod to allow graceful shutdown
                - `Runtime` receive SIGTERM and allow container exit gracefully
                -  pod remain 'Terminating' state in `etcd` until fully stop
            3. after grace period:
                - `kubelet` send SIGKILL to force terminate container
                - `Runtime` kill container and clean up resource
                - `kubelet` update `API Server` after container deleted.
                - `API server` update `etcd`
    - Pod state
        - `Pending`: accepted but not yet created
        - `Running`: bound to node
        - `Succeeded`: exit with status 0, no error
        - `Failed`: all container exit and at least 1 exit with non-0 status
        - `Unknow`: cant communicate with pod
        - `CrashLoopBackOff`: keep start and crashing cycle
    - Define and run pods:
        - use YAML, kind: Pod, 
        - in spec, specify image
        - set container port, label, env variable etc
        - command
            - `kubectl create -f [pod-definition.yml]` - Create a pod (declarative), fail if alr exist
            - `kubectl apply -f [pod-definition.yml]` - Create a pod (declarative) and update if alr exist
            - `kubectl run [podname] --image=busybox -- /bin/sh -c "sleep 3600"` - Run a pod (imperative)
            - `kubectl get pods` - List the running pods
            - `kubectl get pods -o wide` - Same but with more info
            - `kubectl describe pod [podname]` - Show pod info
            - `kubectl get pod [podname] -o yaml > file.yaml` - Extract the pod definition in YAML and save it to a file
            - `kubectl exec -it [podname] -- sh` - Interactive mode
            - `kubectl delete -f [pod-definition.yml]` - Delete a pod
            - `kubectl logs [podname] -c [containername]` - get logs for container
            - `kubectl delete pod [podname]` - Same using the pod's name
                - use option `--grace-period=xx` to set grace period
                - or `--wait=false` to immediate kill
    - init container
        - init a pod which used for dependency before application container run
            - one-time, short-lived setup task
            - DB migration, file system setup, dependency check, config, env setup etc
        - always run to completion, must success before next one start, else keep restart unless restartPolicy=Never
        - Probe are not supported
            - as init container not long lived, and always run to complete, with auto restart logic
    - selectors
        - labels: key value pair to identify
        - selector use label to filter/ select object
        - example: `nodeSelector`,  `selector` etc
    - K8S object kind
        - `Pod`: smallest deployment unit, run container
        - `Deployment`: manage replica set and rolling updates for Pods
        - `Service`: expose Pods to network
        - `ConfigMap`: store non-confidential key-value configuration data
        - `Secret`: sensitive information
        - `PersistentVolumeClaim`: PVC, storage resource
    - multi container pod
        - helper process to help main app
        - pattern
            - sidecar: enhance feature, eg: logging
            - adapter: transform data
            - ambassador: proxy to external services
    - networking concept
        - flat network
        - can communicate to each other:
            - all containers within a pod
                - each pod get IP address, containers in pod share IP address, but different port number
                - container communicate with each using localhost+port
            - all pods with each other
                - need to go through service to communicate
            - all nodes with all pods
        - pods use ephemeral IP
        - service use persistent IP
    - workloads
        - app run on k8s
        - `Pod`: 
            - represent set of running containers
        - `ReplicaSet`: 
        - `Deployment`: manage `ReplicaSet`, stateless
        - `StatefulSet`: like `Deployment`, but stateful, Pod get stable DNS name
        - `DaemonSet`: ensure 1 pod on every selected Node, like logging, monitoring agent
    - ReplicatSet
        - ensure preset number of pods running at the same time
        - recommended way is to use deployment to create
        - spec>template> put metadata and spec of normal pod
        - Command
            - `kubectl apply -f [definition.yaml]` - Create a ReplicaSet
            - `kubectl get rs` - List ReplicaSets
            - `kubectl describe rs [rsName]` - Get info
            - `kubectl delete -f [definition.yaml]` - Delete a ReplicaSet
            - `kubectl delete rs [rsName]` - Same but using the ReplicaSet name
    - Deployments
        - Pod:
            - cant self heal, scale, update, roll back
            - but deployment can
        - manage single pod template
        - deployments create `ReplicaSets` in the background
            - but we let deployment manage it and dont interact with it
        - `ReplicaSets` provide self heal and scaling,  `Deployment` help for rolling updates and rollback
        - in yaml:
            - replicas: number of pod instance
            - revisionHistoryLimit: number of old replicaset to keep
            - strategy:
                - RollingUpdate: cycle update pods 1 by 1.
                - Recreate: kill all pods before new one created. 
        - spec>template> put pod meta data here
        - command:
            - `kubectl create deploy [deploymentName] --image=busybox --replicas=3 --port=80` - The imperative way
            - `kubectl apply -f [definition.yaml]` - Create a deployment
            - `kubectl get deploy` - List deployments
            - `kubectl describe deploy [deploymentName]` - Get info
            - `kubectl get rs` - List replicasets
            - `kubectl delete -f [definition.yaml]` - Delete a deployment
            - `kubectl delete deploy [deploymentName]` - Same but using the deployment name

    

