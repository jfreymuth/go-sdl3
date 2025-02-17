//go:build armbe || arm64be || m68k || mips64p32 || ppc || s390 || shbe || sparc

package sdl

const (
	AudioS16 AudioFormat = AudioS16BE
	AudioS32 AudioFormat = AudioS32BE
	AudioF32 AudioFormat = AudioF32BE
)

const (
	PixelformatRGBA32 PixelFormat = PixelformatRGBA8888
	PixelformatARGB32 PixelFormat = PixelformatARGB8888
	PixelformatBGRA32 PixelFormat = PixelformatBGRA8888
	PixelformatABGR32 PixelFormat = PixelformatABGR8888
	PixelformatRGBX32 PixelFormat = PixelformatRGBX8888
	PixelformatXRGB32 PixelFormat = PixelformatXRGB8888
	PixelformatBGRX32 PixelFormat = PixelformatBGRX8888
	PixelformatXBGR32 PixelFormat = PixelformatXBGR8888
)
