FROM golang:1.16 as build

ENV GOPROXY=https://goproxy.cn,direct
ENV GOPRIVATE=chainmaker.org
# 移动到工作目录：/cm-explorer
WORKDIR /cm-explorer
# 将代码复制到容器中
COPY . .
# 编译成二进制可执行文件app
RUN rm -rf go.sum && make
# 移动到用于存放生成的二进制文件的 /build 目录
WORKDIR /build
RUN cp -r /cm-explorer/bin .
RUN cp -r /cm-explorer/configs .
RUN cp -r /cm-explorer/scripts/* bin/

FROM ubuntu:20.04
RUN rm /bin/sh && ln -s /bin/bash /bin/sh
RUN apt-get update && apt-get install -y vim net-tools tree
ENV TZ "Asia/Shanghai"
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y tzdata && \
    echo $TZ > /etc/timezone && \
    ln -fs /usr/share/zoneinfo/$TZ /etc/localtime && \
    dpkg-reconfigure tzdata -f noninteractive

WORKDIR /chainmaker-explorer-backend
# 将二进制文件从 /build 目录复制到这里
COPY --from=build /build/ .

# 声明服务端口
EXPOSE 9997

# 启动容器时运行的命令
WORKDIR /chainmaker-explorer-backend/bin
ENTRYPOINT ["./chainmaker-explorer.bin", "-config" ]
CMD ["../configs/"]
