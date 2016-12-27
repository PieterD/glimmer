package gli

import (
	"errors"

	"github.com/go-gl/gl/v2.1/gl"
)

type Draw struct {
	mode     DrawMode
	program  *Program
	vao      *VAO
	index    *Buffer
	textures []drawTex
}

type drawTex struct {
	texture *Texture
	unit    int
}

type DrawMode uint32

const (
	POINTS         DrawMode = gl.POINTS
	LINE_STRIP     DrawMode = gl.LINE_STRIP
	LINE_LOOP      DrawMode = gl.LINE_LOOP
	LINES          DrawMode = gl.LINES
	TRIANGLE_STRIP DrawMode = gl.TRIANGLE_STRIP
	TRIANGLE_FAN   DrawMode = gl.TRIANGLE_FAN
	TRIANGLES      DrawMode = gl.TRIANGLES
	PATCHES        DrawMode = gl.PATCHES
)

func NewDraw(mode DrawMode, program *Program, vao *VAO, opts ...DrawOption) (*Draw, error) {
	opt := drawOption{}
	for _, o := range opts {
		o(&opt)
	}

	if opt.index != nil {
		if opt.index.bindpoint != gl.ELEMENT_ARRAY_BUFFER {
			return nil, errors.New("Wrong bindpoint for index buffer: Element buffers must be created with BufferElementArray option")
		}
	}

	return &Draw{
		mode:     mode,
		program:  program,
		vao:      vao,
		index:    opt.index,
		textures: opt.textures,
	}, nil
}

func (draw *Draw) Draw(offset, count int) {
	draw.program.Use()
	draw.vao.Use()
	for _, t := range draw.textures {
		t.texture.Use(t.unit)
	}
	if draw.index == nil {
		gl.DrawArrays(uint32(draw.mode), int32(offset), int32(count))
	} else {
		draw.index.Use()
		gl.DrawElements(uint32(draw.mode), int32(count), draw.index.data.typ, gl.PtrOffset(offset*draw.index.data.siz))
	}
}

type DrawOption func(*drawOption)

type drawOption struct {
	index    *Buffer
	textures []drawTex
}

func DrawIndex(index *Buffer) DrawOption {
	return func(opt *drawOption) {
		opt.index = index
	}
}

func DrawTexture(texture *Texture, unit int) DrawOption {
	return func(opt *drawOption) {
		opt.textures = append(opt.textures, drawTex{
			texture: texture,
			unit:    unit,
		})
	}
}
