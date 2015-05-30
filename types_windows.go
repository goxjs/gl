// +build !js

package gl

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
