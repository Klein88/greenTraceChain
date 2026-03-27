#
# Copyright (C) BABEC. All rights reserved.
# Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

for pid in `ps -ef | grep chainmaker | grep "\-c ../config/{org_id}/chainmaker.yml {node_addr}" | grep -v grep |  awk  '{print $2}'`
do
if [ ! -z ${pid} ];then
    kill $pid
fi
done

enable_dockervm={docker_enable}
docker_go_container_name=VM-GO-{node_addr}

enable_dockerJavaVm={docker_java_enable}
docker_java_container_name=VM-JAVA-{node_addr}

if [ ${enable_dockervm} == "true" ];then
  docker_container_lists=(`docker ps -a | grep ${docker_go_container_name} | awk '{print $1}'`)
  for container_id in ${docker_container_lists[*]}
  do
    docker stop ${container_id}
    docker rm ${container_id}
  done
fi

if [ ${enable_dockerJavaVm} == "true" ];then
  docker_java_container_lists=(`docker ps -a | grep ${docker_java_container_name} | awk '{print $1}'`)
  for container_id in ${docker_java_container_lists[*]}
  do
    docker stop ${container_id}
    docker rm ${container_id}
  done
fi

echo "chainmaker is stopped"
