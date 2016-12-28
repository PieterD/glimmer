package caps

import "github.com/go-gl/gl/v2.1/gl"

type BlendCaps struct {}

var Blend = BlendCaps{}

func (_ BlendCaps) Func(sourceFactor, destinationFactor BlendFactor) {
	gl.BlendFunc(uint32(sourceFactor), uint32(destinationFactor))
}

func (_ BlendCaps) Enable() {
	gl.Enable(gl.BLEND)
}

func (_ BlendCaps) Disable() {
	gl.Disable(gl.BLEND)
}

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
