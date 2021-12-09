package worker_dispatch

import (
	"context"
	"fmt"
	"log"
	"sync"
)

var Driver_Pool chan chan Job_Type

func Schedule_Grailed_User_Search(ctx context.Context) {
	log.Println("Started a Grail search")
	wg := sync.WaitGroup{}

	//TODO associate and add context  to each search with some id to grab from pool
	worker_map := Web_Drivers_Init(1, Driver_Pool, &wg)
	temp := Grailed_Items{wg: &wg}

	for key, item := range worker_map {
		fmt.Println("Searching...")
		//item.Start(key)
		temp.traverse_elm_tree("Raf")
		fmt.Printf("Key: %d :: Value: %+v\n", key, item)
	}

	wg.Wait()
	log.Println("Finished with items")

}
