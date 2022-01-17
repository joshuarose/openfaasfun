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