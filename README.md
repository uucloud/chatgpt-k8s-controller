# ChatGPT Kubernetes Controller

##### Translate to: [简体中文](README_zh.md)

ChatGPT Kubernetes Controller is a controller that quickly launches ChatGPT backend services and facilitates the management of multiple API token backends.

## Introduction

ChatGPT Kubernetes Controller generates the corresponding backend API service by listening to a custom resource definition (CRD) named ChatGPT in Kubernetes cluster.

This project requires Kubernetes version >= 1.25, and lower versions have not been tested.

## Deployment

Ensure that you have kubectl and cluster permissions:

```
git clone https://github.com/your-username/chatgpt-k8s-controller.git
make deploy
```

## Usage

After deploying the ChatGPT Kubernetes Controller, you can use the following steps to access the ChatGPT backend API:

1. Replace the token field in /artifacts/examples/example-chatgpt.yaml with your API key, for example:
```
sed -i 's/XXXXXX/YourAPIKeys/' artifacts/examples/example-chatgpt.yaml
```

2. Create a ChatGPT CRD:
```
kubectl apply -f artifacts/examples/example-chatgpt.yaml
```
3. View the corresponding Pod of the CRD, for example:
```
$ kubectl get pods -l controller=example-chatgpt -owide

IP           NODE        NOMINATED NODE   READINESS GATES
gpt-example-chatgpt-557cb86cf5-9zvpt   1/1     Running   0          42s   10.42.0.19   uucloudvm   <none>           <none>
```
4. Send a request to the backend API, for example:
```
curl -X POST -H "Content-Type: application/json" -d '{"question":"hello"}' http://10.42.0.19:8080

{"answer":"你好！有什么可以帮到你的吗？"}
```
Here, 10.42.0.19 is the IP address of the Pod corresponding to the CRD, which can be replaced according to the actual situation.

5. You can also create a service to avoid accessing the Pod IP directly.

## Updating

After updating the parameters, you can update the ChatGPT Kubernetes Controller by reapplying the corresponding CRD file.

## Uninstalling

Run the following command to uninstall ChatGPT Kubernetes Controller:
```
make undeploy
```

## License

This project is licensed under the terms of the Apache License 2.0. See the LICENSE file for details.