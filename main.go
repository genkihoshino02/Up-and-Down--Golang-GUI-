package main

import (
	// "strconv"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	p := widget.NewProgressBar()
	v := 0.
	b := widget.NewButton("Up!", func() {
		v += 0.1
		if v > 1.0 {
			v = 0.
		}
		p.SetValue(v)
	})
	c := widget.NewButton("Down!", func() {
		v -= 0.1
		if v < 0.0 {
			v = 0.
		}
		p.SetValue(v)
	})
	w.SetContent(
		widget.NewVBox(
			l, p, b, c,
		),
	)
	w.ShowAndRun()
}
