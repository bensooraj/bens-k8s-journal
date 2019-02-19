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
* `docker image ls`: List container image pulled from Docker Hub, Google Container Registry
* `docker ps [-a]` or `docker container ls [-a]`: List containers. Use the `-a` flag to list both running and exited containers.
* 


Useful resources:

* [Docker CheatSheet][1]


[1]: https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet#links