// Demo code for the Table primitive.
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/sedwards2009/nuview"
)

const loremIpsumText = "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet."

func main() {
	setupLogging()
	app := nuview.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	table := nuview.NewTable()
	// table.SetBorders(true)
	table.SetSelectable(true, false)

	lorem := strings.Split(loremIpsumText, " ")
	cols, rows := 20, 40
	word := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			color := tcell.ColorWhite.TrueColor()
			if c < 1 || r < 1 {
				color = tcell.ColorYellow.TrueColor()
			}
			cell := nuview.NewTableCell(fmt.Sprintf("r%d,c%d,%s", r, c, lorem[word]))
			cell.SetTextColor(color)
			cell.SetAlign(nuview.AlignCenter)
			table.SetCell(r, c, cell)
			word = (word + 1) % len(lorem)
		}
	}
	table.Select(0, 0)
	table.SetFixed(1, 1)
	table.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	})
	table.SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed.TrueColor())
		table.SetSelectable(false, false)
	})

	layout := nuview.NewFlex()

	topBox := nuview.NewBox()
	topBox.SetBorder(true)
	topBox.SetBorderAttributes(tcell.AttrBold)
	topBox.SetTitle("Top box")
	layout.AddItem(topBox, 10, 0, false)

	innerLayout := nuview.NewFlex()

	// Create button panel
	buttonPanel := nuview.NewFlex()
	buttonPanel.SetBorder(true)
	buttonPanel.SetBorderAttributes(tcell.AttrBold)
	buttonPanel.SetTitle("Controls")
	buttonPanel.SetDirection(nuview.FlexRow)

	// Create + button
	plusButton := nuview.NewButton("+")
	plusButton.SetSelectedFunc(func() {
		offset := table.GetXScroll()
		table.SetXScroll(offset + 1)
	})

	// Create - button
	minusButton := nuview.NewButton("-")
	minusButton.SetSelectedFunc(func() {
		offset := table.GetXScroll()
		table.SetXScroll(offset - 1)
	})

	// Create + button
	plus10Button := nuview.NewButton("+10")
	plus10Button.SetSelectedFunc(func() {
		offset := table.GetXScroll()
		table.SetXScroll(offset + 10)
	})

	// Create - button
	minus10Button := nuview.NewButton("-10")
	minus10Button.SetSelectedFunc(func() {
		offset := table.GetXScroll()
		table.SetXScroll(offset - 10)
	})

	// Add buttons to button panel
	buttonPanel.AddItem(plusButton, 3, 0, false)
	buttonPanel.AddItem(minusButton, 3, 0, false)
	buttonPanel.AddItem(plus10Button, 3, 0, false)
	buttonPanel.AddItem(minus10Button, 3, 0, false)

	innerLayout.AddItem(buttonPanel, 10, 0, false)
	innerLayout.AddItem(table, 0, 1, true)

	rightBox := nuview.NewBox()
	rightBox.SetBorder(true)
	rightBox.SetBorderAttributes(tcell.AttrBold)
	rightBox.SetTitle("Right box")

	innerLayout.AddItem(rightBox, 10, 0, false)
	innerLayout.SetDirection(nuview.FlexColumn)

	layout.AddItem(innerLayout, 0, 1, true)

	bottomBox := nuview.NewBox()
	bottomBox.SetBorder(true)
	bottomBox.SetBorderAttributes(tcell.AttrBold)
	bottomBox.SetTitle("Bottom box")
	layout.AddItem(bottomBox, 10, 0, false)
	layout.SetDirection(nuview.FlexRow)

	app.SetRoot(layout, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func setupLogging() *os.File {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}
	log.SetOutput(logFile)
	return logFile
}
