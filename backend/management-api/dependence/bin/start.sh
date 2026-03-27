#!/usr/bin/env bash

set -e

info(){
    echo -e "\033[32m$1 \033[0m"
}
info2(){
    echo -e "\033[36m$1 \033[0m"
}
error(){
    echo -e "\033[31m$1 \033[0m"
}

RELEASE_PATH=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

count=0
cd $RELEASE_PATH
for file in `ls $RELEASE_PATH`
    do
        if [ -f $RELEASE_PATH/$file/bin/restart.sh ]; then
            echo
            echo "START ==> " $RELEASE_PATH/$file
            cd $RELEASE_PATH/$file
            #mv log log_bak_$(date +"%Y%m%d%H%M%S") > /dev/null 2>&1
            cd bin && ./restart.sh
            count=$(($count+1))
        fi
    done

echo
info "starting $count node"
sleep 2
finishCount=$(ps -ef | grep "../../chainmaker" |grep -v grep|wc -l)
info "success $finishCount node"

function success() {
sleep 0.5
echo
info 'start success'
info2 '================================================================================='
info2 '   ______    __              _             __  ___            __'
info2 '  / ____/   / /_   ____ _   (_)   ____    /  |/  /  ____ _   / /__  ___    _____'
info2 ' / /       / __ \ / __ `/  / /   / __ \  / /|_/ /  / __ `/  / //_/ / _ \  / ___/'
info2 '/ /___    / / / // /_/ /  / /   / / / / / /  / /  / /_/ /  / ,<   /  __/ / /'
info2 '\____/   /_/ /_/ \__,_/  /_/   /_/ /_/ /_/  /_/   \__,_/  /_/|_|  \___/ /_/'
info2 '================================================================================='
info2 'ChainMaker Version: v3.0.0'
echo 'you can use the cmd show log detail: grep "ERROR\|put block\|all necessary" */log/system.log '
}
function err() {
sleep 0.5
echo
error "start fail."
error 'you can use the cmd show log detail: grep "ERROR\|put block\|all necessary" */log/system.log '
error '                                     cat */bin/panic.log'
}

if [[ $finishCount -ge $count ]]; then
  success
else
  err
fi
