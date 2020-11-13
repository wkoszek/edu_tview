package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {

	app := tview.NewApplication()

	but := tview.NewButton("First button")
	but.SetBackgroundColorActivated(tcell.ColorMaroon)
	but.SetLabelColorActivated(tcell.GetColor("yellow"))

	if err := app.SetRoot(but, false).Run(); err != nil {
		panic(err)
	}
}
