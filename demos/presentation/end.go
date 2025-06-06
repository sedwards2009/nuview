package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/sedwards2009/nuview"
)

// End shows the final slide.
func End(nextSlide func()) (title string, info string, content nuview.Primitive) {
	textView := nuview.NewTextView()
	textView.SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	url := "https://codeberg.org/tslocum/cview"
	fmt.Fprint(textView, url)
	return "End", "", Center(len(url), 1, textView)
}
