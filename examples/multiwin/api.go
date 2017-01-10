package main

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

func Start(opts ...WindowOption) error {
	errChan := make(chan error, 1)
	mainApp = newApplication()
	go func() {
		defer close(errChan)
		_, err := New(opts...)
		if err != nil {
			errChan <- err
		}
	}()
	mainApp.loop()
	return <-errChan
}

func Terminate() {
	mainApp.eventChan <-appEventDestroyApplication{}
}