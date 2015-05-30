// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux darwin
// +build !js

package gl

/*
#ifdef os_linux
#include <GLES2/gl2.h>
#endif
#ifdef os_darwin_arm
#include <OpenGLES/ES2/gl.h>
#endif
#ifdef os_darwin_amd64
#include <OpenGL/gl3.h>
#endif
*/
import "C"

// Enum is equivalent to GLenum, and is normally used with one of the
// constants defined in this package.
type Enum uint32

// Types are defined a structs so that in debug mode they can carry
// extra information, such as a string name. See typesdebug.go.

// Attrib identifies the location of a specific attribute variable.
type Attrib struct {
	Value uint
}

// Program identifies a compiled shader program.
type Program struct {
	Value uint32
}

// Shader identifies a GLSL shader.
type Shader struct {
	Value uint32
}

// Buffer identifies a GL buffer object.
type Buffer struct {
	Value uint32
}

// Framebuffer identifies a GL framebuffer.
type Framebuffer struct {
	Value uint32
}

// A Renderbuffer is a GL object that holds an image in an internal format.
type Renderbuffer struct {
	Value uint32
}

// A Texture identifies a GL texture unit.
type Texture struct {
	Value uint32
}

// Uniform identifies the location of a specific uniform variable.
type Uniform struct {
	Value int32
}

func (v Attrib) c() C.GLuint       { return C.GLuint(v.Value) }
func (v Enum) c() C.GLenum         { return C.GLenum(v) }
func (v Program) c() C.GLuint      { return C.GLuint(v.Value) }
func (v Shader) c() C.GLuint       { return C.GLuint(v.Value) }
func (v Buffer) c() C.GLuint       { return C.GLuint(v.Value) }
func (v Framebuffer) c() C.GLuint  { return C.GLuint(v.Value) }
func (v Renderbuffer) c() C.GLuint { return C.GLuint(v.Value) }
func (v Texture) c() C.GLuint      { return C.GLuint(v.Value) }
func (v Uniform) c() C.GLint       { return C.GLint(v.Value) }

func (v Attrib) Valid() bool       { return v.Value != 0 }
func (v Program) Valid() bool      { return v.Value != 0 }
func (v Shader) Valid() bool       { return v.Value != 0 }
func (v Buffer) Valid() bool       { return v.Value != 0 }
func (v Framebuffer) Valid() bool  { return v.Value != 0 }
func (v Renderbuffer) Valid() bool { return v.Value != 0 }
func (v Texture) Valid() bool      { return v.Value != 0 }
func (v Uniform) Valid() bool      { return v.Value != 0 }
