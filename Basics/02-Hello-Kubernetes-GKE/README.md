## Prep
List the active account and project:
```sh
$ gcloud auth list
   Credentialed Accounts
ACTIVE  ACCOUNT
*       xxxxxxxxxx@gmail.com

To set the active account, run:
    $ gcloud config set account `ACCOUNT`

$ gcloud config list project 
[core]
project = kubernetes-practice-219913

Your active configuration is: [default]
```

List the default/current config values (I wanted the zone and region details):
```sh
$ gcloud config configurations list

NAME     IS_ACTIVE  ACCOUNT               PROJECT                     DEFAULT_ZONE   DEFAULT_REGION
default  True       xxxxxxxxxx@gmail.com  kubernetes-practice-219913  asia-south1-a  asia-south1
```

## Docker operations

Build a docker image for the hello world app:
```sh
$ docker build -t asia.gcr.io/kubernetes-practice-219913/go-hello-world:v1 .
```
Test run the app:
```sh
$ docker run -p 8080:8000 --name my-golang-app asia.gcr.io/kubernetes-practice-219913/go-hello-world:v1
```
All looks good for now:
![Golang Hello World app running from Docker](imgs/docker_hello_world.png)
Push the image to the Google Container Registry
```
$ docker push asia.gcr.io/kubernetes-practice-219913/go-hello-world:v1
```
The image is uploaded to the Google Container Registry as well:
![Google Container Registry screenshot](imgs/gcr_1.png)

Check details of the image (CLI) available in the Google Container registry:
```sh
$ gcloud container images list --repository=asia.gcr.io/kubernetes-practice-219913
NAME
asia.gcr.io/kubernetes-practice-219913/go-hello-world

$ gcloud container images list-tags asia.gcr.io/kubernetes-practice-219913/go-hello-world
DIGEST        TAGS  TIMESTAMP
673ae58f6fda  v1    2019-02-25T09:39:36
```

## Kubernetes

Question: 
> How do I point the kubectl installed on my local machine to the `gcloud` cluster?

Answer (from the documentation):
> When you create a cluster using `gcloud container clusters create`, an entry is automatically added to the `kubeconfig` in your environment, and the current context changes to that cluster. This works only if you are creating the cluster from your local machine.

Spin up the 2 node cluster:
```sh
# n1-standard-1: Standard machine type with 1 vCPU and 3.75 GB of memory.
# zone: I am launching the cluster in the zone same as my project. Also, it is in the same region that the image is also uploaded.
$ gcloud container clusters create hello-world \
                --num-nodes 2 \
                --machine-type n1-standard-1 \
                --zone asia-south1-a
```
Got the following error:
> ERROR: (`gcloud.container.clusters.create`) ResponseError: code=403, message=Kubernetes Engine API is not enabled for this project. Please ensure it is enabled in Google Cloud Console and try again: visit https://console.cloud.google.com/apis/api/container.googleapis.com/overview?xxxxxxxxxxxxxxx to do so.

I enabled the `Kubernetes Engine API` from the UI and ran the command again. The console output looks something like below:
```sh
# Creating cluster hello-world in asia-south1-a... 
# Cluster is being health-checked (master is healthy)...done.                                                 
# Created [https://container.googleapis.com/v1/projects/kubernetes-practice-219913/zones/asia-south1-a/clusters/hello-world].
# To inspect the contents of your cluster, go to: https://console.cloud.google.com/kubernetes/workload_/gcloud/asia-south1-a/hello-world?project=kubernetes-practice-219913

kubeconfig entry generated for hello-world.
NAME         LOCATION       MASTER_VERSION  MASTER_IP      MACHINE_TYPE   NODE_VERSION  NUM_NODES  STATUS
hello-world  asia-south1-a  1.11.7-gke.4    35.244.37.152  n1-standard-1  1.11.7-gke.4  2          RUNNING
```
The `Kubernetes Engine => Clusters` UI looks like the following:
![GKE Cluster View](imgs/gke_1.png)


## Useful resources:

* [`gcloud`: Container Registry | Managing images][1]
* [`gcloud`: Pushing and pulling images][2]
* [`gcloud`: Compute Engine | Machine Types][3]

[1]: https://cloud.google.com/container-registry/docs/managing
[2]: https://cloud.google.com/container-registry/docs/pushing-and-pulling
[3]: https://cloud.google.com/compute/docs/machine-types