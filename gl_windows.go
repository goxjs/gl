// +build !js

package gl

// TODO: Complete implementation of all remaining functions, search for `panic("*: not yet implemented")`.

import (
	"log"
	"strings"
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
)

var ContextWatcher = new(contextWatcher)

type contextWatcher struct {
	initGL bool
}

func (cw *contextWatcher) OnMakeCurrent(context interface{}) {
	if !cw.initGL {
		// Initialise gl bindings using the current context.
		err := gl.Init()
		if err != nil {
			log.Fatalln("gl.Init:", err)
		}
		cw.initGL = true
	}
}
func (contextWatcher) OnDetach() {}

func AttachShader(p Program, s Shader) {
	gl.AttachShader(p.Value, s.Value)
}

func BindAttribLocation(p Program, a Attrib, name string) {
	gl.BindAttribLocation(p.Value, uint32(a.Value), gl.Str(name+"\x00"))
}

func BindBuffer(target Enum, b Buffer) {
	gl.BindBuffer(uint32(target), b.Value)
}

func BindFramebuffer(target Enum, fb Framebuffer) {
	gl.BindFramebuffer(uint32(target), fb.Value)
}

func BindRenderbuffer(target Enum, rb Renderbuffer) {
	gl.BindRenderbuffer(uint32(target), rb.Value)
}

func ActiveTexture(texture Enum) {
	gl.ActiveTexture(uint32(texture))
}

func BindTexture(target Enum, t Texture) {
	gl.BindTexture(uint32(target), t.Value)
}

func BlendColor(red, green, blue, alpha float32) {
	gl.BlendColor(red, green, blue, alpha)
}

func BlendEquation(mode Enum) {
	gl.BlendEquation(uint32(mode))
}

func BlendEquationSeparate(modeRGB, modeAlpha Enum) {
	gl.BlendEquationSeparate(uint32(modeRGB), uint32(modeAlpha))
}

func BlendFunc(sfactor, dfactor Enum) {
	gl.BlendFunc(uint32(sfactor), uint32(dfactor))
}

func BlendFuncSeparate(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha Enum) {
	gl.BlendFuncSeparate(uint32(sfactorRGB), uint32(dfactorRGB), uint32(sfactorAlpha), uint32(dfactorAlpha))
}

func BufferData(target Enum, src []byte, usage Enum) {
	gl.BufferData(uint32(target), int(len(src)), gl.Ptr(&src[0]), uint32(usage))
}

func BufferInit(target Enum, size int, usage Enum) {
	gl.BufferData(uint32(target), size, nil, uint32(usage))
}

func BufferSubData(target Enum, offset int, data []byte) {
	gl.BufferSubData(uint32(target), offset, int(len(data)), gl.Ptr(&data[0]))
}

func CheckFramebufferStatus(target Enum) Enum {
	return Enum(gl.CheckFramebufferStatus(uint32(target)))
}

func Clear(mask Enum) {
	gl.Clear(uint32(mask))
}

func ClearColor(red, green, blue, alpha float32) {
	gl.ClearColor(red, green, blue, alpha)
}

func ClearDepthf(d float32) {
	gl.ClearDepthf(d)
}

func ClearStencil(s int) {
	gl.ClearStencil(int32(s))
}

func ColorMask(red, green, blue, alpha bool) {
	gl.ColorMask(red, green, blue, alpha)
}

func CompileShader(s Shader) {
	gl.CompileShader(s.Value)
}

func CompressedTexImage2D(target Enum, level int, internalformat Enum, width, height, border int, data []byte) {
	gl.CompressedTexImage2D(uint32(target), int32(level), uint32(internalformat), int32(width), int32(height), int32(border), int32(len(data)), gl.Ptr(data))
}

func CompressedTexSubImage2D(target Enum, level, xoffset, yoffset, width, height int, format Enum, data []byte) {
	panic("CompressedTexSubImage2D: not yet implemented")
}

func CopyTexImage2D(target Enum, level int, internalformat Enum, x, y, width, height, border int) {
	panic("CopyTexImage2D: not yet implemented")
}

