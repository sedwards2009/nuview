// Demo code for the Panels primitive.
package main

import (
	"fmt"

	cview "github.com/sedwards2009/nuview"
)

const panelCount = 5

func main() {
	app := cview.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	panels := cview.NewPanels()
	for panel := 0; panel < panelCount; panel++ {
		func(panel int) {
			modal := cview.NewModal()
			modal.SetText(fmt.Sprintf("This is page %d. Choose where to go next.", panel+1))
			modal.AddButtons([]string{"Next", "Quit"})
			modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonIndex == 0 {
					panels.SetCurrentPanel(fmt.Sprintf("panel-%d", (panel+1)%panelCount))
				} else {
					app.Stop()
				}
			})

			panels.AddPanel(fmt.Sprintf("panel-%d", panel), modal, false, panel == 0)
		}(panel)
	}

	app.SetRoot(panels, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
