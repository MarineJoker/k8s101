# k8s101
```bash
➜  k8s101 git:(main) ✗ docklser build -t httpserver:v3 .
➜  k8s101 git:(main) ✗ docker run -d -p 8080:8080 httpserver:v3
f9aa3d4484c7cfaf5d6f7c58b179029a36959e3df3cb17cb671ff95053546691
➜  k8s101 git:(main) ✗ curl 127.0.0.1:8080/localhost/healthz
访问成功%
```
