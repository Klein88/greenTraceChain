#!/bin/bash

BROWSER_BIN="chainmaker-management.bin"
pid=`ps -ef | grep ${BROWSER_BIN} | grep -v grep | awk '{print $2}'`
if [ ! -z ${pid} ];then
    echo "kill -9 $pid"
    kill -9 $pid
else
    echo "$BROWSER_BIN already stopped"
fi