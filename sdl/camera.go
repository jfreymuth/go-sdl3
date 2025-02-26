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
// #cgo noescape SDL_GetNumCameraDrivers
// #cgo nocallback SDL_GetNumCameraDrivers
// #cgo noescape SDL_GetCameraDriver
// #cgo nocallback SDL_GetCameraDriver
// #cgo noescape SDL_GetCurrentCameraDriver
// #cgo nocallback SDL_GetCurrentCameraDriver
// #cgo noescape SDL_GetCameras
// #cgo nocallback SDL_GetCameras
// #cgo noescape SDL_GetCameraSupportedFormats
// #cgo nocallback SDL_GetCameraSupportedFormats
// #cgo noescape SDL_GetCameraName
// #cgo nocallback SDL_GetCameraName
// #cgo noescape SDL_GetCameraPosition
// #cgo nocallback SDL_GetCameraPosition
// #cgo noescape SDL_OpenCamera
// #cgo nocallback SDL_OpenCamera
// #cgo noescape SDL_GetCameraPermissionState
// #cgo nocallback SDL_GetCameraPermissionState
// #cgo noescape SDL_GetCameraID
// #cgo nocallback SDL_GetCameraID
// #cgo noescape SDL_GetCameraProperties
// #cgo nocallback SDL_GetCameraProperties
// #cgo noescape SDL_GetCameraFormat
// #cgo nocallback SDL_GetCameraFormat
// #cgo noescape SDL_AcquireCameraFrame
// #cgo nocallback SDL_AcquireCameraFrame
// #cgo noescape SDL_ReleaseCameraFrame
// #cgo noescape SDL_CloseCamera
// #include <SDL3/SDL.h>
import "C"
import "unsafe"

// # CategoryCamera
//
// Video capture for the SDL library.
//
// This API lets apps read input from video sources, like webcams. Camera
// devices can be enumerated, queried, and opened. Once opened, it will
// provide [Surface] objects as new frames of video come in. These surfaces
// can be uploaded to a [Texture] or processed as pixels in memory.
//
// Several platforms will alert the user if an app tries to access a camera,
// and some will present a UI asking the user if your application should be
// allowed to obtain images at all, which they can deny. A successfully opened
// camera will not provide images until permission is granted. Applications,
// after opening a camera device, can see if they were granted access by
// either polling with the [Camera.PermissionState] function, or waiting
// for an [EventCameraDeviceApproved] or [EventCameraDeviceDenied]
// event. Platforms that don't have any user approval process will report
// approval immediately.
//
// Note that SDL cameras only provide video as individual frames; they will
// not provide full-motion video encoded in a movie file format, although an
// app is free to encode the acquired frames into any format it likes. It also
// does not provide audio from the camera hardware through this API; not only
// do many webcams not have microphones at all, many people--from streamers to
// people on Zoom calls--will want to use a separate microphone regardless of
// the camera. In any case, recorded audio will be available through SDL's
// audio API no matter what hardware provides the microphone.
//
// ## Camera gotchas
//
// Consumer-level camera hardware tends to take a little while to warm up,
// once the device has been opened. Generally most camera apps have some sort
// of UI to take a picture (a button to snap a pic while a preview is showing,
// some sort of multi-second countdown for the user to pose, like a photo
// booth), which puts control in the users' hands, or they are intended to
// stay on for long times (Pokemon Go, etc).
//
// It's not uncommon that a newly-opened camera will provide a couple of
// completely black frames, maybe followed by some under-exposed images. If
// taking a single frame automatically, or recording video from a camera's
// input without the user initiating it from a preview, it could be wise to
// drop the first several frames (if not the first several _seconds_ worth of
// frames!) before using images from a camera.

// This is a unique ID for a camera device for the time it is connected to the
// system, and is never reused for the lifetime of the application.
//
// If the device is disconnected and reconnected, it will get a new ID.
//
// The value 0 is an invalid ID.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CameraID
type CameraID uint32

// The opaque structure used to identify an opened SDL camera.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Camera
type Camera C.struct_SDL_Camera

// The details of an output format for a camera device.
//
// Cameras often support multiple formats; each one will be encapsulated in
// this struct.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CameraSpec
type CameraSpec struct {
	Format               PixelFormat // Frame format
	Colorspace           Colorspace  // Frame colorspace
	Width                int         // Frame width
	Height               int         // Frame height
	FramerateNumerator   int         // Frame rate numerator ((num / denom) == FPS, (denom / num) == duration in seconds)
	FramerateDenominator int         // Frame rate demoninator ((num / denom) == FPS, (denom / num) == duration in seconds)
}

