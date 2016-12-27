package gli

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v2.1/gl"
)

type Extensions struct {
	m map[string]struct{}
}

func GetExtensions() (*Extensions, error) {
	ptr := gl.GetString(gl.EXTENSIONS)
	if ptr == nil {
		panic(fmt.Errorf("GetString(EXTENSIONS) returned nil"))
	}
	extensionString := gl.GoStr(ptr)
	extensionSlice := strings.Split(extensionString, " ")
	extensionMap := make(map[string]struct{}, len(extensionSlice))
	for _, extension := range extensionSlice {
		extensionMap[extension] = struct{}{}
	}
	return &Extensions{
		m: extensionMap,
	}, nil
}

func (extensions *Extensions) Has(ext string) bool {
	_, ok := extensions.m[ext]
	return ok
}
