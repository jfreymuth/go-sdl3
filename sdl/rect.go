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
// #cgo noescape SDL_HasRectIntersection
// #cgo nocallback SDL_HasRectIntersection
// #cgo noescape SDL_GetRectIntersection
// #cgo nocallback SDL_GetRectIntersection
// #cgo noescape SDL_GetRectUnion
// #cgo nocallback SDL_GetRectUnion
// #cgo noescape SDL_GetRectEnclosingPoints
// #cgo nocallback SDL_GetRectEnclosingPoints
// #cgo noescape SDL_GetRectAndLineIntersection
// #cgo nocallback SDL_GetRectAndLineIntersection
// #cgo noescape SDL_HasRectIntersectionFloat
// #cgo nocallback SDL_HasRectIntersectionFloat
// #cgo noescape SDL_GetRectIntersectionFloat
// #cgo nocallback SDL_GetRectIntersectionFloat
// #cgo noescape SDL_GetRectUnionFloat
// #cgo nocallback SDL_GetRectUnionFloat
// #cgo noescape SDL_GetRectEnclosingPointsFloat
// #cgo nocallback SDL_GetRectEnclosingPointsFloat
// #cgo noescape SDL_GetRectAndLineIntersectionFloat
// #cgo nocallback SDL_GetRectAndLineIntersectionFloat
// #include <SDL3/SDL.h>
import "C"
import "unsafe"

// # CategoryRect
//
// Some helper functions for managing rectangles and 2D points, in both
// integer and floating point versions.

// The structure that defines a point (using integers).
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Point
type Point struct {
	X int
	Y int
}

// The structure that defines a point (using floating point values).
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_FPoint
type FPoint struct {
	X float32
	Y float32
}

// A rectangle, with the origin at the upper left (using integers).
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Rect
type Rect struct {
	X int
	Y int
	W int
	H int
}

// A rectangle, with the origin at the upper left (using floating point
// values).
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_FRect
type FRect struct {
	X float32
	Y float32
	W float32
	H float32
}

// Convert a [Rect] to [FRect]
//
// rect: a pointer to a [Rect].
//
// frect: a pointer filled in with the floating point representation of
// `rect`.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RectToFRect
func (r Rect) ToFRect() FRect {
	return FRect{float32(r.X), float32(r.Y), float32(r.W), float32(r.H)}
}

// Determine whether a point resides inside a rectangle.
//
// A point is considered part of a rectangle if p's x and y coordinates are >=
// to the rectangle's top left corner, and < the rectangle's x+w and y+h.
// So a 1x1 rectangle considers point (0,0) as "inside" and (0,1) as not.
//
// p: the point to test.
//
// r: the rectangle to test.
//
// Returns true if p is contained by r, false otherwise.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PointInRect
func (p Point) In(r Rect) bool {
	return p.X >= r.X && p.X < r.X+r.W && p.Y >= r.Y && p.Y < r.Y+r.H
}

// Determine whether a rectangle has no area.
//
// A rectangle is considered "empty" for this function if r is nil, or if
// r's width and/or height are <= 0.
//
// r: the rectangle to test.
//
// Returns true if the rectangle is "empty", false otherwise.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RectEmpty
func (r *Rect) Empty() bool {
	return r == nil || r.W <= 0 || r.H <= 0
}

// Determine whether two rectangles intersect.
//
// a: a [Rect] structure representing the first rectangle.
//
// b: a [Rect] structure representing the second rectangle.
//
// Returns true if there is an intersection, false otherwise.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasRectIntersection
func (a Rect) HasIntersection(b Rect) bool {
	return (bool)(C.SDL_HasRectIntersection(
		&C.SDL_Rect{C.int(a.X), C.int(a.Y), C.int(a.W), C.int(a.H)},
		&C.SDL_Rect{C.int(b.X), C.int(b.Y), C.int(b.W), C.int(b.H)}))
}

// Calculate the intersection of two rectangles.
//
// a: a [Rect] structure representing the first rectangle.
//
// b: a [Rect] structure representing the second rectangle.
//
// result: the intersection of rectangles A and B.
//
// Returns true if there is an intersection, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRectIntersection
func (a Rect) Intersection(b Rect) (Rect, bool) {
	var result C.SDL_Rect
	intersects := C.SDL_GetRectIntersection(
		&C.SDL_Rect{C.int(a.X), C.int(a.Y), C.int(a.W), C.int(a.H)},
		&C.SDL_Rect{C.int(b.X), C.int(b.Y), C.int(b.W), C.int(b.H)}, &result)
	return Rect{int(result.x), int(result.y), int(result.w), int(result.h)}, bool(intersects)
}

// Calculate the union of two rectangles.
//
// a: a [Rect] structure representing the first rectangle.
//
// b: a [Rect] structure representing the second rectangle.
//
// Returns the union of rectangles A and B.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRectUnion
func (a Rect) Union(b Rect) Rect {
	var result C.SDL_Rect
	C.SDL_GetRectUnion(
		&C.SDL_Rect{C.int(a.X), C.int(a.Y), C.int(a.W), C.int(a.H)},
		&C.SDL_Rect{C.int(b.X), C.int(b.Y), C.int(b.W), C.int(b.H)}, &result)
	return Rect{int(result.x), int(result.y), int(result.w), int(result.h)}
}

