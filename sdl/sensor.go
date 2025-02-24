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
// #cgo noescape SDL_GetSensors
// #cgo nocallback SDL_GetSensors
// #cgo noescape SDL_GetSensorNameForID
// #cgo nocallback SDL_GetSensorNameForID
// #cgo noescape SDL_GetSensorTypeForID
// #cgo nocallback SDL_GetSensorTypeForID
// #cgo noescape SDL_GetSensorNonPortableTypeForID
// #cgo nocallback SDL_GetSensorNonPortableTypeForID
// #cgo noescape SDL_OpenSensor
// #cgo nocallback SDL_OpenSensor
// #cgo noescape SDL_GetSensorFromID
// #cgo nocallback SDL_GetSensorFromID
// #cgo noescape SDL_GetSensorProperties
// #cgo nocallback SDL_GetSensorProperties
// #cgo noescape SDL_GetSensorName
// #cgo nocallback SDL_GetSensorName
// #cgo noescape SDL_GetSensorType
// #cgo nocallback SDL_GetSensorType
// #cgo noescape SDL_GetSensorNonPortableType
// #cgo nocallback SDL_GetSensorNonPortableType
// #cgo noescape SDL_GetSensorID
// #cgo nocallback SDL_GetSensorID
// #cgo noescape SDL_GetSensorData
// #cgo nocallback SDL_GetSensorData
// #cgo noescape SDL_CloseSensor
// #cgo noescape SDL_UpdateSensors
// #cgo nocallback SDL_UpdateSensors
// #include <SDL3/SDL.h>
import "C"
import "unsafe"

// # CategorySensor
//
// SDL sensor management.
//
// These APIs grant access to gyros and accelerometers on various platforms.
//
// In order to use these functions, [Init] must have been called with the
// [InitSensor] flag. This causes SDL to scan the system for sensors, and
// load appropriate drivers.

// The opaque structure used to identify an opened SDL sensor.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Sensor
type Sensor C.struct_SDL_Sensor

// This is a unique ID for a sensor for the time it is connected to the
// system, and is never reused for the lifetime of the application.
//
// The value 0 is an invalid ID.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SensorID
type SensorID uint32

// A constant to represent standard gravity for accelerometer sensors.
//
// The accelerometer returns the current acceleration in SI meters per second
// squared. This measurement includes the force of gravity, so a device at
// rest will have an value of StandardGravity away from the center of the
// earth, which is a positive Y value.
//
// This macro is available since SDL 3.2.0.
const StandardGravity = 9.80665

// The different sensors defined by SDL.
//
// Additional sensors may be available, using platform dependent semantics.
//
// Here are the additional Android sensors:
//
// https://developer.android.com/reference/android/hardware/SensorEvent.html#values
//
// Accelerometer sensor notes:
//
// The accelerometer returns the current acceleration in SI meters per second
// squared. This measurement includes the force of gravity, so a device at
// rest will have an value of [StandardGravity] away from the center of the
// earth, which is a positive Y value.
//
//   - values[0]: Acceleration on the x axis
//   - values[1]: Acceleration on the y axis
//   - values[2]: Acceleration on the z axis
//
// For phones and tablets held in natural orientation and game controllers
// held in front of you, the axes are defined as follows:
//
//   - -X ... +X : left ... right
//   - -Y ... +Y : bottom ... top
//   - -Z ... +Z : farther ... closer
//
// The accelerometer axis data is not changed when the device is rotated.
//
// Gyroscope sensor notes:
//
// The gyroscope returns the current rate of rotation in radians per second.
// The rotation is positive in the counter-clockwise direction. That is, an
// observer looking from a positive location on one of the axes would see
// positive rotation on that axis when it appeared to be rotating
// counter-clockwise.
//
//   - values[0]: Angular speed around the x axis (pitch)
//   - values[1]: Angular speed around the y axis (yaw)
//   - values[2]: Angular speed around the z axis (roll)
//
// For phones and tablets held in natural orientation and game controllers
// held in front of you, the axes are defined as follows:
//
//   - -X ... +X : left ... right
//   - -Y ... +Y : bottom ... top
//   - -Z ... +Z : farther ... closer
//
// The gyroscope axis data is not changed when the device is rotated.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SensorType
type SensorType uint32

const (
	SensorUnknown SensorType = iota // Unknown sensor type
	SensorAccel                     // Accelerometer
	SensorGyro                      // Gyroscope
	SensorAccelL                    // Accelerometer for left Joy-Con controller and Wii nunchuk
	SensorGyroL                     // Gyroscope for left Joy-Con controller
	SensorAccelR                    // Accelerometer for right Joy-Con controller
	SensorGyroR                     // Gyroscope for right Joy-Con controller

	SensorInvalid SensorType = 0xFFFFFFFF
)

