# Go bindings for SDL3

This package provides cgo-based bindings for SDL3

## Status

- Most functions are available, exceptions are:
	- some macros and varargs functions that don't really make sense in Go
	- headers that provide functionality that is already available in the Go standard library
- Many features are untested, especially OpenGL and platform specific APIs
- The documentation is copied directly from the C headers and needs to be adjusted

## Why not purego-sdl3

[purego-sdl3](https://github.com/JupiterRider/purego-sdl3) is another SDL3 binding. The approach taken there is interesting and very convenient for cross-compilation, but also experimental and reliant on internal details of the Go runtime. For some use-cases, "traditional" cgo bindings may be more appropriate.