// Calculate a minimal rectangle enclosing a set of points.
//
// If clip is not nil then only points inside of the clipping rectangle are
// considered.
//
// points: an array of Points representing points to be enclosed.
//
// clip: a [Rect] used for clipping or nil to enclose all points.
//
// result: the minimal enclosing rectangle.
//
// Returns true if any points were enclosed or false if all the points were
// outside of the clipping rectangle.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPoints
func GetRectEnclosingPoints(points []Point, clip *Rect) (Rect, bool) {
	var cclip *C.SDL_Rect
	if clip != nil {
		cclip = &C.SDL_Rect{C.int(clip.X), C.int(clip.Y), C.int(clip.W), C.int(clip.H)}
	}
	cpoints := make([]C.SDL_Point, len(points))
	for i, p := range points {
		cpoints[i] = C.SDL_Point{C.int(p.X), C.int(p.Y)}
	}
	var result C.SDL_Rect
	ok := C.SDL_GetRectEnclosingPoints((*C.SDL_Point)(unsafe.SliceData(cpoints)), (C.int)(len(points)), cclip, &result)
	return Rect{int(result.x), int(result.y), int(result.w), int(result.h)}, bool(ok)
}

// Calculate the intersection of a rectangle and line segment.
//
// This function is used to clip a line segment to a rectangle. A line segment
// contained entirely within the rectangle or that does not intersect will
// remain unchanged. A line segment that crosses the rectangle at either or
// both ends will be clipped to the boundary of the rectangle and the new
// coordinates saved in x1, y1, x2, and/or y1 as necessary.
//
// rect: a [Rect] structure representing the rectangle to intersect.
//
// x1: a pointer to the starting X-coordinate of the line.
//
// y1: a pointer to the starting Y-coordinate of the line.
//
// x2: a pointer to the ending X-coordinate of the line.
//
// y2: a pointer to the ending Y-coordinate of the line.
//
// Returns true if there is an intersection, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersection
func (r Rect) GetLineIntersection(x1 *int, y1 *int, x2 *int, y2 *int) bool {
	var cx1, cy1, cx2, cy2 C.int
	ok := C.SDL_GetRectAndLineIntersection(&C.SDL_Rect{C.int(r.X), C.int(r.Y), C.int(r.W), C.int(r.H)}, &cx1, &cy1, &cx2, &cy2)
	*x1 = int(cx1)
	*y1 = int(cy1)
	*x2 = int(cx2)
	*y2 = int(cy2)
	return bool(ok)
}

// Determine whether a point resides inside a floating point rectangle.
//
// A point is considered part of a rectangle if p's x and y coordinates are >=
// to the rectangle's top left corner, and <= the rectangle's x+w and y+h. So a
// 1x1 rectangle considers point (0,0) and (0,1) as "inside" and (0,2) as not.
//
// p: the point to test.
//
// r: the rectangle to test.
//
// Returns true if p is contained by r, false otherwise.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PointInRectFloat
func (p FPoint) In(r FRect) bool {
	return p.X >= r.X && p.X <= r.X+r.W && p.Y >= r.Y && p.Y <= r.Y+r.H
}

// Determine whether a floating point rectangle can contain any point.
//
// A rectangle is considered "empty" for this function if r is nil, or if
// r's width and/or height are < 0.0f.
//
// r: the rectangle to test.
//
// Returns true if the rectangle is "empty", false otherwise.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RectEmptyFloat
func (r *FRect) Empty() bool {
	return r == nil || r.W < 0 || r.H < 0
}

// Determine whether two floating point rectangles are equal, within some
// given epsilon.
//
// Rectangles are considered equal if each of their x,
// y, width and height are within epsilon of each other. If you don't know
// what value to use for epsilon, you should call the [FRect.Equal]
// function instead.
//
// a: the first rectangle to test.
//
// b: the second rectangle to test.
//
// epsilon: the epsilon value for comparison.
//
// Returns true if the rectangles are equal, false otherwise.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RectsEqualEpsilon
func (a FRect) EqualEpsilon(b FRect, epsilon float32) bool {
	dx, dy, dw, dh := a.X-b.X, a.Y-b.Y, a.W-b.W, a.H-b.H
	return max(dx, -dx) <= epsilon &&
		max(dy, -dy) <= epsilon &&
		max(dw, -dw) <= epsilon &&
		max(dh, -dh) <= epsilon
}

// Determine whether two floating point rectangles are equal, within a default
// epsilon.
//
// Rectangles are considered equal if each of their x,
// y, width and height are within SDL_FLT_EPSILON of each other. This is often
// a reasonable way to compare two floating point rectangles and deal with the
// slight precision variations in floating point calculations that tend to pop
// up.
//
// a: the first rectangle to test.
//
// b: the second rectangle to test.
//
// Returns true if the rectangles are equal, false otherwise.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RectsEqualFloat
func (a FRect) Equal(b FRect, epsilon float32) bool {
	return a.EqualEpsilon(b, 0x0.000002p0)
}

