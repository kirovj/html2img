package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	u := launcher.New().Leakless(false).MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("D:\\Go\\html2img\\exmaple\\a.html")
	page.MustWaitLoad().MustScreenshotFullPage("a.png")
}
