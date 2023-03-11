package scheduler

import (
	"github.com/seanbit/nano/internal/log"
	"sync/atomic"
)

type LocalScheduler struct {
	chDie   chan struct{}
	chExit  chan struct{}
	chTasks chan Task
	started int32
	closed  int32
}

func NewLocalScheduler(cap int) ILocalScheduler {
	if cap <= 0 {
		cap = messageQueueBacklog
	}
	scheduler := &LocalScheduler{
		chDie:   make(chan struct{}),
		chExit:  make(chan struct{}),
		chTasks: make(chan Task, cap),
	}
	go scheduler.Scheduling()
	return scheduler
}

func (us *LocalScheduler) Schedule(task Task) {
	us.chTasks <- task
}

func (us *LocalScheduler) Close() {
	if atomic.AddInt32(&us.closed, 1) != 1 {
		return
	}
	close(us.chDie)
	<-us.chExit
	log.Println("User Scheduler stopped")
}

func (us *LocalScheduler) Scheduling() {
	if atomic.AddInt32(&us.started, 1) != 1 {
		return
	}

	defer func() {
		close(us.chExit)
	}()

	for {
		select {
		case f := <-us.chTasks:
			try(f)

		case <-us.chDie:
			return
		}
	}
}
