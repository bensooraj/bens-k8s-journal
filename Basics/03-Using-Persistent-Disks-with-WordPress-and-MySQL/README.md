## Using Persistent Disks with WordPress and MySQL

This journal records me working through the exercise at [Kubernetes Engine Tutorials => Using Persistent Disks with WordPress and MySQL][1].

The objective is to set up a single-replica WordPress deployment and a single-replica MySQL database on a GCE cluster, both using PersistentVolumes (PV) and PersistentVolumeClaims (PVC) to store/persist data beyond their pod's lifecycle.

The default [`StorageClass`][2](I don't really what this means, yet) dynamically creates [persistent disks][3] (I think these are equivalent to [EBS in AWS][4]) and create two PersistentVolumeClaims - one for each Deployment, WordPress and MySQL.

**Important** (Copy-Paste from the tutorial):
> Deployments are designed for stateless workloads. If a Deployment uses a persistent disk, it cannot scale beyond one replica, as persistent disks can be attached to only a single instance at a time in read/write mode. To scale past one instance, use a [StatefulSet][5], which attaches a different disk to each member of the set.

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
$ gcloud container clusters create k8s-wordpress --num-nodes=3

kubeconfig entry generated for k8s-wordpress.
NAME           LOCATION       MASTER_VERSION  MASTER_IP       MACHINE_TYPE   NODE_VERSION  NUM_NODES  STATUS
k8s-wordpress  asia-south1-a  1.11.7-gke.4    35.200.213.246  n1-standard-1  1.11.7-gke.4  3          RUNNING
```

[1]: https://cloud.google.com/kubernetes-engine/docs/tutorials/persistent-disk
[2]: https://kubernetes.io/docs/concepts/storage/storage-classes/
[3]: https://cloud.google.com/persistent-disk/
[4]: https://aws.amazon.com/ebs/
[5]: https://kubernetes.io/docs/tutorials/stateful-application/basic-stateful-set/