// Get a list of currently connected sensors.
//
// Returns a slice of sensor instance IDs or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSensors
func GetSensors() ([]SensorID, error) {
	var count C.int
	sensors := C.SDL_GetSensors(&count)
	if sensors == nil {
		return nil, getError()
	}
	result := make([]SensorID, count)
	for i, s := range unsafe.Slice(sensors, count) {
		result[i] = SensorID(s)
	}
	C.SDL_free(unsafe.Pointer(sensors))
	return result, nil
}

// Get the implementation dependent name of a sensor.
//
// This can be called before any sensors are opened.
//
// id: the sensor instance ID.
//
// Returns the sensor name, or an empty string if id is not valid.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSensorNameForID
func (id SensorID) Name() string {
	return C.GoString(C.SDL_GetSensorNameForID((C.SDL_SensorID)(id)))
}

// Get the type of a sensor.
//
// This can be called before any sensors are opened.
//
// id: the sensor instance ID.
//
// Returns the [SensorType], or [SensorInvalid] if id is
// not valid.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSensorTypeForID
func (id SensorID) Type() SensorType {
	return (SensorType)(C.SDL_GetSensorTypeForID((C.SDL_SensorID)(id)))
}

// Get the platform dependent type of a sensor.
//
// This can be called before any sensors are opened.
//
// id: the sensor instance ID.
//
// Returns the sensor platform dependent type, or -1 if id is not
// valid.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableTypeForID
func (id SensorID) NonPortableType() int {
	return (int)(C.SDL_GetSensorNonPortableTypeForID((C.SDL_SensorID)(id)))
}

// Open a sensor for use.
//
// id: the sensor instance ID.
//
// Returns a [Sensor] object or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_OpenSensor
func OpenSensor(id SensorID) (*Sensor, error) {
	s := (*Sensor)(C.SDL_OpenSensor((C.SDL_SensorID)(id)))
	if s == nil {
		return nil, getError()
	}
	return s, nil
}

// Return the [Sensor] associated with an instance ID.
//
// id: the sensor instance ID.
//
// Returns a [Sensor] object or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSensorFromID
func GetSensorFromID(id SensorID) (*Sensor, error) {
	s := (*Sensor)(C.SDL_GetSensorFromID((C.SDL_SensorID)(id)))
	if s == nil {
		return nil, getError()
	}
	return s, nil
}

// Get the properties associated with a sensor.
//
// sensor: the [Sensor] object.
//
// Returns a valid property ID or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSensorProperties
func (sensor *Sensor) Properties() (PropertiesID, error) {
	props := (PropertiesID)(C.SDL_GetSensorProperties((*C.SDL_Sensor)(sensor)))
	if props == 0 {
		return 0, getError()
	}
	return props, nil
}

// Get the implementation dependent name of a sensor.
//
// sensor: the [Sensor] object.
//
// Returns the sensor name or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSensorName
func (sensor *Sensor) Name() (string, error) {
	name := C.SDL_GetSensorName((*C.SDL_Sensor)(sensor))
	if name == nil {
		return "", getError()
	}
	return C.GoString(name), nil
}

// Get the type of a sensor.
//
// sensor: the [Sensor] object to inspect.
//
// Returns the [SensorType] type, or [SensorInvalid] if sensor is nil.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSensorType
func (sensor *Sensor) Type() SensorType {
	return (SensorType)(C.SDL_GetSensorType((*C.SDL_Sensor)(sensor)))
}

// Get the platform dependent type of a sensor.
//
// sensor: the [Sensor] object to inspect.
//
// Returns the sensor platform dependent type, or -1 if sensor is nil.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableType
func (sensor *Sensor) NonPortableType() int {
	return (int)(C.SDL_GetSensorNonPortableType((*C.SDL_Sensor)(sensor)))
}

// Get the instance ID of a sensor.
//
// sensor: the [Sensor] object to inspect.
//
// Returns the sensor instance ID or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSensorID
func (sensor *Sensor) ID() (SensorID, error) {
	id := (SensorID)(C.SDL_GetSensorID((*C.SDL_Sensor)(sensor)))
	if id == 0 {
		return 0, getError()
	}
	return id, nil
}

// Get the current state of an opened sensor.
//
// The number of values and interpretation of the data is sensor dependent.
//
// sensor: the [Sensor] object to query.
//
// data: a pointer filled with the current sensor state.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSensorData
func (sensor *Sensor) Read(data []float32) error {
	if !C.SDL_GetSensorData((*C.SDL_Sensor)(sensor), (*C.float)(unsafe.SliceData(data)), (C.int)(len(data))) {
		return getError()
	}
	return nil
}

// Close a sensor previously opened with [OpenSensor].
//
// sensor: the [Sensor] object to close.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CloseSensor
func (sensor *Sensor) Close() {
	C.SDL_CloseSensor((*C.SDL_Sensor)(sensor))
}

// Update the current state of the open sensors.
//
// This is called automatically by the event loop if sensor events are
// enabled.
//
// This needs to be called from the thread that initialized the sensor
// subsystem.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UpdateSensors
func UpdateSensors() {
	C.SDL_UpdateSensors()
}
