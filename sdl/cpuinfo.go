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
// #cgo noescape SDL_GetNumLogicalCPUCores
// #cgo nocallback SDL_GetNumLogicalCPUCores
// #cgo noescape SDL_GetCPUCacheLineSize
// #cgo nocallback SDL_GetCPUCacheLineSize
// #cgo noescape SDL_HasAltiVec
// #cgo nocallback SDL_HasAltiVec
// #cgo noescape SDL_HasMMX
// #cgo nocallback SDL_HasMMX
// #cgo noescape SDL_HasSSE
// #cgo nocallback SDL_HasSSE
// #cgo noescape SDL_HasSSE2
// #cgo nocallback SDL_HasSSE2
// #cgo noescape SDL_HasSSE3
// #cgo nocallback SDL_HasSSE3
// #cgo noescape SDL_HasSSE41
// #cgo nocallback SDL_HasSSE41
// #cgo noescape SDL_HasSSE42
// #cgo nocallback SDL_HasSSE42
// #cgo noescape SDL_HasAVX
// #cgo nocallback SDL_HasAVX
// #cgo noescape SDL_HasAVX2
// #cgo nocallback SDL_HasAVX2
// #cgo noescape SDL_HasAVX512F
// #cgo nocallback SDL_HasAVX512F
// #cgo noescape SDL_HasARMSIMD
// #cgo nocallback SDL_HasARMSIMD
// #cgo noescape SDL_HasNEON
// #cgo nocallback SDL_HasNEON
// #cgo noescape SDL_HasLSX
// #cgo nocallback SDL_HasLSX
// #cgo noescape SDL_HasLASX
// #cgo nocallback SDL_HasLASX
// #cgo noescape SDL_GetSystemRAM
// #cgo nocallback SDL_GetSystemRAM
// #cgo noescape SDL_GetSIMDAlignment
// #cgo nocallback SDL_GetSIMDAlignment
// #include <SDL3/SDL.h>
import "C"

// # CategoryCPUInfo
//
// CPU feature detection for SDL.
//
// These functions are largely concerned with reporting if the system has
// access to various SIMD instruction sets, but also has other important info
// to share, such as system RAM size and number of logical CPU cores.
//
// CPU instruction set checks, like SDL_HasSSE() and SDL_HasNEON(), are
// available on all platforms, even if they don't make sense (an ARM processor
// will never have SSE and an x86 processor will never have NEON, for example,
// but these functions still exist and will simply return false in these
// cases).

// A guess for the cacheline size used for padding.
//
// Most x86 processors have a 64 byte cache line. The 64-bit PowerPC
// processors have a 128 byte cache line. We use the larger value to be
// generally safe.
//
// This macro is available since SDL 3.2.0.
const CachelineSize = 128

// Get the number of logical CPU cores available.
//
// Returns the total number of logical CPU cores. On CPUs that include
// technologies such as hyperthreading, the number of logical cores
// may be more than the number of physical cores.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetNumLogicalCPUCores
func GetNumLogicalCPUCores() int {
	return (int)(C.SDL_GetNumLogicalCPUCores())
}

// Determine the L1 cache line size of the CPU.
//
// This is useful for determining multi-threaded structure padding or SIMD
// prefetch sizes.
//
// Returns the L1 cache line size of the CPU, in bytes.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetCPUCacheLineSize
func GetCPUCacheLineSize() int {
	return (int)(C.SDL_GetCPUCacheLineSize())
}

// Determine whether the CPU has AltiVec features.
//
// This always returns false on CPUs that aren't using PowerPC instruction
// sets.
//
// Returns true if the CPU has AltiVec features or false if not.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasAltiVec
func HasAltiVec() bool {
	return (bool)(C.SDL_HasAltiVec())
}

// Determine whether the CPU has MMX features.
//
// This always returns false on CPUs that aren't using Intel instruction sets.
//
// Returns true if the CPU has MMX features or false if not.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasMMX
func HasMMX() bool {
	return (bool)(C.SDL_HasMMX())
}

// Determine whether the CPU has SSE features.
//
// This always returns false on CPUs that aren't using Intel instruction sets.
//
// Returns true if the CPU has SSE features or false if not.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasSSE
func HasSSE() bool {
	return (bool)(C.SDL_HasSSE())
}

