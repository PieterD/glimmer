package gli

import (
	"errors"
	"strings"

	"github.com/go-gl/gl/v2.1/gl"
)

type Extensions struct {
	m map[string]struct{}
}

var extensions *Extensions

func GetExtensions() *Extensions {
	if extensions != nil {
		return extensions
	}
	ptr := gl.GetString(gl.EXTENSIONS)
	if ptr == nil {
		panic(errors.New("GetString(EXTENSIONS) returned nil"))
	}
	extensionString := gl.GoStr(ptr)
	extensionSlice := strings.Split(extensionString, " ")
	extensionMap := make(map[string]struct{}, len(extensionSlice))
	for _, extension := range extensionSlice {
		extensionMap[extension] = struct{}{}
	}
	extensions = &Extensions{
		m: extensionMap,
	}
	return extensions
}

func (extensions *Extensions) Has(ext string) bool {
	_, ok := extensions.m[ext]
	return ok
}
