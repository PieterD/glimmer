package caps

import (
	"errors"

	"github.com/go-gl/gl/v2.1/gl"
)

// Capabilities related to Culling.
type CullCaps struct{}

var Cull = CullCaps{}

// Enable face culling.
func (_ CullCaps) Enable() {
	gl.Enable(gl.CULL_FACE)
}

// Disable face culling.
func (_ CullCaps) Disable() {
	gl.Disable(gl.CULL_FACE)
}

// Set which face to render; front, back, or both.
func (_ CullCaps) Face(front, back bool) {
	if front && back {
		gl.CullFace(gl.FRONT_AND_BACK)
	} else if front {
		gl.CullFace(gl.FRONT)
	} else if back {
		gl.CullFace(gl.BACK)
	} else {
		panic(errors.New("Invalid CullFace setting: no front and no back."))
	}
}

// Set the orientation of front-facing polygons to be clockwise.
func (_ CullCaps) Clockwise() {
	gl.FrontFace(gl.CW)
}

// Set the orientation of front-facing polygons to be counter-clockwise.
func (_ CullCaps) CounterClockwise() {
	gl.FrontFace(gl.CCW)
}