// Determine whether two rectangles intersect with float precision.
//
// a: an [FRect] structure representing the first rectangle.
//
// b: an [FRect] structure representing the second rectangle.
//
// Returns true if there is an intersection, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasRectIntersectionFloat
func (a FRect) HasIntersection(b FRect) bool {
	return (bool)(C.SDL_HasRectIntersectionFloat(
		&C.SDL_FRect{C.float(a.X), C.float(a.Y), C.float(a.W), C.float(a.H)},
		&C.SDL_FRect{C.float(b.X), C.float(b.Y), C.float(b.W), C.float(b.H)}))
}

// Calculate the intersection of two rectangles with float precision.
//
// a: an [FRect] structure representing the first rectangle.
//
// b: an [FRect] structure representing the second rectangle.
//
// result: the intersection of rectangles A and B.
//
// Returns true if there is an intersection, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRectIntersectionFloat
func (a FRect) Intersection(b FRect) (FRect, bool) {
	var result C.SDL_FRect
	intersects := C.SDL_GetRectIntersectionFloat(
		&C.SDL_FRect{C.float(a.X), C.float(a.Y), C.float(a.W), C.float(a.H)},
		&C.SDL_FRect{C.float(b.X), C.float(b.Y), C.float(b.W), C.float(b.H)}, &result)
	return FRect{float32(result.x), float32(result.y), float32(result.w), float32(result.h)}, bool(intersects)
}

// Calculate the union of two rectangles with float precision.
//
// a: an [FRect] structure representing the first rectangle.
//
// b: an [FRect] structure representing the second rectangle.
//
// Returns the union of rectangles A and B.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRectUnionFloat
func (a FRect) Union(b FRect) FRect {
	var result C.SDL_FRect
	C.SDL_GetRectUnionFloat(
		&C.SDL_FRect{C.float(a.X), C.float(a.Y), C.float(a.W), C.float(a.H)},
		&C.SDL_FRect{C.float(b.X), C.float(b.Y), C.float(b.W), C.float(b.H)}, &result)
	return FRect{float32(result.x), float32(result.y), float32(result.w), float32(result.h)}
}

// Calculate a minimal rectangle enclosing a set of points with float
// precision.
//
// If clip is not nil then only points inside of the clipping rectangle are
// considered.
//
// points: an array of [FPoint] structures representing points to be
// enclosed.
//
// clip: an [FRect] used for clipping or nil to enclose all points.
//
// result: an [FRect] structure filled in with the minimal enclosing
// rectangle.
//
// Returns true if any points were enclosed or false if all the points were
// outside of the clipping rectangle.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPointsFloat
func GetRectEnclosingPointsFloat(points []FPoint, clip *FRect) (FRect, bool) {
	var cclip *C.SDL_FRect
	if clip != nil {
		cclip = &C.SDL_FRect{C.float(clip.X), C.float(clip.Y), C.float(clip.W), C.float(clip.H)}
	}
	cpoints := make([]C.SDL_FPoint, len(points))
	for i, p := range points {
		cpoints[i] = C.SDL_FPoint{C.float(p.X), C.float(p.Y)}
	}
	var result C.SDL_FRect
	ok := C.SDL_GetRectEnclosingPointsFloat((*C.SDL_FPoint)(unsafe.SliceData(cpoints)), (C.int)(len(points)), cclip, &result)
	return FRect{float32(result.x), float32(result.y), float32(result.w), float32(result.h)}, bool(ok)
}

// Calculate the intersection of a rectangle and line segment with float
// precision.
//
// This function is used to clip a line segment to a rectangle. A line segment
// contained entirely within the rectangle or that does not intersect will
// remain unchanged. A line segment that crosses the rectangle at either or
// both ends will be clipped to the boundary of the rectangle and the new
// coordinates saved in x1, y1, x2, and/or y2 as necessary.
//
// rect: an [FRect] structure representing the rectangle to intersect.
//
// x1: a pointer to the starting X-coordinate of the line.
//
// y1: a pointer to the starting Y-coordinate of the line.
//
// x2: a pointer to the ending X-coordinate of the line.
//
// y2: a pointer to the ending Y-coordinate of the line.
//
// Returns true if there is an intersection, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersectionFloat
func (r *FRect) GetLineIntersection(x1 *float32, y1 *float32, x2 *float32, y2 *float32) bool {
	var cx1, cy1, cx2, cy2 C.float
	ok := C.SDL_GetRectAndLineIntersectionFloat(&C.SDL_FRect{C.float(r.X), C.float(r.Y), C.float(r.W), C.float(r.H)}, &cx1, &cy1, &cx2, &cy2)
	*x1 = float32(cx1)
	*y1 = float32(cy1)
	*x2 = float32(cx2)
	*y2 = float32(cy2)
	return bool(ok)
}
