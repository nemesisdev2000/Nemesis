package ListenerSetup

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/nemesisdev2000/Nemesis/Client/ClientComms"
)

func SetupTcpListener(a fyne.App) {

	label := widget.NewLabel("")

	setupWindow := a.NewWindow("Configure Tcp Listener")
	setupWindow.Resize(fyne.NewSize(700, 1000))
	hostname := widget.NewEntry()
	hostname.SetPlaceHolder("Enter Hostname of Listener")
	hostname.Resize(fyne.NewSize(250, 30))
	hostname.Move(fyne.NewPos(100, 100))

	port := widget.NewEntry()
	port.SetPlaceHolder("Enter the Bind port")
	port.Resize(fyne.NewSize(250, 30))
	port.Move(fyne.NewPos(100, 150))

	done_btn := widget.NewButton("Done", func() {
		fmt.Println("Hostname : ", hostname.Text)
		fmt.Println("Port : ", port.Text)

		stat := ClientComms.SendTcpListener(port.Text, hostname.Text)
		if stat == true {
			label := widget.NewLabel("")
			label.Text = "Listener Created"
			successPopUp := widget.NewPopUp(label, setupWindow.Canvas())
			successPopUp.ShowAtPosition(fyne.NewPos(300, 300))
		} else {
			label.Move(fyne.NewPos(500, 500))
			label.Text = "Error"
			label.Refresh()
		}
	})

	done_btn.Resize(fyne.NewSize(50, 50))
	done_btn.Move(fyne.NewPos(150, 250))

	content := container.NewWithoutLayout(hostname, port, done_btn, label)
	setupWindow.SetContent(content)
	setupWindow.Show()
}
