package scheduler

import (
	"github.com/seanbit/nano/internal/log"
	"sync/atomic"
)

const (
	UserSchema = "UserScheduler"
)

type UserScheduler struct {
	chDie   chan struct{}
	chExit  chan struct{}
	chTasks chan Task
	started int32
	closed  int32
}

func NewUserScheduler(cap int) *UserScheduler {
	if cap <= 0 {
		cap = messageQueueBacklog
	}
	return &UserScheduler{
		chDie:   make(chan struct{}),
		chExit:  make(chan struct{}),
		chTasks: make(chan Task, cap),
	}
}

func (us *UserScheduler) Schedule(task Task) {
	us.chTasks <- task
}

func (us *UserScheduler) Close() {
	if atomic.AddInt32(&us.closed, 1) != 1 {
		return
	}
	close(us.chDie)
	<-us.chExit
	log.Println("User Scheduler stopped")
}

func (us *UserScheduler) Sched() {
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
