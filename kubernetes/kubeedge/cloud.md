# kubeedge 1.17 cloud 部署

## 一 准备
```shell
hostnamectl set-hostname kmaster1
cat >> /etc/hosts << EOF
192.168.1.200 kmaster1
192.168.1.201 kedge1
EOF
```

## 二 安装 cloud

### 1 patch
```shell
kubectl get daemonset -n kube-system |grep -v NAME |awk '{print $1}' | xargs -n 1 kubectl patch daemonset -n kube-system --type='json' -p='[{"op": "replace","path": "/spec/template/spec/affinity","value":{"nodeAffinity":{"requiredDuringSchedulingIgnoredDuringExecution":{"nodeSelectorTerms":[{"matchExpressions":[{"key":"node-role.kubernetes.io/edge","operator":"DoesNotExist"}]}]}}}}]'

kubectl get daemonset -n kube-flannel |grep -v NAME |awk '{print $1}' | xargs -n 1 kubectl patch daemonset -n kube-flannel --type='json' -p='[{"op": "replace","path": "/spec/template/spec/affinity","value":{"nodeAffinity":{"requiredDuringSchedulingIgnoredDuringExecution":{"nodeSelectorTerms":[{"matchExpressions":[{"key":"node-role.kubernetes.io/edge","operator":"DoesNotExist"}]}]}}}}]'

kubectl get daemonset -n metallb-system |grep -v NAME |awk '{print $1}' | xargs -n 1 kubectl patch daemonset -n metallb-system --type='json' -p='[{"op": "replace","path": "/spec/template/spec/affinity","value":{"nodeAffinity":{"requiredDuringSchedulingIgnoredDuringExecution":{"nodeSelectorTerms":[{"matchExpressions":[{"key":"node-role.kubernetes.io/edge","operator":"DoesNotExist"}]}]}}}}]'
```

### 2 keadm
```shell
wget https://github.com/kubeedge/kubeedge/releases/download/v1.17.0/keadm-v1.17.0-linux-amd64.tar.gz
tar -zxvf keadm-v1.17.0-linux-amd64.tar.gz
mv ./keadm-v1.17.0-linux-amd64/keadm/keadm /usr/local/bin/

# 192.168.1.210 ip-pool中没有被分配的ip地址
keadm init --advertise-address=192.168.1.210 --set iptablesHanager.mode="external" --kubeedge-version=1.17.0
```

### 3 准备
```shell
# 修改 svc 类型为 LoadBalancer
kubectl patch svc cloudcore -n kubeedge -p '{"spec": {"type": "LoadBalancer"}}'

wget https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
kubectl apply -f components.yaml
```

### 4 部署
```shell
kubectl patch deploy metrics-server -n kube-system --type='json' -p='[{"op":"add","path":"/spec/template/spec/containers/0/args/-","value":"--kubelet-insecure-tls"}]'

keadm gettoken
```

### 5 端口转发
```
iptables -t nat -A OUTPUT -p tcp --dport 10350 -j DNAT --to 192.168.1.210:10003
```

### 测试
```shel
cat << EOF > test_all.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx
spec:
  selector:
    matchLabels:
      app: nginx
  replicas: 1
  template:
    metadata:
      labels:
        app: nginx
    spec:
      nodeName: kedge1
      hostNetwork: true  # 使用宿主机网络
      containers:
        - name: nginx
          image: nginx:latest
          ports:
          - containerPort: 80

---

apiVersion: v1
kind: Service
metadata:
  name: nginx-svc
spec:
  type: LoadBalancer
  ports:
    - port: 80
      targetPort: 80
  selector:
    app: nginx
EOF

kubectl apply -f test_all.yaml
kubectl delete -f test_all.yaml
```
