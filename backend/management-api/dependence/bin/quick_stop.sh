#!/usr/bin/env bash

set -e

p1=$1

if [[ "$p1" == "-h" ||  "$p1" == "--help" ]] ; then
  echo "normal stop: ./quick_stop.sh"
  echo "clean  data: ./quick_stop.sh clean"
  exit 0
fi

for file in `ls $RELEASE_PATH`
    do
        if [ -d $file ]; then
            echo "STOP ==> " $RELEASE_PATH/$file
            cd $file/bin && ./stop.sh && cd - > /dev/null
        fi
    done

for pid in `ps -ef | grep "./cmlogagentd" | grep -v grep |  awk  '{print $2}'`
do
if [ ! -z ${pid} ];then
    kill -9 $pid
    echo "kill -9 $pid"
else
    echo "already stopped"
fi
done

if [[ "$p1" == "clean" ]] ; then
  RELEASE_PATH=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)
  cd $RELEASE_PATH
  for file in `ls $RELEASE_PATH`
      do
          if [ -d $file ]; then
              echo "CLEAN ==> " $RELEASE_PATH/$file/data
              cd $file && rm -rf data && cd - > /dev/null
          fi
      done
fi

