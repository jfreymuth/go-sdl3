/*
Portions of this file are copied from the SDL3 headers. The SDL3 license is
reproduced below:

This software is provided 'as-is', without any express or implied
warranty.  In no event will the authors be held liable for any damages
arising from the use of this software.

Permission is granted to anyone to use this software for any purpose,
including commercial applications, and to alter it and redistribute it
freely, subject to the following restrictions:

1. The origin of this software must not be misrepresented; you must not
   claim that you wrote the original software. If you use this software
   in a product, an acknowledgment in the product documentation would be
   appreciated but is not required.
2. Altered source versions must be plainly marked as such, and must not be
   misrepresented as being the original software.
3. This notice may not be removed or altered from any source distribution.
*/

package sdl

// #cgo pkg-config: sdl3
// #cgo noescape SDL_IOFromFile
// #cgo nocallback SDL_IOFromFile
// #cgo nocallback SDL_IOFromMem
// #cgo nocallback SDL_IOFromConstMem
// #cgo noescape SDL_IOFromDynamicMem
// #cgo nocallback SDL_IOFromDynamicMem
// #cgo noescape wrap_SDL_OpenIO
// #cgo nocallback wrap_SDL_OpenIO
// #cgo noescape SDL_CloseIO
// #cgo noescape SDL_GetIOProperties
// #cgo nocallback SDL_GetIOProperties
// #cgo noescape SDL_GetIOStatus
// #cgo nocallback SDL_GetIOStatus
// #cgo noescape SDL_GetIOSize
// #cgo noescape SDL_SeekIO
// #cgo noescape SDL_TellIO
// #cgo noescape SDL_ReadIO
// #cgo noescape SDL_WriteIO
// #cgo noescape SDL_FlushIO
// #cgo noescape SDL_SaveFile_IO
// #cgo noescape SDL_ReadU8
// #cgo noescape SDL_ReadS8
// #cgo noescape SDL_ReadU16LE
// #cgo noescape SDL_ReadS16LE
// #cgo noescape SDL_ReadU16BE
// #cgo noescape SDL_ReadS16BE
// #cgo noescape SDL_ReadU32LE
// #cgo noescape SDL_ReadS32LE
// #cgo noescape SDL_ReadU32BE
// #cgo noescape SDL_ReadS32BE
// #cgo noescape SDL_ReadU64LE
// #cgo noescape SDL_ReadS64LE
// #cgo noescape SDL_ReadU64BE
// #cgo noescape SDL_ReadS64BE
// #cgo noescape SDL_WriteU8
// #cgo noescape SDL_WriteS8
// #cgo noescape SDL_WriteU16LE
// #cgo noescape SDL_WriteS16LE
// #cgo noescape SDL_WriteU16BE
// #cgo noescape SDL_WriteS16BE
// #cgo noescape SDL_WriteU32LE
// #cgo noescape SDL_WriteS32LE
// #cgo noescape SDL_WriteU32BE
// #cgo noescape SDL_WriteS32BE
// #cgo noescape SDL_WriteU64LE
// #cgo noescape SDL_WriteS64LE
// #cgo noescape SDL_WriteU64BE
// #cgo noescape SDL_WriteS64BE
// #include <SDL3/SDL.h>
// extern SDL_IOStream *wrap_SDL_OpenIO(uintptr_t userdata);
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

// # CategoryIOStream
//
// SDL provides an abstract interface for reading and writing data streams. It
// offers implementations for files, memory, etc, and the app can provide
// their own implementations, too.
//
// [IOStream] is not related to the standard C++ iostream class, other than
// both are abstract interfaces to read/write data.

// [IOStream] status, set by a read or write operation.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_IOStatus
type IOStatus uint32

const (
	IOStatusReady     IOStatus = iota // Everything is ready (no errors and not EOF).
	IOStatusError                     // Read or write I/O error
	IOStatusEOF                       // End of file
	IOStatusNotReady                  // Non blocking I/O, not ready
	IOStatusReadonly                  // Tried to write a read-only buffer
	IOStatusWriteonly                 // Tried to read a write-only buffer
)

// https://wiki.libsdl.org/SDL3/SDL_IOStreamInterface
type IOStreamInterface interface {

	// Return the number of bytes in this [IOStream]
	//
	// Returns the total size of the data stream, or -1 on error.
	Size() int64

	// Seek to offset relative to whence, one of stdio's whence values:
	// [io.SeekStart], [io.SeekCurrent], [io.SeekEnd]
	//
	// Returns the final offset in the data stream, or -1 on error.
	Seek(offset int64, whence int) int64

	// Read up to len(p) bytes from the data stream to the area pointed
	// at by p.
	//
	// On an incomplete read, you should return a status other than
	// [IOStatusReady].
	//
	// Returns the number of bytes read
	Read(p []byte) (int, IOStatus)

	// Write exactly len(p) bytes from the area pointed at by p
	// to data stream.
	//
	// On an incomplete write, you should return a status other than
	// [IOStatusReady].
	//
	// Retruns the number of bytes written
	Write(p []byte) (int, IOStatus)

	// If the stream is buffering, make sure the data is written out.
	//
	// On failure, you should return a status other than [IOStatusReady].
	//
	// Returns true if successful or false on write error when flushing data.
	Flush() (IOStatus, bool)

	// Close and free any allocated resources.
	//
	// This does not guarantee file writes will sync to physical media; they
	// can be in the system's file cache, waiting to go to disk.
	//
	// The [IOStream] is still destroyed even if this fails, so clean up anything
	// even if flushing buffers, etc, returns an error.
	//
	// Returns true if successful or false on write error when flushing data.
	Close() bool
}

