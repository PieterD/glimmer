package main

import (
	"github.com/PieterD/glimmer/gli"
	"github.com/PieterD/glimmer/win"
	. "github.com/PieterD/pan"
)

func main() {
	// Create window
	window, err := win.New(
		win.Size(800, 600),
		win.Title("Triangle"))
	Panic(err)
	defer window.Destroy()

	// Create shader program
	program, err := gli.NewProgram(vSource, fSource,
		gli.ProgramArbGeometryShader4(gSource, gli.GEOM_IN_TRIANGLES, gli.GEOM_OUT_TRIANGLE_STRIP, 6))
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

	for !window.ShouldClose() {
		clear.Clear()
		draw.Draw(0, 3)
		window.Swap()
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

var gSource = `
#version 110
#extension GL_ARB_geometry_shader4: enable

varying in vec4 theColor[3];
varying out vec4 color;

void main() {
	int i;
	for(i=0; i<gl_VerticesIn; i++) {
		gl_Position = gl_PositionIn[i];
		color = theColor[i];
		EmitVertex();
	}
	EndPrimitive();

	// Mirror triangle along the x axis
	for(i=0; i<gl_VerticesIn; i++) {
		vec4 pos = gl_PositionIn[i];
		pos.x = -pos.x;
		gl_Position = pos;
		color = theColor[i];
		EmitVertex();
	}
	EndPrimitive();
}
`

var fSource = `
#version 110

varying vec4 color;

void main() {
	gl_FragColor = color;
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
