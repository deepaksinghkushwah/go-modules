package main

import (
	"time"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func showTime(clock *widget.Label) {
	formattedTime := time.Now().Format("03:04:05")
	clock.SetText(formattedTime)
}

func main() {
	a := app.New()
	w := a.NewWindow("Clock")

	clock := widget.NewLabel("")

	w.SetContent(clock)

	go func() {
		t := time.NewTicker(time.Second)
		for range t.C {
			showTime(clock)
		}
	}()

	w.ShowAndRun()
}