//export cb_IOStreamSize
func cb_IOStreamSize(userdata uintptr) C.Sint64 {
	h := cgo.Handle(userdata)
	return C.Sint64(h.Value().(IOStreamInterface).Size())
}

//export cb_IOStreamSeek
func cb_IOStreamSeek(userdata uintptr, offset C.Sint64, whence C.SDL_IOWhence) C.Sint64 {
	h := cgo.Handle(userdata)
	return C.Sint64(h.Value().(IOStreamInterface).Seek(int64(offset), int(whence)))
}

//export cb_IOStreamRead
func cb_IOStreamRead(userdata uintptr, ptr unsafe.Pointer, size C.size_t, status *C.SDL_IOStatus) C.size_t {
	h := cgo.Handle(userdata)
	n, s := h.Value().(IOStreamInterface).Read(unsafe.Slice((*byte)(ptr), size))
	*status = C.SDL_IOStatus(s)
	return C.size_t(n)
}

//export cb_IOStreamWrite
func cb_IOStreamWrite(userdata uintptr, ptr unsafe.Pointer, size C.size_t, status *C.SDL_IOStatus) C.size_t {
	h := cgo.Handle(userdata)
	n, s := h.Value().(IOStreamInterface).Write(unsafe.Slice((*byte)(ptr), size))
	*status = C.SDL_IOStatus(s)
	return C.size_t(n)
}

//export cb_IOStreamFlush
func cb_IOStreamFlush(userdata uintptr, status *C.SDL_IOStatus) C.bool {
	h := cgo.Handle(userdata)
	s, ok := h.Value().(IOStreamInterface).Flush()
	*status = C.SDL_IOStatus(s)
	return C.bool(ok)
}

//export cb_IOStreamClose
func cb_IOStreamClose(userdata uintptr) C.bool {
	h := cgo.Handle(userdata)
	result := h.Value().(IOStreamInterface).Close()
	h.Delete()
	return (C.bool)(result)
}

// The read/write operation structure.
//
// This operates as an opaque handle. There are several APIs to create various
// types of I/O streams, or an app can supply an [IOStreamInterface] to
// [OpenIO] to provide their own stream implementation behind this
// struct's abstract interface.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_IOStream
type IOStream C.struct_SDL_IOStream

// \name IOFrom functions
//
// Functions to create [IOStream] structures from various data streams.

// Use this function to create a new [IOStream] structure for reading from
// and/or writing to a named file.
//
// The mode string is treated roughly the same as in a call to the C
// library's fopen(), even if SDL doesn't happen to use fopen() behind the
// scenes.
//
// Available mode strings:
//
//   - "r": Open a file for reading. The file must exist.
//   - "w": Create an empty file for writing. If a file with the same name
//     already exists its content is erased and the file is treated as a new
//     empty file.
//   - "a": Append to a file. Writing operations append data at the end of the
//     file. The file is created if it does not exist.
//   - "r+": Open a file for update both reading and writing. The file must
//     exist.
//   - "w+": Create an empty file for both reading and writing. If a file with
//     the same name already exists its content is erased and the file is
//     treated as a new empty file.
//   - "a+": Open a file for reading and appending. All writing operations are
//     performed at the end of the file, protecting the previous content to be
//     overwritten. You can reposition (fseek, rewind) the internal pointer to
//     anywhere in the file for reading, but writing operations will move it
//     back to the end of file. The file is created if it does not exist.
//
// **NOTE**: In order to open a file as a binary file, a "b" character has to
// be included in the mode string. This additional "b" character can either
// be appended at the end of the string (thus making the following compound
// modes: "rb", "wb", "ab", "r+b", "w+b", "a+b") or be inserted between the
// letter and the "+" sign for the mixed modes ("rb+", "wb+", "ab+").
// Additional characters may follow the sequence, although they should have no
// effect. For example, "t" is sometimes appended to make explicit the file is
// a text file.
//
// This function supports Unicode filenames, but they must be encoded in UTF-8
// format, regardless of the underlying operating system.
//
// In Android, [IOFromFile] can be used to open content:// URIs. As a
// fallback, [IOFromFile] will transparently open a matching filename in
// the app's `assets`.
//
// Closing the [IOStream] will close SDL's internal file handle.
//
// The following properties may be set at creation time by SDL:
//
//   - [PropIOStreamWindowsHandlePointer]: a pointer, that can be cast
//     to a win32 `HANDLE`, that this [IOStream] is using to access the
//     filesystem. If the program isn't running on Windows, or SDL used some
//     other method to access the filesystem, this property will not be set.
//   - [PropIOStreamStdioFilePointer]: a pointer, that can be cast to a
//     stdio `FILE *`, that this [IOStream] is using to access the filesystem.
//     If SDL used some other method to access the filesystem, this property
//     will not be set. PLEASE NOTE that if SDL is using a different C runtime
//     than your app, trying to use this pointer will almost certainly result in
//     a crash! This is mostly a problem on Windows; make sure you build SDL and
//     your app with the same compiler and settings to avoid it.
//   - [PropIOStreamFileDescriptorNumber]: a file descriptor that this
//     [IOStream] is using to access the filesystem.
//   - [PropIOStreamAndroidAAssetPointer]: a pointer, that can be cast
//     to an Android NDK `AAsset *`, that this [IOStream] is using to access
//     the filesystem. If SDL used some other method to access the filesystem,
//     this property will not be set.
//
// file: a UTF-8 string representing the filename to open.
//
// mode: an ASCII string representing the mode to be used for opening
// the file.
//
// Returns a pointer to the [IOStream] structure that is created or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_IOFromFile
func IOFromFile(file string, mode string) (*IOStream, error) {
	stream := (*IOStream)(C.SDL_IOFromFile(tmpstring(file), tmpstring(mode)))
	if stream == nil {
		return nil, getError()
	}
	return stream, nil
}

