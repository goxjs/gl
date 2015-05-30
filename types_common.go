// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux darwin
// +build !js

package gl

/*
#cgo darwin,amd64  LDFLAGS: -framework OpenGL
#cgo darwin,arm    LDFLAGS: -framework OpenGLES
#cgo linux         LDFLAGS: -lGLESv2

#cgo darwin,amd64  CFLAGS: -Dos_darwin_amd64
#cgo darwin,arm    CFLAGS: -Dos_darwin_arm
#cgo linux         CFLAGS: -Dos_linux

#ifdef os_linux
#include <GLES2/gl2.h>
#endif
#ifdef os_darwin_arm
#include <OpenGLES/ES2/gl.h>
#endif
#ifdef os_darwin_amd64
#include <OpenGL/gl3.h>
#endif

void blendColor(GLfloat r, GLfloat g, GLfloat b, GLfloat a) { glBlendColor(r, g, b, a); }
void clearColor(GLfloat r, GLfloat g, GLfloat b, GLfloat a) { glClearColor(r, g, b, a); }
void clearDepthf(GLfloat d)                                 { glClearDepthf(d); }
void depthRangef(GLfloat n, GLfloat f)                      { glDepthRangef(n, f); }
void sampleCoverage(GLfloat v, GLboolean invert)            { glSampleCoverage(v, invert); }
*/
import "C"

func glBoolean(b bool) C.GLboolean {
	if b {
		return TRUE
	}
	return FALSE
}

// Desktop OpenGL and the ES 2/3 APIs have a very slight difference
// that is imperceptible to C programmers: some function parameters
// use the type Glclampf and some use GLfloat. These two types are
// equivalent in size and bit layout (both are single-precision
// floats), but it plays havoc with cgo. We adjust the types by using
// C wrappers for the problematic functions.

func blendColor(r, g, b, a float32) {
	C.blendColor(C.GLfloat(r), C.GLfloat(g), C.GLfloat(b), C.GLfloat(a))
}
func clearColor(r, g, b, a float32) {
	C.clearColor(C.GLfloat(r), C.GLfloat(g), C.GLfloat(b), C.GLfloat(a))
}
func clearDepthf(d float32)            { C.clearDepthf(C.GLfloat(d)) }
func depthRangef(n, f float32)         { C.depthRangef(C.GLfloat(n), C.GLfloat(f)) }
func sampleCoverage(v float32, i bool) { C.sampleCoverage(C.GLfloat(v), glBoolean(i)) }
