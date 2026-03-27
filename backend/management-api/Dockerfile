FROM golang:1.16 as build

ENV GOPROXY=https://goproxy.cn,direct

# 移动到工作目录：/cm-mgmt
WORKDIR /cm-mgmt
# 将代码复制到容器中
COPY . .
# 编译成二进制可执行文件app
RUN cd src && go build -o ../bin/chainmaker-management.bin
# 移动到用于存放生成的二进制文件的 /build 目录
WORKDIR /build
RUN cp -r /cm-mgmt/bin .
RUN cp -r /cm-mgmt/scripts/* bin/
RUN cp -r /cm-mgmt/configs .
RUN cp -r /cm-mgmt/dependence .

FROM ubuntu:20.04

RUN apt update && apt install -y ca-certificates && update-ca-certificates

WORKDIR /chainmaker-management
# 将二进制文件从 /build 目录复制到这里
COPY --from=build /build/ .

# 声明服务端口
EXPOSE 9999

# 启动容器时运行的命令
WORKDIR /chainmaker-management/bin
ENTRYPOINT ["./chainmaker-management.bin", "-config" ]
CMD ["../configs/"]