const PropIOStreamWindowsHandlePointer = "SDL.iostream.windows.handle"
const PropIOStreamStdioFilePointer = "SDL.iostream.stdio.file"
const PropIOStreamFileDescriptorNumber = "SDL.iostream.file_descriptor"
const PropIOStreamAndroidAAssetPointer = "SDL.iostream.android.aasset"

// Use this function to prepare a read-write memory buffer for use with
// [IOStream].
//
// This function sets up an [IOStream] struct based on a memory area of a
// certain size, for both read and write access.
//
// This memory buffer is not copied by the [IOStream]; the pointer you
// provide must remain valid until you close the stream. Closing the stream
// will not free the original buffer.
//
// If you need to make sure the [IOStream] never writes to the memory
// buffer, you should use [IOFromConstMem] with a read-only buffer of
// memory instead.
//
// The following properties will be set at creation time by SDL:
//
//   - [PropIOStreamMemoryPointer]: this will be a pointer to the start of mem.
//   - [PropIOStreamMemorySizeNumber]: this will be len(mem).
//
// mem: a buffer to feed an [IOStream] stream.
//
// Returns a pointer to a new [IOStream] structure or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_IOFromMem
func IOFromMem(mem []byte) (*IOStream, error) {
	stream := (*IOStream)(C.SDL_IOFromMem(unsafe.Pointer(unsafe.SliceData(mem)), (C.size_t)(len(mem))))
	if stream == nil {
		return nil, getError()
	}
	return stream, nil
}

const PropIOStreamMemoryPointer = "SDL.iostream.memory.base"
const PropIOStreamMemorySizeNumber = "SDL.iostream.memory.size"

// Use this function to prepare a read-only memory buffer for use with
// [IOStream].
//
// This function sets up an [IOStream] struct based on a memory area of a
// certain size. It assumes the memory area is not writable.
//
// Attempting to write to this [IOStream] stream will report an error
// without writing to the memory buffer.
//
// This memory buffer is not copied by the [IOStream]; the pointer you
// provide must remain valid until you close the stream. Closing the stream
// will not free the original buffer.
//
// If you need to write to a memory buffer, you should use [IOFromMem]
// with a writable buffer of memory instead.
//
// The following properties will be set at creation time by SDL:
//
//   - [PropIOStreamMemoryPointer]: this will be a pointer to the start of mem.
//   - [PropIOStreamMemorySizeNumber]: this will be len(mem).
//
// mem: a read-only buffer to feed an [IOStream] stream.
//
// Returns a pointer to a new [IOStream] structure or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_IOFromConstMem
func IOFromConstMem(mem []byte) (*IOStream, error) {
	stream := (*IOStream)(C.SDL_IOFromConstMem(unsafe.Pointer(unsafe.SliceData(mem)), (C.size_t)(len(mem))))
	if stream == nil {
		return nil, getError()
	}
	return stream, nil
}

// Use this function to create an [IOStream] that is backed by dynamically
// allocated memory.
//
// This supports the following properties to provide access to the memory and
// control over allocations:
//
//   - [PropIOStreamDynamicMemoryPointer]: a pointer to the internal
//     memory of the stream. This can be set to NULL to transfer ownership of
//     the memory to the application, which should free the memory with
//     SDL_free(). If this is done, the next operation on the stream must be
//     [IOStream.Close].
//   - [PropIOStreamDynamicChunksizeNumber]: memory will be allocated in
//     multiples of this size, defaulting to 1024.
//
// Returns a pointer to a new [IOStream] structure or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_IOFromDynamicMem
func IOFromDynamicMem() (*IOStream, error) {
	stream := (*IOStream)(C.SDL_IOFromDynamicMem())
	if stream == nil {
		return nil, getError()
	}
	return stream, nil
}

const PropIOStreamDynamicMemoryPointer = "SDL.iostream.dynamic.memory"
const PropIOStreamDynamicChunksizeNumber = "SDL.iostream.dynamic.chunksize"

