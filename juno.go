// main.go

package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"runtime"
)

const (
	width  = 500
	height = 500
)

var (
	vertices = []float32{
		-0.5, 0.5, 0.0,
		0.5, 0.5, 0.0,
		0.5, -0.5, 0.0,
		-0.5, -0.5, 0.0,
	}
	indices = []uint32{
		0, 1, 3,
		1, 2, 3,
	}
)

var (
	vao uint32
	vbo uint32
	ebo uint32

	shaderProgram uint32
)

var (
	projection = mgl32.Perspective(mgl32.DegToRad(45.0), float32(width)/height, 0.1, 10.0)
	camPos     = mgl32.Vec3{0, 0, 3}
	camFront   = mgl32.Vec3{0, 0, -1}
	camUp      = mgl32.Vec3{0, 1, 0}

	deltaTime  float32
	lastFrame  float32
	lastX      float32 = float32(width) / 2.0
	lastY      float32 = float32(height) / 2.0
	yaw        float32 = -90.0
	pitch      float32
	firstMouse = true
)

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()

	initOpenGL()

	for !window.ShouldClose() {
		currentFrame := float32(glfw.GetTime())
		deltaTime = currentFrame - lastFrame
		lastFrame = currentFrame

		processInput(window)

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.UseProgram(shaderProgram)

		view := mgl32.LookAtV(camPos, camPos.Add(camFront), camUp)
		model := mgl32.Ident4()
		transform := projection.Mul4(view).Mul4(model)

		transformLoc := gl.GetUniformLocation(shaderProgram, gl.Str("transform\x00"))
		gl.UniformMatrix4fv(transformLoc, 1, false, &transform[0])

		gl.BindVertexArray(vao)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)

		glfw.PollEvents()
		window.SwapBuffers()
	}
}

func initGlfw() error {
	if err := glfw.Init(); err != nil {
		return err
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	return nil
}
