// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build js,wasm

package gl

import "syscall/js"

type Enum int

type Attrib struct {
	Value int
}

type Program struct {
	js.Value
}

type Shader struct {
	js.Value
}

type Buffer struct {
	js.Value
}

type Framebuffer struct {
	js.Value
}

type Renderbuffer struct {
	js.Value
}

type Texture struct {
	js.Value
}

type Uniform struct {
	js.Value
}

func (v Attrib) Valid() bool       { return v.Value != 0 }
func (v Program) Valid() bool      { return v.Value != js.Null() }
func (v Shader) Valid() bool       { return v.Value != js.Null() }
func (v Buffer) Valid() bool       { return v.Value != js.Null() }
func (v Framebuffer) Valid() bool  { return v.Value != js.Null() }
func (v Renderbuffer) Valid() bool { return v.Value != js.Null() }
func (v Texture) Valid() bool      { return v.Value != js.Null() }
func (v Uniform) Valid() bool      { return v.Value != js.Null() }
