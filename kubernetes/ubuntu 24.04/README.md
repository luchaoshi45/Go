### kubernetes install
https://www.augensten.online/655061ae/index.html#3-1-K8S-%E9%9B%86%E7%BE%A4%E8%BD%AF%E4%BB%B6-apt-%E6%BA%90%E5%87%86%E5%A4%87 <br>
https://www.bilibili.com/video/BV1LC4y1g7wz/?p=3&spm_id_from=pageDriver&vd_source=38ba9d0f671684c07ba19eab096ee0fe <br>

- kubernetes集群主要包含master（主控节点）和node（工作节点）组成，master和node一般是多对多或者一对多的模式，
- 为了学习使用这里选择1个master节点和多个node节点（这里一个节点相当于就是一台Linux系统主机）
![集群](kubernetes/install/框图.png)
#### 1 集群节点准备
```shell
# 配置节点主机名
hostnamectl set-hostname master && bash
hostnamectl set-hostname node1 && bash
hostnamectl set-hostname node2 && bash
```
```shell
# 修改 hosts 文件
vim /etc/hosts
# 添加
192.168.1.100 k8s-master01
192.168.1.101 k8s-worker01
192.168.1.103 k8s-worker02
# 实现三个节点可以相互 ping
```

```shell
# 关闭防火墙
sudo ufw status
sudo ufw disable

# ntpdate 计划时间同步
sudo apt update
apt-get install ntpdate
crontab -e
# 输入 0 0 * * * /usr/sbin/ntpdate ntp.aliyun.com

# 查看 os 内核版本在 5 以上
uname -r 
```

```shell
# 添加网桥过滤及内核转发配置文件
sudo tee /etc/sysctl.d/k8s.conf > /dev/null << EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
EOF

sysctl --system
sysctl net.bridge.bridge-nf-call-ip6tables
sysctl net.bridge.bridge-nf-call-iptables
sysctl net.ipv4.ip_forward
sysctl vm.swappiness

# 加载 br_netfilter 模块
modprobe br_netfilter
lsmod | grep br_netfilter
```

```shell
# 安装 ipset ipvsadm
sudo apt-get install ipset
ipset --version
sudo apt-get install ipvsadm
ipvsadm --version

# 配置 ipvsadm 模块加载
# 添加需要加载的模块
cat << EOF | tee /etc/modules-load.d/ipvs.conf
ip_vs
ip_vs_rr
ip_vs_wrr
ip_vs_sh
nf_conntrack
EOF

创建加载模块脚本文件
cat << EOF | tee /root/ipvs.sh
#!/bin/sh
modprobe -- ip_vs
modprobe -- ip_vs_rr
modprobe -- ip_vs_wrr
modprobe -- ip_vs_sh
modprobe -- nf_conntrack
EOF

执行脚本文件
bash /root/ipvs.sh

# 查看配置是否生效
lsmod | grep ip_vs
#ip_vs_sh               12288  0
#ip_vs_wrr              12288  0
#ip_vs_rr               12288  0
#ip_vs                 225280  6 ip_vs_rr,ip_vs_sh,ip_vs_wrr
#nf_conntrack          200704  1 ip_vs
#nf_defrag_ipv6         24576  2 nf_conntrack,ip_vs
#libcrc32c              12288  2 nf_conntrack,ip_vs
```

```shell
# 禁用 Swap
sudo swapoff -a
sudo sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab

# 永久关闭 Swap
sudo vim /etc/fstab
# 注释 /swap.img       none    swap    sw      0       0
systemctl reboot
free -h
```


#### 2 K8S 集群容器运行时 Containerd 准备
```shell
# 下载指定版本 containerd
wget https://github.com/containerd/containerd/releases/download/v1.7.16/cri-containerd-1.7.16-linux-amd64.tar.gz

# 解压安装
tar xf cri-containerd-1.7.16-linux-amd64.tar.gz -C /

# 验证是否安装成功
which containerd
# /usr/local/bin/containerd
# which runc
# /usr/local/sbin/runc
containerd --version
# containerd github.com/containerd/containerd v1.7.16 83031836b2cf55637d7abf847b17134c51b38e53
runc --version
# runc version 1.1.12
# commit: v1.1.12-0-g51d5e946
# spec: 1.0.2-dev
# go: go1.21.9
# libseccomp: 2.5.5

```

```shell
# 创建配置文件目录
mkdir /etc/containerd
# 生成配置文件
containerd config default > /etc/containerd/config.toml

修改第67行
vim /etc/containerd/config.toml
# sandbox_image = "registry.k8s.io/pause:3.9" 由3.8修改为3.9
# 如果使用阿里云容器镜像仓库，也可以修改为：
# sandbox_image = "registry.aliyuncs.com/google_containers/pause:3.9" 由3.8修改为3.9

修改第139行
vim /etc/containerd/config.toml
# SystemdCgroup = true 由false修改为true

```
```shell
# 设置开机自启动并现在启动
systemctl enable --now containerd

ls /var/run/containerd
# containerd.sock        io.containerd.runtime.v1.linux
# containerd.sock.ttrpc  io.containerd.runtime.v2.task
# 验证其版本
containerd --version
# containerd github.com/containerd/containerd v1.7.16 83031836b2cf55637d7abf847b17134c51b38e53

```

