package gomon

import (
	"log"
	"runtime"
	//	"runtime/pprof"
	"time"
)

const LOG_PREFIX = "GOMON:"

type Monitor struct {
	Id         int64
	watchingGR bool
	memStats   runtime.MemStats
	ticker     time.Ticker
}

// New creates a new instance of gomon monitor
func New() *Monitor {
	monitor := &Monitor{}
	monitor.Id = time.Now().Unix()
	monitor.ticker = *time.NewTicker(time.Second * 5)
	log.Printf("Initializing new gomon monitor. Monitor id: %d", monitor.Id)

	return monitor
}

func (m Monitor) Start() {
	log.Println("Starting gomon")
	go func() {
		for _ = range m.ticker.C {
			m.logGR()
		}
	}()
}

func (m *Monitor) CaptureGoRoutineStats() *Monitor {
	m.watchingGR = true
	return m
}

func (m *Monitor) logGR() {
	if m.watchingGR == true {
		log.Printf("%s Num go routines running: %d", LOG_PREFIX, runtime.NumGoroutine())
	}
}
