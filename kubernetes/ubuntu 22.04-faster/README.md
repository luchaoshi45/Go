# ubuntu 22.04 k8s docker
### https://www.cnblogs.com/guangdelw/p/18222715 <br>
### https://blog.csdn.net/SeeYouGoodBye/article/details/135706243 <br>

## 一 系统初始化 🎰

### 1 base
```shell
sudo su
passwd root
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
192.168.1.201 knode1
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

### 安装containerd
```shell
sudo apt install -y containerd

# 生成containetd的配置文件
sudo mkdir -p /etc/containerd/
containerd config default | sudo tee /etc/containerd/config.toml >/dev/null 2>&1
# 修改/etc/containerd/config.toml，修改SystemdCgroup为true
sudo sed -i "s#SystemdCgroup\ \=\ false#SystemdCgroup\ \=\ true#g" /etc/containerd/config.toml
sudo cat /etc/containerd/config.toml | grep SystemdCgroup

# 修改沙箱镜像源
sudo sed -i "s#registry.k8s.io/pause#registry.cn-hangzhou.aliyuncs.com/google_containers/pause#g" /etc/containerd/config.toml
sudo cat /etc/containerd/config.toml | grep sandbox_image

# 重启containerd
systemctl restart containerd.service


# 创建 containerd.service.d 目录
mkdir /etc/systemd/system/containerd.service.d/

# 创建或编辑文件
cat > /etc/systemd/system/containerd.service.d/http-proxy.conf <<-EOF
[Service]
Environment="HTTP_PROXY=http://127.0.0.1:7897"
Environment="HTTPS_PROXY=http://127.0.0.1:7897"
Environment="NO_PROXY=noproxy_address>"
EOF

# 重启containerd
systemctl daemon-reload
systemctl restart containerd
```

### 安装工具
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

## 三 快照 node 节点 📷
### 1 配置 IP
```shell
# knode1
cat << EOF > /etc/netplan/00-installer-config.yaml
network:
  version: 2
  renderer: networkd
  ethernets:
    ens33:
      dhcp4: yes
      addresses:
        - 192.168.1.201/24
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
### 2 配置 节点 hostname
```shell
# knode1
hostnamectl set-hostname knode1
cat >> /etc/hosts << EOF
192.168.1.200 kmaster1
192.168.1.201 knode1
EOF
```

## 四 初始化 😀

### 主机 生成初始化配置文件
```shell
kubeadm init \
--kubernetes-version=v1.27.6 \
--image-repository registry.aliyuncs.com/google_containers \
--pod-network-cidr=10.244.0.0/16 \
--ignore-preflight-errors=Swap \
--v=10

mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

# 检查是否成功
kubectl get nodes
```

<img src="./k8s安装成功.png" alt="Image" style="width: 800px;">


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

### node join
```shell
# 主机 生成 token
kubeadm token create --print-join-command

# 从机
kubeadm join 192.168.1.200:6443 --token tk1bme.fu50ljfd97u4k920 --discovery-token-ca-cert-hash sha256:b4401951ee9aaf8f8c4c1b13aaa779950875f0e64e4d57060c7f80b2500a8814

kubectl get nodes
kubectl get cs
```

### 安装 flannel
```shell
ctr image pull docker.io/flannel/flannel:v0.25.5
ctr image pull docker.io/flannel/flannel-cni-plugin:v1.5.1-flannel1
kubectl apply -f https://github.com/flannel-io/flannel/releases/latest/download/kube-flannel.yml
```
