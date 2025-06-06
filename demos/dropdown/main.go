// Demo code for the DropDown primitive.
package main

import "github.com/sedwards2009/nuview"

func main() {
	app := nuview.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	dropdown := nuview.NewDropDown()
	dropdown.SetLabel("Select an option (hit Enter): ")
	dropdown.SetOptions(nil,
		nuview.NewDropDownOption("First"),
		nuview.NewDropDownOption("Second"),
		nuview.NewDropDownOption("Third"),
		nuview.NewDropDownOption("Fourth"),
		nuview.NewDropDownOption("Fifth"))

	app.SetRoot(dropdown, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
