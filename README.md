# openfaasfun

This is just me having fun seeing what it would look like to run microfunctions using [openfaas](https://docs.openfaas.com) in [minikube](https://minikube.sigs.k8s.io/docs/)

###Goal
I want to deploy a Kanye West micro-function into a k8s cluster that can run anywhere and get some of the same benefits we see from AWS Lambda.


### Pre-Requisites
- [minikube](https://minikube.sigs.k8s.io/docs/)
- [golang](https://go.dev/doc/install)
- [kubectl]()
- [openfaas-cli](https://docs.openfaas.com/cli/install/)
- [helm](https://helm.sh/docs/helm/helm_install/)


### Process

We are setting up both the platform in k8s on minikube and deploying a micro-service to it. If you've used Lambda, or another faas provider you've likely only done the latter! (Credit to Alex Ellis and his article from 2017 https://faun.pub/getting-started-with-openfaas-on-minikube-634502c7acdf)

1.) Create openfaas namespaces in minikube. 
```bash
kubectl apply -f https://raw.githubusercontent.com/openfaas/faas-netes/master/namespaces.yml
```

2.) Add helm repo for openfaas
```bash
helm repo add openfaas https://openfaas.github.io/faas-netes/
```

3.) Update helm
```bash
helm repo update
```

4.) Generate openfaas password
```bash
export PASSWORD=$(head -c 12 /dev/urandom | shasum| cut -d' ' -f1)
```

5.) Make note of password
```bash
echo $PASSWORD
```
6.) Push password as k8s secret
```bash
kubectl -n openfaas create secret generic basic-auth --from-literal=basic-auth-user=admin --from-literal=basic-auth-password="$PASSWORD"
```

7.) Ensure Minikube Ingress is enabled (skip if you've already done this)
```bash
minikube addons enable ingress
minikube addons enable ingress-dns
```

8.) Update HOSTS file on your OS to setup local dns hostname and add the following line (/etc/hosts on MacOS c:\Windows\System32\Drivers\etc\hosts on Windows) Depending on your OS you may want to restart after.
```
127.0.0.1 gateway.openfaas.local
```

9.) Deploy the helm chart to your cluster
```bash
helm upgrade openfaas --install openfaas/openfaas --namespace openfaas --set functionNamespace=openfaas-fn --set basic_auth=true --set ingress.enabled=true
```

10.) Start the Minikube tunnel for ingress
```bash
minikube tunnel
```

11.) To test navigate to http://gateway.openfaas.local and login with username: admin and the password you generated in step 4

#### Congrats! you've got your own lambda platform, Now let's get Kanye wisdom

12.) Run the port forward for the CLI
```bash
kubectl port-forward -n openfaas svc/gateway 8080:8080 & 
```

13.) Authenticate OpenFAAS CLI against cluster
```bash
faas login --gateway http://localhost:8080 -u admin --password $PASSWORD
```

14.) Login to docker hub to build against your registry (replace environment variables with credentials or set them in EXPORT)
```bash
faas registry-login --username $DOCKER_USERNAME --password $DOCKER_PASSWORD
```

15.) Update kanye.yml, replace joshuarose with your docker username

16.) Deploy Kanye to your cluster
```bash
faas up -f kanye.yml --build-arg GO111MODULE=on
```

17.) Invoke Kanye from localhost:8080
![openfaas screenshot](https://github.com/joshuarose/openfaasfun/img/screenshot.png)