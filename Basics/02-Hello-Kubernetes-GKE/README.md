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

