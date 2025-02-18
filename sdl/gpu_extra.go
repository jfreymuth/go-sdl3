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
import (
	"unsafe"
)

// GPUGenericTransferBufferMapping is a generic variant of
// [GPUTransferBufferMapping].
type GPUGenericTransferBufferMapping[T any] struct {
	Position int
	m        *GPUTransferBufferMapping
	data     []T
}

// NewGPUGenericTransferBufferMapping creates a generic interface for an
// existing mapping.
//
// Both the original and the generic mapping can be used at the same time (but
// not concurrently), but Unmap should only be called on one of then.
func NewGPUGenericTransferBufferMapping[T any](m *GPUTransferBufferMapping) *GPUGenericTransferBufferMapping[T] {
	var t T
	return &GPUGenericTransferBufferMapping[T]{
		m:    m,
		data: unsafe.Slice((*T)(m.p), m.transferBuffer.size/int(unsafe.Sizeof(t))),
	}
}

// Mapped reports wether the mapping is valid.
func (m *GPUGenericTransferBufferMapping[T]) Mapped() bool {
	return m.m.Mapped()
}

// TransferBuffer returns the [GPUTransferBuffer] from which this mapping was
// created.
func (m *GPUGenericTransferBufferMapping[T]) TransferBuffer() *GPUTransferBuffer {
	return m.m.TransferBuffer()
}

// Unmaps a previously mapped transfer buffer.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UnmapGPUTransferBuffer
func (m *GPUGenericTransferBufferMapping[T]) Unmap() {
	m.m.Unmap()
}

// Size returns the size of the mapped memory in units of the generic type
// parameter.
func (m *GPUGenericTransferBufferMapping[T]) Size() int {
	return len(m.data)
}

// WriteAt writes data into the mapped memory region at the given offset.
func (m *GPUGenericTransferBufferMapping[T]) WriteAt(offset int, t []T) {
	if m.m.p == nil {
		return
	}
	copy(m.data[offset:], t)
}

// Write writes a slice into the mapped memory region and advances the current
// position.
func (m *GPUGenericTransferBufferMapping[T]) Write(t []T) {
	if m.m.p == nil {
		return
	}
	m.Position += copy(m.data[m.Position:], t)
}

// WriteOne writes a single element into the mapped memory region and advances
// the current position.
func (m *GPUGenericTransferBufferMapping[T]) WriteOne(t T) {
	if m.m.p == nil {
		return
	}
	m.data[m.Position] = t
	m.Position++
}

// ReadAt reads data from the mapped memory region at the given offset.
func (m *GPUGenericTransferBufferMapping[T]) ReadAt(offset int, t []T) {
	if m.m.p == nil {
		return
	}
	copy(t, m.data[offset:])
}

// Read reads a slice from the mapped memory region and advances the current
// position.
func (m *GPUGenericTransferBufferMapping[T]) Read(t []T) {
	if m.m.p == nil {
		return
	}
	m.Position += copy(t, m.data[m.Position:])
}

// ReadOne reads a single element from the mapped memory region and advances
// the current position.
func (m *GPUGenericTransferBufferMapping[T]) ReadOne() T {
	var t T
	if m.m.p == nil {
		return t
	}
	t = m.data[m.Position]
	m.Position++
	return t
}

// Pushes data to a vertex uniform slot on the command buffer.
//
// Subsequent draw calls will use this uniform data.
//
// cmd: a command buffer.
//
// slotIndex: the vertex uniform slot to push data to.
//
// t: client data to write.
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
// cmd: a command buffer.
//
// slotIndex: the vertex uniform slot to push data to.
//
// t: client data to write.
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
// cmd: a command buffer.
//
// slotIndex: the vertex uniform slot to push data to.
//
// t: client data to write.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PushGPUComputeUniformData
func GPUPushComputeUniform[T any](cmd *GPUCommandBuffer, slotIndex int, t T) {
	C.SDL_PushGPUComputeUniformData((*C.SDL_GPUCommandBuffer)(cmd), (C.Uint32)(slotIndex), unsafe.Pointer(&t), (C.Uint32)(unsafe.Sizeof(t)))
}
