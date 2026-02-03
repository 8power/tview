package main

import (
	"fmt"

	"github.com/8power/tcell/v2"
	"github.com/8power/tview"
)

// End shows the final slide.
func End(nextSlide func()) (title string, content tview.Primitive) {
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetDoneFunc(func(key tcell.Key) {
			nextSlide()
		})
	url := "[:::https://github.com/8power/tview]https://github.com/8power/tview"
	fmt.Fprint(textView, url)
	return "End", Center(tview.TaggedStringWidth(url), 1, textView)
}
