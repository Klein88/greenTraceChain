# the second stage
FROM ubuntu:20.04
RUN rm /bin/sh && ln -s /bin/bash /bin/sh
RUN apt-get update && apt-get install -y vim net-tools tree
ENV TZ "Asia/Shanghai"
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y tzdata && \
    echo $TZ > /etc/timezone && \
    ln -fs /usr/share/zoneinfo/$TZ /etc/localtime && \
    dpkg-reconfigure tzdata -f noninteractive

COPY ./configs /chainmaker-explorer-backend/configs
COPY ./bin /chainmaker-explorer-backend/bin
COPY ./scripts/* /chainmaker-explorer-backend/bin/

# 声明服务端口
EXPOSE 9997

# 启动容器时运行的命令
WORKDIR /chainmaker-explorer-backend/bin
ENTRYPOINT ["./chainmaker-explorer.bin", "-config" ]
CMD ["../configs/"]
