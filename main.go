package main

import (
    "fmt"
    "github.com/disintegration/imaging"
    "github.com/go-gl/gl/v4.1-core/gl"
    "github.com/go-gl/glfw/v3.3/glfw"
    "github.com/go-gl/mathgl/mgl32"
    "image"
    "image/draw"
    "image/jpeg"
    "log"
    "os"
    "runtime"
    "strings"
)

func chk(err error) {
    if err != nil {
        panic(err)
    }
}

const (
    width          = 500
    height         = 500
    SIZEOF_FLOAT32 = 4
    SIZEOF_UINT32
)



func (v *Vertex_Buffer) create_vao()  uint32 {
    var vbo, vao, ebo uint32
    gl.GenVertexArrays(1, &vao)
    gl.GenBuffers(1, &vbo)

    gl.BindVertexArray(vao)
    gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

    gl.BufferData(gl.ARRAY_BUFFER, SIZEOF_FLOAT32*len(v.buffer), gl.Ptr(v.buffer), gl.STATIC_DRAW)

    var offset = 0
    stride := v.get_stride()
    for i  := uint32(0); i < uint32(len(v.attributes)); i++ {
        gl.EnableVertexAttribArray(i)
        gl.VertexAttribPointerWithOffset(i, int32(v.attributes[i].num_floats), gl.FLOAT,
            false, stride, uintptr(offset))
        offset += v.attributes[i].get_size()
    }

    if len(v.indexes) > 0{
        gl.GenBuffers(1, &ebo)
        gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
        gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, SIZEOF_UINT32 * len(v.indexes), gl.Ptr(v.indexes), gl.STATIC_DRAW)
    }
    return vao
}

func (v *Vertex_Buffer) add_indexes(idxs []uint32){
    v.indexes = append(v.indexes, idxs...)
}

func build_vertex_buffer() Vertex_Buffer{
    ret := new_vertex_buffer()
    ret.add_attribute(Vertex_Attribute{
        name:       "pos",
        num_floats: 3,
    })

    ret.add_attribute(Vertex_Attribute{
        name:       "texture_pos",
        num_floats: 2,
    })

    ret.add_vertex([]float32{-0.5, -0.5, -0.5, 0.0, 0.0})
    ret.add_vertex([]float32{0.5, -0.5, -0.5, 1.0, 0.0})
    ret.add_vertex([]float32{0.5, 0.5, -0.5, 1.0, 1.0})
    ret.add_vertex([]float32{0.5, 0.5, -0.5, 1.0, 1.0})
    ret.add_vertex([]float32{-0.5, 0.5, -0.5, 0.0, 1.0})
    ret.add_vertex([]float32{-0.5, -0.5, -0.5, 0.0, 0.0})

    ret.add_vertex([]float32{-0.5, -0.5, 0.5, 0.0, 0.0})
    ret.add_vertex([]float32{0.5, -0.5, 0.5, 1.0, 0.0})
    ret.add_vertex([]float32{0.5, 0.5, 0.5, 1.0, 1.0})
    ret.add_vertex([]float32{0.5, 0.5, 0.5, 1.0, 1.0})
    ret.add_vertex([]float32{-0.5, 0.5, 0.5, 0.0, 1.0})
    ret.add_vertex([]float32{-0.5, -0.5, 0.5, 0.0, 0.0})

    ret.add_vertex([]float32{-0.5, 0.5, 0.5, 1.0, 0.0})
    ret.add_vertex([]float32{-0.5, 0.5, -0.5, 1.0, 1.0})
    ret.add_vertex([]float32{-0.5, -0.5, -0.5, 0.0, 1.0})
    ret.add_vertex([]float32{-0.5, -0.5, -0.5, 0.0, 1.0})
    ret.add_vertex([]float32{-0.5, -0.5, 0.5, 0.0, 0.0})
    ret.add_vertex([]float32{-0.5, 0.5, 0.5, 1.0, 0.0})

    ret.add_vertex([]float32{0.5, 0.5, 0.5, 1.0, 0.0})
    ret.add_vertex([]float32{0.5, 0.5, -0.5, 1.0, 1.0})
    ret.add_vertex([]float32{0.5, -0.5, -0.5, 0.0, 1.0})
    ret.add_vertex([]float32{0.5, -0.5, -0.5, 0.0, 1.0})
    ret.add_vertex([]float32{0.5, -0.5, 0.5, 0.0, 0.0})
    ret.add_vertex([]float32{0.5, 0.5, 0.5, 1.0, 0.0})

    ret.add_vertex([]float32{-0.5, -0.5, -0.5, 0.0, 1.0})
    ret.add_vertex([]float32{0.5, -0.5, -0.5, 1.0, 1.0})
    ret.add_vertex([]float32{0.5, -0.5, 0.5, 1.0, 0.0})
    ret.add_vertex([]float32{0.5, -0.5, 0.5, 1.0, 0.0})
    ret.add_vertex([]float32{-0.5, -0.5, 0.5, 0.0, 0.0})
    ret.add_vertex([]float32{-0.5, -0.5, -0.5, 0.0, 1.0})

    ret.add_vertex([]float32{-0.5, 0.5, -0.5, 0.0, 1.0})
    ret.add_vertex([]float32{0.5, 0.5, -0.5, 1.0, 1.0})
    ret.add_vertex([]float32{0.5, 0.5, 0.5, 1.0, 0.0})
    ret.add_vertex([]float32{0.5, 0.5, 0.5, 1.0, 0.0})
    ret.add_vertex([]float32{-0.5, 0.5, 0.5, 0.0, 0.0})
    ret.add_vertex([]float32{-0.5, 0.5, -0.5, 0.0, 1.0})
    return ret
}


