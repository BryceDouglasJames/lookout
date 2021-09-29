package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	browser "github.com/brycedouglasjames/lookout/web_interface"
	jobs "github.com/brycedouglasjames/lookout/worker_dispatch"
	"github.com/go-rod/rod"
)

type testpool struct {
	singlejob   chan jobs.Job_Type
	queue       chan jobs.Job_Type
	readypool   chan chan jobs.Job_Type
	workers     map[int]browser.Web_Driver_Worker
	dis_stopped sync.WaitGroup
	work_stop   *sync.WaitGroup
	quit        chan bool
}

func main() {
	/*request, err := browser.PageRequest("https://amazon.com")
	defer Display_Error(err)

	document, _ := html.Parse(strings.NewReader(request))
	Node, err := browser.GetDesiredTag(document, "li")
	defer Display_Error(err)

	head := browser.Format(Node)
	fmt.Println(head)

	test := browser.Browser_test()
	fmt.Println(test)*/

	driver_worker_test()

}

type GetImage struct {
	name       string
	wait_group *sync.WaitGroup
}

func (j *GetImage) Do() {
	j.wait_group.Add(1)
	time.Sleep(1 * time.Second)
	page := rod.New().MustConnect().MustPage("https://www.wikipedia.org/")
	page.MustWaitLoad().MustScreenshot("a.png")
	log.Println(page.String())
	j.wait_group.Done()
}

func driver_worker_test() {
	Stop_Group := sync.WaitGroup{}
	Driver_Pool := make(chan chan jobs.Job_Type)
	//ctx := context.Background()
	//Drivers := make([]*jobs.Job_Type, 10, 10)

	worker_map := browser.Web_Drivers_Init(5, Driver_Pool, &Stop_Group)
	temp := GetImage{name: "help", wait_group: &Stop_Group}

	for key, item := range worker_map {
		//item.Start(key)
		temp.Do()
		time.Sleep(time.Millisecond)
		fmt.Printf("Key: %d :: Value: %+v\n", key, item)
	}

}

func Display_Error(err error) {
	fmt.Println(err)
	return
}
