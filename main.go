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
	now := time.Now()
	zoneName, _ := now.Zone()
	// If zoneName == EST, then assume we are in Ann Arbor, MI
	//  and so we want to display mountain view
	timezone = "America/Los_Angeles"
	// If zoneName == PST, then assume we are in Mountain View, CA
	//  so we want to display Ann Arbor, MI
	if zoneName == "PST" {
		timezone = "America/Detroit"
	}

	//systray.SetTemplateIcon(icon.GetIcon(), icon.GetIcon())

	localTime := systray.AddMenuItem("Local time", "Local time")
	hcmcTime := systray.AddMenuItem("Ann Arbor", "America/Detroit")
	sfTime := systray.AddMenuItem("Khan HQ time", "America/Los_Angeles")
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