// Create a custom [IOStream].
//
// Applications do not need to use this function unless they are providing
// their own [IOStream] implementation. If you just need an [IOStream] to
// read/write a common data source, you should use the built-in
// implementations in SDL, like [IOFromFile] or [IOFromMem], etc.
//
// iface: the interface that implements this [IOStream].
//
// Returns a pointer to the allocated memory on success or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_OpenIO
func OpenIO(iface IOStreamInterface) (*IOStream, error) {
	h := cgo.NewHandle(iface)
	stream := (*IOStream)(C.wrap_SDL_OpenIO(C.uintptr_t(h)))
	if stream == nil {
		h.Delete()
		return nil, getError()
	}
	return stream, nil
}

// Close and free an allocated [IOStream] structure.
//
// CloseIO closes and cleans up the [IOStream] stream. It releases any
// resources used by the stream and frees the [IOStream] itself. This
// returns nil on success, or an error if the stream failed to flush to its
// output (e.g. to disk).
//
// Note that if this fails to flush the stream for any reason, this function
// reports an error, but the [IOStream] is still invalid once this function
// returns.
//
// This call flushes any buffered writes to the operating system, but there
// are no guarantees that those writes have gone to physical media; they might
// be in the OS's file cache, waiting to go to disk later. If it's absolutely
// crucial that writes go to disk immediately, so they are definitely stored
// even if the power fails before the file cache would have caught up, one
// should call [IOStream.Flush] before closing. Note that flushing takes time and
// makes the system and your app operate less efficiently, so do so sparingly.
//
// stream: [IOStream] structure to close.
//
// Returns nil on success or an error on failure.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CloseIO
func (stream *IOStream) Close() error {
	if !C.SDL_CloseIO((*C.SDL_IOStream)(stream)) {
		return getError()
	}
	return nil
}

// Get the properties associated with an [IOStream].
//
// stream: a pointer to an [IOStream] structure.
//
// Returns a valid property ID or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetIOProperties
func (stream *IOStream) Properties() (PropertiesID, error) {
	props := (PropertiesID)(C.SDL_GetIOProperties((*C.SDL_IOStream)(stream)))
	if props == 0 {
		return 0, getError()
	}
	return props, nil
}

// Query the stream status of an [IOStream].
//
// This information can be useful to decide if a short read or write was due
// to an error, an EOF, or a non-blocking operation that isn't yet ready to
// complete.
//
// An [IOStream]'s status is only expected to change after a [IOStream.Read] or
// [IOStream.Write] call; don't expect it to change if you just call this query
// function in a tight loop.
//
// stream: the [IOStream] to query.
//
// Returns an [IOStatus] enum with the current state.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetIOStatus
func (stream *IOStream) Status() IOStatus {
	return (IOStatus)(C.SDL_GetIOStatus((*C.SDL_IOStream)(stream)))
}

// Use this function to get the size of the data stream in an [IOStream].
//
// stream: the [IOStream] to get the size of the data stream from.
//
// Returns the size of the data stream in the [IOStream] or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetIOSize
func (stream *IOStream) Size() (int64, error) {
	size := (int64)(C.SDL_GetIOSize((*C.SDL_IOStream)(stream)))
	if size < 0 {
		return 0, getError()
	}
	return size, nil
}

// Seek within an [IOStream] data stream.
//
// This function seeks to byte offset, relative to whence.
//
// `whence` may be any of the following values:
//
//   - [io.SeekStart]: seek from the beginning of data
//   - [io.SeekCurrent]: seek relative to current read point
//   - [io.SeekEnd]: seek relative to the end of data
//
// If this stream can not seek, it will return -1.
//
// stream: a pointer to an [IOStream] structure.
//
// offset: an offset in bytes, relative to whence location; can be
// negative.
//
// whence: any of [io.SeekStart], [io.SeekCurrent], [io.SeekEnd].
//
// Returns the final offset in the data stream after the seek or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SeekIO
func (stream *IOStream) Seek(offset int64, whence int) (int64, error) {
	offset = (int64)(C.SDL_SeekIO((*C.SDL_IOStream)(stream), (C.Sint64)(offset), (C.SDL_IOWhence)(whence)))
	if offset < 0 {
		return 0, getError()
	}
	return offset, nil
}

// Determine the current read/write offset in an [IOStream] data stream.
//
// Tell is actually a wrapper function that calls the [IOStream]'s
// Seek method, with an offset of 0 bytes from [io.SeekCurrent], to
// simplify application development.
//
// stream: an [IOStream] data stream object from which to get the
// current offset.
//
// Returns the current offset in the stream, or -1 if the information can not
// be determined.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_TellIO
func (stream *IOStream) Tell() int64 {
	return (int64)(C.SDL_TellIO((*C.SDL_IOStream)(stream)))
}

// Read from a data source.
//
// This function reads up len(ptr) bytes from the data source to the area
// pointed at by ptr. This function may read less bytes than requested.
//
// This function will return zero when the data stream is completely read, and
// [IOStream.Status] will return [IOStatusEOF]. If zero is returned and
// the stream is not at EOF, [IOStream.Status] will return a different error
// value. In both cases, this method will return a non-nil error
//
// stream: a pointer to an [IOStream] structure.
//
// ptr: a pointer to a buffer to read data into.
//
// Returns the number of bytes read or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReadIO
func (stream *IOStream) Read(ptr []byte) (int, error) {
	n := (int)(C.SDL_ReadIO((*C.SDL_IOStream)(stream), unsafe.Pointer(unsafe.SliceData(ptr)), (C.size_t)(len(ptr))))
	if n == 0 {
		return 0, getError()
	}
	return n, nil
}

