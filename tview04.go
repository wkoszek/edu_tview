package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	var numKeyPresses = 0

	app := tview.NewApplication()

	t0 := tview.NewTextView()
	t0.SetScrollable(true)
	t0.SetTextColor(tcell.GetColor("orange"))
	t0.SetText(fmt.Sprintf("[Start]"))
	t0.SetTextAlign(tview.AlignCenter)

	b0 := tview.NewButton("First button")
	b0.SetBackgroundColor(tcell.GetColor("black"))
	b0.SetBackgroundColorActivated(tcell.GetColor("darkblue"))
	b0.SetLabelColor(tcell.GetColor("orange"))
	b0.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Name() == "Enter" {
			numKeyPresses++
			t0.SetText(fmt.Sprintf("Keypress %d", numKeyPresses))
		}
		if event.Name() == "Esc" {
			app.Stop()
		}
		if event.Name() == "Rune[c]" || event.Name() == "Rune[C]" {
			t0.SetText("hhhh")
		}

		return nil
	})

	flex := tview.NewFlex()
	flex.AddItem(b0, 20, 2, true)
	flex.AddItem(t0, 20, 2, false)

	if err := app.SetRoot(flex, false).Run(); err != nil {
		panic(err)
	}
}
