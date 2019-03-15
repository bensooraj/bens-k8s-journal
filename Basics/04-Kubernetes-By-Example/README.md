## Kubernetes By Example

This is a journal of me walking through the entire [Kubernetes By Example][1] exercises on `Google Kubernetes Engine`. Here's their [GitHub repository][2].

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





[1]: http://kubernetesbyexample.com
[2]: https://github.com/openshift-evangelists/kbe