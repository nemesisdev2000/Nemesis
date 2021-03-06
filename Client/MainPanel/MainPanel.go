package MainPanel

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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

		listenerWindow := a.NewWindow("Listeners")
		listenerWindow.Resize(fyne.NewSize(700, 400))
		listenerWindow.CenterOnScreen()
		fmt.Println("Number of listeners : ", d)

		headers := []string{"Listener ID", "Listener Details"}
		var listenerDetails [][]string
		listenerDetails = append(listenerDetails, headers)

		type fn func(string)
		funcLists := map[string]fn{
			"Stop Listener": ClientComms.StopListener,
		}

		if len(d) > 1 {
			t := MakeTable(d)
			listenerDetails = append(listenerDetails, t...)
			list := widget.NewTable(
				func() (int, int) {
					return len(listenerDetails), len(listenerDetails[0])
				},
				func() fyne.CanvasObject {
					return widget.NewLabel("")
				},
				func(i widget.TableCellID, o fyne.CanvasObject) {
					if strings.Contains(listenerDetails[i.Row][i.Col], "Listener") {
						o.(*widget.Label).SetText(listenerDetails[i.Row][i.Col])
						o.(*widget.Label).TextStyle = fyne.TextStyle{Bold: true}
					}
					o.(*widget.Label).SetText(listenerDetails[i.Row][i.Col])
				})

			var funcName, ID string
			del := widget.NewButton("Ok", func() { funcLists[funcName](ID) })
			list.OnSelected = func(id widget.TableCellID) {
				ID = listenerDetails[id.Row][id.Col]
				drop := widget.NewSelect(
					[]string{"Stop Listener"},
					func(s string) {
						del.Refresh()
						fmt.Println("Label : ", s)
						funcName = s
					})
				editconf := container.NewVBox(drop, del)
				listenerWindow.SetContent(editconf)
			}

			fmt.Println("Details : ", listenerDetails)
			list.SetColumnWidth(0, 300)

			listenerWindow.SetContent(list)
		}
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
