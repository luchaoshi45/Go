# k8s 1.27 单机部署
- ubuntu 22.04
- https://www.cnblogs.com/guangdelw/p/18222715 <br>
- https://blog.csdn.net/SeeYouGoodBye/article/details/135706243 <br>
- https://blog.csdn.net/qq_41076892/article/details/133872947 <br>

## 一 系统初始化 🎰

### 1 root
```shell
sudo su
passwd root
systemctl stop ufw
systemctl disable ufw
apt update
apt install net-tools
apt install openssh-server
sed -i 's/^#\?PermitRootLogin .*/PermitRootLogin yes/' /etc/ssh/sshd_config
systemctl restart ssh
systemctl enable ssh
```

### 2 设置静态 IP 地址
```shell
cat >> /etc/netplan/00-installer-config.yaml << EOF
network:
  version: 2
  renderer: networkd
  ethernets:
    ens33:
      dhcp4: yes
      addresses:
        - 192.168.1.200/24
      routes:
        - to: 0.0.0.0/0
          via: 192.168.1.1
          metric: 100
      nameservers:
        addresses:
          - 8.8.8.8
          - 8.8.4.4
EOF

netplan apply
```

### 3 换源
```shell
cat << EOF > /etc/apt/sources.list
deb http://mirrors.aliyun.com/ubuntu/ jammy main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ jammy-updates main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ jammy-backports main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ jammy-security main restricted universe multiverse

# 源码包（如果不需要，可以去掉）
deb-src http://mirrors.aliyun.com/ubuntu/ jammy main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ jammy-updates main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ jammy-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ jammy-security main restricted universe multiverse
EOF

apt-get update
```

### 4 修改主机名和 hosts
```shell
hostnamectl set-hostname kmaster1
cat >> /etc/hosts << EOF
192.168.1.200 kmaster1
192.168.1.201 kedge1
EOF
```

### 5 关闭 swap
```shell
sed -ri 's/^([^#].*swap.*)$/#\1/' /etc/fstab && grep swap /etc/fstab && swapoff -a && free -h
```

### 6 设置内核参数
```shell
cat >> /etc/sysctl.conf <<EOF
vm.swappiness = 0
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
net.bridge.bridge-nf-call-ip6tables = 1
EOF

cat >> /etc/modules-load.d/neutron.conf <<EOF
br_netfilter
EOF

#加载模块
modprobe  br_netfilter
#让配置生效
sysctl -p
```

##  二 安装 k8s 组件 ✨

### 1 安装 containerd
```shell
apt update
apt install -y ca-certificates curl gnupg lsb-release
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
apt-get update

sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt install apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo apt update
sudo apt-get update
apt install docker-ce docker-ce-cli containerd.io docker-compose -y

cat  << EOF > /etc/docker/daemon.json
{
"registry-mirrors": [
"https://docker.m.daocloud.io",
"https://index.docker.io/v1"
],
 "exec-opts": ["native.cgroupdriver=systemd"],
 "data-root": "/data/docker",
 "log-driver": "json-file",
 "log-opts": {
	 "max-size": "20m",
	 "max-file": "5"
	}
}
EOF
systemctl daemon-reload && systemctl restart docker
systemctl enable docker.service
docker info
```

### 2 安装 kubelet、kubeadm 和 kubectl
```shell
# 更新 apt 依赖
sudo apt-get update && sudo apt-get install -y apt-transport-https ca-certificates curl gpg

# 添加 Kubernetes 的 key
curl -fsSL https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg

# 添加 Kubernetes apt 仓库，使用阿里云镜像源
echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main' | sudo tee /etc/apt/sources.list.d/kubernetes.list

# 更新 apt 索引
sudo apt update

# 查看版本列表
apt-cache madison kubeadm

# 安装 1.27.6 版本的 kubelet、kubeadm 和 kubectl
sudo apt-get install -y kubelet=1.27.6-00 kubeadm=1.27.6-00 kubectl=1.27.6-00

# 锁定版本，不随 apt upgrade 更新
sudo apt-mark hold kubelet kubeadm kubectl

# kubectl 命令补全
sudo apt install -y bash-completion
kubectl completion bash | sudo tee /etc/profile.d/kubectl_completion.sh > /dev/null
. /etc/profile.d/kubectl_completion.sh
```

