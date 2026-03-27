#!/bin/bash

EXPLORER_BIN="../bin/chainmaker-explorer.bin"

pid=`ps -ef | grep ${EXPLORER_BIN} | grep -v grep | awk '{print $2}'`
if [ ! -z ${pid} ];then
    kill -9 $pid
    echo "kill -9 $pid"
else
  echo "nothing"
fi