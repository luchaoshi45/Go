# 快速开始使用 k8s
### https://kubernetes.io/zh-cn/docs/home/
### https://www.bilibili.com/video/BV1Ht421A7zA?p=7&vd_source=38ba9d0f671684c07ba19eab096ee0fe
## 1 Namespace
```shell
cat > namespace.yaml << EOF
apiVersion: v1
kind: Namespace
metadata:
  name: test-ns
EOF

kubectl apply -f namespace.yaml
kubectl get ns
kubectl delete -f namespace.yaml
```

## 2 Pod
```shell
kubectl run nginx-pod --image nginx:1.14.2
kubectl run nginx-pod --image nginx:1.14.2 -n test

kubectl delete pod nginx-pod
kubectl describe pod nginx-pod
kubectl logs nginx-pod -n default
```

```shell
cat > nginx-tomcat-pod.yaml << EOF
apiVersion: v1
kind: Pod
metadata:
  labels:
    run: nginx-tomcat
  name: nginx-tomcat
spec:
  containers:
  - name: nginx
    image: nginx:1.14.2
    ports:
    - containerPort: 80
  - name: tomcat
    image: tomcat
EOF

kubectl apply -f nginx-tomcat-pod.yaml
kubectl delete -f nginx-tomcat-pod.yaml
```

```shell
# 查看 pod ip
kubectl get pods -owide
curl -v 10.244.195.138
curl -v 10.244.195.138:80
curl -v 10.244.195.138:8080

kubectl exec -it nginx-tomcat -c nginx -- sh
kubectl exec -it nginx-tomcat -c tomcat -- sh

# tomcat 执行 证明 Pod 内部 容器之间的网络互通
curl localhost:80
```

## 3 Deployment
Deployment负责创建和更新应用程序的实例，使Pod拥有多副本，自愈，扩缩容等能力。创建Deployment后，Kubemetes Master 将应用程序实例调度
到集群中的各个节点上。如果托管实例的节点关闭或被删除，Deployment控制器会将该实例替换为群集中另一个节点上的实例。这提供了一种自我修复
机制来解决机器故障维护问题。I
```shell
kubectl create deployment deployment-tomcat --image=tomcat:9.0.55
kubectl create deployment nginx-deployment --image=nginx:1.14.2

# 另开窗监听
kubectl get pods -w
kubectl get deploy
kubectl delete pod  nginx-deployment-5d4b6f579d-gvz84

# 册除 deployment
kubectl delete deployment deployment-tomcat
kubectl delete deployment nginx-deployment
# 多副本
kubectl create deployment nginx-deployment --image=nginx:1.14.2 --replicas=3
kubectl describe pods nginx-deployment
```

```shell
cat > nginx-deployment.yaml << EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
EOF

kubectl apply -f nginx-deployment.yaml
kubectl delete -f nginx-deployment.yaml
```

```shell
kubectl scale --replicas=5 deployment/nginx-deployment

kubectl set image deployment/nginx-deployment nginx=nginx:1.16.1 --record
kubectl get deployment nginx-deployment
kubectl describe pod nginx-deployment-595dff4fdb-v6zt5
kubectl rollout undo deployment/nginx-deployment
```

```shell
# 在每个节点上的一个端口（由 Kubernetes 随机选择）上创建一个 NodePort Service，从而暴露你的 Deployment。
kubectl expose deployment nginx-deployment --type=NodePort --port=80
kubectl delete service nginx-deployment

kubectl get svc

node1 192.168.1.201
service 10.97.28.153
pod 10.244.195.154

# pod
curl 10.244.195.154:80
# service
curl 10.97.28.153:80
# node
curl 192.168.1.200:32056
```

## 4 Service

