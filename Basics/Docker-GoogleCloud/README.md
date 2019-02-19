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


Useful resources:

* [Docker CheatSheet][1]


[1]: https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet#links