// Write to an [IOStream] data stream.
//
// This function writes exactly len(ptr) bytes from the area pointed at by ptr
// to the stream. If this fails for any reason, it'll return less than len(ptr)
// to demonstrate how far the write progressed. On success, it returns len(ptr).
//
// On error, this function still attempts to write as much as possible, so it
// might return a positive value less than the requested write size.
//
// The caller can use [IOStream.Status] to determine if the problem is
// recoverable, such as a non-blocking write that can simply be retried later,
// or a fatal error.
//
// stream: a pointer to an [IOStream] structure.
//
// ptr: a pointer to a buffer containing data to write.
//
// Returns the number of bytes written, which will be less than len(p) on
// failure, in which case a non-nil error is returned.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WriteIO
func (stream *IOStream) Write(ptr []byte) (int, error) {
	n := (int)(C.SDL_WriteIO((*C.SDL_IOStream)(stream), unsafe.Pointer(unsafe.SliceData(ptr)), (C.size_t)(len(ptr))))
	if n == 0 {
		return 0, getError()
	}
	return n, nil
}

// Flush any buffered data in the stream.
//
// This function makes sure that any buffered data is written to the stream.
// Normally this isn't necessary but if the stream is a pipe or socket it
// guarantees that any pending data is sent.
//
// stream: [IOStream] structure to flush.
//
// Returns nil on success or an error on failure.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_FlushIO
func (stream *IOStream) Flush() error {
	if !C.SDL_FlushIO((*C.SDL_IOStream)(stream)) {
		return getError()
	}
	return nil
}

// Load all the data from an SDL data stream.
//
// stream: the [IOStream] to read all available data from.
//
// closeio: if true, calls [IOStream.Close] on the stream before returning, even
// in the case of an error.
//
// Returns the data or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_LoadFile_IO
func (stream *IOStream) ReadAll(closeio bool) ([]byte, error) {
	var datasize C.size_t
	ptr := C.SDL_LoadFile_IO((*C.SDL_IOStream)(stream), &datasize, (C.bool)(closeio))
	if ptr == nil {
		return nil, getError()
	}
	data := C.GoBytes(ptr, C.int(datasize))
	C.SDL_free(ptr)
	return data, nil
}

// Save all the data into an SDL data stream.
//
// stream: the [IOStream] to write all data to.
//
// data: the data to be written.
//
// closeio: if true, calls [IOStream.Close] on the stream before returning, even
// in the case of an error.
//
// Returns nil on success or an error on failure.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SaveFile_IO
func (stream *IOStream) WriteAll(data []byte, closeio bool) error {
	if !C.SDL_SaveFile_IO((*C.SDL_IOStream)(stream), unsafe.Pointer(unsafe.SliceData(data)), (C.size_t)(len(data)), (C.bool)(closeio)) {
		return getError()
	}
	return nil
}

// \name Read endian functions
//
// Read an item of the specified endianness and return in native format.

// Use this function to read a byte from an [IOStream].
//
// This function will return zero when the data stream is completely read, and
// [IOStream.Status] will return [IOStatusEOF]. If zero is returned and
// the stream is not at EOF, [IOStream.Status] will return a different error
// value. In both cases, this method will return a non-nil error
//
// stream: the [IOStream] to read from.
//
// Returns the data read or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReadU8
func (stream *IOStream) ReadU8() (byte, error) {
	var value C.Uint8
	if !C.SDL_ReadU8((*C.SDL_IOStream)(stream), &value) {
		return 0, getError()
	}
	return byte(value), nil
}

// Use this function to read a signed byte from an [IOStream].
//
// This function will return zero when the data stream is completely read, and
// [IOStream.Status] will return [IOStatusEOF]. If zero is returned and
// the stream is not at EOF, [IOStream.Status] will return a different error
// value. In both cases, this method will return a non-nil error
//
// stream: the [IOStream] to read from.
//
// Returns the data read or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReadS8
func (stream *IOStream) ReadS8() (int8, error) {
	var value C.Sint8
	if !C.SDL_ReadS8((*C.SDL_IOStream)(stream), &value) {
		return 0, getError()
	}
	return int8(value), nil
}

// Use this function to read 16 bits of little-endian data from an
// [IOStream] and return in native format.
//
// SDL byteswaps the data only if necessary, so the data returned will be in
// the native byte order.
//
// This function will return zero when the data stream is completely read, and
// [IOStream.Status] will return [IOStatusEOF]. If zero is returned and
// the stream is not at EOF, [IOStream.Status] will return a different error
// value. In both cases, this method will return a non-nil error
//
// stream: the stream from which to read data.
//
// Returns the data read or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReadU16LE
func (stream *IOStream) ReadU16LE() (uint16, error) {
	var value C.Uint16
	if !C.SDL_ReadU16LE((*C.SDL_IOStream)(stream), &value) {
		return 0, getError()
	}
	return uint16(value), nil
}

