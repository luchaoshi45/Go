https://www.cnblogs.com/guangdelw/p/18222715 <br>
https://blog.csdn.net/SeeYouGoodBye/article/details/135706243 <br>
### ubuntu 22.04 k8s docker

#### 一 系统初始化

##### 1 root
```shell
#依次开机，设置ip地址，主机名
# 设置为root登录
sudo su
passwd root
```

##### 2 设置静态 IP 地址
```shell
vim /etc/netplan/00-installer-config.yaml

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

netplan apply
```

##### 3 换源
```shell
vim /etc/apt/sources.list

deb http://mirrors.aliyun.com/ubuntu/ jammy main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ jammy-updates main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ jammy-backports main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ jammy-security main restricted universe multiverse

# 源码包（如果不需要，可以去掉）
deb-src http://mirrors.aliyun.com/ubuntu/ jammy main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ jammy-updates main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ jammy-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ jammy-security main restricted universe multiverse

apt-get update
```

##### 4 修改主机名和 hosts
```shell
hostnamectl set-hostname kmaster1
hostnamectl set-hostname knode1
hostnamectl set-hostname knode2
cat >> /etc/hosts << EOF
192.168.1.200 kmaster1
192.168.1.201 knode1
192.168.1.202 knode2
EOF
```

##### 5 配置互信
```shell
ssh-keygen
ssh-copy-id kmaster1
ssh-copy-id knode1
ssh-copy-id knode2
```

##### 6 关闭 swap
```shell
sed -ri 's/^([^#].*swap.*)$/#\1/' /etc/fstab && grep swap /etc/fstab && swapoff -a && free -h
```

##### 7 设置内核参数
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

####  k8s 组件

##### 1 docker
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

cat > /etc/docker/daemon.json <<EOF
{
"registry-mirrors": [
"https://docker.mirrors.ustc.edu.cn",
"https://hub-mirror.c.163.com",
"https://reg-mirror.qiniu.com",
"https://registry.docker-cn.com"
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

systemctl restart docker.service
systemctl enable docker.service
docker info

```

##### 2 安装最新版本的kubeadm、kubelet 和 kubectl
```shell

apt-get update && apt-get install -y apt-transport-https
curl -fsSL https://mirrors.aliyun.com/kubernetes-new/core/stable/v1.30/deb/Release.key |
    gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
echo "deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://mirrors.aliyun.com/kubernetes-new/core/stable/v1.30/deb/ /" |
    tee /etc/apt/sources.list.d/kubernetes.list
    
apt-get update
apt-get install -y kubelet kubeadm kubectl

systemctl enable kubelet
```

##### 3 cri-dockerd
Kubernetes自v1.24移除了对docker-shim的支持，而Docker Engine默认又不支持CRI规范，<br>
因而二者将无法直接完成整合。为此，Mirantis和Docker联合创建了cri-dockerd项目，<br>
用于为Docker Engine提供一个能够支持到CRI规范的垫片，从而能够让Kubernetes基于CRI控制Docker。
```shell
wget https://github.com/Mirantis/cri-dockerd/releases/download/v0.3.14/cri-dockerd_0.3.14.3-0.ubuntu-jammy_amd64.deb
dpkg -i ./cri-dockerd_0.3.14.3-0.ubuntu-jammy_amd64.deb

sed -ri 's@^(.*fd://).*$@\1 --pod-infra-container-image registry.aliyuncs.com/google_containers/pause@' /usr/lib/systemd/system/cri-docker.service
systemctl daemon-reload && systemctl restart cri-docker && systemctl enable cri-docker
```

#### 初始化

##### 主机 生成初始化配置文件
```shell
kubeadm config print init-defaults > kubeadm.yaml
vim kubeadm.yaml


apiVersion: kubeadm.k8s.io/v1beta3
bootstrapTokens:
- groups:
  - system:bootstrappers:kubeadm:default-node-token
  token: abcdef.0123456789abcdef
  ttl: 24h0m0s
  usages:
  - signing
  - authentication
kind: InitConfiguration
localAPIEndpoint:
  # 修改成本master的ip
  advertiseAddress: 192.168.1.200
  bindPort: 6443
nodeRegistration:
  # 修改成cri-dockerd的sock
  criSocket: unix:///run/cri-dockerd.sock
  imagePullPolicy: IfNotPresent
  # 修改成本master的主机名
  name: master
  taints: null
---
apiServer:
  timeoutForControlPlane: 4m0s
apiVersion: kubeadm.k8s.io/v1beta3
certificatesDir: /etc/kubernetes/pki
clusterName: kubernetes
controllerManager: {}
dns: {}
etcd:
  local:
    # 修改etcd的数据目录
    dataDir: /data/etcd
# 修改加速地址
imageRepository: registry.aliyuncs.com/google_containers
kind: ClusterConfiguration
# 修改成具体对应的版本好
kubernetesVersion: 1.30.1
# 如果是多master节点，就需要添加这项，指向代理的地址，这里就设置成master的节点
controlPlaneEndpoint: "master:6443"
networking:
  dnsDomain: cluster.local
  serviceSubnet: 10.96.0.0/12
  # 添加pod的IP地址
  podSubnet: 10.244.0.0/16
scheduler: {}
# 在最后添加上下面两部分
---
apiVersion: kubeproxy.config.k8s.io/v1alpha1
kind: KubeProxyConfiguration
mode: ipvs
---
apiVersion: kubelet.config.k8s.io/v1beta1
kind: KubeletConfiguration
cgroupDriver: systemd


kubeadm init --config=kubeadm.yaml


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


kubeadm init --config=kubeadm.yaml

mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

##### 主机 生成初始化配置文件
```shell
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

kubeadm init \
--kubernetes-version=v1.30.1 \
--image-repository registry.aliyuncs.com/google_containers \
--pod-network-cidr=10.24.0.0/16 \
--ignore-preflight-errors=Swap \
--cri-socket unix:///run/cri-dockerd.sock \
--v=10


mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

# 检查是否成功
kubectl get nodes

```

##### node join
```shell
# 主机 生成 token
kubeadm token create --print-join-command

# 从机
kubeadm join 192.168.1.200:6443 \
--token oga6a9.rj3c68vib60ljqnb \
--discovery-token-ca-cert-hash sha256:61fd2fd1afe24ae90b177d1a93963fc6dcd53b90a0b0848c4e52c532febca40d \
--cri-socket unix:///run/cri-dockerd.sock

kubectl get nodes
kubectl get cs
```

##### 安装pod网络calico
```shell
wget https://docs.projectcalico.org/manifests/calico.yaml
vim calico.yaml 

# 修改 4061 行为 pod 子网地址
- name: CALICO_IPV4POOL_CIDR
  value: "10.244.0.0/16"
  
kubectl apply -f calico.yaml

watch kubectl get pods --all-namespaces -o wide
kubectl get nodes
kubectl get cs

kubectl run busybox --image busybox:1.28 --restart=Never --rm -it busybox -- sh
# 输入 nslookup kubernetes.default.svc.cluster.local
# 退出 exit
```







