package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {

	but := tview.NewButton("First button")
	but.SetBackgroundColor(tcell.GetColor("yellow"))
	but.SetLabelColorActivated(tcell.GetColor("maroon"))

	if err := tview.NewApplication().SetRoot(but, false).Run(); err != nil {
		panic(err)
	}
}