// Determine whether the CPU has SSE2 features.
//
// This always returns false on CPUs that aren't using Intel instruction sets.
//
// Returns true if the CPU has SSE2 features or false if not.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasSSE2
func HasSSE2() bool {
	return (bool)(C.SDL_HasSSE2())
}

// Determine whether the CPU has SSE3 features.
//
// This always returns false on CPUs that aren't using Intel instruction sets.
//
// Returns true if the CPU has SSE3 features or false if not.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasSSE3
func HasSSE3() bool {
	return (bool)(C.SDL_HasSSE3())
}

// Determine whether the CPU has SSE4.1 features.
//
// This always returns false on CPUs that aren't using Intel instruction sets.
//
// Returns true if the CPU has SSE4.1 features or false if not.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasSSE41
func HasSSE41() bool {
	return (bool)(C.SDL_HasSSE41())
}

// Determine whether the CPU has SSE4.2 features.
//
// This always returns false on CPUs that aren't using Intel instruction sets.
//
// Returns true if the CPU has SSE4.2 features or false if not.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasSSE42
func HasSSE42() bool {
	return (bool)(C.SDL_HasSSE42())
}

// Determine whether the CPU has AVX features.
//
// This always returns false on CPUs that aren't using Intel instruction sets.
//
// Returns true if the CPU has AVX features or false if not.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasAVX
func HasAVX() bool {
	return (bool)(C.SDL_HasAVX())
}

// Determine whether the CPU has AVX2 features.
//
// This always returns false on CPUs that aren't using Intel instruction sets.
//
// Returns true if the CPU has AVX2 features or false if not.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasAVX2
func HasAVX2() bool {
	return (bool)(C.SDL_HasAVX2())
}

// Determine whether the CPU has AVX-512F (foundation) features.
//
// This always returns false on CPUs that aren't using Intel instruction sets.
//
// Returns true if the CPU has AVX-512F features or false if not.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasAVX512F
func HasAVX512F() bool {
	return (bool)(C.SDL_HasAVX512F())
}

// Determine whether the CPU has ARM SIMD (ARMv6) features.
//
// This is different from ARM NEON, which is a different instruction set.
//
// This always returns false on CPUs that aren't using ARM instruction sets.
//
// Returns true if the CPU has ARM SIMD features or false if not.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasARMSIMD
func HasARMSIMD() bool {
	return (bool)(C.SDL_HasARMSIMD())
}

// Determine whether the CPU has NEON (ARM SIMD) features.
//
// This always returns false on CPUs that aren't using ARM instruction sets.
//
// Returns true if the CPU has ARM NEON features or false if not.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasNEON
func HasNEON() bool {
	return (bool)(C.SDL_HasNEON())
}

// Determine whether the CPU has LSX (LOONGARCH SIMD) features.
//
// This always returns false on CPUs that aren't using LOONGARCH instruction
// sets.
//
// Returns true if the CPU has LOONGARCH LSX features or false if not.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasLSX
func HasLSX() bool {
	return (bool)(C.SDL_HasLSX())
}

// Determine whether the CPU has LASX (LOONGARCH SIMD) features.
//
// This always returns false on CPUs that aren't using LOONGARCH instruction
// sets.
//
// Returns true if the CPU has LOONGARCH LASX features or false if not.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasLASX
func HasLASX() bool {
	return (bool)(C.SDL_HasLASX())
}

// Get the amount of RAM configured in the system.
//
// Returns the amount of RAM configured in the system in MiB.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSystemRAM
func GetSystemRAM() int {
	return (int)(C.SDL_GetSystemRAM())
}

// Report the alignment this system needs for SIMD allocations.
//
// This will return the minimum number of bytes to which a pointer must be
// aligned to be compatible with SIMD instructions on the current machine. For
// example, if the machine supports SSE only, it will return 16, but if it
// supports AVX-512F, it'll return 64 (etc). This only reports values for
// instruction sets SDL knows about, so if your SDL build doesn't have
// SDL_HasAVX512F(), then it might return 16 for the SSE support it sees and
// not 64 for the AVX-512 instructions that exist but SDL doesn't know about.
// Plan accordingly.
//
// Returns the alignment in bytes needed for available, known SIMD
// instructions.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSIMDAlignment
func GetSIMDAlignment() int {
	return (int)(C.SDL_GetSIMDAlignment())
}
