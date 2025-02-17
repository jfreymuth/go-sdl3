package sdl

// #cgo pkg-config: sdl3
// #cgo noescape SDL_PushGPUVertexUniformData
// #cgo nocallback SDL_PushGPUVertexUniformData
// #cgo noescape SDL_PushGPUFragmentUniformData
// #cgo nocallback SDL_PushGPUFragmentUniformData
// #cgo noescape SDL_PushGPUComputeUniformData
// #cgo nocallback SDL_PushGPUComputeUniformData
// #include <SDL3/SDL.h>
import "C"
import "unsafe"

func GPUTransferSize[T any](m *TransferBufferMapping) int {
	var t T
	return m.transferBuffer.size / int(unsafe.Sizeof(t))
}

func GPUTransferPut[T any](m *TransferBufferMapping, offset int, t T) {
	if m.p == nil {
		return
	}
	unsafe.Slice((*T)(m.p), m.transferBuffer.size/int(unsafe.Sizeof(t)))[offset] = t
}

func GPUTransferPutSlice[T any](m *TransferBufferMapping, offset int, t []T) {
	if m.p == nil {
		return
	}
	copy(unsafe.Slice((*T)(m.p), m.transferBuffer.size/int(unsafe.Sizeof(t)))[offset:], t)
}

func GPUTransferGet[T any](m *TransferBufferMapping, offset int) T {
	var t T
	if m.p == nil {
		return t
	}
	return unsafe.Slice((*T)(m.p), m.transferBuffer.size/int(unsafe.Sizeof(t)))[offset]
}

func GPUTransferGetSlice[T any](m *TransferBufferMapping, offset int, t []T) {
	if m.p == nil {
		return
	}
	copy(t, unsafe.Slice((*T)(m.p), m.transferBuffer.size/int(unsafe.Sizeof(t)))[offset:])
}

// Pushes data to a vertex uniform slot on the command buffer.
//
// Subsequent draw calls will use this uniform data.
//
// command_buffer: a command buffer.
//
// slot_index: the vertex uniform slot to push data to.
//
// data: client data to write.
//
// length: the length of the data to write.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PushGPUVertexUniformData
func GPUPushVertexUniform[T any](cmd *GPUCommandBuffer, slotIndex int, t T) {
	C.SDL_PushGPUVertexUniformData((*C.SDL_GPUCommandBuffer)(cmd), (C.Uint32)(slotIndex), unsafe.Pointer(&t), (C.Uint32)(unsafe.Sizeof(t)))
}

// Pushes data to a fragment uniform slot on the command buffer.
//
// Subsequent draw calls will use this uniform data.
//
// command_buffer: a command buffer.
//
// slot_index: the fragment uniform slot to push data to.
//
// data: client data to write.
//
// length: the length of the data to write.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PushGPUFragmentUniformData
func GPUPushFragmentUniform[T any](cmd *GPUCommandBuffer, slotIndex int, t T) {
	C.SDL_PushGPUFragmentUniformData((*C.SDL_GPUCommandBuffer)(cmd), (C.Uint32)(slotIndex), unsafe.Pointer(&t), (C.Uint32)(unsafe.Sizeof(t)))
}

// Pushes data to a uniform slot on the command buffer.
//
// Subsequent draw calls will use this uniform data.
//
// command_buffer: a command buffer.
//
// slot_index: the uniform slot to push data to.
//
// data: client data to write.
//
// length: the length of the data to write.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PushGPUComputeUniformData
func GPUPushComputeUniform[T any](cmd *GPUCommandBuffer, slotIndex int, t T) {
	C.SDL_PushGPUComputeUniformData((*C.SDL_GPUCommandBuffer)(cmd), (C.Uint32)(slotIndex), unsafe.Pointer(&t), (C.Uint32)(unsafe.Sizeof(t)))
}