// The position of camera in relation to system device.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CameraPosition
type CameraPosition uint32

const (
	CameraPositionUnknown CameraPosition = iota
	CameraPositionFrontFacing
	CameraPositionBackFacing
)

// Use this function to get the number of built-in camera drivers.
//
// This function returns a hardcoded number. This never returns a negative
// value; if there are no drivers compiled into this build of SDL, this
// function returns zero. The presence of a driver in this list does not mean
// it will function, it just means SDL is capable of interacting with that
// interface. For example, a build of SDL might have v4l2 support, but if
// there's no kernel support available, SDL's v4l2 driver would fail if used.
//
// By default, SDL tries all drivers, in its preferred order, until one is
// found to be usable.
//
// Returns the number of built-in camera drivers.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetNumCameraDrivers
func GetNumCameraDrivers() int {
	return (int)(C.SDL_GetNumCameraDrivers())
}

// Use this function to get the name of a built in camera driver.
//
// The list of camera drivers is given in the order that they are normally
// initialized by default; the drivers that seem more reasonable to choose
// first (as far as the SDL developers believe) are earlier in the list.
//
// The names of drivers are all simple, low-ASCII identifiers, like "v4l2",
// "coremedia" or "android". These never have Unicode characters, and are not
// meant to be proper names.
//
// index: the index of the camera driver; the value ranges from 0 to
// [GetNumCameraDrivers] - 1.
//
// Returns the name of the camera driver at the requested index, or an empty
// string if an invalid index was specified.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetCameraDriver
func GetCameraDriver(index int) string {
	return C.GoString(C.SDL_GetCameraDriver((C.int)(index)))
}

// Get the name of the current camera driver.
//
// The names of drivers are all simple, low-ASCII identifiers, like "v4l2",
// "coremedia" or "android". These never have Unicode characters, and are not
// meant to be proper names.
//
// Returns the name of the current camera driver or an empty string if no driver
// has been initialized.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetCurrentCameraDriver
func GetCurrentCameraDriver() string {
	return C.GoString(C.SDL_GetCurrentCameraDriver())
}

// Get a list of currently connected camera devices.
//
// Returns a slice of camera instance IDs or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetCameras
func GetCameras() ([]CameraID, error) {
	var count C.int
	cameras := C.SDL_GetCameras(&count)
	if cameras == nil {
		return nil, getError()
	}
	result := make([]CameraID, count)
	for i, k := range unsafe.Slice(cameras, count) {
		result[i] = CameraID(k)
	}
	C.SDL_free(unsafe.Pointer(cameras))
	return result, nil
}

// Get the list of native formats/sizes a camera supports.
//
// This returns a list of all formats and frame sizes that a specific camera
// can offer. This is useful if your app can accept a variety of image formats
// and sizes and so want to find the optimal spec that doesn't require
// conversion.
//
// This function isn't strictly required; if you call [OpenCamera] with a
// nil spec, SDL will choose a native format for you, and if you instead
// specify a desired format, it will transparently convert to the requested
// format on your behalf.
//
// Note that it's legal for a camera to supply an empty list. This is what
// will happen on Emscripten builds, since that platform won't tell _anything_
// about available cameras until you've opened one, and won't even tell if
// there _is_ a camera until the user has given you permission to check
// through a scary warning popup.
//
// devid: the camera device instance ID to query.
//
// Returns a slice of pointers to [CameraSpec] or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetCameraSupportedFormats
func (id CameraID) SupportedFormats() ([]*CameraSpec, error) {
	var count C.int
	formats := C.SDL_GetCameraSupportedFormats((C.SDL_CameraID)(id), &count)
	if formats == nil {
		return nil, getError()
	}
	result := make([]*CameraSpec, count)
	for i, k := range unsafe.Slice(formats, count) {
		result[i] = &CameraSpec{
			PixelFormat(k.format),
			Colorspace(k.colorspace),
			int(k.width),
			int(k.height),
			int(k.framerate_numerator),
			int(k.framerate_denominator),
		}
	}
	C.SDL_free(unsafe.Pointer(formats))
	return result, nil
}

