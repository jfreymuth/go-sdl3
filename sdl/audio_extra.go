package sdl

// #cgo pkg-config: sdl3
// #cgo noescape SDL_PutAudioStreamData
// #cgo noescape SDL_GetAudioStreamData
// #include <SDL3/SDL.h>
import "C"
import "unsafe"

// PutDataInt16 is a variant of [AudioStream.PutData] that accepts an int16
// slice. This should only be used if the stream's format is [AudioS16]
//
// https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamData
func (stream *AudioStream) PutDataInt16(buf []int16) error {
	if !C.SDL_PutAudioStreamData((*C.SDL_AudioStream)(stream), unsafe.Pointer(unsafe.SliceData(buf)), (C.int)(len(buf)*2)) {
		return getError()
	}
	return nil
}

// PutDataInt32 is a variant of [AudioStream.PutData] that accepts an int32
// slice. This should only be used if the stream's format is [AudioS32]
//
// https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamData
func (stream *AudioStream) PutDataInt32(buf []int32) error {
	if !C.SDL_PutAudioStreamData((*C.SDL_AudioStream)(stream), unsafe.Pointer(unsafe.SliceData(buf)), (C.int)(len(buf)*4)) {
		return getError()
	}
	return nil
}

// PutDataFloat is a variant of [AudioStream.PutData] that accepts a float32
// slice. This should only be used if the stream's format is [AudioF32]
//
// https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamData
func (stream *AudioStream) PutDataFloat(buf []float32) error {
	if !C.SDL_PutAudioStreamData((*C.SDL_AudioStream)(stream), unsafe.Pointer(unsafe.SliceData(buf)), (C.int)(len(buf)*4)) {
		return getError()
	}
	return nil
}

// GetDataInt16 is a variant of [AudioStream.GetData] that accepts an int16
// slice. This should only be used if the stream's format is [AudioS16]
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamData
func (stream *AudioStream) GetDataInt16(buf []int16) (int, error) {
	n := C.SDL_GetAudioStreamData((*C.SDL_AudioStream)(stream), unsafe.Pointer(unsafe.SliceData(buf)), (C.int)(len(buf)*2))
	if n == -1 {
		return 0, getError()
	}
	return int(n) / 2, nil
}

// GetDataInt32 is a variant of [AudioStream.GetData] that accepts an int32
// slice. This should only be used if the stream's format is [AudioS32]
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamData
func (stream *AudioStream) GetDataInt32(buf []int32) (int, error) {
	n := C.SDL_GetAudioStreamData((*C.SDL_AudioStream)(stream), unsafe.Pointer(unsafe.SliceData(buf)), (C.int)(len(buf)*4))
	if n == -1 {
		return 0, getError()
	}
	return int(n) / 4, nil
}

// GetDataFloat is a variant of [AudioStream.GetData] that accepts a float32
// slice. This should only be used if the stream's format is [AudioF32]
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamData
func (stream *AudioStream) GetDataFloat(buf []float32) (int, error) {
	n := C.SDL_GetAudioStreamData((*C.SDL_AudioStream)(stream), unsafe.Pointer(unsafe.SliceData(buf)), (C.int)(len(buf)*4))
	if n == -1 {
		return 0, getError()
	}
	return int(n) / 4, nil
}
