```bash
admin:~/repo/k8s101/k8s# kubectl get pod |grep httpserver-probe
httpserver-probe                    1/1     Running   0          5m35s
admin:~/repo/k8s101/k8s# kubectl get pod  httpserver-probe -oyaml|grep resources
{"apiVersion":"v1","kind":"Pod","metadata":{"annotations":{},"name":"httpserver-probe","namespace":"default"},
"spec":{"containers":[{"image":"httpserver:v6","name":"http-probe","readinessProbe":{"httpGet":{"path":"/localhost/healthz","port":8080},"initialDelaySeconds":30,"periodSeconds":5,"successThreshold":2},
"resources":{"limits":{"memory":"800Mi"},"requests":{"memory":"500Mi"}}}],"hostNetwork":true},"status":{"qosClass":"Burstable"}}
```