// Use this function to read 16 bits of little-endian data from an
// [IOStream] and return in native format.
//
// SDL byteswaps the data only if necessary, so the data returned will be in
// the native byte order.
//
// This function will return zero when the data stream is completely read, and
// [IOStream.Status] will return [IOStatusEOF]. If zero is returned and
// the stream is not at EOF, [IOStream.Status] will return a different error
// value. In both cases, this method will return a non-nil error
//
// stream: the stream from which to read data.
//
// Returns the data read or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReadS16LE
func (stream *IOStream) ReadS16LE() (int16, error) {
	var value C.Sint16
	if !C.SDL_ReadS16LE((*C.SDL_IOStream)(stream), &value) {
		return 0, getError()
	}
	return int16(value), nil
}

// Use this function to read 16 bits of big-endian data from an [IOStream]
// and return in native format.
//
// SDL byteswaps the data only if necessary, so the data returned will be in
// the native byte order.
//
// This function will return zero when the data stream is completely read, and
// [IOStream.Status] will return [IOStatusEOF]. If zero is returned and
// the stream is not at EOF, [IOStream.Status] will return a different error
// value. In both cases, this method will return a non-nil error
//
// stream: the stream from which to read data.
//
// Returns the data read or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReadU16BE
func (stream *IOStream) ReadU16BE() (uint16, error) {
	var value C.Uint16
	if !C.SDL_ReadU16BE((*C.SDL_IOStream)(stream), &value) {
		return 0, getError()
	}
	return uint16(value), nil
}

// Use this function to read 16 bits of big-endian data from an [IOStream]
// and return in native format.
//
// SDL byteswaps the data only if necessary, so the data returned will be in
// the native byte order.
//
// This function will return zero when the data stream is completely read, and
// [IOStream.Status] will return [IOStatusEOF]. If zero is returned and
// the stream is not at EOF, [IOStream.Status] will return a different error
// value. In both cases, this method will return a non-nil error
//
// stream: the stream from which to read data.
//
// Returns the data read or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReadS16BE
func (stream *IOStream) ReadS16BE() (int16, error) {
	var value C.Sint16
	if !C.SDL_ReadS16BE((*C.SDL_IOStream)(stream), &value) {
		return 0, getError()
	}
	return int16(value), nil
}

// Use this function to read 32 bits of little-endian data from an
// [IOStream] and return in native format.
//
// SDL byteswaps the data only if necessary, so the data returned will be in
// the native byte order.
//
// This function will return zero when the data stream is completely read, and
// [IOStream.Status] will return [IOStatusEOF]. If zero is returned and
// the stream is not at EOF, [IOStream.Status] will return a different error
// value. In both cases, this method will return a non-nil error
//
// stream: the stream from which to read data.
//
// Returns the data read or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReadU32LE
func (stream *IOStream) ReadU32LE() (uint32, error) {
	var value C.Uint32
	if !C.SDL_ReadU32LE((*C.SDL_IOStream)(stream), &value) {
		return 0, getError()
	}
	return uint32(value), nil
}

// Use this function to read 32 bits of little-endian data from an
// [IOStream] and return in native format.
//
// SDL byteswaps the data only if necessary, so the data returned will be in
// the native byte order.
//
// This function will return zero when the data stream is completely read, and
// [IOStream.Status] will return [IOStatusEOF]. If zero is returned and
// the stream is not at EOF, [IOStream.Status] will return a different error
// value. In both cases, this method will return a non-nil error
//
// stream: the stream from which to read data.
//
// Returns the data read or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReadS32LE
func (stream *IOStream) ReadS32LE() (int32, error) {
	var value C.Sint32
	if !C.SDL_ReadS32LE((*C.SDL_IOStream)(stream), &value) {
		return 0, getError()
	}
	return int32(value), nil
}

// Use this function to read 32 bits of big-endian data from an [IOStream]
// and return in native format.
//
// SDL byteswaps the data only if necessary, so the data returned will be in
// the native byte order.
//
// This function will return zero when the data stream is completely read, and
// [IOStream.Status] will return [IOStatusEOF]. If zero is returned and
// the stream is not at EOF, [IOStream.Status] will return a different error
// value. In both cases, this method will return a non-nil error
//
// stream: the stream from which to read data.
//
// Returns the data read or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReadU32BE
func (stream *IOStream) ReadU32BE() (uint32, error) {
	var value C.Uint32
	if !C.SDL_ReadU32BE((*C.SDL_IOStream)(stream), &value) {
		return 0, getError()
	}
	return uint32(value), nil
}

// Use this function to read 32 bits of big-endian data from an [IOStream]
// and return in native format.
//
// SDL byteswaps the data only if necessary, so the data returned will be in
// the native byte order.
//
// This function will return zero when the data stream is completely read, and
// [IOStream.Status] will return [IOStatusEOF]. If zero is returned and
// the stream is not at EOF, [IOStream.Status] will return a different error
// value. In both cases, this method will return a non-nil error
//
// stream: the stream from which to read data.
//
// Returns the data read or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReadS32BE
func (stream *IOStream) ReadS32BE() (int32, error) {
	var value C.Sint32
	if !C.SDL_ReadS32BE((*C.SDL_IOStream)(stream), &value) {
		return 0, getError()
	}
	return int32(value), nil
}

