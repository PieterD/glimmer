package gli

import (
	"image"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
)

func init() {
	runtime.LockOSThread()
}

func Viewport(v image.Rectangle) {
	gl.Viewport(int32(v.Min.X), int32(v.Min.Y), int32(v.Max.X), int32(v.Max.Y))
}
