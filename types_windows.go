// +build !js
// +build !gldebug

package gl

import (
	"fmt"

	"golang.org/x/mobile/f32"
)

type Enum uint32

type Attrib struct {
	Value uint
}

type Program struct {
	Value uint32
}

type Shader struct {
	Value uint32
}

type Buffer struct {
	Value uint32
}

type Framebuffer struct {
	Value uint32
}

type Renderbuffer struct {
	Value uint32
}

type Texture struct {
	Value uint32
}

type Uniform struct {
	Value int32
}

func (v Attrib) Valid() bool       { return v.Value != 0 }
func (v Program) Valid() bool      { return v.Value != 0 }
func (v Shader) Valid() bool       { return v.Value != 0 }
func (v Buffer) Valid() bool       { return v.Value != 0 }
func (v Framebuffer) Valid() bool  { return v.Value != 0 }
func (v Renderbuffer) Valid() bool { return v.Value != 0 }
func (v Texture) Valid() bool      { return v.Value != 0 }
func (v Uniform) Valid() bool      { return v.Value != 0 }

func (v Attrib) String() string       { return fmt.Sprintf("Attrib(%d)", v.Value) }
func (v Program) String() string      { return fmt.Sprintf("Program(%d)", v.Value) }
func (v Shader) String() string       { return fmt.Sprintf("Shader(%d)", v.Value) }
func (v Buffer) String() string       { return fmt.Sprintf("Buffer(%d)", v.Value) }
func (v Framebuffer) String() string  { return fmt.Sprintf("Framebuffer(%d)", v.Value) }
func (v Renderbuffer) String() string { return fmt.Sprintf("Renderbuffer(%d)", v.Value) }
func (v Texture) String() string      { return fmt.Sprintf("Texture(%d)", v.Value) }
func (v Uniform) String() string      { return fmt.Sprintf("Uniform(%d)", v.Value) }

func (u Uniform) WriteAffine(a *f32.Affine) {
	var m [9]float32
	m[0*3+0] = a[0][0]
	m[0*3+1] = a[1][0]
	m[0*3+2] = 0
	m[1*3+0] = a[0][1]
	m[1*3+1] = a[1][1]
	m[1*3+2] = 0
	m[2*3+0] = a[0][2]
	m[2*3+1] = a[1][2]
	m[2*3+2] = 1
	UniformMatrix3fv(u, m[:])
}

func (u Uniform) WriteMat4(p *f32.Mat4) {
	var m [16]float32
	m[0*4+0] = p[0][0]
	m[0*4+1] = p[1][0]
	m[0*4+2] = p[2][0]
	m[0*4+3] = p[3][0]
	m[1*4+0] = p[0][1]
	m[1*4+1] = p[1][1]
	m[1*4+2] = p[2][1]
	m[1*4+3] = p[3][1]
	m[2*4+0] = p[0][2]
	m[2*4+1] = p[1][2]
	m[2*4+2] = p[2][2]
	m[2*4+3] = p[3][2]
	m[3*4+0] = p[0][3]
	m[3*4+1] = p[1][3]
	m[3*4+2] = p[2][3]
	m[3*4+3] = p[3][3]
	UniformMatrix4fv(u, m[:])
}

func (u Uniform) WriteVec4(v *f32.Vec4) {
	Uniform4f(u, v[0], v[1], v[2], v[3])
}
