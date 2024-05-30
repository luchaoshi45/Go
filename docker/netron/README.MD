### netron

#### 修改 [Dockerfile](Dockerfile)

#### 构建和运行
```shell
# 构建 Docker 镜像 netron 是镜像名
docker build -t netron docker/netron
# 查看 Docker 镜像 
docker images
# 运行 Docker 容器
docker run -d -p 8080:8080 netron
# 查看运行的 Docker 容器
docker ps
```
```shell
# 运行 Docker 容器返回的实例ID 25d83155194420580129980bfd339097d49a0aa85ba2ef605603791cea4a1aa2
# 检查镜像启动日志
docker logs 25d83155194420580129980bfd339097d49a0aa85ba2ef605603791cea4a1aa2

```

#### Netron 界面
http://localhost:8080 <br>
http://10.181.105.230:8080 <br>