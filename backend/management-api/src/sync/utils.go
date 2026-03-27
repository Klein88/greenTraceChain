/*
Package sync comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package sync

// Task task
type Task struct {
	f func() error
}

// NewTask new Task
func NewTask(f func() error) *Task {
	return &Task{
		f: f,
	}
}
func (t *Task) execute() {
	err := t.f()
	if err != nil {
		log.Error("Task execute err :", err.Error())
	}
}

// Pool worker pool
type Pool struct {
	workerNum  int
	EntryChan  chan *Task
	workerChan chan *Task
}

// NewPool new Pool
func NewPool(num int) *Pool {
	return &Pool{
		workerNum:  num,
		EntryChan:  make(chan *Task),
		workerChan: make(chan *Task),
	}
}

func (p *Pool) worker() {
	for task := range p.workerChan {
		task.execute()
	}
}

// Run worker run
func (p *Pool) Run() {
	for i := 0; i < p.workerNum; i++ {
		go p.worker()
	}
	for task := range p.EntryChan {
		p.workerChan <- task
	}
}

// SingletonSync singleton sync
type SingletonSync struct {
	SyncStart bool
}
