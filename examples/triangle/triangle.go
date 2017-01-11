package main

import (
	"github.com/PieterD/glimmer/gli"
	"github.com/PieterD/glimmer/win"
	. "github.com/PieterD/pan"
)

func main() {
	Panic(win.Start(
		win.Size(800, 600),
		win.Title("Triangle"),
		win.Func(myMain)))
}

func myMain(window *win.Window) {
	// Create shader program
	program, err := gli.NewProgram(vSource, fSource)
	Panic(err)
	defer program.Delete()

	// Create Vertex Buffer Object for pixel data
	vbo, err := gli.NewBuffer(vertexData)
	Panic(err)
	defer vbo.Delete()

	// Create Vertex Array Object
	vao, err := gli.NewVAO()
	Panic(err)
	defer vao.Delete()

	// Set Vertex Array attributes
	vao.Enable(2, vbo, program.Attrib("position"))
	vao.Enable(3, vbo, program.Attrib("color"),
		gli.VAOOffset(6))

	// Configure a drawing method
	draw, err := gli.NewDraw(gli.TRIANGLES, program, vao)
	Panic(err)

	// Configure a clearing method
	clear, err := gli.NewClear(gli.ClearColor(0, 0, 0, 1))
	Panic(err)

	for {
		ie := window.Poll()
		if ie == nil {
			clear.Clear()
			draw.Draw(0, 3)
			window.Swap()
			continue
		}
		switch ie.(type) {
		case win.EventClose:
			return
		}
	}
}

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
