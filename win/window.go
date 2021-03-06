package win

import (
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Represents a single window.
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
	gl.Init()
	f(w)
}

// Swap the front and back buffers.
func (w *Window) Swap() {
	w.gw.SwapBuffers()
}

// Return this window's event channel.
func (w *Window) Events() <-chan interface{} {
	return w.ev
}

// Try to read a single event off the event channel, or return nil if none are buffered.
func (w *Window) Poll() interface{} {
	for {
		select {
		case ie := <-w.Events():
			return ie
		default:
			return nil
		}
	}
}

// Close this window.
func (w *Window) Destroy() {
	close(w.closing)
	<-w.closed
}

func (w *Window) initialize() {
	w.closing = make(chan struct{})
	w.closed = make(chan struct{})
	w.pev = make(chan interface{}, 10)
	w.ev = make(chan interface{}, 1000)
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
		w.pev <- EventResize{
			Width:  width,
			Height: height,
		}
	})

	w.gw.SetCloseCallback(func(gw *glfw.Window) {
		w.pev <- EventClose{}
	})

	w.gw.SetKeyCallback(func(gw *glfw.Window, key glfw.Key, scanCode int, action glfw.Action, mod glfw.ModifierKey) {
		w.pev <- EventKey{
			Key:      Key(key),
			ScanCode: scanCode,
			Action:   Action(action),
			Mod:      Mod(mod),
		}
	})

	w.gw.SetCharCallback(func(gw *glfw.Window, key rune) {
		w.pev <- EventChar{
			Char: key,
		}
	})

	w.gw.SetCursorPosCallback(func(gw *glfw.Window, x, y float64) {
		w.pev <- EventMousePos{
			X: int(x),
			Y: int(y),
		}
	})

	w.gw.SetMouseButtonCallback(func(gw *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
		w.pev <- EventMouseButton{
			Button: Button(button),
			Action: Action(action),
			Mod:    Mod(mod),
		}
	})
}
