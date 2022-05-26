package MainPanel

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func ShowMainWindow(w fyne.Window, a fyne.App) {
	window := a.NewWindow("New Window")

	window.Resize(fyne.NewSize(1000, 700))
	window.CenterOnScreen()

	menuItem1 := fyne.NewMenuItem("New", func() { fmt.Println("New button pressed") })
	menuItem2 := fyne.NewMenuItem("Open", func() { fmt.Println("Open button pressed") })

	newmenu1 := fyne.NewMenu("File", menuItem1, menuItem2)
	newmenu2 := fyne.NewMenu("Listener", menuItem1, menuItem2)
	newmenu3 := fyne.NewMenu("Payloads", menuItem1, menuItem2)
	menuItem1.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Types", func() { fmt.Println("Pressed Types") }),
		fyne.NewMenuItem("Setup", func() { fmt.Println("Pressed Types") }),
		fyne.NewMenuItem("Create", func() { fmt.Println("Pressed Types") }),
	)

	menu := fyne.NewMainMenu(newmenu1, newmenu2, newmenu3)

	window.SetMainMenu(menu)
	w.Close()
	window.Show()
}
