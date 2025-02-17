package sdl

// #cgo pkg-config: sdl3
// #cgo noescape SDL_PutAudioStreamData
// #cgo noescape SDL_GetAudioStreamData
// #include <SDL3/SDL.h>
import "C"
import "unsafe"

// https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamData
func (stream *AudioStream) PutDataInt16(buf []int16) error {
	if !C.SDL_PutAudioStreamData((*C.SDL_AudioStream)(stream), unsafe.Pointer(unsafe.SliceData(buf)), (C.int)(len(buf)*2)) {
		return getError()
	}
	return nil
}

// https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamData
func (stream *AudioStream) PutDataInt32(buf []int32) error {
	if !C.SDL_PutAudioStreamData((*C.SDL_AudioStream)(stream), unsafe.Pointer(unsafe.SliceData(buf)), (C.int)(len(buf)*4)) {
		return getError()
	}
	return nil
}

// https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamData
func (stream *AudioStream) PutDataFloat(buf []float32) error {
	if !C.SDL_PutAudioStreamData((*C.SDL_AudioStream)(stream), unsafe.Pointer(unsafe.SliceData(buf)), (C.int)(len(buf)*4)) {
		return getError()
	}
	return nil
}

// https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamData
func (stream *AudioStream) GetDataInt16(buf []int16) (int, error) {
	n := C.SDL_GetAudioStreamData((*C.SDL_AudioStream)(stream), unsafe.Pointer(unsafe.SliceData(buf)), (C.int)(len(buf)*2))
	if n == -1 {
		return 0, getError()
	}
	return int(n), nil
}

// https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamData
func (stream *AudioStream) GetDataInt32(buf []int32) (int, error) {
	n := C.SDL_GetAudioStreamData((*C.SDL_AudioStream)(stream), unsafe.Pointer(unsafe.SliceData(buf)), (C.int)(len(buf)*4))
	if n == -1 {
		return 0, getError()
	}
	return int(n), nil
}

// https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamData
func (stream *AudioStream) GetDataFloat(buf []float32) (int, error) {
	n := C.SDL_GetAudioStreamData((*C.SDL_AudioStream)(stream), unsafe.Pointer(unsafe.SliceData(buf)), (C.int)(len(buf)*4))
	if n == -1 {
		return 0, getError()
	}
	return int(n), nil
}
