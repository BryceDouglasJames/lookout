package worker_dispatch

import (
	"fmt"
	"log"
	"time"

	//"time"

	//"reflect"
	"runtime"
	"sync"

	//worker "github.com/brycedouglasjames/lookout/worker_dispatch"

	"github.com/go-rod/rod"
)

type Web_Driver_Worker struct {
	ID int
	// will need eventually ctx           context.Context
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
			fmt.Println(buf)
			log.Println("panicing!")
			w.Exit <- true
		}
	}()
	job.Do()
}

func (w *Web_Driver_Worker) Start(Job_Type interface{}) {
	w.Waiting.Wait()
	w.Waiting.Add(1)
	fmt.Println("PROGRESS................")
	time.Sleep(1 * time.Second)
	switch t := Job_Type.(type) {
	case GetImage:
		w.Job_Queue(t)
	default:
		fmt.Printf("ERROR %#v\n", t)
	}
	//w.Waiting.Done()

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
