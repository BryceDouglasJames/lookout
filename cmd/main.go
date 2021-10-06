package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	obj_pool "github.com/brycedouglasjames/lookout/pool"
	worker "github.com/brycedouglasjames/lookout/worker_dispatch"
)

var Driver_Pool *obj_pool.Driver_Pool

func main() {
	scheduler := worker.CreateScheduler()

	thisCtx := context.Background()
	c1 := context.WithValue(thisCtx, "quick", 1)
	processor_trigger, _ := scheduler.Add_Process(c1, worker.Schedule_Grailed_User_Search, time.Second*15, true)

	time.AfterFunc(10*time.Second, func() {
		fmt.Println("Triggered grailed search")
		processor_trigger <- true
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	scheduler.Stop()
}

func Display_Error(err error) {
	fmt.Println(err)
	return
}
