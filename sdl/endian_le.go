//go:build 386 || amd64 || amd64p32 || alpha || arm || arm64 || loong64 || mipsle || mips64le || mips64p32le || nios2 || ppc64le || riscv || riscv64 || sh || wasm

package sdl

const (
	AudioS16 AudioFormat = AudioS16LE
	AudioS32 AudioFormat = AudioS32LE
	AudioF32 AudioFormat = AudioF32LE
)

const (
	PixelformatRGBA32 PixelFormat = PixelformatABGR8888
	PixelformatARGB32 PixelFormat = PixelformatBGRA8888
	PixelformatBGRA32 PixelFormat = PixelformatARGB8888
	PixelformatABGR32 PixelFormat = PixelformatRGBA8888
	PixelformatRGBX32 PixelFormat = PixelformatXBGR8888
	PixelformatXRGB32 PixelFormat = PixelformatBGRX8888
	PixelformatBGRX32 PixelFormat = PixelformatXRGB8888
	PixelformatXBGR32 PixelFormat = PixelformatRGBX8888
)