#### 2 安装 Docker 和 cri-dockerd
```shell
sudo apt install -y apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt update
sudo apt install -y docker-ce docker-ce-cli containerd.io
sudo systemctl enable docker
sudo systemctl start docker
docker --version

# 修改 cgroup 方式
sudo vim /etc/docker/daemon.json

# 添加
{
  "exec-opts": ["native.cgroupdriver=systemd"]
}

sudo systemctl restart docker
docker info | grep -i cgroup
# 看到 Cgroup Driver: systemd
```

```shell
# 安装
sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common
wget https://github.com/Mirantis/cri-dockerd/releases/download/v0.3.14/cri-dockerd_0.3.14.3-0.ubuntu-jammy_amd64.deb
sudo dpkg -i cri-dockerd_0.3.14.3-0.ubuntu-jammy_amd64.deb

# 修改配置 
# systemctl status cri-docker 查看 load 文件
vim /usr/lib/systemd/system/cri-docker.service
# 添加 --pod-infra-container-image=registry.k8s.io/pause:3.9
# --pod-infra-container-image=registry.aliyuncs.com/google_containers/pause:3.9

# 启动 
systemctl enable --now cri-docker

# 验证 看到 cri-dockerd.sock
ls /var/run/
```



#### 安装 Docker 和 cri-dockerd
![kubeadm kubelet kubectl](kubeadm%20kubelet%20kubectl.png)

```shell
# 安装 kubeadm, kubelet 和 kubectl
curl -fsSL https://pkgs.k8s.io/core:/stable/v1.30/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
curl -fsSL https://mirrors.aliyun.com/kubernetes-new/core/stable/v1.30/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
echo "deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.30/deb/ /" | sudo tee /etc/apt/sources.list.d/kubernetes.list
echo "deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://mirrors.aliyun.com/kubernetes-new/core/stable/v1.30/deb/ /" | sudo tee /etc/apt/sources.list.d/kubernetes.list
sudo apt-get update
ls /etc/apt/sources.list.d/
# 显示 kubernetes.list  ubuntu.sources  ubuntu.sources.curtin.orig
```

```shell
# 查看软件列表
apt-cache policy kubeadm
apt-cache policy kubelet
apt-cache policy kubectl

# 查看软件列表及其依赖关系
apt-cache showpkg kubeadm
# 查看软件版本
apt-cache madison kubeadm
# 默认安装
apt-get install -y kubelet kubeadm kubectl

# 安装指定版本
apt-get install -y kubelet kubeadm kubectl
# 锁定版本，防止后期自动更新
apt-mark hold kubelet kubeadm kubectl

# 解锁版本，执行更新
# spt-mark unhold kubelet kubeadm kubectl
```

```shell  配置 kubelet
vim /etc/sysconfig/kubelet
# 添加 KUBELET_EXTRA_ARGS="--cgroup-driver=systemd"

# 设置 kubelet 为开机自启即可，由于没有生成配置文件，集群初始化后自动启动
systemctl enable kubelet
```


#### 使用部署配置文件初始化K8S集群
```shell
advertiseAddress: 192.168.1.100
name: master
criSocket: unix:///var/run/cri-dockerd.sock
podSubnet: 10.244.0.0/16
imageRepository: registry.aliyuncs.com/google_containers
kind: Kubeletconfiguration
apiVersion: kubelet.config.k8s.io/v1beta3
cgroupDriver:systemd




apiVersion: kubeadm.k8s.io/v1beta3
kind: InitConfiguration
bootstrapTokens:
- groups:
  - system:bootstrappers:kubeadm:default-node-token
  token: abcdef.0123456789abcdef
  ttl: 24h0m0s
  usages:
  - signing
  - authentication
localAPIEndpoint:
  advertiseAddress: 192.168.1.100
  bindPort: 6443
nodeRegistration:
  criSocket: unix:///var/run/cri-dockerd.sock
  imagePullPolicy: IfNotPresent
  name: k8s-master01
  taints: null
---
apiVersion: kubeadm.k8s.io/v1beta3
kind: ClusterConfiguration
apiServer:
  timeoutForControlPlane: 4m0s
certificatesDir: /etc/kubernetes/pki
clusterName: kubernetes
controllerManager: {}
dns: {}
etcd:
  local:
    dataDir: /var/lib/etcd
imageRepository: registry.aliyuncs.com/google_containers
kubernetesVersion: 1.30.0
networking:
  dnsDomain: cluster.local
  serviceSubnet: 10.96.0.0/12
  podSubnet: 10.244.0.0/16
scheduler: {}
---
apiVersion: kubelet.config.k8s.io/v1beta1
kind: KubeletConfiguration
cgroupDriver: systemd

```

```shell
# 拉取镜像
sudo kubeadm config images pull --config kubeadm-config.yaml
crictl images
``
# 初始化
sudo kubeadm init --config=kubeadm-config.yaml

kubeadm config images pull --image-repository registry.aliyuncs.com/google_containers --cri-socket unix:///var/run/cri-dockerd.sock
kubeadm init --kubernetes-version=1.30.0 --pod-network-cidr=10.244.0.0/16 --apiserver-advertise-address=192.168.1.100 --cri-socket unix:///var/run/cri-dockerd.sock


# 配置
mkdir -p $HOME/.kube
cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
chown $(id -u):$(id -g) $HOME/.kube/config

# 验证
kubectl get nodes


# 可能遇到的错误解决
sudo rm /etc/kubernetes/manifests/kube-apiserver.yaml
sudo rm /etc/kubernetes/manifests/kube-controller-manager.yaml
sudo rm /etc/kubernetes/manifests/kube-scheduler.yaml
sudo rm /etc/kubernetes/manifests/etcd.yaml
sudo systemctl stop kubelet
```