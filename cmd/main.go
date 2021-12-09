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

var (
	Driver_Pool *obj_pool.Driver_Pool
)

func main() {

	/*
	*	This is a quick example of how to launch a crawler process.
	*	Generate context with key value pair
	*	Add process to scheduler
	*	Create appropriate data channels
	*	To stop scheduler, send exit channel from os ^c
	 */

	scheduler := worker.CreateScheduler()
	//TODO make type system for scheduler
	//DO NOT USE BUILT IN TYPES FOR KEY
	c1 := context.WithValue(context.Background(), "quick", 1)
	processor_trigger, _ := scheduler.Add_Process(c1, worker.Schedule_Grailed_User_Search, time.Second*15, true)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Triggered grailed search")
		processor_trigger <- true
	})

	log_writer(time.Now(), "Made a grailed search... meta.")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

func Display_Error(err error) {
	fmt.Println(err)
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
