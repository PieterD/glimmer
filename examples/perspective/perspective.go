package main

import (
	"fmt"
	"time"

	"github.com/PieterD/glimmer/caps"
	"github.com/PieterD/glimmer/gli"
	"github.com/PieterD/glimmer/mat"
	"github.com/PieterD/glimmer/win"
	. "github.com/PieterD/pan"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	window, err := win.New(
		win.Size(800, 600),
		win.Title("Perspective"),
		win.Poll())
	Panic(err)
	defer window.Destroy()

	program, err := gli.NewProgram(vSource, fSource)
	Panic(err)
	defer program.Delete()

	vbo, err := gli.NewBuffer(vertexData)
	Panic(err)
	defer vbo.Delete()

	vao, err := gli.NewVAO()
	Panic(err)
	defer vao.Delete()

	vao.Enable(3, vbo, program.Attrib("position"),
		gli.VAOStride(6))
	vao.Enable(3, vbo, program.Attrib("color"),
		gli.VAOOffset(3),
		gli.VAOStride(6))

	draw, err := gli.NewDraw(gli.TRIANGLES, program, vao)
	Panic(err)

	clear, err := gli.NewClear(
		gli.ClearColor(0, 0, 0, 1),
		gli.ClearDepth())
	Panic(err)

	perspectiveUniform := program.Uniform("perspectiveMatrix")

	pmat := mat.PerspectiveMatrix(1.0, 3.0, 1.0, 800, 600)
	perspectiveUniform.SetFloat(pmat[:]...)

	offsetUniform := program.Uniform("offset")

	var x, y, xStep, yStep float64
	x = 0.5
	y = 0.5
	window.Glfw().SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		switch action {
		case glfw.Press:
			fmt.Printf("press\n")
			switch key {
			case glfw.KeyLeft:
				xStep = -0.1
			case glfw.KeyRight:
				xStep = 0.1
			case glfw.KeyUp:
				yStep = 0.1
			case glfw.KeyDown:
				yStep = -0.1
			}
		case glfw.Release:
			fmt.Printf("release\n")
			switch key {
			case glfw.KeyLeft, glfw.KeyRight:
				xStep = 0.0
			case glfw.KeyUp, glfw.KeyDown:
				yStep = 0.0
			case glfw.KeyEscape:
				window.Close()
			}
		}
	})

	caps.Cull.Enable()
	caps.Cull.Face(false, true)
	caps.Cull.Clockwise()

	caps.Depth.Enable()
	caps.Depth.Func(caps.DF_LESS_EQUAL)
	//caps.Depth.Range(1.0, 3.0)

	prev := time.Now()
	for !window.ShouldClose() {
		cur := time.Now()
		diff := cur.Sub(prev)
		x += xStep * diff.Seconds()
		y += yStep * diff.Seconds()
		prev = cur
		offsetUniform.SetFloat(float32(x), float32(y))

		clear.Clear()
		draw.Draw(0, 36)
		window.Swap()
	}
}

var vSource = `
#version 110

attribute vec3 position;
attribute vec3 color;
varying vec4 theColor;
uniform vec2 offset;
uniform mat4 perspectiveMatrix;

void main() {
	vec4 cameraPos = vec4(position, 1.0) + vec4(offset, 0.0, 0.0);
	gl_Position = perspectiveMatrix * cameraPos;
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
	0.25, 0.25, -1.25,
	0.0, 0.0, 1.0,
	0.25, -0.25, -1.25,
	0.0, 0.0, 1.0,
	-0.25, 0.25, -1.25,
	0.0, 0.0, 1.0,

	0.25, -0.25, -1.25,
	0.0, 0.0, 1.0,
	-0.25, -0.25, -1.25,
	0.0, 0.0, 1.0,
	-0.25, 0.25, -1.25,
	0.0, 0.0, 1.0,

	0.25, 0.25, -2.75,
	0.8, 0.8, 0.8,
	-0.25, 0.25, -2.75,
	0.8, 0.8, 0.8,
	0.25, -0.25, -2.75,
	0.8, 0.8, 0.8,

	0.25, -0.25, -2.75,
	0.8, 0.8, 0.8,
	-0.25, 0.25, -2.75,
	0.8, 0.8, 0.8,
	-0.25, -0.25, -2.75,
	0.8, 0.8, 0.8,

	-0.25, 0.25, -1.25,
	0.0, 1.0, 0.0,
	-0.25, -0.25, -1.25,
	0.0, 1.0, 0.0,
	-0.25, -0.25, -2.75,
	0.0, 1.0, 0.0,

	-0.25, 0.25, -1.25,
	0.0, 1.0, 0.0,
	-0.25, -0.25, -2.75,
	0.0, 1.0, 0.0,
	-0.25, 0.25, -2.75,
	0.0, 1.0, 0.0,

	0.25, 0.25, -1.25,
	0.5, 0.5, 0.0,
	0.25, -0.25, -2.75,
	0.5, 0.5, 0.0,
	0.25, -0.25, -1.25,
	0.5, 0.5, 0.0,

	0.25, 0.25, -1.25,
	0.5, 0.5, 0.0,
	0.25, 0.25, -2.75,
	0.5, 0.5, 0.0,
	0.25, -0.25, -2.75,
	0.5, 0.5, 0.0,

	0.25, 0.25, -2.75,
	1.0, 0.0, 0.0,
	0.25, 0.25, -1.25,
	1.0, 0.0, 0.0,
	-0.25, 0.25, -1.25,
	1.0, 0.0, 0.0,

	0.25, 0.25, -2.75,
	1.0, 0.0, 0.0,
	-0.25, 0.25, -1.25,
	1.0, 0.0, 0.0,
	-0.25, 0.25, -2.75,
	1.0, 0.0, 0.0,

	0.25, -0.25, -2.75,
	0.0, 1.0, 1.0,
	-0.25, -0.25, -1.25,
	0.0, 1.0, 1.0,
	0.25, -0.25, -1.25,
	0.0, 1.0, 1.0,

	0.25, -0.25, -2.75,
	0.0, 1.0, 1.0,
	-0.25, -0.25, -2.75,
	0.0, 1.0, 1.0,
	-0.25, -0.25, -1.25,
	0.0, 1.0, 1.0,
}
