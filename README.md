# nuview - Terminal-based user interface toolkit

This package is a fork of [cview](https://codeberg.org/tslocum/cview) which was a fork of [tview](https://github.com/rivo/tview).
See [FORK.md](FORK.md) for more information.

## Demo

[![Recording of presentation demo](cview.svg)](https://github.com/sedwards2009/nuview/tree/master/demos/presentation)

## Features

Available widgets:

- __Input forms__ (including __input/password fields__, __drop-down selections__, __checkboxes__, and __buttons__)
- Navigable multi-color __text views__
- Selectable __lists__ with __context menus__
- Modal __dialogs__
- Horizontal and vertical __progress bars__
- __Grid__, __Flexbox__ and __tabbed panel layouts__
- Sophisticated navigable __table views__
- Flexible __tree views__
- Draggable and resizable __windows__
- An __application__ wrapper

Widgets may be customized and extended to suit any application.

[Mouse support](#hdr-Mouse_Support) is available.


## Installation

```bash
go get github.com/sedwards2009/nuview
```

## Hello World

This basic example creates a TextView titled "Hello, World!" and displays it in your terminal:

```go
package main

import (
	"github.com/sedwards2009/nuview"
)

func main() {
	app := nuview.NewApplication()

	tv := nuview.NewTextView()
	tv.SetBorder(true)
	tv.SetTitle("Hello, world!")
	tv.SetText("Lorem ipsum dolor sit amet")

	app.SetRoot(tv, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
```

Demo are available in the [demos](demos) directory.

For a presentation highlighting the features of this package, compile and run
the program in the [demos/presentation](demos/presentation) directory.

## Dependencies

This package is based on [github.com/gdamore/tcell](https://github.com/gdamore/tcell)
(and its dependencies) and [github.com/rivo/uniseg](https://github.com/rivo/uniseg).

## Support

PR and issues can done up at https://github.com/sedwards2009/nuview .
