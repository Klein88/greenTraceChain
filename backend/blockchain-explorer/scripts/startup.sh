#!/bin/bash

EXPLORER_BIN="../bin/chainmaker-explorer.bin"
CONFIG_PATH="../configs/"

nohup ./${EXPLORER_BIN} -config ${CONFIG_PATH} >output 2>&1 &