package main

import (
	"errors"
	"fmt"

	. "github.com/PieterD/pan"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var mainApp *application

type application struct {
	started   chan struct{}
	eventChan chan interface{}

	id      uint64
	windows map[*Window]struct{}
}

func newApplication() *application {
	return &application{
		started:   make(chan struct{}),
		eventChan: make(chan interface{}, 1),
		id:        1,
		windows:   make(map[*Window]struct{}),
	}
}

func (app *application) loop() {
	err := glfw.Init()
	Panic(err)
	defer glfw.Terminate()
	close(app.started)

	app.mainLoop()
	for w := range app.windows {
		w.Destroy()
	}
	app.closeLoop()
}

func (app *application) mainLoop() {
	started := false

	for !started || len(app.windows) > 0 {
		glfw.WaitEvents()
		select {
		case ie := <-app.eventChan:
			switch e := ie.(type) {
			case appEventNewWindow:
				started = true
				w, err := app.loopNewWindow(e.opt)
				e.ret <- appEventNewWindowReturn{w: w, err: err}
			case appEventDestroyWindow:
				app.loopDestroyWindow(e)
			case appEventDestroyApplication:
				return
			}
		default:
		}
	}
}

func (app *application) closeLoop() {
	for len(app.windows) > 0 {
		glfw.PollEvents()
		select {
		case ie := <-app.eventChan:
			switch e := ie.(type) {
			case appEventNewWindow:
				e.ret <- appEventNewWindowReturn{err: errors.New("Not creating new window: Application is closing")}
			case appEventDestroyWindow:
				app.loopDestroyWindow(e)
			}
		default:
		}
	}
}

func (app *application) loopNewWindow(opt windowOption) (*Window, error) {
	resizable := glfw.False
	if opt.resizable {
		resizable = glfw.True
	}

	glfw.WindowHint(glfw.Resizable, resizable)
	glfw.WindowHint(glfw.ContextVersionMajor, opt.majorVersion)
	glfw.WindowHint(glfw.ContextVersionMinor, opt.minorVersion)
	gw, err := glfw.CreateWindow(opt.width, opt.height, opt.title, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating glfw window: %v", err)
	}
	w := &Window{
		id: app.id,
		gw: gw,
	}
	app.id++
	w.initialize()
	app.windows[w] = struct{}{}
	return w, nil
}

func (app *application) newWindow(opt windowOption) (*Window, error) {
	<-app.started
	ret := make(chan appEventNewWindowReturn)
	app.eventChan <- appEventNewWindow{
		opt: opt,
		ret: ret,
	}
	glfw.PostEmptyEvent()
	rv := <-ret
	return rv.w, rv.err
}

func (app *application) loopDestroyWindow(e appEventDestroyWindow) {
	e.w.gw.Destroy()
	close(e.w.pev)
	delete(app.windows, e.w)
}

func (app *application) destroyWindow(w *Window) {
	<-app.started
	app.eventChan <- appEventDestroyWindow{w: w}
	glfw.PostEmptyEvent()
}