func main() {
    runtime.LockOSThread()

    window := initGlfw()
    defer glfw.Terminate()
    program := initOpenGL()

    vbuf := build_vertex_buffer()
    vao := vbuf.create_vao()
    gl.UseProgram(program)
    gl.Enable(gl.DEPTH_TEST)
    gl.DepthFunc(gl.LESS)

    _, err := newTexture0("./IMG_4401.jpg")
    chk(err)
    gl.Uniform1i(gl.GetUniformLocation(program, gl.Str("tex\x00")), 0)

    model := mgl32.Ident4()
    //model = mgl32.HomogRotate3DX(mgl32.DegToRad(-40.0))
    view := mgl32.Ident4()
    view = view.Mul4(mgl32.Translate3D(0,0, -1))
    projection := mgl32.Perspective(mgl32.RadToDeg(45), 800.0 / 600.0, 0, 100.0)



    uniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
    gl.UniformMatrix4fv(uniform, 1, false, &model[0])

    uniform = gl.GetUniformLocation(program, gl.Str("view\x00"))
    gl.UniformMatrix4fv(uniform, 1, false, &view[0])

    uniform = gl.GetUniformLocation(program, gl.Str("projection\x00"))
    gl.UniformMatrix4fv(uniform, 1, false, &projection[0])



    for !window.ShouldClose() {
        err := gl.GetError()
        if err > 0{
            fmt.Println(err)
        }

        gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
        gl.ClearColor(0.2, 0.3, 0.3, 1.0)

        model = model.Mul4(mgl32.HomogRotate3DX(mgl32.DegToRad(1)))
        uniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
        gl.UniformMatrix4fv(uniform, 1, false, &model[0])

        gl.UseProgram(program)
        gl.BindVertexArray(vao)
        gl.DrawArrays(gl.TRIANGLES, 0, vbuf.get_len())
        //gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.Ptr(nil))

        window.SwapBuffers()
        glfw.PollEvents()
    }
}


func compileShader(source string, shaderType uint32) (uint32, error) {
    shader := gl.CreateShader(shaderType)

    csources, free := gl.Strs(source)
    gl.ShaderSource(shader, 1, csources, nil)
    free()
    gl.CompileShader(shader)

    var status int32
    gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
    if status == gl.FALSE {
        var logLength int32
        gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

        log := strings.Repeat("\x00", int(logLength+1))
        gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

        return 0, fmt.Errorf("failed to compile %v: %v", source, log)
    }
    return shader, nil
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw() *glfw.Window {
    if err := glfw.Init(); err != nil {
        panic(err)
    }
    glfw.WindowHint(glfw.Resizable, glfw.False)
    glfw.WindowHint(glfw.ContextVersionMajor, 3)
    glfw.WindowHint(glfw.ContextVersionMinor, 3)
    glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
    glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

    window, err := glfw.CreateWindow(width, height, "dice roller", nil, nil)
    chk(err)
    window.MakeContextCurrent()

    return window
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() uint32 {
    if err := gl.Init(); err != nil {
        panic(err)
    }
    version := gl.GoStr(gl.GetString(gl.VERSION))
    log.Println("OpenGL version", version)

    vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
    chk(err)

    fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
    chk(err)

    prog := gl.CreateProgram()
    gl.AttachShader(prog, vertexShader)
    gl.AttachShader(prog, fragmentShader)
    gl.LinkProgram(prog)
    return prog
}

func newTexture0(file string) (uint32, error) {
    imgFile, err := os.Open(file)
    if err != nil {
        return 0, fmt.Errorf("texture %q not found on disk: %v", file, err)
    }
    img,  err := jpeg.Decode(imgFile)
    chk(err)
    img = imaging.Rotate180(img)
    rgba := image.NewRGBA(img.Bounds())
    if rgba.Stride != rgba.Rect.Size().X*4 {
        return 0, fmt.Errorf("unsupported stride")
    }
    draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)


    var texture uint32
    gl.GenTextures(1, &texture)
    gl.ActiveTexture(gl.TEXTURE0)
    gl.BindTexture(gl.TEXTURE_2D, texture)
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
    gl.TexImage2D(
        gl.TEXTURE_2D,
        0,
        gl.RGBA,
        int32(rgba.Rect.Size().X),
        int32(rgba.Rect.Size().Y),
        0,
        gl.RGBA,
        gl.UNSIGNED_BYTE,
        gl.Ptr(rgba.Pix))

    return texture, nil
}