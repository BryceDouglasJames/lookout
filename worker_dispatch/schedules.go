package worker_dispatch

import (
	"context"
	"fmt"
	"log"
	"sync"
)

var Driver_Pool chan chan Job_Type

func Schedule_Grailed_User_Search(ctx context.Context) {
	log.Println("Started a search on user search")
	wg := sync.WaitGroup{}

	worker_map := Web_Drivers_Init(1, Driver_Pool, &wg)
	temp := GetImage{name: "help", wait_group: &wg}

	for key, item := range worker_map {
		item.Start(key)
		temp.traverse_elm_tree("Raf")
		fmt.Printf("Key: %d :: Value: %+v\n", key, item)
	}

	wg.Wait()
	log.Println("Finsihed with items")

}
