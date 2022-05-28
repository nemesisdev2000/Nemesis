package MainPanel

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/nemesisdev2000/Nemesis/Client/ClientComms"
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
	showListeners := fyne.NewMenuItem("Show Listeners", func() {
		d := ClientComms.ShowListeners()
		//_ = ClientComms.ShowListeners()

		listenerWindow := a.NewWindow("Listeners")
		listenerWindow.Resize(fyne.NewSize(700, 400))

		//content := container.NewMax()
		listenerIDHead := widget.NewLabel("Listener ID")
		connectionHead := widget.NewLabel("Listener Details")
		connectionHead.Wrapping = fyne.TextWrapWord

		spacer := layout.NewSpacer()
		spacer.Resize(fyne.NewSize(30, 40))
		spacer.Move(fyne.NewPos(1, 1))

		//listenerID := container.NewBorder(container.NewHBox(listenerIDHead, spacer, widget.NewSeparator(), connectionHead), nil, nil, nil, content)

		listenerDetails := MakeTable(d)

		fmt.Println("Details : ", listenerDetails)

		contain := container.NewHBox(listenerIDHead, widget.NewSeparator(), connectionHead)
		details := container.NewBorder(contain, nil, nil, nil)

		listenerWindow.SetContent(details)
		listenerWindow.Show()
	})
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

func MakeTable(d []string) [][]string {
	var t [][]string
	var v []string

	for i := 0; i < len(d); i++ {
		if len(d[i]) > 0 {
			if strings.Contains(d[i], ":") {
				v = append(v, d[i])
				t = append(t, v)
				v = nil
			} else {
				v = append(v, d[i])
			}
		}
	}
	return t
}
