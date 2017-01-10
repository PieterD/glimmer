package win

import (
	"fmt"
	"runtime"
	"sync/atomic"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	runtime.LockOSThread()
}

// The Window type wraps up the GLFW Window.
// Glimmer only supports one window at a time.
type Window struct {
	w     *glfw.Window
	close int32
	poll  bool
}

// Create a new Window with the given options.
func New(opts ...WindowOption) (*Window, error) {
	opt := windowOption{
		resizable:    false,
		majorVersion: 2,
		minorVersion: 0,
		width:        800,
		height:       600,
		title:        "",
		poll:         false,
	}
	for _, o := range opts {
		o(&opt)
	}

	resizable := glfw.False
	if opt.resizable {
		resizable = glfw.True
	}
	err := glfw.Init()
	if err != nil {
		return nil, fmt.Errorf("Error initializing glfw: %v", err)
	}
	glfw.WindowHint(glfw.Resizable, resizable)
	glfw.WindowHint(glfw.ContextVersionMajor, opt.majorVersion)
	glfw.WindowHint(glfw.ContextVersionMinor, opt.minorVersion)
	w, err := glfw.CreateWindow(opt.width, opt.height, opt.title, nil, nil)
	if err != nil {
		glfw.Terminate()
		return nil, fmt.Errorf("Error creating glfw window: %v", err)
	}
	w.MakeContextCurrent()
	err = gl.Init()
	if err != nil {
		w.Destroy()
		glfw.Terminate()
		return nil, fmt.Errorf("Error initializing opengl: %v", err)
	}
	glfw.SwapInterval(1)
	return &Window{
		w:    w,
		poll: opt.poll,
	}, nil
}

// Destroy the given window.
//
// This should be the very last Glimmer call you make.
func (window *Window) Destroy() {
	window.w.Destroy()
	glfw.Terminate()
}

// Mark the window for closing. After this, ShouldClose will return true.
//
// This may be called from any Goroutine.
func (window *Window) Close() {
	atomic.StoreInt32(&window.close, 1)
}

// Return true if Window.Close was called, or if the window was closed by the user.
func (window *Window) ShouldClose() bool {
	if atomic.LoadInt32(&window.close) == 1 {
		return true
	}
	return window.w.ShouldClose()
}

// Return the GLFW 3.2 Window handle.
func (window *Window) Glfw() *glfw.Window {
	return window.w
}

// Post an empty event in the GLFW event queue, waking up Window.Swap as fast as it can.
// You may have to wait for the vsync.
// It does nothing if the Poll option was provided at Window creation.
//
// This may be called from any Goroutine.
func (window *Window) WakeUp() {
	glfw.PostEmptyEvent()
}

// Swap the front and back buffers, and pump the event queue.
func (window *Window) Swap() {
	window.w.SwapBuffers()
	if window.poll {
		glfw.PollEvents()
	} else {
		glfw.WaitEvents()
	}
}

type windowOption struct {
	resizable    bool
	majorVersion int
	minorVersion int
	width        int
	height       int
	title        string
	poll         bool
}

// Options for Window creation.
type WindowOption func(opt *windowOption)

// Provide the window size in pixels.
//
// Default: 800x600
func Size(width, height int) WindowOption {
	return func(opt *windowOption) {
		opt.width = width
		opt.height = height
	}
}

// Provide the window title.
func Title(title string) WindowOption {
	return func(opt *windowOption) {
		opt.title = title
	}
}

// Make the window resizable.
func Resizable() WindowOption {
	return func(opt *windowOption) {
		opt.resizable = true
	}
}

// Set the minimum OpenGL context version.
//
// Default: 2.0
func Version(major, minor int) WindowOption {
	return func(opt *windowOption) {
		opt.majorVersion = major
		opt.minorVersion = minor
	}
}

// Set the window to poll for events, instead of waiting.
//
// By default, Window.Swap will block after swapping the front and back buffers, waiting for events.
// With this, Window.Swap will pump all waiting events, and then return immediately without waiting.
func Poll() WindowOption {
	return func(opt *windowOption) {
		opt.poll = true
	}
}