// Use this function to read 64 bits of little-endian data from an
// [IOStream] and return in native format.
//
// SDL byteswaps the data only if necessary, so the data returned will be in
// the native byte order.
//
// This function will return zero when the data stream is completely read, and
// [IOStream.Status] will return [IOStatusEOF]. If zero is returned and
// the stream is not at EOF, [IOStream.Status] will return a different error
// value. In both cases, this method will return a non-nil error
//
// stream: the stream from which to read data.
//
// Returns the data read or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReadU64LE
func (stream *IOStream) ReadU64LE() (uint64, error) {
	var value C.Uint64
	if !C.SDL_ReadU64LE((*C.SDL_IOStream)(stream), &value) {
		return 0, getError()
	}
	return uint64(value), nil
}

// Use this function to read 64 bits of little-endian data from an
// [IOStream] and return in native format.
//
// SDL byteswaps the data only if necessary, so the data returned will be in
// the native byte order.
//
// This function will return zero when the data stream is completely read, and
// [IOStream.Status] will return [IOStatusEOF]. If zero is returned and
// the stream is not at EOF, [IOStream.Status] will return a different error
// value. In both cases, this method will return a non-nil error
//
// stream: the stream from which to read data.
//
// Returns the data read or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReadS64LE
func (stream *IOStream) ReadS64LE() (int64, error) {
	var value C.Sint64
	if !C.SDL_ReadS64LE((*C.SDL_IOStream)(stream), &value) {
		return 0, getError()
	}
	return int64(value), nil
}

// Use this function to read 64 bits of big-endian data from an [IOStream]
// and return in native format.
//
// SDL byteswaps the data only if necessary, so the data returned will be in
// the native byte order.
//
// This function will return zero when the data stream is completely read, and
// [IOStream.Status] will return [IOStatusEOF]. If zero is returned and
// the stream is not at EOF, [IOStream.Status] will return a different error
// value. In both cases, this method will return a non-nil error
//
// stream: the stream from which to read data.
//
// Returns the data read or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReadU64BE
func (stream *IOStream) ReadU64BE() (uint64, error) {
	var value C.Uint64
	if !C.SDL_ReadU64BE((*C.SDL_IOStream)(stream), &value) {
		return 0, getError()
	}
	return uint64(value), nil
}

// Use this function to read 64 bits of big-endian data from an [IOStream]
// and return in native format.
//
// SDL byteswaps the data only if necessary, so the data returned will be in
// the native byte order.
//
// This function will return zero when the data stream is completely read, and
// [IOStream.Status] will return [IOStatusEOF]. If zero is returned and
// the stream is not at EOF, [IOStream.Status] will return a different error
// value. In both cases, this method will return a non-nil error
//
// stream: the stream from which to read data.
//
// Returns the data read or an error.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReadS64BE
func (stream *IOStream) ReadS64BE() (int64, error) {
	var value C.Sint64
	if !C.SDL_ReadS64BE((*C.SDL_IOStream)(stream), &value) {
		return 0, getError()
	}
	return int64(value), nil
}

// \name Write endian functions
//
// Write an item of native format to the specified endianness.

// Use this function to write a byte to an [IOStream].
//
// stream: the [IOStream] to write to.
//
// Returns nil on successful write or an error on failure.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WriteU8
func (stream *IOStream) WriteU8(value byte) error {
	if !C.SDL_WriteU8((*C.SDL_IOStream)(stream), (C.Uint8)(value)) {
		return getError()
	}
	return nil
}

// Use this function to write a signed byte to an [IOStream].
//
// stream: the [IOStream] to write to.
//
// Returns nil on successful write or an error on failure.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WriteS8
func (stream *IOStream) WriteS8(value int8) error {
	if !C.SDL_WriteS8((*C.SDL_IOStream)(stream), (C.Sint8)(value)) {
		return getError()
	}
	return nil
}

// Use this function to write 16 bits in native format to an [IOStream] as
// little-endian data.
//
// SDL byteswaps the data only if necessary, so the application always
// specifies native format, and the data written will be in little-endian
// format.
//
// stream: the stream to which data will be written.
//
// Returns nil on successful write or an error on failure.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WriteU16LE
func (stream *IOStream) WriteU16LE(value uint16) error {
	if !C.SDL_WriteU16LE((*C.SDL_IOStream)(stream), (C.Uint16)(value)) {
		return getError()
	}
	return nil
}

// Use this function to write 16 bits in native format to an [IOStream] as
// little-endian data.
//
// SDL byteswaps the data only if necessary, so the application always
// specifies native format, and the data written will be in little-endian
// format.
//
// stream: the stream to which data will be written.
//
// Returns nil on successful write or an error on failure.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WriteS16LE
func (stream *IOStream) WriteS16LE(value int16) error {
	if !C.SDL_WriteS16LE((*C.SDL_IOStream)(stream), (C.Sint16)(value)) {
		return getError()
	}
	return nil
}

