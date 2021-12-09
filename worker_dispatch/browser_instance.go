package worker_dispatch

import (
	"context"
	"log"
	"runtime"
	"sync"

	//worker "github.com/brycedouglasjames/lookout/worker_dispatch"

	"github.com/go-rod/rod"
)

type Web_Driver_Worker struct {
	ID            int
	ctx           context.Context
	Waiting       *sync.WaitGroup
	Master_signal chan chan Job_Type
	Job           chan Job_Type
	Exit          chan bool
}

//This function iniztializes n amount of drivers. This object is sent
//to the object pool and designated to each client.
func Web_Drivers_Init(size int, master_queue chan chan Job_Type, done *sync.WaitGroup) map[int]Web_Driver_Worker {
	drivers := make(map[int]Web_Driver_Worker)
	for i := 0; i < size; i++ {
		conn := &Web_Driver_Worker{
			ID:            i,
			Waiting:       done,
			Master_signal: master_queue,
			Job:           make(chan Job_Type),
			Exit:          make(chan bool),
		}
		drivers[i] = *conn
	}

	return drivers

}

func (w *Web_Driver_Worker) Job_Queue(job Job_Type) {
	go func() {
		if fail := recover(); fail != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Println("panicing!")
			w.Exit <- true
		}

	}()
	job.Do()
}

func (w *Web_Driver_Worker) Start(id int) {
	go func() {
		w.Waiting.Wait()
		w.Waiting.Add(1)
		for {
			w.Master_signal <- w.Job
			select {
			case job := <-w.Job:
				w.Job_Queue(job)
			case <-w.Exit:
				//w.Stop()
				w.Waiting.Done()
				return
			}
		}
	}()
}

func (w *Web_Driver_Worker) Stop(id int) {
}

func Browser_test() string {
	page := rod.New().MustConnect().MustPage("https://www.wikipedia.org/")
	page.MustWaitLoad().MustScreenshot("a.png")
	return page.String()
}

func (w *Web_Driver_Worker) getID() int {
	return w.ID
}
