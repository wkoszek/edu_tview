package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	b0 := tview.NewButton("First button")
	b0.SetBackgroundColor(tcell.GetColor("black"))
	b0.SetBackgroundColorActivated(tcell.GetColor("darkblue"))
	b0.SetLabelColor(tcell.GetColor("orange"))

	b1 := tview.NewButton("Second button")
	b1.SetBackgroundColor(tcell.GetColor("black"))
	b1.SetBackgroundColorActivated(tcell.GetColor("darkred"))
	b1.SetLabelColor(tcell.GetColor("orange"))

	flex := tview.NewFlex()
	flex.AddItem(b0, 20, 2, true)
	flex.AddItem(b1, 20, 2, false)

	idx := 0
	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			idx++

			if idx%2 == 0 {
				app.SetFocus(b0)
			}
			if idx%2 == 1 {
				app.SetFocus(b1)
			}
		}

		return nil
	})

	if err := app.SetRoot(flex, false).Run(); err != nil {
		panic(err)
	}
}
