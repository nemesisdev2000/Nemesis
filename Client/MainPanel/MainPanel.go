package MainPanel

import (
	"fmt"

	"fyne.io/fyne/v2"

	"github.com/nemesisdev2000/Nemesis/Client/SetupMenus/ListenerSetup"
)

func ShowMainWindow(w fyne.Window, a fyne.App) {
	window := a.NewWindow("Nemesis")

	window.Resize(fyne.NewSize(1000, 700))
	window.CenterOnScreen()

	menuItem1 := fyne.NewMenuItem("New", func() { fmt.Println("New button pressed") })
	menuItem2 := fyne.NewMenuItem("Open", func() { fmt.Println("Open button pressed") })

	newmenu1 := fyne.NewMenu("File", menuItem1, menuItem2)

	listenerConfigure := fyne.NewMenuItem("Configure Listener", func() { ListenerSetup.SetupTcpListener(a) })
	showListeners := fyne.NewMenuItem("Show Listeners", func() { fmt.Println("Shows configured listeners") })
	ListenerMenu := fyne.NewMenu("Listener", listenerConfigure, showListeners)

	newmenu3 := fyne.NewMenu("Payloads", menuItem1, menuItem2)
	menuItem1.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Types", func() { fmt.Println("Pressed Types") }),
		fyne.NewMenuItem("Setup", func() { fmt.Println("Pressed Types") }),
		fyne.NewMenuItem("Create", func() { fmt.Println("Pressed Types") }),
	)

	menu := fyne.NewMainMenu(newmenu1, ListenerMenu, newmenu3)

	window.SetMainMenu(menu)
	w.Close()
	window.Show()
}
