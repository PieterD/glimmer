package gli

import "github.com/go-gl/gl/v2.1/gl"

type Clear struct {
	clearOption
}

func NewClear(opts ...ClearOption) (*Clear, error) {
	clear := Clear{}
	for _, o := range opts {
		o(&clear.clearOption)
	}
	return &clear, nil
}

func (clear *Clear) Clear() {
	bits := uint32(0)
	if clear.color {
		bits |= gl.COLOR_BUFFER_BIT
		gl.ClearColor(clear.r, clear.g, clear.b, clear.a)
	}
	if clear.depth {
		bits |= gl.DEPTH_BUFFER_BIT
	}
	if clear.stencil {
		bits |= gl.STENCIL_BUFFER_BIT
	}
	gl.Clear(bits)
}

type ClearOption func(*clearOption)

type clearOption struct {
	color      bool
	r, g, b, a float32
	depth      bool
	stencil    bool
}

func ClearColor(r, g, b, a float32) ClearOption {
	return func(opt *clearOption) {
		opt.color = true
		opt.r = r
		opt.g = g
		opt.b = b
		opt.a = a
	}
}

func ClearDepth() ClearOption {
	return func(opt *clearOption) {
		opt.depth = true
	}
}

func ClearStencil() ClearOption {
	return func(opt *clearOption) {
		opt.stencil = true
	}
}
