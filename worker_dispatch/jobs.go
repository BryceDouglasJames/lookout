package worker_dispatch

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
	"github.com/go-rod/rod/lib/input"
)

//functions that can be called from the outside...duh
type Job_Type interface {
	Do()
}

/******TAKE SCRRENSHOT OF PAGE******/
type GetImage struct {
	wg *sync.WaitGroup
}

func (job *GetImage) Do() {
	job.wg.Add(1)
	time.Sleep(1 * time.Second)
	page := rod.New().MustConnect().MustPage("https://www.wikipedia.com/")
	page.MustWaitLoad().MustScreenshot("a.png")
	log.Println(page.String())
	job.wg.Done()
}

/*************************************/

/*******SEARCH GRAILED FOR ITEM BASED SEARCH*******/
type Grailed_Items struct {
	wg *sync.WaitGroup
}

func (job *Grailed_Items) traverse_elm_tree(s string) {
	job.wg.Add(1)

	type Item struct {
		name string
		link string
	}
	list := make(map[int]Item)
	page := rod.New().MustConnect().MustPage("https://google.com/")
	page.MustEmulate(devices.IPhone6or7or8Plus)
	search := page.MustWaitLoad().MustElementX("/html/body/div[1]/div[3]/form/div[1]/div[1]/div[1]/div/div[2]/input")
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
	itemView := page.MustWaitLoad().MustElementX("/html/body/div[3]/div[7]/div/div/div[3]/div[2]/div")
	for i := 0; i < 3; i++ {
		itemView.Page().Mouse.MustScroll(0, 10000)
	}

	//HERE....WE CAN PASS ALL THE LINKS FROM THE JOB TO THE WEB INTERFACE.
	//WE DON'T NEED TO SCRAPE HERE, WE CAN JUST SCROLL, STOP, SEND DOC OVER AND CONTINUE SCROLLING
	//TODO

	items := itemView.MustElements(".feed-item")
	fmt.Println(items)
	for key, item := range items {
		if item.MustHas(".listing-title") {
			itemLink := item.MustElement("a").MustProperty("href")
			itemName := item.MustElement(".listing-title")
			temp := &Item{
				name: itemName.MustText(),
				link: itemLink.Str(),
			}
			list[key] = *temp
			fmt.Printf("ITEM #%d %+v\n", key, temp)
			itemView.Timeout(time.Millisecond * 30)
		}
	}
	page.Close()
	job.wg.Done()
}

/********************************************/
