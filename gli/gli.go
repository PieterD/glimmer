package gli

import (
	"runtime"
)

func init() {
	runtime.LockOSThread()
}
