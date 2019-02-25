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








## Useful resources:

* [`gcloud`: Container Registry | Managing images][1]
* [`gcloud`: Pushing and pulling images][2]

[1]: https://cloud.google.com/container-registry/docs/managing
[2]: https://cloud.google.com/container-registry/docs/pushing-and-pulling