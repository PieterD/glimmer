package main

import (
	"time"

	. "github.com/PieterD/glimmer/examples/shared"
	"github.com/PieterD/glimmer/gli"
	"github.com/PieterD/glimmer/win"
)

var vSource = `
#version 110

attribute vec2 position;
attribute vec3 color;
uniform float alpha;
varying vec4 theColor;

void main() {
	gl_Position = vec4(position, 0.0, 1.0);
	theColor = vec4(color * alpha, 1.0);
}
`

var fSource = `
#version 110

varying vec4 theColor;

void main() {
	gl_FragColor = theColor;
}
`

var vertexData = []float32{
	0.75, 0.75,
	1.0, 0.0, 0.0,
	0.75, -0.75,
	0.0, 1.0, 0.0,
	-0.75, -0.75,
	0.0, 0.0, 1.0,
	-0.75, 0.75,
	1.0, 1.0, 1.0,
}

var indexData = []byte{
	0, 1, 2,
	0, 2, 3,
}

func main() {
	window, err := win.New(
		win.Size(800, 600),
		win.Title("Square"),
		// Tell the window to Poll for events rather than wait
		win.Poll())
	Panic(err)
	defer window.Destroy()

	program, err := gli.NewProgram(vSource, fSource)
	Panic(err)
	defer program.Delete()

	vbo, err := gli.NewBuffer(vertexData)
	Panic(err)
	defer vbo.Delete()

	// Create VBO for element indices
	idx, err := gli.NewBuffer(indexData,
		gli.BufferElementArray())
	Panic(err)
	defer idx.Delete()

	vao, err := gli.NewVAO()
	Panic(err)
	defer vao.Delete()

	vao.Enable(2, vbo, program.Attrib("position"),
		gli.VAOStride(5))
	vao.Enable(3, vbo, program.Attrib("color"),
		gli.VAOOffset(2),
		gli.VAOStride(5))

	// Fetch uniform from program
	alpha := program.Uniform("alpha")

	draw, err := gli.NewDraw(gli.TRIANGLES, program, vao,
		// Set the index buffer
		gli.DrawIndex(idx))
	Panic(err)

	clear, err := gli.NewClear(gli.ClearColor(0, 0, 0, 1))
	Panic(err)

	for !window.ShouldClose() {
		// Set the uniform
		alpha.SetFloat(float32(time.Now().Nanosecond()) / 1000000000)

		clear.Clear()
		draw.Draw(0, 6)
		window.Swap()
	}
}
