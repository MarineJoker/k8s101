```bash
admin:~/repo/k8s101/k8s# kubectl apply -f graceful-start.yaml
pod/httpserver-probe created
admin:~/repo/k8s101/k8s# kubectl apply -f service.yaml
service/httpserver-basic created
admin:~/repo/k8s101/k8s# kubectl get pod |grep httpserver-probe
httpserver-probe                    1/1     Running   0          5m35s
admin:~/repo/k8s101/k8s# kubectl get service httpserver-basic
NAME               TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
httpserver-basic   ClusterIP   10.105.91.189   <none>        80/TCP    40s
admin:~/repo/k8s101/k8s# curl 10.105.91.189
Welcome to the homepage!
```
