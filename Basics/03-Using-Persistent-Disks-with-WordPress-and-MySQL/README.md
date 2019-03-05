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

Creating a GKE cluster using `gcloud` automatically makes an entry in the kubconfig file and also set the current context for `kubectl`. Let's verify:
```sh
# Get all clusters the local kubectl instance has accesses to. Pay attention the listing gke_kubernetes-practice-219913_asia-south1-a_k8s-wordpress
$ kubectl config get-clusters
NAME
docker-for-desktop-cluster
gke_kubernetes-practice-219913_asia-south1-a_k8s-wordpress
minikube

# Checkout all the the existing context. You can use `gcloud config view` as well.
$ kubectl config get-contexts
CURRENT   NAME                                                         CLUSTER                                                      AUTHINFO                                                     NAMESPACE
          docker-for-desktop                                           docker-for-desktop-cluster                                   docker-for-desktop                                           
*         gke_kubernetes-practice-219913_asia-south1-a_k8s-wordpress   gke_kubernetes-practice-219913_asia-south1-a_k8s-wordpress   gke_kubernetes-practice-219913_asia-south1-a_k8s-wordpress   
          minikube                                                     minikube                                                     minikube                                                     

# And the current context is set to:
$ kubectl config current-context
gke_kubernetes-practice-219913_asia-south1-a_k8s-wordpress
```
To use an existing GKE cluster, use the following command `gcloud container clusters get-credentials [cluster-name]`.

### Create `PersistentVolumes` and `PersistentVolumeClaims`

Important points to remember:
* When a `PersistentVolumeClaim` is created, if there is no existing `PersistentVolume` for it to bind to, k8s dynamically provisions a new `PersistentVolume` based on the `StorageClass` configuration.
* When a `StorageClass` is not specified in the `PersistentVolumeClaim`, the cluster's default `StorageClass` is used instead.
* GKE(GKE) has a default `StorageClass` installed that dynamically provisions `PersistentVolumes` backed by [persistent disks][3].

The `mysql-volumeclaim.yaml` file would like the following:
```yaml
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: mysql-volumeclaim
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 200Gi
```
and the `wordpress-volumeclaim.yaml`:
```yaml
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: wordpress-volumeclaim
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 200Gi
```
Now:
```sh
# deploy the two PersistentVolumeClaim manifests
$ kubectl apply -f pvcs/mysql-volumeclaim.yaml
persistentvolumeclaim/mysql-volumeclaim created

$ kubectl apply -f pvcs/wordpress-volumeclaim.yaml
persistentvolumeclaim/wordpress-volumeclaim created

# List out the [PersistentVolumeClaim]s
$ kubectl get pvc
NAME                    STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
mysql-volumeclaim       Bound    pvc-8cb222da-3f7d-11e9-bbe6-42010aa001a3   200Gi      RWO            standard       12m
wordpress-volumeclaim   Bound    pvc-9174b903-3f7d-11e9-bbe6-42010aa001a3   200Gi      RWO            standard       12m

# And of course, the volumes themselves
$ kubectl get persistentvolumes
NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                           STORAGECLASS   REASON   AGE
pvc-8cb222da-3f7d-11e9-bbe6-42010aa001a3   200Gi      RWO            Delete           Bound    default/mysql-volumeclaim       standard                25m
pvc-9174b903-3f7d-11e9-bbe6-42010aa001a3   200Gi      RWO            Delete           Bound    default/wordpress-volumeclaim   standard                25m
```

### Set up MySQL

Get ready with the secrets, this will be passed on to the MySQL docker container.
```sh
# Create a secret for MySQL DB root password:
$ kubectl create secret generic mysql --from-literal=password=dbpassword
secret/mysql created

# List the secrets
$ kubectl get secrets
NAME                  TYPE                                  DATA   AGE
default-token-lk5xg   kubernetes.io/service-account-token   3      1h
mysql                 Opaque                                1      28s
```

[1]: https://cloud.google.com/kubernetes-engine/docs/tutorials/persistent-disk
[2]: https://kubernetes.io/docs/concepts/storage/storage-classes/
[3]: https://cloud.google.com/persistent-disk/
[4]: https://aws.amazon.com/ebs/
[5]: https://kubernetes.io/docs/tutorials/stateful-application/basic-stateful-set/