```shell
cat << EOF > tomcat-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: tomcat-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: tomcat
  template:
    metadata:
      labels:
        app: tomcat
    spec:
      containers:
      - name: tomcat
        image: tomcat:latest
        ports:
        - containerPort: 8080
EOF

cat << EOF > tomcat-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: tomcat-service
spec:
  selector:
    app: tomcat
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
  type: NodePort  # 可以是 NodePort, LoadBalancer 或 ClusterIP
EOF

kubectl apply -f tomcat-deployment.yaml
kubectl apply -f tomcat-service.yaml

kubectl delete -f tomcat-deployment.yaml
kubectl delete -f tomcat-service.yaml
```

## 5 Volume
```shell
cat << EOF > nginx-deployment-pvc.yaml
# PersistentVolume (PV) 定义
apiVersion: v1
kind: PersistentVolume
metadata:
  name: nginx-pv
spec:
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  storageClassName: manual
  hostPath:
    path: "/mnt/data"

---
# PersistentVolumeClaim (PVC) 定义
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: nginx-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
  storageClassName: manual

---
# Nginx Deployment 定义
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:latest
        ports:
        - containerPort: 80
        volumeMounts:
        - mountPath: "/usr/share/nginx/html"
          name: nginx-storage
      volumes:
      - name: nginx-storage
        persistentVolumeClaim:
          claimName: nginx-pvc
EOF

kubectl apply -f nginx-deployment-pvc.yaml
kubectl delete -f nginx-deployment-pvc.yaml

kubectl get pv
kubectl get pvc
kubectl get deployments
kubectl get pods
```

```shell
kubectl describe pv nginx-pvc

kubectl get pods -owie
# pvc 的节点上执行
cat << EOF > /mnt/data/index.html
Holle
EOF

curl 10.24.195.150
```

## 6 ConfigMap

```shell
cat << EOF > redis-config.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: redis-config
data:
  redis.conf: |
    bind 0.0.0.0
    port 6379
EOF

cat << EOF > redis-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: redis
spec:
  replicas: 1
  selector:
    matchLabels:
      app: redis
  template:
    metadata:
      labels:
        app: redis
    spec:
      containers:
      - name: redis
        image: redis:latest
        ports:
        - containerPort: 6379
        volumeMounts:
        - name: redis-config-volume
          mountPath: /usr/local/etc/redis/
          subPath: redis.conf
      volumes:
      - name: redis-config-volume
        configMap:
          name: redis-config
EOF

kubectl apply -f redis-config.yaml
kubectl apply -f redis-deployment.yaml

kubectl delete -f redis-config.yaml
kubectl delete -f redis-deployment.yaml
```

```shell
kubectl get pods
# 进入 pod
kubectl exec -it redis-5b5c997dfb-j92np -- /bin/bash

# 进入 redis
kubectl exec -it redis-5b5c997dfb-j92np -- redis-cli

KEYS *
SET mykey "Hello"
exit
```

## 7 Secret

```shell
# 创建Redis密码的Secret
kubectl create secret generic redis-password --from-literal=password=lcs

cat << EOF > redis-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: redis
spec:
  replicas: 1
  selector:
    matchLabels:
      app: redis
  template:
    metadata:
      labels:
        app: redis
    spec:
      containers:
        - name: redis
          image: redis:latest
          ports:
            - containerPort: 6379
          env:
            - name: REDIS_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: redis-password
                  key: password
EOF

kubectl apply -f redis-deployment.yaml
kubectl delete -f redis-deployment.yaml
```

```shell
kubectl get pods
# 进入 pod
kubectl exec -it redis-5b5c997dfb-j92np -- /bin/bash

# 进入 redis
kubectl exec -it redis-768c986f9d-tp9t4 -- redis-cli -a lcs

KEYS *
SET mykey "Hello"
exit
```

## 8 Ingress
```shell
https://github.com/kubernetes/ingress-nginx
# yaml 安装
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.35.0/deploy/static/provider/baremetal/deploy.yaml

# helm 安装
kubectl create namespace ingress-nginx

apt install snap
snap install helm --classic
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update

helm install ingress-nginx ingress-nginx/ingress-nginx \
    --namespace ingress-nginx \
    --set controller.service.type=LoadBalancer

# 验证安装 
kubectl get pods -n ingress-nginx
kubectl get svc -n ingress-nginx
```
