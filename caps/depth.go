package caps

import "github.com/go-gl/gl/v2.1/gl"

type DepthCaps struct{}

var Depth = DepthCaps{}

func (_ DepthCaps) Enable() {
	gl.Enable(gl.DEPTH)
}

func (_ DepthCaps) Disable() {
	gl.Disable(gl.DEPTH)
}

func (_ DepthCaps) Func(depthFunc DepthFunc) {
	gl.DepthFunc(uint32(depthFunc))
}

func (_ DepthCaps) Mask(option bool) {
	gl.DepthMask(option)
}

func (_ DepthCaps) Range(near, far float32) {
	gl.DepthRange(float64(near), float64(far))
}

type DepthFunc uint32

const (
	DF_NEVER         DepthFunc = gl.NEVER
	DF_LESS          DepthFunc = gl.LESS
	DF_EQUAL         DepthFunc = gl.EQUAL
	DF_LESS_EQUAL    DepthFunc = gl.LEQUAL
	DF_GREATER       DepthFunc = gl.GREATER
	DF_NOT_EQUAL     DepthFunc = gl.NOTEQUAL
	DF_GREATER_EQUAL DepthFunc = gl.GEQUAL
	DF_ALWAYS        DepthFunc = gl.ALWAYS
)
