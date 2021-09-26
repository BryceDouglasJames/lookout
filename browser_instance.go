package lookout

import (
	"github.com/go-rod/rod"
)

func Browser_test() string {
	page := rod.New().MustConnect().MustPage("https://www.wikipedia.org/")
	page.MustWaitLoad().MustScreenshot("a.png")
	return page.String()
}
