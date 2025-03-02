# Go bindings for SDL3

This package provides cgo-based bindings for SDL3

## Status

- Relatively untested, feedback/bug reports are very welcome
- Breaking API changes may happen
- Most functions are available, exceptions are:
	- some macros and varargs functions that don't really make sense in Go
	- headers that provide functionality that is already available in the Go standard library
	- Feel free to open an issue if you think a particular function or header should be implemented
- The documentation is copied from the C headers and slightly adjusted (cross-references, different parameters etc.), but does not follow Go conventions.
  Reworking it to conform to godoc style would be a huge effort with questionable benefit

## Why not purego-sdl3

[purego-sdl3](https://github.com/JupiterRider/purego-sdl3) is another SDL3 binding. The approach taken there is interesting and very convenient for cross-compilation, but also experimental and reliant on internal details of the Go runtime. For some use-cases, "traditional" cgo bindings may be more appropriate.