### 3 cri-dockerd
```shell
wget https://github.com/Mirantis/cri-dockerd/releases/download/v0.3.14/cri-dockerd_0.3.14.3-0.ubuntu-jammy_amd64.deb
dpkg -i ./cri-dockerd_0.3.14.3-0.ubuntu-jammy_amd64.deb

sed -ri 's@^(.*fd://).*$@\1 --pod-infra-container-image registry.aliyuncs.com/google_containers/pause@' /usr/lib/systemd/system/cri-docker.service
systemctl daemon-reload && systemctl restart cri-docker && systemctl enable cri-docker
```


##  三 安装 k8s
### 1 主机 生成初始化配置文件
```shell
kubeadm init \
--kubernetes-version=v1.27.6 \
--image-repository registry.aliyuncs.com/google_containers \
--pod-network-cidr=10.244.0.0/16 \
--ignore-preflight-errors=Swap \
--cri-socket unix:///run/cri-dockerd.sock \
--v=10

mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

# 检查是否成功
kubectl get nodes
```

```shell
# 删除重装 😭
systemctl stop kubelet
systemctl stop flanneld
systemctl stop etcd
systemctl stop kube-apiserver
systemctl stop kube-controller-manager
systemctl stop kube-scheduler
sudo kubeadm reset
sudo kubeadm init
sudo rm -rf /etc/kubernetes/
sudo rm -rf /var/lib/etcd/
rm -rf /etc/cni/net.d
rm -rf /var/lib/cni/
rm -rf /var/lib/kubelet/*
iptables -F
iptables -t nat -F
iptables -t mangle -F
iptables -X
rm -rf /var/lib/etcd
```

### 2 node join
```shell
# 主机 生成 token
kubeadm token create --print-join-command

# 从机
kubeadm join 192.168.1.200:6443 --token tk1bme.fu50ljfd97u4k920 --discovery-token-ca-cert-hash sha256:b4401951ee9aaf8f8c4c1b13aaa779950875f0e64e4d57060c7f80b2500a8814 \
--cri-socket unix:///run/cri-dockerd.sock

kubectl get nodes
kubectl get cs
```

### 3 安装 flannel
```shell
docker image pull docker.io/flannel/flannel:v0.25.5
docker image pull docker.io/flannel/flannel-cni-plugin:v1.5.1-flannel1
kubectl apply -f https://github.com/flannel-io/flannel/releases/latest/download/kube-flannel.yml
```

```shell
# 污点
kubectl describe node kmaster1 |grep Taints
kubectl taint nodes kmaster1 node-role.kubernetes.io/control-plane:NoSchedule-

vi /etc/kubernetes/manifests/kube-apiserver.yaml

# 在spec.containers.command的最后面加上
# - --service-node-port-range=1-65535

# 重启
systemctl daemon-reload
systemctl restart kubelet
```

### 4 部署 metallb
```shell
kubectl edit configmap -n kube-system kube-proxy
# strictARP: true
# mode: "ipvs"

kubectl rollout restart daemonset kube-proxy -n kube-system

wget https://raw.githubusercontent.com/metallb/metallb/v0.13.5/config/manifests/metallb-native.yaml
wget https://raw.githubusercontent.com/metallb/metallb/v0.14.5/config/manifests/metallb-native.yaml
kubectl apply -f metallb-native.yaml

cat << EOF > ip-pool.yaml
# ip-pool.yaml
apiVersion: metallb.io/v1beta1
kind: IPAddressPool
metadata:
  name: ip-pool
  namespace: metallb-system
spec:
  addresses:
    - 192.168.1.210-192.168.1.220 # 根据虚拟机的ip地址来配置 这些ip地址可以分配给k8s中的服务
EOF

cat << EOF > advertise.yaml
# advertise.yaml
apiVersion: metallb.io/v1beta1
kind: L2Advertisement
metadata:
  name: l2adver
  namespace: metallb-system
spec:
  ipAddressPools: # 如果不配置则会通告所有的IP池地址
    - ip-pool
EOF

kubectl apply -f ip-pool.yaml
kubectl apply -f advertise.yaml

kubectl get ipaddresspool -n metallb-system
```
```shell
# 测试机器ok
cat << EOF > test_metallb.yaml
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
      containers:
        - name: nginx
          image: nginx:latest

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

kubectl apply -f test_metallb.yaml

kubectl delete -f test_metallb.yaml
```