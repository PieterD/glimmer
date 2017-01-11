package main

import (
	"fmt"

	"github.com/PieterD/glimmer/win"
	. "github.com/PieterD/pan"
)

func main() {
	err := win.Start(
		win.Size(800, 600),
		win.Version(2, 0),
		win.Title("Main window"),
		win.Func(run1))
	Panic(err)
}

func run1(w *win.Window) {
	for ie := range w.Events() {
		fmt.Printf("  RUN1 Event: %#v\n", ie)
		switch e := ie.(type) {
		case win.EventClose:
			return
		case win.EventKey:
			if e.Action.Release() {
				if e.Key == win.KeyN {
					win.New(
						win.Title("run2"),
						win.Func(run2))
				} else if e.Key == win.KeyQ {
					win.Terminate()
				}
			}
		}
	}
}

func run2(w *win.Window) {
	for e := range w.Events() {
		fmt.Printf("  RUN2 Event: %#v\n", e)
		switch e.(type) {
		case win.EventClose:
			return
		}
	}
}
