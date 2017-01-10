package main

import "runtime"

// Thread safety in GLFW:
// http://www.glfw.org/docs/latest/intro_guide.html

func init() {
	runtime.LockOSThread()
}

// Create a new window with the given options, returning the window handle and possibly an error.
func New(opts ...WindowOption) (*Window, error) {
	opt := applyWindowOptions(opts...)
	w, err := mainApp.newWindow(opt)
	if err != nil {
		return nil, err
	}
	go func() {
		defer mainApp.destroyWindow(w)
		w.run(opt.f)
	}()
	return w, nil
}

// Start the main loop, and run the provided function after this has been done.
// If f returns an error, Main will return this error.
func Main(f func() error) error {
	errChan := make(chan error, 1)
	mainApp = newApplication()
	go func() {
		defer close(errChan)
		<-mainApp.started
		err := f()
		if err != nil {
			errChan <- err
			Terminate()
		}
	}()
	mainApp.loop()
	return <-errChan
}

// Start the main loop, and create a window with the given options.
func Start(opts ...WindowOption) error {
	return Main(func() error {
		_, err := New(opts...)
		return err
	})
}

// Terminate the application, closing all windows and the main loop.
func Terminate() {
	mainApp.eventChan <- appEventDestroyApplication{}
}
