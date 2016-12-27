package win

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Window struct {
	poll bool
	w    *glfw.Window
}

func New(opts ...WindowOption) (*Window, error) {
	opt := windowOption{
		resizable:    false,
		majorVersion: 2,
		minorVersion: 1,
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

func (window *Window) Destroy() {
	window.w.Destroy()
	glfw.Terminate()
}

func (window *Window) ShouldClose() bool {
	return window.w.ShouldClose()
}

func (window *Window) Glfw() *glfw.Window {
	return window.w
}

func (window *Window) WakeUp() {
	glfw.PostEmptyEvent()
}

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

type WindowOption func(opt *windowOption)

func Size(width, height int) WindowOption {
	return func(opt *windowOption) {
		opt.width = width
		opt.height = height
	}
}

func Title(title string) WindowOption {
	return func(opt *windowOption) {
		opt.title = title
	}
}

func Resizable() WindowOption {
	return func(opt *windowOption) {
		opt.resizable = true
	}
}

func Version(major, minor int) WindowOption {
	return func(opt *windowOption) {
		opt.majorVersion = major
		opt.minorVersion = minor
	}
}

func Poll() WindowOption {
	return func(opt *windowOption) {
		opt.poll = true
	}
}
