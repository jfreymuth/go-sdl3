package sdl

// #cgo pkg-config: sdl3
// #include <SDL3/SDL.h>
import "C"
import "unsafe"

func tmpstring(s string) *C.char {
	if s == "" {
		return nil
	}
	p := make([]byte, len(s)+1)
	copy(p, s)
	return (*C.char)(unsafe.Pointer(&p[0]))
}
