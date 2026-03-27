VM_GO_IMAGE_NAME="chainmakerofficial/chainmaker-vm-engine:v3.0.1"

function parse_yaml {
   local prefix=$2
   local s='[[:space:]]*' w='[a-zA-Z0-9_]*' fs=$(echo @|tr @ '\034')
   sed -ne "s|^\($s\):|\1|" \
        -e "s|^\($s\)\($w\)$s:$s[\"']\(.*\)[\"']$s\$|\1$fs\2$fs\3|p" \
        -e "s|^\($s\)\($w\)$s:$s\(.*\)$s\$|\1$fs\2$fs\3|p"  $1 |
   awk -F$fs '{
      indent = length($1)/2;
      vname[indent] = $2;
      for (i in vname) {if (i > indent) {delete vname[i]}}
      if (length($3) > 0) {
         vn=""; for (i=0; i<indent; i++) {vn=(vn)(vname[i])("_")}
         printf("%s%s%s=\"%s\"\n", "'$prefix'",vn, $2, $3);
      }
   }'
}

config_file="../config/{org_id}/chainmaker.yml"
# config_file="../../config/wx-org1-solo/chainmaker.yml"
eval $(parse_yaml "$config_file" "chainmaker_")

mount_path=$chainmaker_vm_go_data_mount_path
log_path=$chainmaker_vm_go_log_mount_path
if [[ "${mount_path:0:1}" != "/" ]];then
  mount_path=$(pwd)/$mount_path
fi
if [[ "${log_path:0:1}" != "/" ]];then
  log_path=$(pwd)/$log_path
fi

mkdir -p "$mount_path"
mkdir -p "$log_path"

CON=`docker image ls $VM_GO_IMAGE_NAME | wc -l`  #‘redis:latest'根据镜像和版本自己修改
if [ $CON -eq 1 ]  #CON取值为2表示镜像存在，为1镜像不存在
then
docker pull $VM_GO_IMAGE_NAME  #镜像存在时执行此行命令
fi

docker run -itd \
  --net=host \
  -v "$mount_path":/mount \
  -v "$log_path":/log \
  -e CHAIN_RPC_PROTOCOL="1" \
  -e CHAIN_RPC_PORT="$chainmaker_vm_go_contract_engine_port" \
  -e SANDBOX_RPC_PORT="$chainmaker_vm_common_runtime_server_port" \
  -e MAX_SEND_MSG_SIZE="$chainmaker_vm_common_max_send_msg_size" \
  -e MAX_RECV_MSG_SIZE="$chainmaker_vm_common_max_recv_msg_size" \
  -e MAX_CONN_TIMEOUT="$chainmaker_vm_common_dial_timeout" \
  -e MAX_ORIGINAL_PROCESS_NUM="$chainmaker_vm_go_max_concurrency" \
  -e DOCKERVM_CONTRACT_ENGINE_LOG_LEVEL="$chainmaker_vm_go_log_level" \
  -e DOCKERVM_SANDBOX_LOG_LEVEL="$chainmaker_vm_go_log_level" \
  -e DOCKERVM_LOG_IN_CONSOLE="$chainmaker_vm_go_log_in_console" \
  --name VM-GO-{node_addr} \
  --privileged $VM_GO_IMAGE_NAME \
  > /dev/null

echo "start docker vm service container succeed:  VM-GO-{node_addr}"
