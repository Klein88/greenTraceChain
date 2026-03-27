CURRENT_PATH=$(pwd)
PROJECT_PATH=$(dirname "${CURRENT_PATH}")

docker run  \
 -itd \
 --restart=always \
 -p 9997:9997 \
 -v $PROJECT_PATH/configs:/chainmaker-explorer-backend/configs \
 -v $PROJECT_PATH/bin:/chainmaker-explorer-backend/bin \
 --privileged=true \
 --name=explorer-backend \
 --shm-size=4G \
 --log-opt max-size=1000m \
 --log-opt max-file=3 \
 golang:1.16 \
 bash -c "cd /chainmaker-explorer-backend/bin && ./chainmaker-explorer.bin -config ../configs/"