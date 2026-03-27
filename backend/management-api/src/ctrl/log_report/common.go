/*
Package log_report comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package log_report

import (
	loggers "management_backend/src/logger"
	"management_backend/src/sync"
	"time"
)

var (
	log = loggers.GetLogger(loggers.ModuleWeb)
)

const (
	// NO_AUTO no auto
	NO_AUTO = iota
	// AUTO auto
	AUTO
)

// TickerMap ticker map
var TickerMap = map[string]*Ticker{}

// Ticker ticker
type Ticker struct {
	stopCh     chan struct{}
	tickerTime int
}

// NewTicker newTicker
func NewTicker(tickerTime int) *Ticker {
	return &Ticker{
		stopCh:     make(chan struct{}),
		tickerTime: tickerTime,
	}
}

// Start ticker start
func (tickerUp *Ticker) Start(chainId string) {
	TickerMap[chainId] = tickerUp
	go func() {
		ticker := time.NewTicker(time.Hour * time.Duration(tickerUp.tickerTime))
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				//定时上报链信息
				err := sync.ReportChainData(chainId)
				log.Error(err)

			case <-tickerUp.stopCh:
				//停止定时任务，并杀死进程
				return
			}
		}
	}()
}

// StopTicker stopTicker
func (tickerUp *Ticker) StopTicker(chainId string) {
	delete(TickerMap, chainId)
	close(tickerUp.stopCh)
}
