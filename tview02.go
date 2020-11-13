package main

import (
	"fmt"

	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	b := tview.NewButton("First button")
	b.SetBackgroundColor(tcell.GetColor("black"))
	b.SetBackgroundColorActivated(tcell.GetColor("darkblue"))
	b.SetLabelColor(tcell.GetColor("orange"))

	enterCount := 0
	b.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			fmt.Printf("ENTER %d\n", enterCount)
			enterCount++
			time.Sleep(100 * time.Millisecond)
		}

		return nil
	})

	if err := app.SetRoot(b, false).Run(); err != nil {
		panic(err)
	}
}
