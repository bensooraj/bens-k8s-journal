## Kubernetes By Example

This is a journal of me walking through the entire [Kubernetes By Example][1] exercises on `Google Kubernetes Engine`. Here's their [GitHub repository][2].

1. [Check config details](#check-config-details)
2. [Spin-up a k8s cluster (GKE)](#spin-up-a-k8s-cluster-(gke))
3. [Pods](#pods)
4. [Labels](#labels)

### Check config details
```sh
# List the active accounts:
$ gcloud auth list
   Credentialed Accounts
ACTIVE  ACCOUNT
*       xxyyzz@gmail.com

To set the active account, run:
    $ gcloud config set account `ACCOUNT`

# Checkout the project we are currently in
$ gcloud config list project 
[core]
project = kubernetes-practice-219913

Your active configuration is: [default]

# List the default/current config values (I wanted the zone and region details):
$ gcloud config configurations list
NAME     IS_ACTIVE  ACCOUNT               PROJECT                     DEFAULT_ZONE   DEFAULT_REGION
default  True       xxyyzz@gmail.com  kubernetes-practice-219913  asia-south1-a  asia-south1
```

### Spin-up a k8s cluster (GKE)

```sh
# Create a 3-node cluster and set kubectl context
$ gcloud container clusters create k8s-by-example --num-nodes=3

# Creating cluster k8s-by-example in asia-south1-a... Cluster is being health-checked (master is healthy)...done.                                              
# Created [https://container.googleapis.com/v1/projects/kubernetes-practice-219913/zones/asia-south1-a/clusters/k8s-by-example].
# To inspect the contents of your cluster, go to: https://console.cloud.google.com/kubernetes/workload_/gcloud/asia-south1-a/k8s-by-example?project=kubernetes-practice-219913

kubeconfig entry generated for k8s-by-example.
NAME            LOCATION       MASTER_VERSION  MASTER_IP       MACHINE_TYPE   NODE_VERSION  NUM_NODES  STATUS
k8s-by-example  asia-south1-a  1.11.7-gke.4    35.200.190.186  n1-standard-1  1.11.7-gke.4  3          RUNNING
```

Creating a GKE cluster using `gcloud` automatically makes an entry in the kubconfig file and also set the current context for `kubectl`.

### Pods

> A pod is a collection of containers sharing a network and mount namespace and is the basic unit of deployment in Kubernetes. All containers in a pod are scheduled on the same node.

A dry-run `kubectl run sise --image=mhausenblas/simpleservice:0.5.0 --port=9876 --dry-run=true -o yaml`, gives the following yaml output:

```yaml
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    run: sise
  name: sise
spec:
  replicas: 1
  selector:
    matchLabels:
      run: sise
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        run: sise
    spec:
      containers:
      - image: mhausenblas/simpleservice:0.5.0
        name: sise
        ports:
        - containerPort: 9876
        resources: {}
status: {}
```

Let's run the pod using the image `mhausenblas/simpleservice:0.5.0`:
```sh
$ kubectl run sise --image=mhausenblas/simpleservice:0.5.0 --port=9876
kubectl run --generator=deployment/apps.v1beta1 is DEPRECATED and will be removed in a future version. Use kubectl create instead.
deployment.apps/sise created

# List out the pod
$ kubectl get po -o wide
NAME                   READY   STATUS    RESTARTS   AGE   IP          NODE                                            NOMINATED NODE
sise-bf8d99689-qgkkk   1/1     Running   0          43s   10.12.1.6   gke-k8s-by-example-default-pool-f7f7edae-09cs   <none>

# Grab the IP address
$ kubectl describe pods sise-bf8d99689-qgkkk | grep IP
IP:                 10.12.1.6

# Get inside the pod and access the API using the IP address.
# This is accessible from the cluster as well
$ kubectl exec -it sise-bf8d99689-qgkkk sh
> curl localhost:9876/info
{"host": "localhost:9876", "version": "0.5.0", "from": "127.0.0.1"}# 

> curl 10.12.1.6:9876/info
{"host": "10.12.1.6:9876", "version": "0.5.0", "from": "10.12.1.6"}# 

# List the deployments
$ kubectl get deployments.
NAME   DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
sise   1         1         1            1           20m

# And delete it
$ kubectl delete deployments sise
deployment.extensions "sise" deleted
```

#### Using a configuration file
```sh
# Apply a configuration to a resource by filename or stdin. The resource name must be specified. This resource will be created if it doesn't exist yet. JSON and YAML formats are accepted.
$ kubectl apply -f pod/pod.yaml
pod/twocontainers created

# List the pods
$ kubectl get pods -o wide
NAME            READY   STATUS    RESTARTS   AGE   IP          NODE                                            NOMINATED NODE
twocontainers   2/2     Running   0          1m    10.12.1.7   gke-k8s-by-example-default-pool-f7f7edae-09cs   <none>

# Get inside the container named 'shell' within the pod named 'twocontainers'
$ kubectl exec -it twocontainers -c shell -- bash
[root@twocontainers /]# curl localhost:9876/info
{"host": "localhost:9876", "version": "0.5.0", "from": "127.0.0.1"}

[root@twocontainers /]# curl 10.12.1.7:9876/info
{"host": "10.12.1.7:9876", "version": "0.5.0", "from": "10.12.1.7"}

# Clean up
$ kubectl delete pods twocontainers
pod "twocontainers" deleted
```

Creating pods with resource limits
```sh
# in the constraint-pod.yaml file:
      resources:
        limits:
          memory: "64Mi" 
          cpu: "500m"

# Create the pod
kubectl apply -f pod/constraint-pod.yaml

# List the pods
$ kubectl get pods
NAME                    READY   STATUS    RESTARTS   AGE
containers-constraint   1/1     Running   0          14m

# Clean up
$ kubectl delete pods containers-constraint
pod "containers-constraint" deleted
```

### Labels

> Labels are the mechanism you use to organize Kubernetes objects. A label is a key-value pair with certain restrictions concerning length and allowed values but without any pre-defined meaning.

```sh
# Create the pod using labels/labels-1.yaml
$ kubectl create -f labels/labels-1.yaml
pod/labelex created

# Check the pods created
$ kubectl get pods
NAME      READY   STATUS             RESTARTS   AGE
labelex   0/1     ImagePullBackOff   0          31s
```

Oops! Looks like I made some mistake while specifying the image for the container. Let me checkout what went wrong using the `describe` command:
```sh
$ kubectl describe pods labelex
Name:               labelex
Namespace:          default
Priority:           0
PriorityClassName:  <none>
Node:               gke-k8s-by-example-default-pool-41076e94-4n53/10.160.0.12
Start Time:         Fri, 15 Mar 2019 18:12:37 +0530
Labels:             env=development
Annotations:        <none>
Status:             Pending
IP:                 10.12.2.8
.
.
.
Events:
  Type     Reason          Age                   From                                                    Message
  ----     ------          ----                  ----                                                    -------
  Normal   Scheduled       5m10s                 default-scheduler                                       Successfully assigned default/labelex to gke-k8s-by-example-default-pool-41076e94-4n53
  Normal   SandboxChanged  5m1s (x2 over 5m3s)   kubelet, gke-k8s-by-example-default-pool-41076e94-4n53  Pod sandbox changed, it will be killed and re-created.
  Normal   Pulling         4m10s (x3 over 5m9s)  kubelet, gke-k8s-by-example-default-pool-41076e94-4n53  pulling image "mhausenblas/simpleservice:0.5."
  Warning  Failed          4m5s (x3 over 5m4s)   kubelet, gke-k8s-by-example-default-pool-41076e94-4n53  Failed to pull image "mhausenblas/simpleservice:0.5.": rpc error: code = Unknown desc = Error response from daemon: manifest for mhausenblas/simpleservice:0.5. not found
  Warning  Failed          4m5s (x3 over 5m4s)   kubelet, gke-k8s-by-example-default-pool-41076e94-4n53  Error: ErrImagePull
  Normal   BackOff         3m26s (x7 over 5m2s)  kubelet, gke-k8s-by-example-default-pool-41076e94-4n53  Back-off pulling image "mhausenblas/simpleservice:0.5."
  Warning  Failed          3s (x19 over 5m2s)    kubelet, gke-k8s-by-example-default-pool-41076e94-4n53  Error: ImagePullBackOff
```

Skimming through the `Events` section I found:

`Failed to pull image "mhausenblas/simpleservice:0.5.": rpc error: code = Unknown desc = Error response from daemon: manifest for mhausenblas/simpleservice:0.5. not found`


Lol! I mentioned the wrong image name(`mhausenblas/simpleservice:0.5.` instead of `mhausenblas/simpleservice:0.5.0`). Let me correct that and apply the changes:

```sh
# This time the image is successfully pulled
$ kubectl describe pods labelex
# Events:
#  Type     Reason          Age                   From                                                    Message
#   ----     ------          ----                  ----                                                    -------
#  Normal   Pulled          68s                   kubelet, gke-k8s-by-example-default-pool-41076e94-4n53  Successfully pulled image "mhausenblas/simpleservice:0.5.0"

# List the pod created
$ kubectl get pods
NAME      READY   STATUS    RESTARTS   AGE
labelex   1/1     Running   0          11m

# Show the labels as well
$ kubectl get pods --show-labels
NAME      READY   STATUS    RESTARTS   AGE   LABELS
labelex   1/1     Running   0          15m   env=development

# Filter by the label now
$ kubectl get pods -l env=development
NAME      READY   STATUS    RESTARTS   AGE
labelex   1/1     Running   0          16m

# Add a label to the pod
$ kubectl label pods labelex ownwer=bensooraj
pod/labelex labeled

# List them out again
$ kubectl get pods --show-labels
NAME      READY   STATUS    RESTARTS   AGE   LABELS
labelex   1/1     Running   0          17m   env=development,ownwer=bensooraj

# Filter by the new label.
$ kubectl get pods --selector ownwer=bensooraj
NAME      READY   STATUS    RESTARTS   AGE
labelex   1/1     Running   0          19m

```
I am really sorry for the spelling mistake with the label `ownwer=bensooraj`. It hurts my eyes.

Anyways, `--selector` and `-l` mean the same thing.

#### Set based selectors

> Kubernetes objects also support set-based selectors

We will launch another pod that has two labels (env=production and owner=bensooraj)
```sh
# Create a new pod using labels/labels-2.yaml
$ kubectl apply -f labels/labels-2.yaml

# List out all the pods along with the labels
$ kubectl get pods --show-labels
NAME       READY   STATUS    RESTARTS   AGE   LABELS
labelex    1/1     Running   0          57m   env=development,ownwer=bensooraj
labelex2   1/1     Running   0          2m    env=production,owner=bensooraj

# Let's get fancy here with selecting the labels
$ kubectl get pods --show-labels -l 'env in (development)'
NAME      READY   STATUS    RESTARTS   AGE   LABELS
labelex   1/1     Running   0          57m   env=development,ownwer=bensooraj

# The following lists all pods that are either labelled with env=development or with env=production
$ kubectl get pods --show-labels -l 'env in (development, production)'
NAME       READY   STATUS    RESTARTS   AGE   LABELS
labelex    1/1     Running   0          57m   env=development,ownwer=bensooraj
labelex2   1/1     Running   0          3m    env=production,owner=bensooraj
```

I can even delete pods like that:
```sh
$ kubectl delete pods -l 'env in (development, production)'
pod "labelex" deleted
pod "labelex2" deleted

# You can see them getting terminated
$ kubectl get pods -w
NAME       READY   STATUS        RESTARTS   AGE
labelex    1/1     Terminating   0          61m
labelex2   1/1     Terminating   0          6m34s
```

[1]: http://kubernetesbyexample.com
[2]: https://github.com/openshift-evangelists/kbe