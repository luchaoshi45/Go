### minio

#### shell
```shell
# 拉取 minio 镜像
docker pull minio/minio
# 查看 minio 镜像
docker images
```

```shell
# 运行 Docker 容器

# 最近更新的命令
docker run -d \
  -p 9000:9000 \
  -p 9001:9001 \
  --name minio1 \
  -v /home/minio/data:/data \
  -e "MINIO_ROOT_USER=AKIAIOSFODNN7EXAMPLE" \
  -e "MINIO_ROOT_PASSWORD=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY" \
  minio/minio server /data --console-address ":9001"

# 查看运行的 Docker 容器
docker ps

# 停止删除 Docker 容器
docker stop ec19d5a05a47
docker rm ec19d5a05a47
```

```shell
# 运行 Docker 容器返回的实例ID 97217674751d221031d0a95a9afd4a1c0439aaaba5f1b423e44263a2a46ec131
# 检查镜像启动日志
docker logs 97217674751d221031d0a95a9afd4a1c0439aaaba5f1b423e44263a2a46ec131

```

#### minio 界面
###### 用户名：AKIAIOSFODNN7EXAMPLE
###### 密码：wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
http://10.181.105.230:9000 <br>
http://10.181.105.230:9001 <br>
http://10.181.105.230:38259 <br>

API: http://172.17.0.7:9000  http://127.0.0.1:9000 <br>
WebUI: http://172.17.0.7:38259 http://127.0.0.1:38259 <br>
Docs: https://min.io/docs/minio/linux/index.html <br>
