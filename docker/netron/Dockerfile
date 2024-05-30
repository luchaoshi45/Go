# 使用 Ubuntu 20.04 作为基础镜像
FROM ubuntu:20.04

# 更新包管理器，并安装 Python 和 pip
RUN apt-get update && apt-get install -y python3 python3-pip

# 使用 pip 安装 Netron
RUN pip3 install netron

# 暴露 Netron 的默认端口
EXPOSE 8080

# 运行 Netron 并将其绑定到 0.0.0.0 以便外部访问
CMD ["netron", "--host", "0.0.0.0", "--port", "8080"]
