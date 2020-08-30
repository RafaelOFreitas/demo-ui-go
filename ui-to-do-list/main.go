package main

import (
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	// create rect for window
	windowRect := sciter.NewRect(0, 0, 800, 400)

	// create sciter window object
	win, _ := window.New(sciter.SW_TITLEBAR|sciter.SW_MAIN, windowRect)

	// set title for window
	win.SetTitle("Test window")

	// load index.html file in window
	win.LoadFile("./ui-to-do-list/index.html")

	// lauch sciter window
	win.Show()

	// run sciter application
	win.Run()
}
