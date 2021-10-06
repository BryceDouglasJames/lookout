package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	browser "github.com/brycedouglasjames/lookout/web_interface"
	jobs "github.com/brycedouglasjames/lookout/worker_dispatch"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
	"github.com/go-rod/rod/lib/input"
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
	page := rod.New().MustConnect().MustPage("https://www.wikipedia.com/")
	page.MustWaitLoad().MustScreenshot("a.png")
	log.Println(page.String())
	j.wait_group.Done()
}

func (j *GetImage) traverse_elm_tree(s string) {
	j.wait_group.Add(1)

	type Item struct {
		name string
		link string
	}
	list := make(map[int]Item)

	page := rod.New().MustConnect().MustPage("https://google.com/")
	page.MustEmulate(devices.IPhone6or7or8Plus)
	search := page.MustWaitLoad().MustElement("input")
	search.MustInput("grailed")
	page.Keyboard.MustPress(input.Enter)
	page.MustWaitLoad().MustElementR("a", "grailed").MustClick()
	page.Timeout(2 * time.Second)
	searchbar := page.MustElementX("/html/body/div[3]/div/div[1]/header/div/div/div/div[1]/div")
	searchbar.MustElement("input").MustInput(s)
	searchbar.MustElementR("button", "Search").MustTap()
	page.Keyboard.MustPress(input.Escape)
	searchbar.MustElement("input").MustInput(s)
	searchbar.MustElementR("button", "Search").MustTap().MustClick()

	//for i := 0; i < 100; i++ {
	//	page.MustWaitLoad().Mouse.MustScroll(0, -100)
	//}

	itemView := page.MustWaitLoad().MustElementX("/html/body/div[3]/div[7]/div/div/div[3]/div[2]/div")
	items := itemView.MustElements(".feed-item")
	for key, item := range items {

		if item.MustHas(".listing-title") {
			itemLink := item.MustElement("a").MustProperty("href")
			itemName := item.MustElement(".listing-title")
			temp := &Item{
				name: itemName.MustText(),
				link: itemLink.Str(),
			}
			list[key] = *temp
			fmt.Printf("%+v\n", temp)
			itemView.Timeout(time.Millisecond * 30)
		}

	}

	page.Close()
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
		item.Start(key)
		temp.traverse_elm_tree("Raf")
		fmt.Printf("Key: %d :: Value: %+v\n", key, item)
	}

}

func Display_Error(err error) {
	fmt.Println(err)
	return
}
