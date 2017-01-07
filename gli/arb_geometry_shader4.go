package gli

import (
	"errors"

	"github.com/go-gl/gl/v2.1/gl"
)

type programOptionArbGeometryShader4 struct {
	use               bool
	source            string
	inType            GeometryInputType
	outType           GeometryOutputType
	numOutputVertices int
}

type GeometryInputType uint32

const (
	GEOM_IN_POINTS              GeometryInputType = gl.POINTS
	GEOM_IN_LINES               GeometryInputType = gl.LINES
	GEOM_IN_LINES_ADJACENCY     GeometryInputType = gl.LINES_ADJACENCY_ARB
	GEOM_IN_TRIANGLES           GeometryInputType = gl.TRIANGLES
	GEOM_IN_TRIANGLES_ADJACENCY GeometryInputType = gl.TRIANGLES_ADJACENCY_ARB
)

type GeometryOutputType uint32

const (
	GEOM_OUT_POINTS         GeometryOutputType = gl.POINTS
	GEOM_OUT_LINE_STRIP     GeometryOutputType = gl.LINE_STRIP
	GEOM_OUT_TRIANGLE_STRIP GeometryOutputType = gl.TRIANGLE_STRIP
)

func ProgramArbGeometryShader4(source string, inType GeometryInputType, outType GeometryOutputType, numOutputVertices int) ProgramOption {
	return func(opt *programOption) {
		opt.arbGeometryShader4.use = true
		opt.arbGeometryShader4.source = source
		opt.arbGeometryShader4.inType = inType
		opt.arbGeometryShader4.outType = outType
		opt.arbGeometryShader4.numOutputVertices = numOutputVertices
	}
}

func (opt programOptionArbGeometryShader4) enable(id uint32) error {
	if opt.use {
		if !GetExtensions().GL_ARB_geometry_shader4() {
			return errors.New("GL_ARB_geometry_shader4 requested, but extension is not available")
		}
		geomId, err := newShader(opt.source, gl.GEOMETRY_SHADER_ARB)
		if err != nil {
			return err
		}
		defer gl.DeleteShader(geomId)
		gl.AttachShader(id, geomId)
		gl.ProgramParameteriARB(id, gl.GEOMETRY_INPUT_TYPE_ARB, int32(opt.inType))
		gl.ProgramParameteriARB(id, gl.GEOMETRY_OUTPUT_TYPE_ARB, int32(opt.outType))
		gl.ProgramParameteriARB(id, gl.GEOMETRY_VERTICES_OUT_ARB, int32(opt.numOutputVertices))
	}
	return nil
}
