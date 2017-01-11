package win

func applyWindowOptions(opts ...WindowOption) windowOption {
	opt := windowOption{
		resizable:    false,
		majorVersion: 2,
		minorVersion: 0,
		width:        800,
		height:       600,
		title:        "",
		f:            nil,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

type windowOption struct {
	resizable    bool
	majorVersion int
	minorVersion int
	width        int
	height       int
	title        string
	f            func(w *Window)
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

func Func(f func(w *Window)) WindowOption {
	return func(opt *windowOption) {
		opt.f = f
	}
}
