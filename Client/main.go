package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/nemesisdev2000/Nemesis/Client/ClientComms"
	"github.com/nemesisdev2000/Nemesis/Client/MainPanel"
)

func main() {
	a := app.New()

	w := a.NewWindow("Client Application")
	w.Resize(fyne.NewSize(400, 400))
	w.CenterOnScreen()

	label := widget.NewLabel("")

	username := widget.NewEntry()
	username.SetPlaceHolder("Username...")
	username.Resize(fyne.NewSize(250, 30))
	username.Move(fyne.NewPos(40, 100))

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password...")
	password.Resize(fyne.NewSize(250, 30))
	password.Move(fyne.NewPos(40, 150))

	login_btn := widget.NewButton("Login", func() {
		stat := ClientComms.Login(username.Text, password.Text)
		if stat {
			MainPanel.ShowMainWindow(w, a)
		} else {
			label.Text = "Wrong Credentials"
			label.Refresh()
		}
	})

	login_btn.Resize(fyne.NewSize(150, 30))
	login_btn.Move(fyne.NewPos(40, 250))
	content := container.NewWithoutLayout(username, password, login_btn, label)

	w.SetContent(content)
	w.ShowAndRun()
}
