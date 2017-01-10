package main

import (
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
)

type Window struct {
	id      uint64
	gw      *glfw.Window
	closing chan struct{}
	closed  chan struct{}
	pev     chan interface{}
	ev      chan interface{}
}

func (w *Window) run(f func(*Window)) {
	runtime.LockOSThread()
	w.gw.MakeContextCurrent()
	glfw.SwapInterval(1)
	f(w)
}

func (w *Window) Events() <-chan interface{} {
	return w.ev
}

func (w *Window) Destroy() {
	close(w.closing)
	<- w.closed
}

func (w *Window) initialize() {
	w.closing = make(chan struct{})
	w.closed = make(chan struct{})
	w.pev = make(chan interface{}, 100)
	w.ev = make(chan interface{})
	go func() {
		in := w.pev
		out := w.ev
		out = nil
		// TODO: Buffer this so we never block?
		var stored interface{}
		defer close(w.closed)
		defer close(w.ev)
		for {
			select {
			case e := <-in:
				stored = e
				in = nil
				out = w.ev
			case out <- stored:
				stored = nil
				in = w.pev
				out = nil
			case <-w.closing:
				go func() {
					for range w.pev {
					}
				}()
				return
			}
		}
	}()

	w.gw.SetSizeCallback(func(gw *glfw.Window, width, height int) {
		w.ev <- EventResize{
			Width:  width,
			Height: height,
		}
	})

	w.gw.SetCloseCallback(func(gw *glfw.Window) {
		w.ev <- EventClose{}
	})

	w.gw.SetKeyCallback(func(gw *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	})

	w.gw.SetCharCallback(func(gw *glfw.Window, key rune) {
		w.ev <- EventChar{
			Char: key,
		}
	})

	w.gw.SetCursorPosCallback(func(gw *glfw.Window, x, y float64) {
		w.ev <- EventMousePos{
			X: int(x),
			Y: int(y),
		}
	})

	w.gw.SetMouseButtonCallback(func(gw *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {

	})
}
