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
	Cancellations []context.CancelFunc
	Flags         []chan bool
}

func CreateScheduler() *Scheduler {
	//eventually, we will want to pull a schedule record with everything in it. But for now, let's just keep it simple.
	return &Scheduler{
		wg:            new(sync.WaitGroup),
		OngoingJobs:   make([]*job, 0),
		Cancellations: make([]context.CancelFunc, 0),
		Flags:         make([]chan bool, 0),
	}

}

/*
*	When adding a process, you can return n channels
*	to control it's flow through the processor
 */
func (s *Scheduler) Add_Process(ctx context.Context, j job, runtime time.Duration, isActive bool) (chan bool, chan bool) {
	//TODO: Create doc lmao
	ctx, cancel := context.WithCancel(ctx)
	s.Cancellations = append(s.Cancellations, cancel)
	schedule_trigger := make(chan bool)
	active_trigger := make(chan bool)
	s.Flags = append(s.Flags, schedule_trigger)
	s.wg.Add(1)

	//run routine through processor.
	go s.Processor(ctx, j, runtime, s.OngoingJobs, s.Cancellations, schedule_trigger, active_trigger)
	return schedule_trigger, active_trigger
}

func (s *Scheduler) Processor(ctx context.Context, j job, runtime time.Duration, active_queue []*job, Destroy_queue []context.CancelFunc,
	schedule_t chan bool, active_t chan bool) {

	clock := time.NewTicker(runtime)
	run := make(chan bool, 1)
	run <- true

	//Channel routine flags
	for {
		select {
		case <-run:
			s.Run_Process(ctx, j, true)

		case <-ctx.Done():
			s.wg.Done()
			clock.Stop()
			run <- false
			return

		default:
			time.Sleep(time.Second)
		}

	}
}

func (s *Scheduler) Run_Process(ctx context.Context, j job, active bool) {
	//trigger that the search is active
	if active {
		j(ctx)
	}

	//pass context to driver

	//
}

func (s *Scheduler) StopAll() {
	for _, cancel := range s.Cancellations {
		cancel()
	}
	s.wg.Wait()
}

//SAVE FOR LATER -- SENDS KILL PROCESS SIGNAL TO OS
/*fmt.Println("I AM STOPPING")
kill := exec.Command("kill -9 $(pgrep -f [PROCESS])")
stdout, err := kill.Output()
if err != nil {
	fmt.Println("Error killing crawler")
	return
}
fmt.Printf("%v\n", stdout)*/
