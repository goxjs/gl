# gl [![Build Status](https://travis-ci.org/goxjs/gl.svg?branch=master)](https://travis-ci.org/goxjs/gl) [![GoDoc](https://godoc.org/github.com/goxjs/gl?status.svg)](https://godoc.org/github.com/goxjs/gl)

Package gl is a Go cross-platform binding for OpenGL, with an OpenGL ES 2-like API.

It currently supports OS X, Linux via OpenGL ES 2 backend, and all modern desktop
and mobile browsers via WebGL 1 backend. Windows support is upcoming.

This is a fork of golang.org/x/mobile/gl/... packages with [CL 8793](https://go-review.googlesource.com/8793)
merged in. This package may change as that CL is reviewed, and hopefully eventually deleted once
the CL is merged and golang.org/x/mobile/gl/... can be used.

Installation
------------

```bash
go get -u github.com/goxjs/gl/...
go get -u -d -tags=js github.com/goxjs/gl/...
```
