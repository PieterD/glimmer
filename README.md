# glimmer

[![GoDoc](https://godoc.org/github.com/PieterD/glimmer?status.svg)](https://godoc.org/github.com/PieterD/glimmer)
[![Go Report](https://goreportcard.com/badge/github.com/PieterD/glimmer)](https://goreportcard.com/report/github.com/PieterD/glimmer)

Basic opengl utilities, mostly for personal use.

API Is not set in stone.

## Packages

### [win](win)

The win package deals with GLFW window creation and maintenance.

### [gli](gli)

The gli package contains most of the OpenGL Objects; buffers, VAOs, shaders, etc.

### [caps](caps)

The caps package concerns itself with Capabilities (see [glEnable](https://www.opengl.org/sdk/docs/man2/xhtml/glEnable.xml)), since most of these are quite seperate from OpenGL Objects. 

## Examples, in order of complexity

### [Triangle](examples/triangle)
- Window creation
- Shader compilation
- Buffers and VAOs
- Drawing and Clearing

### [Square](examples/square)
- Polling events
- Element buffer and indexed rendering
- Uniforms
- Capabilities (blending)

### [Texture](examples/texture)
- Textures

### [Perspective](examples/perspective)
- Perspective matrix
- Depth buffer
- Face culling
- Keyboard interactivity
- Rudimentary camera panning
- Close on escape

### [Geometry Shader](examples/geometry_arb4)
- Extensions
- Early, highly portable Geometry shader (GL_ARB_geometry_shader4)