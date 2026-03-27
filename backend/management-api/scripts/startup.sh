#!/bin/bash

CONFIG_PATH="../configs/"
BROWSER_BIN="chainmaker-management.bin"
go build -o ${BROWSER_BIN} ../src
./shutdown.sh
nohup ./${BROWSER_BIN} -config ${CONFIG_PATH} >output 2>&1 &