package caps

import "github.com/go-gl/gl/v2.1/gl"

type CullCaps struct {}

var Cull = CullCaps{}

func (_ CullCaps) Enable() {
	gl.Enable(gl.CULL_FACE)
}

func (_ CullCaps) Disable() {
	gl.Disable(gl.CULL_FACE)
}

func (_ CullCaps) Face(front, back bool) {
	if front && back {
		gl.CullFace(gl.FRONT_AND_BACK)
	} else if front {
		gl.CullFace(gl.FRONT)
	} else if back {
		gl.CullFace(gl.BACK)
	}
}

func (_ CullCaps) Clockwise() {
	gl.FrontFace(gl.CW)
}

func (_ CullCaps) CounterClockwise() {
	gl.FrontFace(gl.CCW)
}