// Get the human-readable device name for a camera.
//
// id: the camera device instance ID.
//
// Returns a human-readable device name or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetCameraName
func (id CameraID) Name() (string, error) {
	name := C.SDL_GetCameraName((C.SDL_CameraID)(id))
	if name == nil {
		return "", getError()
	}
	return C.GoString(name), nil
}

// Get the position of the camera in relation to the system.
//
// Most platforms will report [CameraPositionUnknown], but mobile devices, like
// phones, can often make a distinction between cameras on the front of the
// device (that points towards the user, for taking "selfies") and cameras on
// the back (for filming in the direction the user is facing).
//
// id: the camera device instance ID.
//
// Returns the position of the camera on the system hardware.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetCameraPosition
func (id CameraID) Position() CameraPosition {
	return (CameraPosition)(C.SDL_GetCameraPosition((C.SDL_CameraID)(id)))
}

// Open a video recording device (a "camera").
//
// You can open the device with any reasonable spec, and if the hardware can't
// directly support it, it will convert data seamlessly to the requested
// format. This might incur overhead, including scaling of image data.
//
// If you would rather accept whatever format the device offers, you can pass
// a nil spec here and it will choose one for you (and you can use
// [Surface]'s conversion/scaling functions directly if necessary).
//
// You can call [Camera.Format] to get the actual data format if passing
// a nil spec here. You can see the exact specs a device can support without
// conversion with [CameraID.SupportedFormats].
//
// SDL will not attempt to emulate framerate; it will try to set the hardware
// to the rate closest to the requested speed, but it won't attempt to limit
// or duplicate frames artificially; call [Camera.Format] to see the
// actual framerate of the opened the device, and check your timestamps if
// this is crucial to your app!
//
// Note that the camera is not usable until the user approves its use! On some
// platforms, the operating system will prompt the user to permit access to
// the camera, and they can choose Yes or No at that point. Until they do, the
// camera will not be usable. The app should either wait for an
// [EventCameraDeviceApproved] (or [EventCameraDeviceDenied]) event,
// or poll [Camera.PermissionState] occasionally until it returns
// non-zero. On platforms that don't require explicit user approval (and
// perhaps in places where the user previously permitted access), the approval
// event might come immediately, but it might come seconds, minutes, or hours
// later!
//
// id: the camera device instance ID.
//
// spec: the desired format for data the device will provide. Can be
// nil.
//
// Returns a [Camera] object or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_OpenCamera
func OpenCamera(id CameraID, spec *CameraSpec) (*Camera, error) {
	var cspec *C.SDL_CameraSpec
	if spec != nil {
		cspec = &C.SDL_CameraSpec{
			C.SDL_PixelFormat(spec.Format),
			C.SDL_Colorspace(spec.Colorspace),
			C.int(spec.Width),
			C.int(spec.Height),
			C.int(spec.FramerateNumerator),
			C.int(spec.FramerateDenominator),
		}
	}
	c := (*Camera)(C.SDL_OpenCamera((C.SDL_CameraID)(id), cspec))
	if c == nil {
		return nil, getError()
	}
	return c, nil
}

// Query if camera access has been approved by the user.
//
// Cameras will not function between when the device is opened by the app and
// when the user permits access to the hardware. On some platforms, this
// presents as a popup dialog where the user has to explicitly approve access;
// on others the approval might be implicit and not alert the user at all.
//
// This function can be used to check the status of that approval. It will
// return 0 if still waiting for user response, 1 if the camera is approved
// for use, and -1 if the user denied access.
//
// Instead of polling with this function, you can wait for a
// [EventCameraDeviceApproved] (or [EventCameraDeviceDenied]) event
// in the standard SDL event loop, which is guaranteed to be sent once when
// permission to use the camera is decided.
//
// If a camera is declined, there's nothing to be done but call
// [Camera.Close] to dispose of it.
//
// camera: the opened camera device to query.
//
// Returns -1 if user denied access to the camera, 1 if user approved access,
// 0 if no decision has been made yet.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetCameraPermissionState
func (camera *Camera) PermissionState() int {
	return (int)(C.SDL_GetCameraPermissionState((*C.SDL_Camera)(camera)))
}

// Get the instance ID of an opened camera.
//
// camera: a [Camera] to query.
//
// Returns the instance ID of the specified camera on success or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetCameraID
func (camera *Camera) ID() (CameraID, error) {
	id := (CameraID)(C.SDL_GetCameraID((*C.SDL_Camera)(camera)))
	if id == 0 {
		return 0, getError()
	}
	return id, nil
}

