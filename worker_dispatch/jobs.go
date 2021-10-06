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

type Job_Type interface {
	Do()
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

	itemView := page.MustWaitLoad().MustElementX("/html/body/div[3]/div[7]/div/div/div[3]/div[2]/div")

	for i := 0; i < 10; i++ {
		itemView.Page().Mouse.MustScroll(0, 10000)
	}

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
			fmt.Printf("ITEM #%d %+v\n", key, temp)
			itemView.Timeout(time.Millisecond * 30)

		}

	}

	name := "aaa.png"
	page.MustScreenshot(name)
	page.Close()
	j.wait_group.Done()

}
