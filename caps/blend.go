package caps

import "github.com/go-gl/gl/v2.1/gl"

// Capabilities related to Blending.
type BlendCaps struct{}

var Blend = BlendCaps{}

// Enable blending.
func (_ BlendCaps) Enable() {
	gl.Enable(gl.BLEND)
}

// Disable blending.
func (_ BlendCaps) Disable() {
	gl.Disable(gl.BLEND)
}

// Set the blend function.
func (_ BlendCaps) Func(sourceFactor, destinationFactor BlendFactor) {
	gl.BlendFunc(uint32(sourceFactor), uint32(destinationFactor))
}

// The blend function to pass to Blend.Func.
type BlendFactor uint32

const (
	BF_ZERO                     BlendFactor = gl.ZERO
	BF_ONE                      BlendFactor = gl.ONE
	BF_SRC_COLOR                BlendFactor = gl.SRC_COLOR
	BF_ONE_MINUS_SRC_COLOR      BlendFactor = gl.ONE_MINUS_SRC_COLOR
	BF_DST_COLOR                BlendFactor = gl.DST_COLOR
	BF_ONE_MINUS_DST_COLOR      BlendFactor = gl.ONE_MINUS_DST_COLOR
	BF_SRC_ALPHA                BlendFactor = gl.SRC_ALPHA
	BF_ONE_MINUS_SRC_ALPHA      BlendFactor = gl.ONE_MINUS_SRC_ALPHA
	BF_DST_ALPHA                BlendFactor = gl.DST_ALPHA
	BF_ONE_MINUS_DST_ALPHA      BlendFactor = gl.ONE_MINUS_DST_ALPHA
	BF_CONSTANT_COLOR           BlendFactor = gl.CONSTANT_COLOR
	BF_ONE_MINUS_CONSTANT_COLOR BlendFactor = gl.ONE_MINUS_CONSTANT_COLOR
	BF_CONSTANT_ALPHA           BlendFactor = gl.CONSTANT_ALPHA
	BF_ONE_MINUS_CONSTANT_ALPHA BlendFactor = gl.ONE_MINUS_CONSTANT_ALPHA
	BF_SRC_ALPHA_SATURATE       BlendFactor = gl.SRC_ALPHA_SATURATE
)
