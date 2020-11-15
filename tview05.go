package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	var checkboxes [6][6]*tview.Checkbox

	app := tview.NewApplication()

	grid := tview.NewGrid()
	grid.SetSize(6, 6, 1, 1)
	grid.SetGap(1, 1)
	for ri := 0; ri < 6; ri++ {
		for ci := 0; ci < 6; ci++ {
			chkb := tview.NewCheckbox()
			checkboxes[ri][ci] = chkb
			grid.AddItem(chkb, ri, ci, 1, 1, 1, 1, false)
		}
	}

	rl := 0
	cl := 0
	grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Name() == "Esc" {
			app.Stop()
		}

		rmove := 0
		if event.Name() == "Up" {
			rmove = -1
		}
		if event.Name() == "Down" {
			rmove = 1
		}

		cmove := 0
		if event.Name() == "Left" {
			cmove = -1
		}
		if event.Name() == "Right" {
			cmove = 1
		}

		rl = CordRotate((rl + rmove), 0, 6)
		cl = CordRotate((cl + cmove), 0, 6)

		app.SetFocus(checkboxes[rl][cl])

		return event
	})

	if err := app.SetRoot(grid, false).Run(); err != nil {
		panic(err)
	}
}

func CordRotate(val, min, max int) int {
	if val < min {
		val = max - 1
	}
	if val > max-1 {
		val = min
	}
	return val
}