func CopyTexSubImage2D(target Enum, level, xoffset, yoffset, x, y, width, height int) {
	panic("CopyTexSubImage2D: not yet implemented")
}

func CreateBuffer() Buffer {
	var b Buffer
	gl.GenBuffers(1, &b.Value)
	return b
}

func CreateFramebuffer() Framebuffer {
	var b Framebuffer
	gl.GenFramebuffers(1, &b.Value)
	return b
}

func CreateProgram() Program {
	return Program{Value: uint32(gl.CreateProgram())}
}

func CreateRenderbuffer() Renderbuffer {
	var b Renderbuffer
	gl.GenRenderbuffers(1, &b.Value)
	return b
}

func CreateShader(ty Enum) Shader {
	return Shader{Value: uint32(gl.CreateShader(uint32(ty)))}
}

func CreateTexture() Texture {
	var t Texture
	gl.GenTextures(1, &t.Value)
	return t
}

func CullFace(mode Enum) {
	gl.CullFace(uint32(mode))
}

func DeleteBuffer(v Buffer) {
	gl.DeleteBuffers(1, &v.Value)
}

func DeleteFramebuffer(v Framebuffer) {
	gl.DeleteFramebuffers(1, &v.Value)
}

func DeleteProgram(p Program) {
	gl.DeleteProgram(p.Value)
}

func DeleteRenderbuffer(v Renderbuffer) {
	gl.DeleteRenderbuffers(1, &v.Value)
}

func DeleteShader(s Shader) {
	gl.DeleteShader(s.Value)
}

func DeleteTexture(v Texture) {
	gl.DeleteTextures(1, &v.Value)
}

func DepthFunc(fn Enum) {
	gl.DepthFunc(uint32(fn))
}

func DepthMask(flag bool) {
	gl.DepthMask(flag)
}

func DepthRangef(n, f float32) {
	gl.DepthRangef(n, f)
}

func DetachShader(p Program, s Shader) {
	gl.DetachShader(p.Value, s.Value)
}

func Disable(cap Enum) {
	gl.Disable(uint32(cap))
}

func DisableVertexAttribArray(a Attrib) {
	gl.DisableVertexAttribArray(uint32(a.Value))
}

func DrawArrays(mode Enum, first, count int) {
	gl.DrawArrays(uint32(mode), int32(first), int32(count))
}

func DrawElements(mode Enum, count int, ty Enum, offset int) {
	gl.DrawElements(uint32(mode), int32(count), uint32(ty), gl.PtrOffset(offset))
}

func Enable(cap Enum) {
	gl.Enable(uint32(cap))
}

func EnableVertexAttribArray(a Attrib) {
	gl.EnableVertexAttribArray(uint32(a.Value))
}

func Finish() {
	gl.Finish()
}

func Flush() {
	gl.Flush()
}

func FramebufferRenderbuffer(target, attachment, rbTarget Enum, rb Renderbuffer) {
	gl.FramebufferRenderbuffer(uint32(target), uint32(attachment), uint32(rbTarget), rb.Value)
}

func FramebufferTexture2D(target, attachment, texTarget Enum, t Texture, level int) {
	panic("FramebufferTexture2D: not yet implemented")
}

func FrontFace(mode Enum) {
	gl.FrontFace(uint32(mode))
}

func GenerateMipmap(target Enum) {
	gl.GenerateMipmap(uint32(target))
}

func GetActiveAttrib(p Program, index uint32) (name string, size int, ty Enum) {
	var length, si int32
	var typ uint32
	name = strings.Repeat("\x00", 256)
	cname := gl.Str(name)
	gl.GetActiveAttrib(p.Value, uint32(index), int32(len(name)-1), &length, &si, &typ, cname)
	name = name[:strings.IndexRune(name, 0)]
	return name, int(si), Enum(typ)
}

func GetActiveUniform(p Program, index uint32) (name string, size int, ty Enum) {
	var length, si int32
	var typ uint32
	name = strings.Repeat("\x00", 256)
	cname := gl.Str(name)
	gl.GetActiveUniform(p.Value, uint32(index), int32(len(name)-1), &length, &si, &typ, cname)
	name = name[:strings.IndexRune(name, 0)]
	return name, int(si), Enum(typ)
}

