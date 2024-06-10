# 快速开始使用 k8s

## 1 Namespace
```shell
cat << EOF >> namespace.yaml 
apiVersion: v1
kind: Namespace
metadata:
  name: test-ns
EOF
```

## 2 Pod
```shell
kubectl run mynginx --image nginx:1.14.2
kubectl run mynginx --image nginx:1.14.2 -n test
kubectl get pods -owide

kubectl delete pod mynginx
kubectl run mynginx --image nginx:1.14.2
kubectl describe pod mynginx
```