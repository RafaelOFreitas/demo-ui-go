package main

import (
	"time"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	// create rect for window
	windowRect := sciter.NewRect(0, 0, 800, 800)

	// create sciter window object
	win, _ := window.New(sciter.SW_TITLEBAR|sciter.SW_MAIN, windowRect)

	// set title for window
	win.SetTitle("Test window")

	// load index.html file in window
	win.LoadFile("./ui-splash-screen/index.html")

	showSplashScreen(win)

	// run sciter application
	win.Run()
}

func showSplashScreen(mainScreen *window.Window) {
	rect := sciter.NewRect(0, 0, 500, 500)

	splashScreen, _ := window.New(sciter.SW_TOOL, rect)

	splashScreen.LoadFile("./ui-splash-screen/splash.html")

	splashScreen.DefineFunction("load_main_screen", func(...*sciter.Value) *sciter.Value {
		mainScreen.Show()
		return nil
	})

	splashScreen.Show()

	go func() {
		time.Sleep(6 * time.Second)
		splashScreen.Call("close_splash", sciter.NewValue("ok"))
	}()
}