func GetAttachedShaders(p Program) []Shader {
	panic("GetAttachedShaders: not yet implemented")
}

func GetAttribLocation(p Program, name string) Attrib {
	return Attrib{Value: uint(gl.GetAttribLocation(p.Value, gl.Str(name+"\x00")))}
}

func GetBooleanv(dst []bool, pname Enum) {
	panic("GetBooleanv: not yet implemented")
}

func GetFloatv(dst []float32, pname Enum) {
	panic("GetFloatv: not yet implemented")
}

func GetIntegerv(pname Enum, data []int32) {
	panic("GetIntegerv: not yet implemented")
}

func GetInteger(pname Enum) int {
	var data int32
	gl.GetIntegerv(uint32(pname), &data)
	return int(data)
}

func GetBufferParameteri(target, pname Enum) int {
	var params int32
	gl.GetBufferParameteriv(uint32(target), uint32(pname), &params)
	return int(params)
}

func GetError() Enum {
	return Enum(gl.GetError())
}

func GetFramebufferAttachmentParameteri(target, attachment, pname Enum) int {
	panic("GetFramebufferAttachmentParameteri: not yet implemented")
}

func GetProgrami(p Program, pname Enum) int {
	var result int32
	gl.GetProgramiv(p.Value, uint32(pname), &result)
	return int(result)
}

func GetProgramInfoLog(p Program) string {
	var logLength int32
	gl.GetProgramiv(p.Value, gl.INFO_LOG_LENGTH, &logLength)
	if logLength == 0 {
		return ""
	}

	logBuffer := make([]uint8, logLength)
	gl.GetProgramInfoLog(p.Value, logLength, nil, &logBuffer[0])
	return gl.GoStr(&logBuffer[0])
}

func GetRenderbufferParameteri(target, pname Enum) int {
	panic("GetRenderbufferParameteri: not yet implemented")
}

func GetShaderi(s Shader, pname Enum) int {
	var result int32
	gl.GetShaderiv(s.Value, uint32(pname), &result)
	return int(result)
}

func GetShaderInfoLog(s Shader) string {
	var logLength int32
	gl.GetShaderiv(s.Value, gl.INFO_LOG_LENGTH, &logLength)
	if logLength == 0 {
		return ""
	}

	logBuffer := make([]uint8, logLength)
	gl.GetShaderInfoLog(s.Value, logLength, nil, &logBuffer[0])
	return gl.GoStr(&logBuffer[0])
}

func GetShaderPrecisionFormat(shadertype, precisiontype Enum) (rangeLow, rangeHigh, precision int) {
	panic("GetShaderPrecisionFormat: not yet implemented")
}

func GetShaderSource(s Shader) string {
	panic("GetShaderSource: not yet implemented")
}

func GetString(pname Enum) string {
	return gl.GoStr(gl.GetString(uint32(pname)))
}

func GetTexParameterfv(dst []float32, target, pname Enum) {
	panic("GetTexParameterfv: not yet implemented")
}

func GetTexParameteriv(dst []int32, target, pname Enum) {
	panic("GetTexParameteriv: not yet implemented")
}

func GetUniformfv(dst []float32, src Uniform, p Program) {
	panic("GetUniformfv: not yet implemented")
}

func GetUniformiv(dst []int32, src Uniform, p Program) {
	panic("GetUniformiv: not yet implemented")
}

func GetUniformLocation(p Program, name string) Uniform {
	return Uniform{Value: gl.GetUniformLocation(p.Value, gl.Str(name+"\x00"))}
}

func GetVertexAttribf(src Attrib, pname Enum) float32 {
	panic("GetVertexAttribf: not yet implemented")
}

func GetVertexAttribfv(dst []float32, src Attrib, pname Enum) {
	panic("GetVertexAttribfv: not yet implemented")
}

func GetVertexAttribi(src Attrib, pname Enum) int32 {
	panic("GetVertexAttribi: not yet implemented")
}