// Use this function to write 16 bits in native format to an [IOStream] as
// big-endian data.
//
// SDL byteswaps the data only if necessary, so the application always
// specifies native format, and the data written will be in big-endian format.
//
// stream: the stream to which data will be written.
//
// Returns nil on successful write or an error on failure.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WriteU16BE
func (stream *IOStream) WriteU16BE(value uint16) error {
	if !C.SDL_WriteU16BE((*C.SDL_IOStream)(stream), (C.Uint16)(value)) {
		return getError()
	}
	return nil
}

// Use this function to write 16 bits in native format to an [IOStream] as
// big-endian data.
//
// SDL byteswaps the data only if necessary, so the application always
// specifies native format, and the data written will be in big-endian format.
//
// stream: the stream to which data will be written.
//
// Returns nil on successful write or an error on failure.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WriteS16BE
func (stream *IOStream) WriteS16BE(value int16) error {
	if !C.SDL_WriteS16BE((*C.SDL_IOStream)(stream), (C.Sint16)(value)) {
		return getError()
	}
	return nil
}

// Use this function to write 32 bits in native format to an [IOStream] as
// little-endian data.
//
// SDL byteswaps the data only if necessary, so the application always
// specifies native format, and the data written will be in little-endian
// format.
//
// stream: the stream to which data will be written.
//
// Returns nil on successful write or an error on failure.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WriteU32LE
func (stream *IOStream) WriteU32LE(value uint32) error {
	if !C.SDL_WriteU32LE((*C.SDL_IOStream)(stream), (C.Uint32)(value)) {
		return getError()
	}
	return nil
}

// Use this function to write 32 bits in native format to an [IOStream] as
// little-endian data.
//
// SDL byteswaps the data only if necessary, so the application always
// specifies native format, and the data written will be in little-endian
// format.
//
// stream: the stream to which data will be written.
//
// Returns nil on successful write or an error on failure.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WriteS32LE
func (stream *IOStream) WriteS32LE(value int32) error {
	if !C.SDL_WriteS32LE((*C.SDL_IOStream)(stream), (C.Sint32)(value)) {
		return getError()
	}
	return nil
}

// Use this function to write 32 bits in native format to an [IOStream] as
// big-endian data.
//
// SDL byteswaps the data only if necessary, so the application always
// specifies native format, and the data written will be in big-endian format.
//
// stream: the stream to which data will be written.
//
// Returns nil on successful write or an error on failure.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WriteU32BE
func (stream *IOStream) WriteU32BE(value uint32) error {
	if !C.SDL_WriteU32BE((*C.SDL_IOStream)(stream), (C.Uint32)(value)) {
		return getError()
	}
	return nil
}

// Use this function to write 32 bits in native format to an [IOStream] as
// big-endian data.
//
// SDL byteswaps the data only if necessary, so the application always
// specifies native format, and the data written will be in big-endian format.
//
// stream: the stream to which data will be written.
//
// Returns nil on successful write or an error on failure.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WriteS32BE
func (stream *IOStream) WriteS32BE(value int32) error {
	if !C.SDL_WriteS32BE((*C.SDL_IOStream)(stream), (C.Sint32)(value)) {
		return getError()
	}
	return nil
}

// Use this function to write 64 bits in native format to an [IOStream] as
// little-endian data.
//
// SDL byteswaps the data only if necessary, so the application always
// specifies native format, and the data written will be in little-endian
// format.
//
// stream: the stream to which data will be written.
//
// Returns nil on successful write or an error on failure.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WriteU64LE
func (stream *IOStream) WriteU64LE(value uint64) error {
	if !C.SDL_WriteU64LE((*C.SDL_IOStream)(stream), (C.Uint64)(value)) {
		return getError()
	}
	return nil
}

// Use this function to write 64 bits in native format to an [IOStream] as
// little-endian data.
//
// SDL byteswaps the data only if necessary, so the application always
// specifies native format, and the data written will be in little-endian
// format.
//
// stream: the stream to which data will be written.
//
// Returns nil on successful write or an error on failure.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WriteS64LE
func (stream *IOStream) WriteS64LE(value int64) error {
	if !C.SDL_WriteS64LE((*C.SDL_IOStream)(stream), (C.Sint64)(value)) {
		return getError()
	}
	return nil
}

// Use this function to write 64 bits in native format to an [IOStream] as
// big-endian data.
//
// SDL byteswaps the data only if necessary, so the application always
// specifies native format, and the data written will be in big-endian format.
//
// stream: the stream to which data will be written.
//
// Returns nil on successful write or an error on failure.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WriteU64BE
func (stream *IOStream) WriteU64BE(value uint64) error {
	if !C.SDL_WriteU64BE((*C.SDL_IOStream)(stream), (C.Uint64)(value)) {
		return getError()
	}
	return nil
}

// Use this function to write 64 bits in native format to an [IOStream] as
// big-endian data.
//
// SDL byteswaps the data only if necessary, so the application always
// specifies native format, and the data written will be in big-endian format.
//
// stream: the stream to which data will be written.
//
// Returns nil on successful write or an error on failure.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WriteS64BE
func (stream *IOStream) WriteS64BE(value int64) error {
	if !C.SDL_WriteS64BE((*C.SDL_IOStream)(stream), (C.Sint64)(value)) {
		return getError()
	}
	return nil
}
