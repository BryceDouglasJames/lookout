package worker_dispatch

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

var Driver_Pool chan chan Job_Type

func Schedule_Ping(ctx context.Context, args interface{}) {
	wg := sync.WaitGroup{}

	switch args.(type) {
	case string:
		fmt.Printf("%v is an interface string\n", args)
	default:
		fmt.Println("THIS IS EMPTY")
	}

	t := fmt.Sprintf("%v", args)
	fmt.Printf("\n\n\nPinging website %v,,,\n", t)

	temp := Pinger{
		ctx:  ctx,
		wg:   &wg,
		link: t,
		//timout: 10*time.Minute,
	}
	temp.Do()
	wg.Wait()

	//log.Println("finished grabbing screenshot.")
}

func Schedule_Grailed_User_Search(ctx context.Context, args interface{}) {
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

func Generate_Root_Search(ctx context.Context, args interface{}) {
	log.Println("Started General Search ...")

	wg := sync.WaitGroup{}

	//TODO associate and add context  to each search with some id to grab from pool
	worker_map := Web_Drivers_Init(1, Driver_Pool, &wg)

	/*type favContextKey string
		f := func(ctx context.Context, k favContextKey) {
			if v := ctx.Value(k); v != nil {
				fmt.Println("found value:", v)
				return
			}
			fmt.Println("key not found:", k)
	}
		k := favContextKey("TEST")
		f(ctx, k)
	*/

	switch t := ctx.Value("TEST").(type) {
	case string:
		temp := GetImage{
			id:  "GETIMAGE",
			wg:  &wg,
			URL: string(t),
		}
		for key, item := range worker_map {
			fmt.Println("Searching...")
			time.Sleep(2 * time.Second)
			item.Start(temp)
			fmt.Printf("Key: %d :: Value: %+v\n", key, item)
		}
	default:
		fmt.Println("What is this?")
	}
	log.Println("Finished with regular search")
}