func GetVertexAttribiv(dst []int32, src Attrib, pname Enum) {
	panic("GetVertexAttribiv: not yet implemented")
}

func Hint(target, mode Enum) {
	panic("Hint: not yet implemented")
}

func IsBuffer(b Buffer) bool {
	panic("IsBuffer: not yet implemented")
}

func IsEnabled(cap Enum) bool {
	panic("IsEnabled: not yet implemented")
}

func IsFramebuffer(fb Framebuffer) bool {
	panic("IsFramebuffer: not yet implemented")
}

func IsProgram(p Program) bool {
	panic("IsProgram: not yet implemented")
}

func IsRenderbuffer(rb Renderbuffer) bool {
	panic("IsRenderbuffer: not yet implemented")
}

func IsShader(s Shader) bool {
	panic("IsShader: not yet implemented")
}

func IsTexture(t Texture) bool {
	panic("IsTexture: not yet implemented")
}

func LineWidth(width float32) {
	panic("LineWidth: not yet implemented")
}

func LinkProgram(p Program) {
	gl.LinkProgram(p.Value)
}

func PixelStorei(pname Enum, param int32) {
	panic("PixelStorei: not yet implemented")
}

func PolygonOffset(factor, units float32) {
	gl.PolygonOffset(factor, units)
}

func ReadPixels(dst []byte, x, y, width, height int, format, ty Enum) {
	gl.ReadPixels(int32(x), int32(y), int32(width), int32(height), uint32(format), uint32(ty), gl.Ptr(&dst[0]))
}

func ReleaseShaderCompiler() {
	panic("ReleaseShaderCompiler: not yet implemented")
}

func RenderbufferStorage(target, internalFormat Enum, width, height int) {
	gl.RenderbufferStorage(uint32(target), uint32(internalFormat), int32(width), int32(height))
}

func SampleCoverage(value float32, invert bool) {
	gl.SampleCoverage(value, invert)
}

func Scissor(x, y, width, height int32) {
	gl.Scissor(x, y, width, height)
}

func ShaderSource(s Shader, src string) {
	glsource := gl.Str(src + "\x00")
	gl.ShaderSource(s.Value, 1, &glsource, nil)
}

func StencilFunc(fn Enum, ref int, mask uint32) {
	panic("StencilFunc: not yet implemented")
}

func StencilFuncSeparate(face, fn Enum, ref int, mask uint32) {
	panic("StencilFuncSeparate: not yet implemented")
}

func StencilMask(mask uint32) {
	panic("StencilMask: not yet implemented")
}

func StencilMaskSeparate(face Enum, mask uint32) {
	panic("StencilMaskSeparate: not yet implemented")
}

func StencilOp(fail, zfail, zpass Enum) {
	panic("StencilOp: not yet implemented")
}

func StencilOpSeparate(face, sfail, dpfail, dppass Enum) {
	panic("StencilOpSeparate: not yet implemented")
}

func TexImage2D(target Enum, level int, width, height int, format Enum, ty Enum, data []byte) {
	p := unsafe.Pointer(nil)
	if len(data) > 0 {
		p = gl.Ptr(&data[0])
	}
	gl.TexImage2D(uint32(target), int32(level), int32(format), int32(width), int32(height), 0, uint32(format), uint32(ty), p)
}

func TexSubImage2D(target Enum, level int, x, y, width, height int, format, ty Enum, data []byte) {
	gl.TexSubImage2D(uint32(target), int32(level), int32(x), int32(y), int32(width), int32(height), uint32(format), uint32(ty), gl.Ptr(&data[0]))
}

func TexParameterf(target, pname Enum, param float32) {
	panic("TexParameterf: not yet implemented")
}

func TexParameterfv(target, pname Enum, params []float32) {
	panic("TexParameterfv: not yet implemented")
}

func TexParameteri(target, pname Enum, param int) {
	gl.TexParameteri(uint32(target), uint32(pname), int32(param))
}

func TexParameteriv(target, pname Enum, params []int32) {
	panic("TexParameteriv: not yet implemented")
}

