package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"

	//"fyne.io/fyne/theme"

	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")

	tabs := widget.NewTabContainer(
		widget.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		widget.NewTabItem("Tab 2", widget.NewButton("World!", func() {
			go showAnother(myApp)
		})))

	//widget.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab"))

	tabs.SetTabLocation(widget.TabLocationLeading)

	myWindow.SetContent(tabs)

	myWindow.ShowAndRun()
}

func showAnother(a fyne.App) {

	win := a.NewWindow("My Another Window")

	dialog.ShowConfirm("Hello", "This is sample message", callback, win)
	win.SetContent(widget.NewLabel("This is second window"))
	win.Resize(fyne.NewSize(800, 600))
	win.Show()
}

func callback(s bool) {
	if s == true {
		fmt.Println("You clock on yes")
	} else {
		fmt.Println("You clock on no")
	}
}
