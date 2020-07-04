package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"

	//"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")

	tabs := widget.NewTabContainer(
		widget.NewTabItem("Tab 1", widget.NewLabel("Hello")),
<<<<<<< HEAD
		widget.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)
=======
		widget.NewTabItem("Tab 2", widget.NewButton("World!", func() {
			go showAnother(myApp)
		})))
>>>>>>> 43219543ccc1cfbd89837cf07485a3acd451bdef

	//widget.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab"))

	tabs.SetTabLocation(widget.TabLocationLeading)

	myWindow.SetContent(tabs)

	myWindow.ShowAndRun()
}

func showAnother(a fyne.App) {
	win := a.NewWindow("My Another Window")
	win.SetContent(widget.NewLabel("This is second window"))
	win.Resize(fyne.NewSize(800, 600))
	win.Show()
}
