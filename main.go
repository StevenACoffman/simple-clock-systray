package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/getlantern/systray"
)

var timezone string

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	// set initial default time zone
	timezone = "America/Los_Angeles"

	//systray.SetTemplateIcon(icon.GetIcon(), icon.GetIcon())

	localTime := systray.AddMenuItem("Local time", "Local time")
	hcmcTime := systray.AddMenuItem("Ann Arbor", "America/Detroit")
	sfTime := systray.AddMenuItem("San Fransisco time", "America/Los_Angeles")
	sydTime := systray.AddMenuItem("Sydney time", "Australia/Sydney")
	gdlTime := systray.AddMenuItem("Guadalajara time", "America/Mexico_City")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quits this app")

	go func() {
		<-mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	go func() {
		for {
			systray.SetTitle(getClockTime(timezone))
			systray.SetTooltip(timezone + " timezone")
			time.Sleep(1 * time.Second) // we could probably wake up less often?
		}
	}()

	go func() {
		for {
			select {
			case <-localTime.ClickedCh:
				timezone = "Local"
			case <-hcmcTime.ClickedCh:
				timezone = "Asia/Ho_Chi_Minh"
			case <-sydTime.ClickedCh:
				timezone = "Australia/Sydney"
			case <-gdlTime.ClickedCh:
				timezone = "America/Mexico_City"
			case <-sfTime.ClickedCh:
				timezone = "America/Los_Angeles"
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// Cleaning stuff here.
}

func getClockTime(tz string) string {
	t := time.Now()
	utc, _ := time.LoadLocation(tz)

	return t.In(utc).Format(time.Kitchen)
}
