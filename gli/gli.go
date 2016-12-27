package gli

import (
	"image"
	"runtime"

	"fmt"

	"strconv"

	"github.com/go-gl/gl/v2.1/gl"
)

func init() {
	runtime.LockOSThread()
}

func Viewport(v image.Rectangle) {
	gl.Viewport(int32(v.Min.X), int32(v.Min.Y), int32(v.Max.X), int32(v.Max.Y))
}

func Version() string {
	ptr := gl.GetString(gl.VERSION)
	if ptr == nil {
		panic(fmt.Errorf("GetString(VERSION) returned nil"))
	}
	return gl.GoStr(ptr)
}

func ShaderVersion() string {
	ptr := gl.GetString(gl.SHADING_LANGUAGE_VERSION)
	if ptr == nil {
		panic(fmt.Errorf("GetString(SHADING_LANGUAGE_VERSION) returned nil"))
	}
	return gl.GoStr(ptr)
}

func IsShaderVersionSupported(requested string) bool {
	r, err := strconv.ParseFloat(requested, 64)
	if err != nil {
		return false
	}
	hardware := ShaderVersion()
	h, err := strconv.ParseFloat(hardware, 64)
	if err != nil {
		panic(fmt.Errorf("Opengl returned an unparsable shader version '%s': %v", hardware, err))
	}
	return r <= h
}
