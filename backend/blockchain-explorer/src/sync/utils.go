/*
Package sync comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package sync

import (
	"math/rand"
	"sync"
)

var (
	globalChainMaxHeightMp sync.Map //chainid - blockheight(string-> int64)
)

// getMaxHeight
// @desc 获取链存储区块的高度
// @param chainId 链id
// @return int64
func getMaxHeight(chainId string) int64 {
	height, _ := globalChainMaxHeightMp.Load(chainId)
	return height.(int64)
}

// getMaxHeight
// @desc 设置链存储区块的高度
// @param chainId 链id
// @param int64 区块高度
func setMaxHeight(chainId string, height int64) {
	globalChainMaxHeightMp.Store(chainId, height)
}

// nolint
func randomSleepTime() int {
	return rand.Intn(20) + 200

}

// Task task
type Task struct {
	f func() error
}

// NewTask task
func NewTask(f func() error) *Task {
	return &Task{
		f: f,
	}
}

// 执行进程
func (t *Task) execute() {
	err := t.f()
	if err != nil {
		log.Errorf("task exec fail:%v", err)
	}
}

// Pool  a pool
type Pool struct {
	workerNum  int
	EntryChan  chan *Task
	workerChan chan *Task
}

// NewPool new pool
func NewPool(num int) *Pool {
	return &Pool{
		workerNum:  num,
		EntryChan:  make(chan *Task),
		workerChan: make(chan *Task),
	}
}

// task exe
func (p *Pool) worker() {
	for task := range p.workerChan {
		task.execute()
	}
}

// Run a work
func (p *Pool) Run() {
	for i := 0; i < p.workerNum; i++ {
		go p.worker()
	}
	for task := range p.EntryChan {
		p.workerChan <- task
	}
}

// SingletonSync singleton
type SingletonSync struct {
	SyncStart bool
}

var singlesync *SingletonSync
var once sync.Once

// GetSync get sync
func GetSync() *SingletonSync {
	once.Do(func() {
		singlesync = &SingletonSync{
			SyncStart: false,
		}
	})
	return singlesync
}

// GetSyncStart get
// @desc
// @param ${param}
// @return bool
func (singletonSync *SingletonSync) GetSyncStart() bool {
	return singlesync.SyncStart
}

// SetSyncStart s
// @desc
// @param ${param}
func (singletonSync *SingletonSync) SetSyncStart(syncStart bool) {
	singlesync.SyncStart = syncStart
}