func Uniform1f(dst Uniform, v float32) {
	gl.Uniform1f(dst.Value, v)
}

func Uniform1fv(dst Uniform, src []float32) {
	panic("Uniform1fv: not yet implemented")
}

func Uniform1i(dst Uniform, v int) {
	gl.Uniform1i(dst.Value, int32(v))
}

func Uniform1iv(dst Uniform, src []int32) {
	panic("Uniform1iv: not yet implemented")
}

func Uniform2f(dst Uniform, v0, v1 float32) {
	panic("Uniform2f: not yet implemented")
}

func Uniform2fv(dst Uniform, src []float32) {
	panic("Uniform2fv: not yet implemented")
}

func Uniform2i(dst Uniform, v0, v1 int) {
	panic("Uniform2i: not yet implemented")
}

func Uniform2iv(dst Uniform, src []int32) {
	panic("Uniform2iv: not yet implemented")
}

func Uniform3f(dst Uniform, v0, v1, v2 float32) {
	gl.Uniform3f(dst.Value, v0, v1, v2)
}

func Uniform3fv(dst Uniform, src []float32) {
	panic("Uniform3fv: not yet implemented")
}

func Uniform3i(dst Uniform, v0, v1, v2 int32) {
	panic("Uniform3i: not yet implemented")
}

func Uniform3iv(dst Uniform, src []int32) {
	panic("Uniform3iv: not yet implemented")
}

func Uniform4f(dst Uniform, v0, v1, v2, v3 float32) {
	gl.Uniform4f(dst.Value, v0, v1, v2, v3)
}

func Uniform4fv(dst Uniform, src []float32) {
	gl.Uniform4fv(dst.Value, int32(len(src)/4), &src[0])
}

func Uniform4i(dst Uniform, v0, v1, v2, v3 int32) {
	panic("Uniform4i: not yet implemented")
}

func Uniform4iv(dst Uniform, src []int32) {
	panic("Uniform4iv: not yet implemented")
}

func UniformMatrix2fv(dst Uniform, src []float32) {
	panic("UniformMatrix2fv: not yet implemented")
}

func UniformMatrix3fv(dst Uniform, src []float32) {
	gl.UniformMatrix3fv(dst.Value, int32(len(src)/(3*3)), false, &src[0])
}

func UniformMatrix4fv(dst Uniform, src []float32) {
	gl.UniformMatrix4fv(dst.Value, int32(len(src)/(4*4)), false, &src[0])
}

func UseProgram(p Program) {
	gl.UseProgram(p.Value)
}

func ValidateProgram(p Program) {
	gl.ValidateProgram(p.Value)
}

func VertexAttrib1f(dst Attrib, x float32) {
	panic("VertexAttrib1f: not yet implemented")
}

func VertexAttrib1fv(dst Attrib, src []float32) {
	panic("VertexAttrib1fv: not yet implemented")
}

func VertexAttrib2f(dst Attrib, x, y float32) {
	panic("VertexAttrib2f: not yet implemented")
}

func VertexAttrib2fv(dst Attrib, src []float32) {
	panic("VertexAttrib2fv: not yet implemented")
}

func VertexAttrib3f(dst Attrib, x, y, z float32) {
	panic("VertexAttrib3f: not yet implemented")
}

func VertexAttrib3fv(dst Attrib, src []float32) {
	panic("VertexAttrib3fv: not yet implemented")
}

func VertexAttrib4f(dst Attrib, x, y, z, w float32) {
	panic("VertexAttrib4f: not yet implemented")
}

func VertexAttrib4fv(dst Attrib, src []float32) {
	panic("VertexAttrib4fv: not yet implemented")
}

func VertexAttribPointer(dst Attrib, size int, ty Enum, normalized bool, stride, offset int) {
	gl.VertexAttribPointer(uint32(dst.Value), int32(size), uint32(ty), normalized, int32(stride), gl.PtrOffset(offset))
}

func Viewport(x, y, width, height int) {
	gl.Viewport(int32(x), int32(y), int32(width), int32(height))
}
