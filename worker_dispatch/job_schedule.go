package worker_dispatch

import (
	"context"
	"sync"
	"time"
)

type job func(ctx context.Context)
type Scheduler struct {
	wg            *sync.WaitGroup
	OngoingJobs   []*job
	cancellations []context.CancelFunc
	Flags         []chan bool
}

func CreateScheduler() *Scheduler {
	//eventually, we will want to pull a schedule record with everything in it. But for now, let's just keep it simple.
	return &Scheduler{
		wg:            new(sync.WaitGroup),
		OngoingJobs:   make([]*job, 0),
		cancellations: make([]context.CancelFunc, 0),
		Flags:         make([]chan bool, 0),
	}

}

func (s *Scheduler) Add_Process(ctx context.Context, j job, runtime time.Duration, isActive bool) (chan bool, chan bool) {
	ctx, cancel := context.WithCancel(ctx)
	s.cancellations = append(s.cancellations, cancel)

	schedule_trigger := make(chan bool)
	active_trigger := make(chan bool)

	s.Flags = append(s.Flags, schedule_trigger)

	s.wg.Add(1)
	go s.Processor(ctx, j, runtime, isActive, schedule_trigger, active_trigger)
	return schedule_trigger, active_trigger
}

func (s *Scheduler) Run_Process(ctx context.Context, j job, isActive bool) {

	if isActive {
		j(ctx)
	}
}

func (s *Scheduler) Stop() {
	for _, channel := range s.Flags {
		channel <- true
	}
}

func (s *Scheduler) Processor(ctx context.Context, j job, runtime time.Duration, isActive bool, s_t chan bool, a_t chan bool) {
	clock := time.NewTicker(runtime)
	run := make(chan bool, 1)
	run <- true
	for {
		select {
		//...
		case <-run:
			s.Run_Process(ctx, j, isActive)

		case <-ctx.Done():
			s.wg.Done()
			clock.Stop()
			run <- false
			return
		}
	}
}
