// Demo code for the Table primitive.
package main

import (
	"log"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	cview "github.com/sedwards2009/nuview"
)

const loremIpsumText = "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet."

func main() {
	setupLogging()
	app := cview.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	table := cview.NewTable()
	table.SetBorders(true)
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
			cell := cview.NewTableCell(lorem[word])
			cell.SetTextColor(color)
			cell.SetAlign(cview.AlignCenter)
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

	app.SetRoot(table, true)
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
