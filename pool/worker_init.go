package pool

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	//"os/signal"
	"sync"
	"time"

	//obj_pool "github.com/brycedouglasjames/lookout/pool"
	worker "github.com/brycedouglasjames/lookout/worker_dispatch"
)

type key string

var (
	Pool    *Driver_Pool
	DEFAULT key = "TEST"
)

/*func init() {


	*	This is a quick example of how to launch a crawler process.
	*	Generate context with key value pair
	*	Add process to scheduler
	*	Create appropriate data channels
	*	To stop scheduler, send exit channel from os ^c


	Launch_Prompt()
}*/

func Display_Error(err error) {
	fmt.Println(err)
}

func Launch_Prompt() {
	fmt.Printf("\n\nHello %s! Welcome to my crawler thing \n*** type -help for commands ***\n\n", "User")
	Start(os.Stdin, os.Stdout, "User")

}

func Ingest_Url(url string) string {
	scheduler := worker.CreateScheduler()
	//TODO make type system for scheduler

	//***************DO NOT USE BUILT IN TYPES FOR KEY TO AVOID COLLISIONS**************
	c1 := context.WithValue(context.Background(), DEFAULT, url)

	//initizalize driver pool with signal queue
	queue_wg := new(sync.WaitGroup)
	Job_Queue := make(chan chan worker.Job_Type)
	worker.Web_Drivers_Init(1, Job_Queue, queue_wg)
	var arg interface{} = url
	//append(scheduler.OngoingJobs, worker[0].Job)

	//processor_trigger <- true
	processor_trigger, _ := scheduler.Add_Process(c1, worker.Schedule_Ping, time.Minute*10, true, arg)
	processor_trigger <- true

	/*time.AfterFunc(5*time.Second, func() {
		processor_trigger <- true
	})
	time.AfterFunc(5*time.Minute, func() {
		Kill_s <- true
	})*/
	//scheduler.StopAll()
	//log_writer(time.Now(), "Triggered search on "+url+" finsihed in ...")

	//Kill_Processor(scheduler) d
	//Launch_Prompt()
	return "DONE"
}

func Start(in io.Reader, out io.Writer, name string) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, name+" "+">>")
		scanned := scanner.Scan()
		if !scanned || scanner.Text() == "" {
			return
		}
		line := scanner.Text()
		if line == "-help" {
			fmt.Print("HOW TO USE: \n\tstring URL: Get Screenshot of page, save as a.png.\n\t-q: quit\n\n")
		} else if line == "-q" {
			os.Exit(3)
		} else {
			var key interface{} = line
			var val string = key.(string)
			answer := Ingest_Url(val)
			if answer != "DONE" {
				break
			}
		}
	}
}

/*
func Kill_Processor(scheduler *worker.Scheduler) {
	c1 := context.WithValue(context.Background(), string("KILL"), "kill")
	processor_trigger, _ := scheduler.Add_Process(c1, nil, time.Second*15, true)
	processor_trigger <- true
}

func log_writer(time time.Time, message string) {
	file, err := os.OpenFile("activity.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		os.Create("activity.log")
		fmt.Println("ACTIVITY.LOG CREATED, PLEASE RUN AGAIN")
	}
	defer file.Close()
	s := fmt.Sprintf("[%v]:\t%s\n", time, message)
	file.WriteString(s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %v into file.", s)
	file.Close()
}
*/
