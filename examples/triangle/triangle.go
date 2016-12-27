package main

import (
	. "github.com/PieterD/glimmer/examples/shared"
	"github.com/PieterD/glimmer/gli"
	"github.com/PieterD/glimmer/win"
)

var vSource = `
#version 110

attribute vec2 position;
attribute vec3 color;
varying vec4 theColor;

void main() {
	gl_Position = vec4(position, 0.0, 1.0);
	theColor = vec4(color, 1.0);
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
	0.75, -0.75,
	-0.75, -0.75,
	1.0, 0.0, 0.0,
	0.0, 1.0, 0.0,
	0.0, 0.0, 1.0,
}

func main() {
	// Create window
	window, err := win.New(
		win.Size(800, 600),
		win.Title("Triangle"))
	Panic(err)
	defer window.Destroy()

	// Create shader program
	program, err := gli.NewProgram(vSource, fSource)
	Panic(err)
	defer program.Delete()

	// Create Vertex Buffer Object
	vbo, err := gli.NewBuffer(vertexData)
	Panic(err)
	defer vbo.Delete()

	// Create Vertex Array Object
	vao, err := gli.NewVAO()
	Panic(err)
	defer vao.Delete()

	vao.Enable(2, vbo, program.Attrib("position"))
	vao.Enable(3, vbo, program.Attrib("color"),
		gli.VAOOffset(6))

	draw, err := gli.NewDraw(gli.TRIANGLES, program, vao)
	Panic(err)

	gli.ClearColor(0.0, 0.0, 0.0, 1.0)

	for !window.ShouldClose() {
		gli.Clear()
		draw.Draw(0, 3)
		window.Swap()
	}
}
