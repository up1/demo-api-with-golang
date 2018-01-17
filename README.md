# demo-api-with-golang

How to deploy ?

```
$docker build -t api:1.0 .
$docker run --rm -p 8080:8080  api:1.0
$cat deployment.yaml > tmp.yaml && echo "---" >> tmp.yaml && cat ingress.yaml >> tmp.yaml && echo "---" >> tmp.yaml && cat service.yaml >> tmp.yaml
$kubectl delete -f tmp.yaml
$kubectl apply -f tmp.yaml
```
