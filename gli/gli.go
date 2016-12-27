package gli

import (
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
)

func init() {
	runtime.LockOSThread()
}

func Clear() {
	//TODO: Allow more stuff.
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func ClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}