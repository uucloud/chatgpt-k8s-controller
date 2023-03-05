# ChatGPT Kubernetes Controller
快速拉起ChatGPT后端服务的controller，方便管理多API Token后端

## 介绍
controller通过监听一种CRD(ChatGPT)，自动生成对应的后端API服务。

kubernetes >= 1.25可使用，更小的版本没有测试过

## 部署
确保拥有kubectl及集群权限:

```
git clone https://github.com/your-username/chatgpt-k8s-controller.git
make deploy
```

## 使用方式
可使用你的APIKeys替换`/artifacts/examples/example-chatgpt.yaml`中的`token`
例如：
```
sed -i 's/XXXXXX/YourAPIKeys/' artifacts/examples/example-chatgpt.yaml
```

创建chatgpt CRD:
```
kubectl apply -f artifacts/examples/example-chatgpt.yaml
```

这时集群中应该会有一个该CRD对应的pod，例如：
```
$ kubectl get pods -l controller=example-chatgpt -owide
NAME                                   READY   STATUS    RESTARTS   AGE   IP           NODE        NOMINATED NODE   READINESS GATES
gpt-example-chatgpt-557cb86cf5-9zvpt   1/1     Running   0          42s   10.42.0.19   uucloudvm   <none>           <none>
```

请求后端访问ChatGPT API
```
curl -X POST -H "Content-Type: application/json" -d '{"question":"hello"}' http://10.42.0.19:8080

{"answer":"你好！有什么可以帮到你的吗？"}
```

你也可以通过创建一个service来避免直接访问pod ip

## 更新
修改参数后重新apply对应CRD即可

## 卸载
```
make undeploy
```

## 许可证

本项目基于 Apache License 2.0 许可证。有关详细信息，请参见 LICENSE 文件。