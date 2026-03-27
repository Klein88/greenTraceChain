#!/usr/bin/env bash

set -e

RELEASE_PATH=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

cd $RELEASE_PATH
for file in `ls $RELEASE_PATH`
    do
        if [ -d $file ]; then
            echo "START ==> " $RELEASE_PATH/$file
            cd $file/bin && ./restart.sh && cd - > /dev/null
        fi
    done

sleep 1
nohup ./cmlogagentd -laddr=0.0.0.0:22301 -cache-size=10 -up-lines=10 -down-lines=10 -node-dirs={node_paths} > cmlogagentd.log 2>&1 &
echo "logagent is startting, pls check log..."
