### `gcloud` SDK commands:

List the active account name:
```sh
$ gcloud auth list
   Credentialed Accounts
ACTIVE  ACCOUNT
*       xyz@gmail.com

To set the active account, run:
    $ gcloud config set account `ACCOUNT`
```

List the project ID
```sh
$ gcloud config list project 
[core]
project = kubernetes-practice-77777

Your active configuration is: [default]
```

### Test your Docker installation
```sh
$ docker run hello-world
.
..
Hello from Docker!
This message shows that your installation appears to be working correctly.
..
.
```

### Docker commands
```sh
# List container image pulled from Docker Hub, Google Container Registry
$ docker image ls

# List containers. Use the `-a` flag to list both running and exited containers.
$ docker ps [-a]
$ docker container ls [-a]

# Build a docker image. The -t is to name and tag an image with the name:tag syntax.
# The defailt tag is latest
$ docker build -t go-hello-world:0.1 .

# Run the image we've just built
# Checkout the service running using: curl http://localhost:8080
$ docker run -p 8080:8000 --name my-golang-app go-hello-world:0.1

# Stop and remove the running container
$ docker stop my-golang-app && docker rm my-golang-app
# Or
$ docker rm my-golang-app -f

# Container logs. Use -f for trailing/streaming logs
$ docker logs [-f] [container_id]
# Or
$ docker logs [-f] [container_name]


# Peep into a running container
$ docker exec -it [container_id] bash

# To check the container's metadata
$ docker inspect [container_id]
# Or
$ docker inspect [container_name]

# To filter out specifics
$ docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' [container_id]
```

### Google Container Registry

Command for configuring your docker to use `gcloud` as a credential helper:
```sh
$ gcloud auth configure-docker
```

GCR container tagging format: `[hostname]/[projectid]/[image]:[tag]`
* `[hostname]` => gcr.io
* `[project-id]` => project ID
* `[image]` => image name
* `[tag]` => any string of choice. Defaults to "latest".

```sh
# Get the project ID
$ gcloud config list project

# Re-tag the image
$ docker tag go-hello-world:0.1 gcr.io/kubernetes-practice-77777/go-hello-world:0.1

# Push the image to GCR
docker push gcr.io/kubernetes-practice-77777/go-hello-world:0.1

# CLEAN UP
# Stop and remove all containers:
$ docker stop $(docker ps -q)
$ docker rm $(docker ps -aq)

# Remove all locallly cached images
$ docker rmi go-hello-world:0.1 gcr.io/kubernetes-practice-77777/go-hello-world:0.1
$ docker rmi golang:1.8

# Execute the following command with caution. It will remove all remaining images.
# This may/may not be ideal depending on your setup/requirement.
$ docker rmi $(docker images -aq)

# BACK TO GCR
# Pull the image back from GCR to your local cache
$ docker pull gcr.io/kubernetes-practice-77777/go-hello-world:0.1
$ docker run -p 8080:8000 -d gcr.io/[project-id]/node-app:0.2 --name my-golang-app

# Check the service running at port 8000
$ curl http://localhost:8000
```

### Managing images in the Google Container Registry

Generic commands (there are lots more, I am only listing what I have used)
```sh
# To list the images that are in one of your host locations:
$ gcloud container images list --repository=[HOSTNAME]/[PROJECT-ID]

# To list an image's truncated digests and tags:
$ gcloud container images list-tags [HOSTNAME]/[PROJECT-ID]/[IMAGE]

# To delete an image (identified by its tag, and it has multiple tags) 
# from one of your Container Registry repositories:
$ gcloud container images delete [HOSTNAME]/[PROJECT-ID]/[IMAGE]:[TAG] --force-delete-tags

```


## Useful resources:

* [Docker CheatSheet][1] - Unofficial
* [`gcloud`: Managing images][2]
* [`gcloud`: Pushing and pulling images][3]

[1]: https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet#links
[2]: https://cloud.google.com/container-registry/docs/managing
[3]: https://cloud.google.com/container-registry/docs/pushing-and-pulling