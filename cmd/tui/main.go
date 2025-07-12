package main

import (
	"github.com/rivo/tview"
)

func main() {

	app := tview.NewApplication()
	form := tview.NewForm().
		AddInputField("User name", "", 20, nil, nil).
		AddPasswordField("Password", "", 10, '*', nil).
		AddButton("Login", nil).
		AddButton("Dont have a account? Signup", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle(" Welcome to Kasragay social media ").SetTitleAlign(tview.AlignCenter)
	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
