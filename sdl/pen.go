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
// #include <SDL3/SDL.h>
import "C"

// # CategoryPen
//
// SDL pen event handling.
//
// SDL provides an API for pressure-sensitive pen (stylus and/or eraser)
// handling, e.g., for input and drawing tablets or suitably equipped mobile /
// tablet devices.
//
// To get started with pens, simply handle EventPen* events. When a pen
// starts providing input, SDL will assign it a unique [PenID], which will
// remain for the life of the process, as long as the pen stays connected.
//
// Pens may provide more than simple touch input; they might have other axes,
// such as pressure, tilt, rotation, etc.

// SDL pen instance IDs.
//
// Zero is used to signify an invalid/null device.
//
// These show up in pen events when SDL sees input from them. They remain
// consistent as long as SDL can recognize a tool to be the same pen; but if a
// pen physically leaves the area and returns, it might get a new ID.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PenID
type PenID uint32

// The [MouseID] for mouse events simulated with pen input.
//
// This macro is available since SDL 3.2.0.
const PenMouseID MouseID = 0xFFFFFFFE

// The [TouchID] for touch events simulated with pen input.
//
// This macro is available since SDL 3.2.0.
const PenTouchID TouchID = 0xFFFFFFFE

// Pen input flags, as reported by various pen events' PenState field.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PenInputFlags
type PenInputFlags uint32

const (
	PenInputDown      PenInputFlags = (1 << 0)  // pen is pressed down
	PenInputButton1   PenInputFlags = (1 << 1)  // button 1 is pressed
	PenInputButton2   PenInputFlags = (1 << 2)  // button 2 is pressed
	PenInputButton3   PenInputFlags = (1 << 3)  // button 3 is pressed
	PenInputButton4   PenInputFlags = (1 << 4)  // button 4 is pressed
	PenInputButton5   PenInputFlags = (1 << 5)  // button 5 is pressed
	PenInputEraserTip PenInputFlags = (1 << 30) // eraser tip is used
)

// Pen axis indices.
//
// These are the valid values for the Axis field in [PenAxisEvent]. All
// axes are either normalised to 0..1 or report a (positive or negative) angle
// in degrees, with 0.0 representing the centre. Not all pens/backends support
// all axes: unsupported axes are always zero.
//
// To convert angles for tilt and rotation into vector representation, use
// [math.Sin] on the XTilt, YTilt, or Rotation component, for example:
//
//	math.Sin(float64(xtilt) * math.Pi / 180)
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PenAxis
type PenAxis uint32

const (
	PenAxisPressure           PenAxis = iota // Pen pressure.  Unidirectional: 0 to 1.0
	PenAxisXTilt                             // Pen horizontal tilt angle.  Bidirectional: -90.0 to 90.0 (left-to-right).
	PenAxisYTilt                             // Pen vertical tilt angle.  Bidirectional: -90.0 to 90.0 (top-to-down).
	PenAxisDistance                          // Pen distance to drawing surface.  Unidirectional: 0.0 to 1.0
	PenAxisRotation                          // Pen barrel rotation.  Bidirectional: -180 to 179.9 (clockwise, 0 is facing up, -180.0 is facing down).
	PenAxisSlider                            // Pen finger wheel or slider (e.g., Airbrush Pen).  Unidirectional: 0 to 1.0
	PenAxisTangentialPressure                // Pressure from squeezing the pen ("barrel pressure").
	PenAxisCount              = iota         // Total known pen axis types in this version of SDL. This number may grow in future releases!
)
