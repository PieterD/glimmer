package main

import (
	_ "image/png"

	"github.com/PieterD/glimmer/gli"
	"github.com/PieterD/glimmer/win"
	. "github.com/PieterD/pan"
)

func main() {
	Panic(win.Start(
		win.Size(800, 600),
		win.Title("Texture"),
		win.Func(myMain)))
}

func myMain(window *win.Window) {
	program, err := gli.NewProgram(vSource, fSource)
	Panic(err)
	defer program.Delete()

	vbo, err := gli.NewBuffer(vertexData)
	Panic(err)
	defer vbo.Delete()

	idx, err := gli.NewBuffer(indexData,
		gli.BufferElementArray())
	Panic(err)
	defer idx.Delete()

	vao, err := gli.NewVAO()
	Panic(err)
	defer vao.Delete()

	vao.Enable(2, vbo, program.Attrib("position"),
		gli.VAOStride(4))
	vao.Enable(3, vbo, program.Attrib("texCoord"),
		gli.VAOOffset(2),
		gli.VAOStride(4))

	// Set sampler uniform to texture unit 3
	texUniform := program.Uniform("tex")
	texUniform.SetSampler(3)

	// Load image and create texture
	img, err := gli.LoadImage("../opengl_logo.png")
	Panic(err)
	tex, err := gli.NewTexture(img,
		gli.TextureFilter(gli.LINEAR, gli.LINEAR),
		gli.TextureWrap(gli.CLAMP_TO_EDGE, gli.CLAMP_TO_EDGE))
	Panic(err)

	draw, err := gli.NewDraw(gli.TRIANGLES, program, vao,
		gli.DrawIndex(idx),
		// Add texture to draw command on texture unit 3
		gli.DrawTexture(tex, 3))
	Panic(err)

	clear, err := gli.NewClear(gli.ClearColor(0, 0, 0, 1))
	Panic(err)

	for {
		ie := window.Poll()
		if ie == nil {
			clear.Clear()
			draw.Draw(0, 6)
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
attribute vec2 texCoord;
varying vec4 theColor;
varying vec2 theTexCoord;

void main() {
	gl_Position = vec4(position, 0.0, 1.0);
	theTexCoord = texCoord;
}
`

var fSource = `
#version 110

varying vec4 theColor;
varying vec2 theTexCoord;
uniform sampler2D tex;

void main() {
	gl_FragColor = texture2D(tex, theTexCoord);
}
`

var vertexData = []float32{
	0.75, 0.75,
	1, 0,
	0.75, -0.75,
	1, 1,
	-0.75, -0.75,
	0, 1,
	-0.75, 0.75,
	0, 0,
}

var indexData = []byte{
	0, 1, 2,
	0, 2, 3,
}