// Get the properties associated with an opened camera.
//
// camera: the [Camera] obtained from [OpenCamera]().
//
// Returns a valid property ID or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetCameraProperties
func (camera *Camera) Properties() (PropertiesID, error) {
	props := (PropertiesID)(C.SDL_GetCameraProperties((*C.SDL_Camera)(camera)))
	if props == 0 {
		return 0, getError()
	}
	return props, nil
}

// Get the spec that a camera is using when generating images.
//
// Note that this might not be the native format of the hardware, as SDL might
// be converting to this format behind the scenes.
//
// If the system is waiting for the user to approve access to the camera, as
// some platforms require, this will return an error, but this isn't necessarily
// a fatal error; you should either wait for an
// [EventCameraDeviceApproved] (or [EventCameraDeviceDenied]) event,
// or poll [Camera.PermissionState] occasionally until it returns
// non-zero.
//
// camera: opened camera device.
//
// Returns the [CameraSpec] or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetCameraFormat
func (camera *Camera) Format() (*CameraSpec, error) {
	var spec C.SDL_CameraSpec
	if !C.SDL_GetCameraFormat((*C.SDL_Camera)(camera), &spec) {
		return nil, getError()
	}
	return &CameraSpec{
		PixelFormat(spec.format),
		Colorspace(spec.colorspace),
		int(spec.width),
		int(spec.height),
		int(spec.framerate_numerator),
		int(spec.framerate_denominator),
	}, nil
}

// Acquire a frame.
//
// The frame is a memory pointer to the image data, whose size and format are
// given by the spec requested when opening the device.
//
// This is a non blocking API. If there is a frame available, a non-nil
// surface is returned, and a non-zero timestamp.
//
// Note that an error case can also return nil, but a nil by itself is
// normal and just signifies that a new frame is not yet available. Note that
// even if a camera device fails outright (a USB camera is unplugged while in
// use, etc), SDL will send an event separately to notify the app, but
// continue to provide blank frames at ongoing intervals until
// [Camera.Close] is called, so real failure here is almost always an out
// of memory condition.
//
// After use, the frame should be released with [Camera.ReleaseFrame]. If
// you don't do this, the system may stop providing more video!
//
// Do not call [Surface.Destroy] on the returned surface! It must be given
// back to the camera subsystem with [Camera.ReleaseFrame]!
//
// If the system is waiting for the user to approve access to the camera, as
// some platforms require, this will return nil (no frames available); you
// should either wait for an [EventCameraDeviceApproved] (or
// [EventCameraDeviceDenied]) event, or poll
// [Camera.PermissionState] occasionally until it returns non-zero.
//
// camera: opened camera device.
//
// timestampNS: the frame's timestamp, or 0 on error.
//
// Returns a new frame of video on success, nil if none is currently
// available.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AcquireCameraFrame
func (camera *Camera) AcquireFrame() (frame *Surface, timestampNS uint64) {
	s := C.SDL_AcquireCameraFrame((*C.SDL_Camera)(camera), (*C.Uint64)(&timestampNS))
	if s == nil {
		return nil, timestampNS
	}
	return &Surface{s, nil}, timestampNS
}

// Release a frame of video acquired from a camera.
//
// Let the back-end re-use the internal buffer for camera.
//
// This function _must_ be called only on surface objects returned by
// [Camera.AcquireFrame]. This function should be called as quickly as
// possible after acquisition, as SDL keeps a small FIFO queue of surfaces for
// video frames; if surfaces aren't released in a timely manner, SDL may drop
// upcoming video frames from the camera.
//
// If the app needs to keep the surface for a significant time, they should
// make a copy of it and release the original.
//
// The app should not use the surface again after calling this function;
// assume the surface is freed and the pointer is invalid.
//
// camera: opened camera device.
//
// frame: the video frame surface to release.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReleaseCameraFrame
func (camera *Camera) ReleaseFrame(frame *Surface) {
	C.SDL_ReleaseCameraFrame((*C.SDL_Camera)(camera), frame.internal)
}

// Use this function to shut down camera processing and close the camera
// device.
//
// camera: opened camera device.
//
// It is safe to call this function from any thread, but no
// thread may reference `device` once this function is called.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CloseCamera
func (camera *Camera) Close() {
	C.SDL_CloseCamera((*C.SDL_Camera)(camera))
}
