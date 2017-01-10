package main

import (
	"fmt"
	"runtime"
)

// Thread safety in GLFW:
// http://www.glfw.org/docs/latest/intro_guide.html

func init() {
	runtime.LockOSThread()
}

func main() {
	err := Start(
		Size(800, 600),
		Version(2, 0),
		Title("Main window"),
		Func(run1))
	if err != nil {
		fmt.Printf("Final error: %v\n", err)
	}
}

func run1(w *Window) {
	for ie := range w.Events() {
		fmt.Printf("  RUN1 Event: %#v\n", ie)
		switch e := ie.(type) {
		case EventClose:
			return
		case EventChar:
			if e.Char == 'n' {
				New(
					Title("run2"),
					Func(run2))
			} else if e.Char == 'q' {
				Terminate()
			}
		}
	}
}

func run2(w *Window) {
	for e := range w.Events() {
		fmt.Printf("  RUN2 Event: %#v\n", e)
		switch e.(type) {
		case EventClose:
			return
		}
	}
}
