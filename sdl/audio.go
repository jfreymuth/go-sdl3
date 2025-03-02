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
// #cgo noescape SDL_GetNumAudioDrivers
// #cgo nocallback SDL_GetNumAudioDrivers
// #cgo noescape SDL_GetAudioDriver
// #cgo nocallback SDL_GetAudioDriver
// #cgo noescape SDL_GetCurrentAudioDriver
// #cgo nocallback SDL_GetCurrentAudioDriver
// #cgo noescape SDL_GetAudioPlaybackDevices
// #cgo nocallback SDL_GetAudioPlaybackDevices
// #cgo noescape SDL_GetAudioRecordingDevices
// #cgo nocallback SDL_GetAudioRecordingDevices
// #cgo noescape SDL_GetAudioDeviceName
// #cgo nocallback SDL_GetAudioDeviceName
// #cgo noescape SDL_GetAudioDeviceFormat
// #cgo nocallback SDL_GetAudioDeviceFormat
// #cgo noescape SDL_GetAudioDeviceChannelMap
// #cgo nocallback SDL_GetAudioDeviceChannelMap
// #cgo noescape SDL_OpenAudioDevice
// #cgo nocallback SDL_OpenAudioDevice
// #cgo noescape SDL_IsAudioDevicePhysical
// #cgo nocallback SDL_IsAudioDevicePhysical
// #cgo noescape SDL_IsAudioDevicePlayback
// #cgo nocallback SDL_IsAudioDevicePlayback
// #cgo noescape SDL_PauseAudioDevice
// #cgo nocallback SDL_PauseAudioDevice
// #cgo noescape SDL_ResumeAudioDevice
// #cgo nocallback SDL_ResumeAudioDevice
// #cgo noescape SDL_AudioDevicePaused
// #cgo nocallback SDL_AudioDevicePaused
// #cgo noescape SDL_GetAudioDeviceGain
// #cgo nocallback SDL_GetAudioDeviceGain
// #cgo noescape SDL_SetAudioDeviceGain
// #cgo nocallback SDL_SetAudioDeviceGain
// #cgo noescape SDL_CloseAudioDevice
// #cgo noescape SDL_BindAudioStreams
// #cgo nocallback SDL_BindAudioStreams
// #cgo noescape SDL_BindAudioStream
// #cgo nocallback SDL_BindAudioStream
// #cgo noescape SDL_UnbindAudioStreams
// #cgo nocallback SDL_UnbindAudioStreams
// #cgo noescape SDL_UnbindAudioStream
// #cgo nocallback SDL_UnbindAudioStream
// #cgo noescape SDL_GetAudioStreamDevice
// #cgo nocallback SDL_GetAudioStreamDevice
// #cgo noescape SDL_CreateAudioStream
// #cgo nocallback SDL_CreateAudioStream
// #cgo noescape SDL_GetAudioStreamProperties
// #cgo nocallback SDL_GetAudioStreamProperties
// #cgo noescape SDL_GetAudioStreamFormat
// #cgo nocallback SDL_GetAudioStreamFormat
// #cgo noescape SDL_SetAudioStreamFormat
// #cgo nocallback SDL_SetAudioStreamFormat
// #cgo noescape SDL_GetAudioStreamFrequencyRatio
// #cgo nocallback SDL_GetAudioStreamFrequencyRatio
// #cgo noescape SDL_SetAudioStreamFrequencyRatio
// #cgo nocallback SDL_SetAudioStreamFrequencyRatio
// #cgo noescape SDL_GetAudioStreamGain
// #cgo nocallback SDL_GetAudioStreamGain
// #cgo noescape SDL_SetAudioStreamGain
// #cgo nocallback SDL_SetAudioStreamGain
// #cgo noescape SDL_GetAudioStreamInputChannelMap
// #cgo nocallback SDL_GetAudioStreamInputChannelMap
// #cgo noescape SDL_GetAudioStreamOutputChannelMap
// #cgo nocallback SDL_GetAudioStreamOutputChannelMap
// #cgo noescape SDL_SetAudioStreamInputChannelMap
// #cgo nocallback SDL_SetAudioStreamInputChannelMap
// #cgo noescape SDL_SetAudioStreamOutputChannelMap
// #cgo nocallback SDL_SetAudioStreamOutputChannelMap
// #cgo noescape SDL_PutAudioStreamData
// #cgo noescape SDL_GetAudioStreamData
// #cgo noescape SDL_GetAudioStreamAvailable
// #cgo nocallback SDL_GetAudioStreamAvailable
// #cgo noescape SDL_GetAudioStreamQueued
// #cgo nocallback SDL_GetAudioStreamQueued
// #cgo noescape SDL_FlushAudioStream
// #cgo nocallback SDL_FlushAudioStream
// #cgo noescape SDL_ClearAudioStream
// #cgo nocallback SDL_ClearAudioStream
// #cgo noescape SDL_PauseAudioStreamDevice
// #cgo nocallback SDL_PauseAudioStreamDevice
// #cgo noescape SDL_ResumeAudioStreamDevice
// #cgo nocallback SDL_ResumeAudioStreamDevice
// #cgo noescape SDL_AudioStreamDevicePaused
// #cgo nocallback SDL_AudioStreamDevicePaused
// #cgo noescape SDL_LockAudioStream
// #cgo nocallback SDL_LockAudioStream
// #cgo noescape SDL_UnlockAudioStream
// #cgo nocallback SDL_UnlockAudioStream
// #cgo noescape SDL_DestroyAudioStream
// #cgo noescape SDL_LoadWAV
// #cgo nocallback SDL_LoadWAV
// #cgo noescape SDL_LoadWAV_IO
// #cgo noescape SDL_MixAudio
// #cgo nocallback SDL_MixAudio
// #cgo noescape SDL_ConvertAudioSamples
// #cgo nocallback SDL_ConvertAudioSamples
// #cgo noescape SDL_GetAudioFormatName
// #cgo nocallback SDL_GetAudioFormatName
// #cgo noescape SDL_GetSilenceValueForFormat
// #cgo nocallback SDL_GetSilenceValueForFormat
// #include <SDL3/SDL.h>
// extern bool wrap_SDL_SetAudioStreamPutCallback(SDL_AudioStream *stream, uintptr_t userdata);
// extern bool wrap_SDL_SetAudioStreamGetCallback(SDL_AudioStream *stream, uintptr_t userdata);
// extern SDL_AudioStream *wrap_SDL_OpenAudioDeviceStream(SDL_AudioDeviceID devid, const SDL_AudioSpec *spec, uintptr_t userdata);
// extern bool wrap_SDL_SetAudioPostmixCallback(SDL_AudioDeviceID devid, uintptr_t userdata);
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

// # CategoryAudio
//
// Audio functionality for the SDL library.
//
// All audio in SDL3 revolves around [AudioStream]. Whether you want to play
// or record audio, convert it, stream it, buffer it, or mix it, you're going
// to be passing it through an audio stream.
//
// Audio streams are quite flexible; they can accept any amount of data at a
// time, in any supported format, and output it as needed in any other format,
// even if the data format changes on either side halfway through.
//
// An app opens an audio device and binds any number of audio streams to it,
// feeding more data to the streams as available. When the device needs more
// data, it will pull it from all bound streams and mix them together for
// playback.
//
// Audio streams can also use an app-provided callback to supply data
// on-demand, which maps pretty closely to the SDL2 audio model.
//
// SDL also provides a simple .WAV loader in [LoadWAV] (and [LoadWAV_IO]
// if you aren't reading from a file) as a basic means to load sound data into
// your program.
//
// ## Logical audio devices
//
// In SDL3, opening a physical device (like a SoundBlaster 16 Pro) gives you a
// logical device ID that you can bind audio streams to. In almost all cases,
// logical devices can be used anywhere in the API that a physical device is
// normally used. However, since each device opening generates a new logical
// device, different parts of the program (say, a VoIP library, or
// text-to-speech framework, or maybe some other sort of mixer on top of SDL)
// can have their own device opens that do not interfere with each other; each
// logical device will mix its separate audio down to a single buffer, fed to
// the physical device, behind the scenes. As many logical devices as you like
// can come and go; SDL will only have to open the physical device at the OS
// level once, and will manage all the logical devices on top of it
// internally.
//
// One other benefit of logical devices: if you don't open a specific physical
// device, instead opting for the default, SDL can automatically migrate those
// logical devices to different hardware as circumstances change: a user
// plugged in headphones? The system default changed? SDL can transparently
// migrate the logical devices to the correct physical device seamlessly and
// keep playing; the app doesn't even have to know it happened if it doesn't
// want to.
//
// ## Simplified audio
//
// As a simplified model for when a single source of audio is all that's
// needed, an app can use [OpenAudioDeviceStream], which is a single
// function to open an audio device, create an audio stream, bind that stream
// to the newly-opened device, and (optionally) provide a callback for
// obtaining audio data. When using this function, the primary interface is
// the [AudioStream] and the device handle is mostly hidden away; destroying
// a stream created through this function will also close the device, stream
// bindings cannot be changed, etc. One other quirk of this is that the device
// is started in a _paused_ state and must be explicitly resumed; this is
// partially to offer a clean migration for SDL2 apps and partially because
// the app might have to do more setup before playback begins; in the
// non-simplified form, nothing will play until a stream is bound to a device,
// so they start _unpaused_.
//
// ## Channel layouts
//
// Audio data passing through SDL is uncompressed PCM data, interleaved. One
// can provide their own decompression through an MP3, etc, decoder, but SDL
// does not provide this directly. Each interleaved channel of data is meant
// to be in a specific order.
//
// Abbreviations:
//
//   - FRONT = single mono speaker
//   - FL = front left speaker
//   - FR = front right speaker
//   - FC = front center speaker
//   - BL = back left speaker
//   - BR = back right speaker
//   - SR = surround right speaker
//   - SL = surround left speaker
//   - BC = back center speaker
//   - LFE = low-frequency speaker
//
// These are listed in the order they are laid out in memory, so "FL, FR"
// means "the front left speaker is laid out in memory first, then the front
// right, then it repeats for the next audio frame".
//
//   - 1 channel (mono) layout: FRONT
//   - 2 channels (stereo) layout: FL, FR
//   - 3 channels (2.1) layout: FL, FR, LFE
//   - 4 channels (quad) layout: FL, FR, BL, BR
//   - 5 channels (4.1) layout: FL, FR, LFE, BL, BR
//   - 6 channels (5.1) layout: FL, FR, FC, LFE, BL, BR (last two can also be
//     SL, SR)
//   - 7 channels (6.1) layout: FL, FR, FC, LFE, BC, SL, SR
//   - 8 channels (7.1) layout: FL, FR, FC, LFE, BL, BR, SL, SR
//
// This is the same order as DirectSound expects, but applied to all
// platforms; SDL will swizzle the channels as necessary if a platform expects
// something different.
//
// [AudioStream] can also be provided channel maps to change this ordering
// to whatever is necessary, in other audio processing scenarios.

// Mask of bits in an [AudioFormat] that contains the format bit size.
//
// Generally one should use [AudioFormat.BitSize] instead of this macro directly.
//
// This macro is available since SDL 3.2.0.
const AudioMaskBitSize AudioFormat = 0xFF

// Mask of bits in an [AudioFormat] that contain the floating point flag.
//
// Generally one should use [AudioFormat.Float] instead of this macro directly.
//
// This macro is available since SDL 3.2.0.
const AudioMaskFloat AudioFormat = 1 << 8

// Mask of bits in an [AudioFormat] that contain the bigendian flag.
//
// Generally one should use [AudioFormat.BigEndian] or [AudioFormat.LittleEndian]
// instead of this macro directly.
//
// This macro is available since SDL 3.2.0.
const AudioMaskBigEndian AudioFormat = 1 << 12

// Mask of bits in an [AudioFormat] that contain the signed data flag.
//
// Generally one should use [AudioFormat.Signed] instead of this macro directly.
//
// This macro is available since SDL 3.2.0.
const AudioMaskSigned AudioFormat = 1 << 15

// Define an [AudioFormat] value.
//
// SDL does not support custom audio formats, so this macro is not of much use
// externally, but it can be illustrative as to what the various bits of an
// [AudioFormat] mean.
//
// For example, [AudioS32LE] looks like this:
//
//	DefineAudioFormat(true, false, false, 32)
//
// signed: 1 for signed data, 0 for unsigned data.
//
// bigendian: 1 for bigendian data, 0 for littleendian data.
//
// flt: 1 for floating point data, 0 for integer data.
//
// size: number of bits per sample.
//
// Returns a format value in the style of [AudioFormat].
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DEFINE_AUDIO_FORMAT
func DefineAudioFormat(signed, bigendian, float bool, size int) AudioFormat {
	f := AudioFormat(size) & AudioMaskBitSize
	if signed {
		f |= AudioMaskSigned
	}
	if bigendian {
		f |= AudioMaskBigEndian
	}
	if float {
		f |= AudioMaskFloat
	}
	return f
}

// Audio format.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AudioFormat
type AudioFormat uint32

const (
	AudioUnknown AudioFormat = 0x0000 //  Unspecified audio format
	AudioU8      AudioFormat = 0x0008 //  Unsigned 8-bit samples
	AudioS8      AudioFormat = 0x8008 //  Signed 8-bit samples
	AudioS16LE   AudioFormat = 0x8010 //  Signed 16-bit samples
	AudioS16BE   AudioFormat = 0x9010 //  As above, but big-endian byte order
	AudioS32LE   AudioFormat = 0x8020 //  32-bit integer samples
	AudioS32BE   AudioFormat = 0x9020 //  As above, but big-endian byte order
	AudioF32LE   AudioFormat = 0x8120 //  32-bit floating point samples
	AudioF32BE   AudioFormat = 0x9120 //  As above, but big-endian byte order
)

// Retrieve the size, in bits, from an [AudioFormat].
//
// For example, [AudioS16].BitSize() returns 16.
//
// x: an [AudioFormat] value.
//
// Returns data size in bits.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AUDIO_BITSIZE
func (a AudioFormat) BitSize() int {
	return int(a & AudioMaskBitSize)
}

// Retrieve the size, in bytes, from an [AudioFormat].
//
// For example, [AudioS16].ByteSize() returns 2.
//
// x: an [AudioFormat] value.
//
// Returns data size in bytes.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AUDIO_BYTESIZE
func (a AudioFormat) ByteSize() int {
	return a.BitSize() / 8
}

// Determine if an [AudioFormat] represents floating point data.
//
// For example, [AudioS16].Float() returns false.
//
// x: an [AudioFormat] value.
//
// Returns non-zero if format is floating point, zero otherwise.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AUDIO_ISFLOAT
func (a AudioFormat) Float() bool {
	return a&AudioMaskFloat != 0
}

// Determine if an [AudioFormat] represents bigendian data.
//
// For example, [AudioS16LE].BigEndian() returns false.
//
// x: an [AudioFormat] value.
//
// Returns non-zero if format is bigendian, zero otherwise.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AUDIO_ISBIGENDIAN
func (a AudioFormat) BigEndian() bool {
	return a&AudioMaskBigEndian != 0
}

// Determine if an [AudioFormat] represents littleendian data.
//
// For example, [AudioS16BE].LittleEndian() returns false.
//
// x: an [AudioFormat] value.
//
// Returns non-zero if format is littleendian, zero otherwise.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AUDIO_ISLITTLEENDIAN
func (a AudioFormat) LittleEndian() bool {
	return a&AudioMaskBigEndian == 0
}

// Determine if an [AudioFormat] represents signed data.
//
// For example, [AudioU8].Signed() returns false.
//
// x: an [AudioFormat] value.
//
// Returns non-zero if format is signed, zero otherwise.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AUDIO_ISSIGNED
func (a AudioFormat) Signed() bool {
	return a&AudioMaskSigned != 0
}

// Determine if an [AudioFormat] represents integer data.
//
// For example, [AudioF32].Int() returns false.
//
// x: an [AudioFormat] value.
//
// Returns non-zero if format is integer, zero otherwise.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AUDIO_ISINT
func (a AudioFormat) Int() bool {
	return a&AudioMaskFloat == 0
}

// Determine if an [AudioFormat] represents unsigned data.
//
// For example, [AudioS16].Unsigned() returns false.
//
// x: an [AudioFormat] value.
//
// Returns non-zero if format is unsigned, zero otherwise.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AUDIO_ISUNSIGNED
func (a AudioFormat) Unsigned() bool {
	return a&AudioMaskSigned == 0
}

// SDL Audio Device instance IDs.
//
// Zero is used to signify an invalid/null device.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AudioDeviceID
type AudioDeviceID uint32

// A value used to request a default playback audio device.
//
// Several functions that require an [AudioDeviceID] will accept this value
// to signify the app just wants the system to choose a default device instead
// of the app providing a specific one.
//
// This macro is available since SDL 3.2.0.
const AudioDeviceDefaultPlayback AudioDeviceID = 0xFFFFFFFF

// A value used to request a default recording audio device.
//
// Several functions that require an [AudioDeviceID] will accept this value
// to signify the app just wants the system to choose a default device instead
// of the app providing a specific one.
//
// This macro is available since SDL 3.2.0.
const AudioDeviceDefaultRecording AudioDeviceID = 0xFFFFFFFE

// Format specifier for audio data.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AudioSpec
type AudioSpec struct {
	Format   AudioFormat //  Audio data format
	Channels int         //  Number of channels: 1 mono, 2 stereo, etc
	Freq     int         //  sample rate: sample frames per second
}

// Calculate the size of each audio frame (in bytes) from an [AudioSpec].
//
// This reports on the size of an audio sample frame: stereo Sint16 data (2
// channels of 2 bytes each) would be 4 bytes per frame, for example.
//
// x: an [AudioSpec] to query.
//
// Returns the number of bytes used per sample frame.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AUDIO_FRAMESIZE
func (a AudioSpec) Framesize() int {
	return a.Format.ByteSize() * a.Channels
}

// The opaque handle that represents an audio stream.
//
// [AudioStream] is an audio conversion interface.
//
//   - It can handle resampling data in chunks without generating artifacts,
//     when it doesn't have the complete buffer available.
//   - It can handle incoming data in any variable size.
//   - It can handle input/output format changes on the fly.
//   - It can remap audio channels between inputs and outputs.
//   - You push data as you have it, and pull it when you need it
//   - It can also function as a basic audio data queue even if you just have
//     sound that needs to pass from one place to another.
//   - You can hook callbacks up to them when more data is added or requested,
//     to manage data on-the-fly.
//
// Audio streams are the core of the SDL3 audio interface. You create one or
// more of them, bind them to an opened audio device, and feed data to them
// (or for recording, consume data from them).
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AudioStream
type AudioStream C.struct_SDL_AudioStream

// Use this function to get the number of built-in audio drivers.
//
// This function returns a hardcoded number. This never returns a negative
// value; if there are no drivers compiled into this build of SDL, this
// function returns zero. The presence of a driver in this list does not mean
// it will function, it just means SDL is capable of interacting with that
// interface. For example, a build of SDL might have esound support, but if
// there's no esound server available, SDL's esound driver would fail if used.
//
// By default, SDL tries all drivers, in its preferred order, until one is
// found to be usable.
//
// Returns the number of built-in audio drivers.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetNumAudioDrivers
func GetNumAudioDrivers() int {
	return (int)(C.SDL_GetNumAudioDrivers())
}

// Use this function to get the name of a built in audio driver.
//
// The list of audio drivers is given in the order that they are normally
// initialized by default; the drivers that seem more reasonable to choose
// first (as far as the SDL developers believe) are earlier in the list.
//
// The names of drivers are all simple, low-ASCII identifiers, like "alsa",
// "coreaudio" or "wasapi". These never have Unicode characters, and are not
// meant to be proper names.
//
// index: the index of the audio driver; the value ranges from 0 to
// [GetNumAudioDrivers] - 1.
//
// Returns the name of the audio driver at the requested index, or an empty
// string if an invalid index was specified.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioDriver
func GetAudioDriver(index int) string {
	return C.GoString(C.SDL_GetAudioDriver((C.int)(index)))
}

// Get the name of the current audio driver.
//
// The names of drivers are all simple, low-ASCII identifiers, like "alsa",
// "coreaudio" or "wasapi". These never have Unicode characters, and are not
// meant to be proper names.
//
// Returns the name of the current audio driver or an empty string if no driver
// has been initialized.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetCurrentAudioDriver
func GetCurrentAudioDriver() string {
	return C.GoString(C.SDL_GetCurrentAudioDriver())
}

// Get a list of currently-connected audio playback devices.
//
// This returns of list of available devices that play sound, perhaps to
// speakers or headphones ("playback" devices). If you want devices that
// record audio, like a microphone ("recording" devices), use
// [GetAudioRecordingDevices] instead.
//
// This only returns a list of physical devices; it will not have any device
// IDs returned by [OpenAudioDevice].
//
// Returns a 0 terminated array of device instance IDs or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioPlaybackDevices
func GetAudioPlaybackDevices() ([]AudioDeviceID, error) {
	var count C.int
	ptr := C.SDL_GetAudioPlaybackDevices(&count)
	if ptr == nil {
		return nil, getError()
	}
	slice := unsafe.Slice(ptr, count)
	result := make([]AudioDeviceID, count)
	for i, a := range slice {
		result[i] = AudioDeviceID(a)
	}
	C.SDL_free(unsafe.Pointer(ptr))
	return result, nil
}

// Get a list of currently-connected audio recording devices.
//
// This returns of list of available devices that record audio, like a
// microphone ("recording" devices). If you want devices that play sound,
// perhaps to speakers or headphones ("playback" devices), use
// [GetAudioPlaybackDevices] instead.
//
// This only returns a list of physical devices; it will not have any device
// IDs returned by [OpenAudioDevice].
//
// Returns a 0 terminated array of device instance IDs, or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioRecordingDevices
func GetAudioRecordingDevices() ([]AudioDeviceID, error) {
	var count C.int
	ptr := C.SDL_GetAudioRecordingDevices(&count)
	if ptr == nil {
		return nil, getError()
	}
	slice := unsafe.Slice(ptr, count)
	result := make([]AudioDeviceID, count)
	for i, a := range slice {
		result[i] = AudioDeviceID(a)
	}
	C.SDL_free(unsafe.Pointer(ptr))
	return result, nil
}

// Get the human-readable name of a specific audio device.
//
// devid: the instance ID of the device to query.
//
// Returns the name of the audio device, or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceName
func (devid AudioDeviceID) Name() string {
	return C.GoString(C.SDL_GetAudioDeviceName((C.SDL_AudioDeviceID)(devid)))
}

// Get the current audio format of a specific audio device.
//
// For an opened device, this will report the format the device is currently
// using. If the device isn't yet opened, this will report the device's
// preferred format (or a reasonable default if this can't be determined).
//
// You may also specify [AudioDeviceDefaultPlayback] or
// [AudioDeviceDefaultRecording] here, which is useful for getting a
// reasonable recommendation before opening the system-recommended default
// device.
//
// You can also use this to request the current device buffer size. This is
// specified in sample frames and represents the amount of data SDL will feed
// to the physical hardware in each chunk. This can be converted to
// milliseconds of audio with the following equation:
//
//	ms = (int) ((((Sint64) frames) * 1000) / spec.freq);
//
// Buffer size is only important if you need low-level control over the audio
// playback timing. Most apps do not need this.
//
// devid: the instance ID of the device to query.
//
// spec: on return, will be filled with device details.
//
// sampleFrames: device buffer size, in sample frames.
//
// Returns the device details and the device buffer size, in sample frames,
// or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceFormat
func (devid AudioDeviceID) Format() (spec AudioSpec, sampleFrames int, err error) {
	var s C.SDL_AudioSpec
	var sf C.int
	if !C.SDL_GetAudioDeviceFormat((C.SDL_AudioDeviceID)(devid), &s, &sf) {
		return AudioSpec{}, 0, getError()
	}
	return AudioSpec{AudioFormat(s.format), int(s.channels), int(s.freq)}, int(sf), nil
}

// Get the current channel map of an audio device.
//
// Channel maps are optional; most things do not need them, instead passing
// data in the [order that SDL expects](CategoryAudio#channel-layouts).
//
// Audio devices usually have no remapping applied. This is represented by
// returning nil, and does not signify an error.
//
// devid: the instance ID of the device to query.
//
// Returns an array of the current channel mapping, with as many elements as
// the current output spec's channels, or nil if default.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceChannelMap
func (devid AudioDeviceID) ChannelMap() []int {
	var count C.int
	ptr := C.SDL_GetAudioDeviceChannelMap((C.SDL_AudioDeviceID)(devid), &count)
	if ptr == nil {
		return nil
	}
	slice := unsafe.Slice(ptr, count)
	result := make([]int, count)
	for i, c := range slice {
		result[i] = int(c)
	}
	C.SDL_free(unsafe.Pointer(ptr))
	return result
}

// Open a specific audio device.
//
// You can open both playback and recording devices through this function.
// Playback devices will take data from bound audio streams, mix it, and send
// it to the hardware. Recording devices will feed any bound audio streams
// with a copy of any incoming data.
//
// An opened audio device starts out with no audio streams bound. To start
// audio playing, bind a stream and supply audio data to it. Unlike SDL2,
// there is no audio callback; you only bind audio streams and make sure they
// have data flowing into them (however, you can simulate SDL2's semantics
// fairly closely by using [OpenAudioDeviceStream] instead of this
// function).
//
// If you don't care about opening a specific device, pass a devid of either
// [AudioDeviceDefaultPlayback] or
// [AudioDeviceDefaultRecording]. In this case, SDL will try to pick
// the most reasonable default, and may also switch between physical devices
// seamlessly later, if the most reasonable default changes during the
// lifetime of this opened device (user changed the default in the OS's system
// preferences, the default got unplugged so the system jumped to a new
// default, the user plugged in headphones on a mobile device, etc). Unless
// you have a good reason to choose a specific device, this is probably what
// you want.
//
// You may request a specific format for the audio device, but there is no
// promise the device will honor that request for several reasons. As such,
// it's only meant to be a hint as to what data your app will provide. Audio
// streams will accept data in whatever format you specify and manage
// conversion for you as appropriate. [AudioDeviceID.Format] can tell you
// the preferred format for the device before opening and the actual format
// the device is using after opening.
//
// It's legal to open the same device ID more than once; each successful open
// will generate a new logical [AudioDeviceID] that is managed separately
// from others on the same physical device. This allows libraries to open a
// device separately from the main app and bind its own streams without
// conflicting.
//
// It is also legal to open a device ID returned by a previous call to this
// function; doing so just creates another logical device on the same physical
// device. This may be useful for making logical groupings of audio streams.
//
// This function returns the opened device ID on success. This is a new,
// unique [AudioDeviceID] that represents a logical device.
//
// Some backends might offer arbitrary devices (for example, a networked audio
// protocol that can connect to an arbitrary server). For these, as a change
// from SDL2, you should open a default device ID and use an SDL hint to
// specify the target if you care, or otherwise let the backend figure out a
// reasonable default. Most backends don't offer anything like this, and often
// this would be an end user setting an environment variable for their custom
// need, and not something an application should specifically manage.
//
// When done with an audio device, possibly at the end of the app's life, one
// should call [AudioDeviceID.Close] on the returned device id.
//
// devid: the device instance id to open, or
// [AudioDeviceDefaultPlayback] or
// [AudioDeviceDefaultRecording] for the most reasonable
// default device.
//
// spec: the requested device configuration. Can be nil to use
// reasonable defaults.
//
// Returns the device ID or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_OpenAudioDevice
func OpenAudioDevice(devid AudioDeviceID, spec *AudioSpec) (AudioDeviceID, error) {
	var cspec *C.SDL_AudioSpec
	if spec != nil {
		cspec = &C.SDL_AudioSpec{C.SDL_AudioFormat(spec.Format), C.int(spec.Channels), C.int(spec.Freq)}
	}
	id := (AudioDeviceID)(C.SDL_OpenAudioDevice((C.SDL_AudioDeviceID)(devid), cspec))
	if id == 0 {
		return 0, getError()
	}
	return id, nil
}

// Determine if an audio device is physical (instead of logical).
//
// An [AudioDeviceID] that represents physical hardware is a physical
// device; there is one for each piece of hardware that SDL can see. Logical
// devices are created by calling [OpenAudioDevice] or
// [OpenAudioDeviceStream], and while each is associated with a physical
// device, there can be any number of logical devices on one physical device.
//
// For the most part, logical and physical IDs are interchangeable--if you try
// to open a logical device, SDL understands to assign that effort to the
// underlying physical device, etc. However, it might be useful to know if an
// arbitrary device ID is physical or logical. This function reports which.
//
// This function may return either true or false for invalid device IDs.
//
// devid: the device ID to query.
//
// Returns true if devid is a physical device, false if it is logical.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_IsAudioDevicePhysical
func (devid AudioDeviceID) Physical() bool {
	return (bool)(C.SDL_IsAudioDevicePhysical((C.SDL_AudioDeviceID)(devid)))
}

// Determine if an audio device is a playback device (instead of recording).
//
// This function may return either true or false for invalid device IDs.
//
// devid: the device ID to query.
//
// Returns true if devid is a playback device, false if it is recording.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_IsAudioDevicePlayback
func (devid AudioDeviceID) Playback() bool {
	return (bool)(C.SDL_IsAudioDevicePlayback((C.SDL_AudioDeviceID)(devid)))
}

// Use this function to pause audio playback on a specified device.
//
// This function pauses audio processing for a given device. Any bound audio
// streams will not progress, and no audio will be generated. Pausing one
// device does not prevent other unpaused devices from running.
//
// Unlike in SDL2, audio devices start in an _unpaused_ state, since an app
// has to bind a stream before any audio will flow. Pausing a paused device is
// a legal no-op.
//
// Pausing a device can be useful to halt all audio without unbinding all the
// audio streams. This might be useful while a game is paused, or a level is
// loading, etc.
//
// Physical devices can not be paused or unpaused, only logical devices
// created through [OpenAudioDevice] can be.
//
// dev: a device opened by [OpenAudioDevice].
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PauseAudioDevice
func (dev AudioDeviceID) Pause() error {
	if !C.SDL_PauseAudioDevice((C.SDL_AudioDeviceID)(dev)) {
		return getError()
	}
	return nil
}

// Use this function to unpause audio playback on a specified device.
//
// This function unpauses audio processing for a given device that has
// previously been paused with [AudioDevice.Pause]. Once unpaused, any
// bound audio streams will begin to progress again, and audio can be
// generated.
//
// Unlike in SDL2, audio devices start in an _unpaused_ state, since an app
// has to bind a stream before any audio will flow. Unpausing an unpaused
// device is a legal no-op.
//
// Physical devices can not be paused or unpaused, only logical devices
// created through [OpenAudioDevice] can be.
//
// dev: a device opened by [OpenAudioDevice].
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ResumeAudioDevice
func (dev AudioDeviceID) Resume() error {
	if !C.SDL_ResumeAudioDevice((C.SDL_AudioDeviceID)(dev)) {
		return getError()
	}
	return nil
}

// Use this function to query if an audio device is paused.
//
// Unlike in SDL2, audio devices start in an _unpaused_ state, since an app
// has to bind a stream before any audio will flow.
//
// Physical devices can not be paused or unpaused, only logical devices
// created through [OpenAudioDevice] can be. Physical and invalid device
// IDs will report themselves as unpaused here.
//
// dev: a device opened by [OpenAudioDevice].
//
// Returns true if device is valid and paused, false otherwise.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AudioDevicePaused
func (dev AudioDeviceID) Paused() bool {
	return (bool)(C.SDL_AudioDevicePaused((C.SDL_AudioDeviceID)(dev)))
}

// Get the gain of an audio device.
//
// The gain of a device is its volume; a larger gain means a louder output,
// with a gain of zero being silence.
//
// Audio devices default to a gain of 1.0f (no change in output).
//
// Physical devices may not have their gain changed, only logical devices, and
// this function will always return -1.0f when used on physical devices.
//
// devid: the audio device to query.
//
// Returns the gain of the device or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceGain
func (devid AudioDeviceID) Gain() (float32, error) {
	gain := (float32)(C.SDL_GetAudioDeviceGain((C.SDL_AudioDeviceID)(devid)))
	if gain == -1 {
		return -1, getError()
	}
	return gain, nil
}

// Change the gain of an audio device.
//
// The gain of a device is its volume; a larger gain means a louder output,
// with a gain of zero being silence.
//
// Audio devices default to a gain of 1.0f (no change in output).
//
// Physical devices may not have their gain changed, only logical devices, and
// this function will always return false when used on physical devices. While
// it might seem attractive to adjust several logical devices at once in this
// way, it would allow an app or library to interfere with another portion of
// the program's otherwise-isolated devices.
//
// This is applied, along with any per-audiostream gain, during playback to
// the hardware, and can be continuously changed to create various effects. On
// recording devices, this will adjust the gain before passing the data into
// an audiostream; that recording audiostream can then adjust its gain further
// when outputting the data elsewhere, if it likes, but that second gain is
// not applied until the data leaves the audiostream again.
//
// devid: the audio device on which to change gain.
//
// gain: the gain. 1.0f is no change, 0.0f is silence.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread, as it holds
// a stream-specific mutex while running.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetAudioDeviceGain
func (devid AudioDeviceID) SetGain(gain float32) error {
	if !C.SDL_SetAudioDeviceGain((C.SDL_AudioDeviceID)(devid), (C.float)(gain)) {
		return getError()
	}
	return nil
}

// Close a previously-opened audio device.
//
// The application should close open audio devices once they are no longer
// needed.
//
// This function may block briefly while pending audio data is played by the
// hardware, so that applications don't drop the last buffer of data they
// supplied if terminating immediately afterwards.
//
// devid: an audio device id previously returned by [OpenAudioDevice].
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CloseAudioDevice
func (devid AudioDeviceID) Close() {
	C.SDL_CloseAudioDevice((C.SDL_AudioDeviceID)(devid))
}

// Bind a list of audio streams to an audio device.
//
// Audio data will flow through any bound streams. For a playback device, data
// for all bound streams will be mixed together and fed to the device. For a
// recording device, a copy of recorded data will be provided to each bound
// stream.
//
// Audio streams can only be bound to an open device. This operation is
// atomic--all streams bound in the same call will start processing at the
// same time, so they can stay in sync. Also: either all streams will be bound
// or none of them will be.
//
// It is an error to bind an already-bound stream; it must be explicitly
// unbound first.
//
// Binding a stream to a device will set its output format for playback
// devices, and its input format for recording devices, so they match the
// device's settings. The caller is welcome to change the other end of the
// stream's format at any time with [AudioStream.SetFormat].
//
// devid: an audio device to bind a stream to.
//
// streams: an array of audio streams to bind.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BindAudioStreams
func (devid AudioDeviceID) BindAudioStreams(streams []*AudioStream) error {
	if !C.SDL_BindAudioStreams((C.SDL_AudioDeviceID)(devid), (**C.SDL_AudioStream)(unsafe.Pointer(unsafe.SliceData(streams))), (C.int)(len(streams))) {
		return getError()
	}
	return nil
}

// Bind a single audio stream to an audio device.
//
// This is a convenience function, equivalent to calling
// [AudioDeviceID.BindAudioStreams] with a single stream.
//
// devid: an audio device to bind a stream to.
//
// stream: an audio stream to bind to a device.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BindAudioStream
func (devid AudioDeviceID) BindAudioStream(stream *AudioStream) error {
	if !C.SDL_BindAudioStream((C.SDL_AudioDeviceID)(devid), (*C.SDL_AudioStream)(stream)) {
		return getError()
	}
	return nil
}

// Unbind a list of audio streams from their audio devices.
//
// The streams being unbound do not all have to be on the same device. All
// streams on the same device will be unbound atomically (data will stop
// flowing through all unbound streams on the same device at the same time).
//
// Unbinding a stream that isn't bound to a device is a legal no-op.
//
// streams: an array of audio streams to unbind. Can be nil or contain nil.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStreams
func UnbindAudioStreams(streams []*AudioStream) {
	C.SDL_UnbindAudioStreams((**C.SDL_AudioStream)(unsafe.Pointer(unsafe.SliceData(streams))), (C.int)(len(streams)))
}

// Unbind a single audio stream from its audio device.
//
// This is a convenience function, equivalent to calling
// [UnbindAudioStreams] with a single stream.
//
// stream: an audio stream to unbind from a device.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStream
func (stream *AudioStream) Unbind() {
	C.SDL_UnbindAudioStream((*C.SDL_AudioStream)(stream))
}

// Query an audio stream for its currently-bound device.
//
// This reports the audio device that an audio stream is currently bound to.
//
// If not bound, or invalid, this returns zero, which is not a valid device
// ID.
//
// stream: the audio stream to query.
//
// Returns the bound audio device, or 0 if not bound or invalid.
//
// It is safe to call this function from any thread.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamDevice
func (stream *AudioStream) Device() AudioDeviceID {
	return (AudioDeviceID)(C.SDL_GetAudioStreamDevice((*C.SDL_AudioStream)(stream)))
}

// Create a new audio stream.
//
// src: the format details of the input audio.
//
// dst: the format details of the output audio.
//
// Returns a new audio stream or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateAudioStream
func CreateAudioStream(src AudioSpec, dst AudioSpec) (*AudioStream, error) {
	a := (*AudioStream)(C.SDL_CreateAudioStream(
		&C.SDL_AudioSpec{C.SDL_AudioFormat(src.Format), C.int(src.Channels), C.int(src.Freq)},
		&C.SDL_AudioSpec{C.SDL_AudioFormat(dst.Format), C.int(dst.Channels), C.int(dst.Freq)}))
	if a == nil {
		return nil, getError()
	}
	return a, nil
}

// Get the properties associated with an audio stream.
//
// stream: the [AudioStream] to query.
//
// Returns a valid property ID or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamProperties
func (stream *AudioStream) Properties() (PropertiesID, error) {
	props := (PropertiesID)(C.SDL_GetAudioStreamProperties((*C.SDL_AudioStream)(stream)))
	if props == 0 {
		return 0, getError()
	}
	return props, nil
}

// Query the current format of an audio stream.
//
// stream: the [AudioStream] to query.
//
// Returns the input and output audio formats or an error on failure.
//
// It is safe to call this function from any thread, as it holds
// a stream-specific mutex while running.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFormat
func (stream *AudioStream) Format() (src AudioSpec, dst AudioSpec, err error) {
	var srcSpec, dstSpec C.SDL_AudioSpec
	if !C.SDL_GetAudioStreamFormat((*C.SDL_AudioStream)(stream), &srcSpec, &dstSpec) {
		return AudioSpec{}, AudioSpec{}, getError()
	}
	return AudioSpec{AudioFormat(srcSpec.format), int(srcSpec.channels), int(srcSpec.freq)},
		AudioSpec{AudioFormat(dstSpec.format), int(dstSpec.channels), int(dstSpec.freq)}, nil
}

// Change the input and output formats of an audio stream.
//
// Future calls to and [AudioStream.Available] and [AudioStream.GetData]
// will reflect the new format, and future calls to [AudioStream.PutData]
// must provide data in the new input formats.
//
// Data that was previously queued in the stream will still be operated on in
// the format that was current when it was added, which is to say you can put
// the end of a sound file in one format to a stream, change formats for the
// next sound file, and start putting that new data while the previous sound
// file is still queued, and everything will still play back correctly.
//
// If a stream is bound to a device, then the format of the side of the stream
// bound to a device cannot be changed (src_spec for recording devices,
// dst_spec for playback devices). Attempts to make a change to this side will
// be ignored, but this will not report an error. The other side's format can
// be changed.
//
// stream: the stream the format is being changed.
//
// src_spec: the new format of the audio input; if nil, it is not
// changed.
//
// dst_spec: the new format of the audio output; if nil, it is not
// changed.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread, as it holds
// a stream-specific mutex while running.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFormat
func (stream *AudioStream) SetFormat(src, dst *AudioSpec) error {
	var csrc, cdst *C.SDL_AudioSpec
	if src != nil {
		csrc = &C.SDL_AudioSpec{C.SDL_AudioFormat(src.Format), C.int(src.Channels), C.int(src.Freq)}
	}
	if dst != nil {
		cdst = &C.SDL_AudioSpec{C.SDL_AudioFormat(dst.Format), C.int(dst.Channels), C.int(dst.Freq)}
	}
	if !C.SDL_SetAudioStreamFormat((*C.SDL_AudioStream)(stream), csrc, cdst) {
		return getError()
	}
	return nil
}

// Get the frequency ratio of an audio stream.
//
// stream: the [AudioStream] to query.
//
// Returns the frequency ratio of the stream or an error.
//
// It is safe to call this function from any thread, as it holds
// a stream-specific mutex while running.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFrequencyRatio
func (stream *AudioStream) FrequencyRatio() (float32, error) {
	fr := (float32)(C.SDL_GetAudioStreamFrequencyRatio((*C.SDL_AudioStream)(stream)))
	if fr == 0 {
		return 0, getError()
	}
	return fr, nil
}

// Change the frequency ratio of an audio stream.
//
// The frequency ratio is used to adjust the rate at which input data is
// consumed. Changing this effectively modifies the speed and pitch of the
// audio. A value greater than 1.0 will play the audio faster, and at a higher
// pitch. A value less than 1.0 will play the audio slower, and at a lower
// pitch.
//
// This is applied during [AudioStream.GetData], and can be continuously
// changed to create various effects.
//
// stream: the stream the frequency ratio is being changed.
//
// ratio: the frequency ratio. 1.0 is normal speed. Must be between 0.01
// and 100.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread, as it holds
// a stream-specific mutex while running.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFrequencyRatio
func (stream *AudioStream) SetFrequencyRatio(ratio float32) error {
	if !C.SDL_SetAudioStreamFrequencyRatio((*C.SDL_AudioStream)(stream), (C.float)(ratio)) {
		return getError()
	}
	return nil
}

// Get the gain of an audio stream.
//
// The gain of a stream is its volume; a larger gain means a louder output,
// with a gain of zero being silence.
//
// Audio streams default to a gain of 1.0f (no change in output).
//
// stream: the [AudioStream] to query.
//
// Returns the gain of the stream or an error.
//
// It is safe to call this function from any thread, as it holds
// a stream-specific mutex while running.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamGain
func (stream *AudioStream) Gain() (float32, error) {
	gain := (float32)(C.SDL_GetAudioStreamGain((*C.SDL_AudioStream)(stream)))
	if gain == -1 {
		return -1, getError()
	}
	return gain, nil
}

// Change the gain of an audio stream.
//
// The gain of a stream is its volume; a larger gain means a louder output,
// with a gain of zero being silence.
//
// Audio streams default to a gain of 1.0f (no change in output).
//
// This is applied during [AudioStream.GetData], and can be continuously
// changed to create various effects.
//
// stream: the stream on which the gain is being changed.
//
// gain: the gain. 1.0f is no change, 0.0f is silence.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread, as it holds
// a stream-specific mutex while running.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGain
func (stream *AudioStream) SetGain(gain float32) error {
	if !C.SDL_SetAudioStreamGain((*C.SDL_AudioStream)(stream), (C.float)(gain)) {
		return getError()
	}
	return nil
}

// Get the current input channel map of an audio stream.
//
// Channel maps are optional; most things do not need them, instead passing
// data in the [order that SDL expects](CategoryAudio#channel-layouts).
//
// Audio streams default to no remapping applied. This is represented by
// returning nil, and does not signify an error.
//
// stream: the [AudioStream] to query.
//
// Returns an array of the current channel mapping, with as many elements as
// the current output spec's channels, or nil if default.
//
// It is safe to call this function from any thread, as it holds
// a stream-specific mutex while running.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamInputChannelMap
func (stream *AudioStream) InputChannelMap() []int {
	var count C.int
	ptr := C.SDL_GetAudioStreamInputChannelMap((*C.SDL_AudioStream)(stream), &count)
	if ptr == nil {
		return nil
	}
	slice := unsafe.Slice(ptr, count)
	result := make([]int, count)
	for i, c := range slice {
		result[i] = int(c)
	}
	C.SDL_free(unsafe.Pointer(ptr))
	return result
}

// Get the current output channel map of an audio stream.
//
// Channel maps are optional; most things do not need them, instead passing
// data in the [order that SDL expects](CategoryAudio#channel-layouts).
//
// Audio streams default to no remapping applied. This is represented by
// returning nil, and does not signify an error.
//
// stream: the [AudioStream] to query.
//
// Returns an array of the current channel mapping, with as many elements as
// the current output spec's channels, or nil if default.
//
// It is safe to call this function from any thread, as it holds
// a stream-specific mutex while running.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamOutputChannelMap
func (stream *AudioStream) OutputChannelMap() []int {
	var count C.int
	ptr := C.SDL_GetAudioStreamOutputChannelMap((*C.SDL_AudioStream)(stream), &count)
	if ptr == nil {
		return nil
	}
	slice := unsafe.Slice(ptr, count)
	result := make([]int, count)
	for i, c := range slice {
		result[i] = int(c)
	}
	C.SDL_free(unsafe.Pointer(ptr))
	return result
}

// Set the current input channel map of an audio stream.
//
// Channel maps are optional; most things do not need them, instead passing
// data in the [order that SDL expects](CategoryAudio#channel-layouts).
//
// The input channel map reorders data that is added to a stream via
// [AudioStream.PutData]. Future calls to [AudioStream.PutData] must provide
// data in the new channel order.
//
// Each item in the array represents an input channel, and its value is the
// channel that it should be remapped to. To reverse a stereo signal's left
// and right values, you'd have an array of { 1, 0 }. It is legal to remap
// multiple channels to the same thing, so { 1, 1 } would duplicate the
// right channel to both channels of a stereo signal. An element in the
// channel map set to -1 instead of a valid channel will mute that channel,
// setting it to a silence value.
//
// You cannot change the number of channels through a channel map, just
// reorder/mute them.
//
// Data that was previously queued in the stream will still be operated on in
// the order that was current when it was added, which is to say you can put
// the end of a sound file in one order to a stream, change orders for the
// next sound file, and start putting that new data while the previous sound
// file is still queued, and everything will still play back correctly.
//
// Audio streams default to no remapping applied. Passing a nil channel map
// is legal, and turns off remapping.
//
// SDL will copy the channel map; the caller does not have to save this array
// after this call.
//
// If len(chmap) is not equal to the current number of channels in the audio
// stream's format, this will fail. This is a safety measure to make sure a
// race condition hasn't changed the format while this call is setting the
// channel map.
//
// Unlike attempting to change the stream's format, the input channel map on a
// stream bound to a recording device is permitted to change at any time; any
// data added to the stream from the device after this call will have the new
// mapping, but previously-added data will still have the prior mapping.
//
// stream: the [AudioStream] to change.
//
// chmap: the new channel map, nil to reset to default.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread, as it holds
// a stream-specific mutex while running. Don't change the
// stream's format to have a different number of channels from a
// a different thread at the same time, though!
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamInputChannelMap
func (stream *AudioStream) SetInputChannelMap(chmap []int) error {
	var cchmap []C.int
	if chmap != nil {
		cchmap = make([]C.int, len(chmap))
		for i, c := range chmap {
			cchmap[i] = C.int(c)
		}
	}
	if !C.SDL_SetAudioStreamInputChannelMap((*C.SDL_AudioStream)(stream), unsafe.SliceData(cchmap), C.int(len(chmap))) {
		return getError()
	}
	return nil
}

// Set the current output channel map of an audio stream.
//
// Channel maps are optional; most things do not need them, instead passing
// data in the [order that SDL expects](CategoryAudio#channel-layouts).
//
// The output channel map reorders data that leaving a stream via
// [AudioStream.GetData].
//
// Each item in the array represents an input channel, and its value is the
// channel that it should be remapped to. To reverse a stereo signal's left
// and right values, you'd have an array of { 1, 0 }. It is legal to remap
// multiple channels to the same thing, so { 1, 1 } would duplicate the
// right channel to both channels of a stereo signal. An element in the
// channel map set to -1 instead of a valid channel will mute that channel,
// setting it to a silence value.
//
// You cannot change the number of channels through a channel map, just
// reorder/mute them.
//
// The output channel map can be changed at any time, as output remapping is
// applied during [AudioStream.GetData].
//
// Audio streams default to no remapping applied. Passing a nil channel map
// is legal, and turns off remapping.
//
// SDL will copy the channel map; the caller does not have to save this array
// after this call.
//
// If len(chmap) is not equal to the current number of channels in the audio
// stream's format, this will fail. This is a safety measure to make sure a
// race condition hasn't changed the format while this call is setting the
// channel map.
//
// Unlike attempting to change the stream's format, the output channel map on
// a stream bound to a recording device is permitted to change at any time;
// any data added to the stream after this call will have the new mapping, but
// previously-added data will still have the prior mapping. When the channel
// map doesn't match the hardware's channel layout, SDL will convert the data
// before feeding it to the device for playback.
//
// stream: the [AudioStream] to change.
//
// chmap: the new channel map, nil to reset to default.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread, as it holds
// a stream-specific mutex while running. Don't change the
// stream's format to have a different number of channels from a
// a different thread at the same time, though!
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamOutputChannelMap
func (stream *AudioStream) SetOutputChannelMap(chmap []int) error {
	var cchmap []C.int
	if chmap != nil {
		cchmap = make([]C.int, len(chmap))
		for i, c := range chmap {
			cchmap[i] = C.int(c)
		}
	}
	if !C.SDL_SetAudioStreamOutputChannelMap((*C.SDL_AudioStream)(stream), unsafe.SliceData(cchmap), C.int(len(chmap))) {
		return getError()
	}
	return nil
}

// Add data to the stream.
//
// This data must match the format/channels/samplerate specified in the latest
// call to [AudioStream.SetFormat], or the format specified when creating the
// stream if it hasn't been changed.
//
// Note that this call simply copies the unconverted data for later. This is
// different than SDL2, where data was converted during the Put call and the
// Get call would just dequeue the previously-converted data.
//
// stream: the stream the audio data is being added to.
//
// buf: the audio data to add.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread, but if the
// stream has a callback set, the caller might need to manage
// extra locking.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamData
func (stream *AudioStream) PutData(buf []byte) error {
	if !C.SDL_PutAudioStreamData((*C.SDL_AudioStream)(stream), unsafe.Pointer(unsafe.SliceData(buf)), (C.int)(len(buf))) {
		return getError()
	}
	return nil
}

// Get converted/resampled data from the stream.
//
// The input/output data format/channels/samplerate is specified when creating
// the stream, and can be changed after creation by calling
// [AudioStream.SetFormat].
//
// Note that any conversion and resampling necessary is done during this call,
// and [AudioStream.PutData] simply queues unconverted data for later. This
// is different than SDL2, where that work was done while inputting new data
// to the stream and requesting the output just copied the converted data.
//
// stream: the stream the audio is being requested from.
//
// buf: a buffer to fill with audio data.
//
// Returns the number of bytes read from the stream or an error.
//
// It is safe to call this function from any thread, but if the
// stream has a callback set, the caller might need to manage
// extra locking.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamData
func (stream *AudioStream) GetData(buf []byte) (int, error) {
	n := C.SDL_GetAudioStreamData((*C.SDL_AudioStream)(stream), unsafe.Pointer(unsafe.SliceData(buf)), (C.int)(len(buf)))
	if n == -1 {
		return 0, getError()
	}
	return int(n), nil
}

// Get the number of converted/resampled bytes available.
//
// The stream may be buffering data behind the scenes until it has enough to
// resample correctly, so this number might be lower than what you expect, or
// even be zero. Add more data or flush the stream if you need the data now.
//
// If the stream has so much data that it would overflow an int, the return
// value is clamped to a maximum value, but no queued data is lost; if there
// are gigabytes of data queued, the app might need to read some of it with
// [AudioStream.GetData] before this function's return value is no longer
// clamped.
//
// stream: the audio stream to query.
//
// Returns the number of converted/resampled bytes available or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamAvailable
func (stream *AudioStream) Available() (int, error) {
	avail := C.SDL_GetAudioStreamAvailable((*C.SDL_AudioStream)(stream))
	if avail == -1 {
		return 0, getError()
	}
	return int(avail), nil
}

// Get the number of bytes currently queued.
//
// This is the number of bytes put into a stream as input, not the number that
// can be retrieved as output. Because of several details, it's not possible
// to calculate one number directly from the other. If you need to know how
// much usable data can be retrieved right now, you should use
// [AudioStream.Available] and not this function.
//
// Note that audio streams can change their input format at any time, even if
// there is still data queued in a different format, so the returned byte
// count will not necessarily match the number of _sample frames_ available.
// Users of this API should be aware of format changes they make when feeding
// a stream and plan accordingly.
//
// Queued data is not converted until it is consumed by
// [AudioStream.GetData], so this value should be representative of the exact
// data that was put into the stream.
//
// If the stream has so much data that it would overflow an int, the return
// value is clamped to a maximum value, but no queued data is lost; if there
// are gigabytes of data queued, the app might need to read some of it with
// [AudioStream.GetData] before this function's return value is no longer
// clamped.
//
// stream: the audio stream to query.
//
// Returns the number of bytes queued or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamQueued
func (stream *AudioStream) Queued() (int, error) {
	queued := C.SDL_GetAudioStreamQueued((*C.SDL_AudioStream)(stream))
	if queued == -1 {
		return 0, getError()
	}
	return int(queued), nil
}

// Tell the stream that you're done sending data, and anything being buffered
// should be converted/resampled and made available immediately.
//
// It is legal to add more data to a stream after flushing, but there may be
// audio gaps in the output. Generally this is intended to signal the end of
// input, so the complete output becomes available.
//
// stream: the audio stream to flush.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_FlushAudioStream
func (stream *AudioStream) Flush() error {
	if !C.SDL_FlushAudioStream((*C.SDL_AudioStream)(stream)) {
		return getError()
	}
	return nil
}

// Clear any pending data in the stream.
//
// This drops any queued data, so there will be nothing to read from the
// stream until more is added.
//
// stream: the audio stream to clear.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ClearAudioStream
func (stream *AudioStream) Clear() error {
	if !C.SDL_ClearAudioStream((*C.SDL_AudioStream)(stream)) {
		return getError()
	}
	return nil
}

// Use this function to pause audio playback on the audio device associated
// with an audio stream.
//
// This function pauses audio processing for a given device. Any bound audio
// streams will not progress, and no audio will be generated. Pausing one
// device does not prevent other unpaused devices from running.
//
// Pausing a device can be useful to halt all audio without unbinding all the
// audio streams. This might be useful while a game is paused, or a level is
// loading, etc.
//
// stream: the audio stream associated with the audio device to pause.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PauseAudioStreamDevice
func (stream *AudioStream) PauseDevice() error {
	if !C.SDL_PauseAudioStreamDevice((*C.SDL_AudioStream)(stream)) {
		return getError()
	}
	return nil
}

// Use this function to unpause audio playback on the audio device associated
// with an audio stream.
//
// This function unpauses audio processing for a given device that has
// previously been paused. Once unpaused, any bound audio streams will begin
// to progress again, and audio can be generated.
//
// stream: the audio stream associated with the audio device to resume.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ResumeAudioStreamDevice
func (stream *AudioStream) ResumeDevice() error {
	if !C.SDL_ResumeAudioStreamDevice((*C.SDL_AudioStream)(stream)) {
		return getError()
	}
	return nil
}

// Use this function to query if an audio device associated with a stream is
// paused.
//
// Unlike in SDL2, audio devices start in an _unpaused_ state, since an app
// has to bind a stream before any audio will flow.
//
// stream: the audio stream associated with the audio device to query.
//
// Returns true if device is valid and paused, false otherwise.
//
// It is safe to call this function from any thread.
//
// https://wiki.libsdl.org/SDL3/SDL_AudioStreamDevicePaused
func (stream *AudioStream) DevicePaused() bool {
	return (bool)(C.SDL_AudioStreamDevicePaused((*C.SDL_AudioStream)(stream)))
}

// Lock an audio stream for serialized access.
//
// Each [AudioStream] has an internal mutex it uses to protect its data
// structures from threading conflicts. This function allows an app to lock
// that mutex, which could be useful if registering callbacks on this stream.
//
// One does not need to lock a stream to use in it most cases, as the stream
// manages this lock internally. However, this lock is held during callbacks,
// which may run from arbitrary threads at any time, so if an app needs to
// protect shared data during those callbacks, locking the stream guarantees
// that the callback is not running while the lock is held.
//
// As this is just a wrapper over SDL_LockMutex for an internal lock; it has
// all the same attributes (recursive locks are allowed, etc).
//
// stream: the audio stream to lock.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_LockAudioStream
func (stream *AudioStream) Lock() error {
	if !C.SDL_LockAudioStream((*C.SDL_AudioStream)(stream)) {
		return getError()
	}
	return nil
}

// Unlock an audio stream for serialized access.
//
// This unlocks an audio stream after a call to [AudioStream.Lock].
//
// stream: the audio stream to unlock.
//
// Returns nil on success or an error on failure.
//
// You should only call this from the same thread that
// previously called [AudioStream.Lock].
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UnlockAudioStream
func (stream *AudioStream) Unlock() error {
	if !C.SDL_UnlockAudioStream((*C.SDL_AudioStream)(stream)) {
		return getError()
	}
	return nil
}

// A callback that fires when data passes through an [AudioStream].
//
// Apps can (optionally) register a callback with an audio stream that is
// called when data is added with [AudioStream.PutData], or requested with
// [AudioStream.GetData].
//
// Two values are offered here: one is the amount of additional data needed to
// satisfy the immediate request (which might be zero if the stream already
// has enough data queued) and the other is the total amount being requested.
// In a Get call triggering a Put callback, these values can be different. In
// a Put call triggering a Get callback, these values are always the same.
//
// Byte counts might be slightly overestimated due to buffering or resampling,
// and may change from call to call.
//
// This callback is not required to do anything. Generally this is useful for
// adding/reading data on demand, and the app will often put/get data as
// appropriate, but the system goes on with the data currently available to it
// if this callback does nothing.
//
// stream: the SDL audio stream associated with this callback.
//
// additionalAmount: the amount of data, in bytes, that is needed right
// now.
//
// totalAmount: the total amount of data requested, in bytes, that is
// requested or available.
//
// This callbacks may run from any thread, so if you need to
// protect shared data, you should use [AudioStream.Lock] to
// serialize access; this lock will be held before your callback
// is called, so your callback does not need to manage the lock
// explicitly.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AudioStreamCallback
type AudioStreamCallback func(stream *AudioStream, additionalAmount int, totalAmount int)

//export cb_AudioStreamCallback
func cb_AudioStreamCallback(userdata uintptr, stream *C.SDL_AudioStream, additionalAmount C.int, totalAmount C.int) {
	h := cgo.Handle(userdata)
	h.Value().(AudioStreamCallback)((*AudioStream)(stream), int(additionalAmount), int(totalAmount))
}

// Set a callback that runs when data is requested from an audio stream.
//
// This callback is called _before_ data is obtained from the stream, giving
// the callback the chance to add more on-demand.
//
// The callback can (optionally) call [AudioStream.PutData] to add more
// audio to the stream during this call; if needed, the request that triggered
// this callback will obtain the new data immediately.
//
// The callback's `approx_request` argument is roughly how many bytes of
// _unconverted_ data (in the stream's input format) is needed by the caller,
// although this may overestimate a little for safety. This takes into account
// how much is already in the stream and only asks for any extra necessary to
// resolve the request, which means the callback may be asked for zero bytes,
// and a different amount on each call.
//
// The callback is not required to supply exact amounts; it is allowed to
// supply too much or too little or none at all. The caller will get what's
// available, up to the amount they requested, regardless of this callback's
// outcome.
//
// Clearing or flushing an audio stream does not call this callback.
//
// This function obtains the stream's lock, which means any existing callback
// (get or put) in progress will finish running before setting the new
// callback.
//
// Setting a nil function turns off the callback.
//
// stream: the audio stream to set the new callback on.
//
// callback: the new callback function to call when data is requested
// from the stream.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGetCallback
func (stream *AudioStream) SetGetCallback(callback AudioStreamCallback) {
	if stream == nil {
		panic("sdl: stream is nil")
	}
	if callback == nil {
		C.SDL_SetAudioStreamGetCallback((*C.SDL_AudioStream)(stream), nil, nil)
		if props, err := stream.Properties(); err == nil {
			props.Clear("go.audiostream.callback.get")
		}
		return
	}
	h := cgo.NewHandle(callback)
	C.wrap_SDL_SetAudioStreamGetCallback((*C.SDL_AudioStream)(stream), C.uintptr_t(h))
	if props, err := stream.Properties(); err == nil {
		props.SetPointerWithCleanup("go.audiostream.callback.get", uintptr(h), propHandleCleanup)
	}
}

// Set a callback that runs when data is added to an audio stream.
//
// This callback is called _after_ the data is added to the stream, giving the
// callback the chance to obtain it immediately.
//
// The callback can (optionally) call [AudioStream.GetData] to obtain audio
// from the stream during this call.
//
// The callback's `approx_request` argument is how many bytes of _converted_
// data (in the stream's output format) was provided by the caller, although
// this may underestimate a little for safety. This value might be less than
// what is currently available in the stream, if data was already there, and
// might be less than the caller provided if the stream needs to keep a buffer
// to aid in resampling. Which means the callback may be provided with zero
// bytes, and a different amount on each call.
//
// The callback may call [AudioStream.Available] to see the total amount
// currently available to read from the stream, instead of the total provided
// by the current call.
//
// The callback is not required to obtain all data. It is allowed to read less
// or none at all. Anything not read now simply remains in the stream for
// later access.
//
// Clearing or flushing an audio stream does not call this callback.
//
// This function obtains the stream's lock, which means any existing callback
// (get or put) in progress will finish running before setting the new
// callback.
//
// Setting a nil function turns off the callback.
//
// stream: the audio stream to set the new callback on.
//
// callback: the new callback function to call when data is added to the
// stream.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SetAudioStreamPutCallback
func (stream *AudioStream) SetPutCallback(callback AudioStreamCallback) {
	if stream == nil {
		panic("sdl: stream is nil")
	}
	if callback == nil {
		C.SDL_SetAudioStreamPutCallback((*C.SDL_AudioStream)(stream), nil, nil)
		if props, err := stream.Properties(); err == nil {
			props.Clear("go.audiostream.callback.put")
		}
		return
	}
	h := cgo.NewHandle(callback)
	C.wrap_SDL_SetAudioStreamPutCallback((*C.SDL_AudioStream)(stream), C.uintptr_t(h))
	if props, err := stream.Properties(); err == nil {
		props.SetPointerWithCleanup("go.audiostream.callback.put", uintptr(h), propHandleCleanup)
	}
}

// Free an audio stream.
//
// This will release all allocated data, including any audio that is still
// queued. You do not need to manually clear the stream first.
//
// If this stream was bound to an audio device, it is unbound during this
// call. If this stream was created with [OpenAudioDeviceStream], the audio
// device that was opened alongside this stream's creation will be closed,
// too.
//
// stream: the audio stream to destroy.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DestroyAudioStream
func (stream *AudioStream) Destroy() {
	C.SDL_DestroyAudioStream((*C.SDL_AudioStream)(stream))
}

// Convenience function for straightforward audio init for the common case.
//
// If all your app intends to do is provide a single source of PCM audio, this
// function allows you to do all your audio setup in a single call.
//
// This is also intended to be a clean means to migrate apps from SDL2.
//
// This function will open an audio device, create a stream and bind it.
// Unlike other methods of setup, the audio device will be closed when this
// stream is destroyed, so the app can treat the returned [AudioStream] as
// the only object needed to manage audio playback.
//
// Also unlike other functions, the audio device begins paused. This is to map
// more closely to SDL2-style behavior, since there is no extra step here to
// bind a stream to begin audio flowing. The audio device should be resumed
// with [AudioStream.ResumeDevice]
//
// This function works with both playback and recording devices.
//
// The spec parameter represents the app's side of the audio stream. That
// is, for recording audio, this will be the output format, and for playing
// audio, this will be the input format. If spec is nil, the system will
// choose the format, and the app can use [AudioStream.Format] to obtain
// this information later.
//
// If you don't care about opening a specific audio device, you can (and
// probably _should_), use [AudioDeviceDefaultPlayback] for playback and
// [AudioDeviceDefaultRecording] for recording.
//
// One can optionally provide a callback function; if nil, the app is
// expected to queue audio data for playback (or unqueue audio data if
// capturing). Otherwise, the callback will begin to fire once the device is
// unpaused.
//
// Destroying the returned stream with [AudioStream.Destroy] will also close
// the audio device associated with this stream.
//
// devid: an audio device to open, or [AudioDeviceDefaultPlayback]
// or [AudioDeviceDefaultRecording].
//
// spec: the audio stream's data format. Can be nil.
//
// callback: a callback where the app will provide new data for
// playback, or receive new data for recording. Can be nil,
// in which case the app will need to call
// [AudioStream.PutData] or [AudioStream.GetData] as
// necessary.
//
// Returns an audio stream on success, ready to use, an error.
// When done with this stream, call [AudioStream.Destroy] to free resources and
// close the device.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_OpenAudioDeviceStream
func OpenAudioDeviceStream(devid AudioDeviceID, spec *AudioSpec, callback AudioStreamCallback) (*AudioStream, error) {
	var h cgo.Handle
	if callback != nil {
		h = cgo.NewHandle(callback)
	}
	var cspec *C.SDL_AudioSpec
	if spec != nil {
		cspec = &C.SDL_AudioSpec{C.SDL_AudioFormat(spec.Format), C.int(spec.Channels), C.int(spec.Freq)}
	}
	stream := (*AudioStream)(C.wrap_SDL_OpenAudioDeviceStream((C.SDL_AudioDeviceID)(devid), cspec, C.uintptr_t(h)))
	if stream == nil {
		if callback != nil {
			h.Delete()
		}
		return nil, getError()
	}
	if callback != nil {
		props, _ := stream.Properties()
		if devid.Playback() {
			props.SetPointerWithCleanup("go.audiostream.callback.get", uintptr(h), propHandleCleanup)
		} else {
			props.SetPointerWithCleanup("go.audiostream.callback.put", uintptr(h), propHandleCleanup)
		}
	}
	return stream, nil
}

// A callback that fires when data is about to be fed to an audio device.
//
// This is useful for accessing the final mix, perhaps for writing a
// visualizer or applying a final effect to the audio data before playback.
//
// This callback should run as quickly as possible and not block for any
// significant time, as this callback delays submission of data to the audio
// device, which can cause audio playback problems.
//
// The postmix callback _must_ be able to handle any audio data format
// specified in spec, which can change between callbacks if the audio device
// changed. However, this only covers frequency and channel count; data is
// always provided here in [AudioF32] format.
//
// The postmix callback runs _after_ logical device gain and audiostream gain
// have been applied, which is to say you can make the output data louder at
// this point than the gain settings would suggest.
//
// spec: the current format of audio that is to be submitted to the
// audio device.
//
// buffer: the buffer of audio samples to be submitted. The callback can
// inspect and/or modify this data. This slice must not be retained, it may
// become invalid after the callback returns.
//
// This will run from a background thread owned by SDL. The
// application is responsible for locking resources the callback
// touches that need to be protected.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AudioPostmixCallback
type AudioPostmixCallback func(spec AudioSpec, buffer []float32)

//export cb_AudioPostmixCallback
func cb_AudioPostmixCallback(userdata C.uintptr_t, spec *C.SDL_AudioSpec, buffer *C.float, buflen C.int) {
	h := cgo.Handle(userdata)
	h.Value().(AudioPostmixCallback)(AudioSpec{AudioF32, int(spec.channels), int(spec.freq)}, unsafe.Slice((*float32)(buffer), buflen/4))
}

// Set a callback that fires when data is about to be fed to an audio device.
//
// This is useful for accessing the final mix, perhaps for writing a
// visualizer or applying a final effect to the audio data before playback.
//
// The buffer is the final mix of all bound audio streams on an opened device;
// this callback will fire regularly for any device that is both opened and
// unpaused. If there is no new data to mix, either because no streams are
// bound to the device or all the streams are empty, this callback will still
// fire with the entire buffer set to silence.
//
// This callback is allowed to make changes to the data; the contents of the
// buffer after this call is what is ultimately passed along to the hardware.
//
// The callback is always provided the data in float format (values from -1.0f
// to 1.0f), but the number of channels or sample rate may be different than
// the format the app requested when opening the device; SDL might have had to
// manage a conversion behind the scenes, or the playback might have jumped to
// new physical hardware when a system default changed, etc. These details may
// change between calls. Accordingly, the size of the buffer might change
// between calls as well.
//
// This callback can run at any time, and from any thread; if you need to
// serialize access to your app's data, you should provide and use a mutex or
// other synchronization device.
//
// All of this to say: there are specific needs this callback can fulfill, but
// it is not the simplest interface. Apps should generally provide audio in
// their preferred format through an [AudioStream] and let SDL handle the
// difference.
//
// This function is extremely time-sensitive; the callback should do the least
// amount of work possible and return as quickly as it can. The longer the
// callback runs, the higher the risk of audio dropouts or other problems.
//
// This function will block until the audio device is in between iterations,
// so any existing callback that might be running will finish before this
// function sets the new callback and returns.
//
// Setting a nil callback function disables any previously-set callback.
//
// devid: the ID of an opened audio device.
//
// callback: a callback function to be called. Can be nil.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetAudioPostmixCallback
func (devid AudioDeviceID) SetPostmixCallback(callback AudioPostmixCallback) error {
	if callback == nil {
		if !C.SDL_SetAudioPostmixCallback((C.SDL_AudioDeviceID)(devid), nil, nil) {
			return getError()
		}
		return nil
	}
	if !C.wrap_SDL_SetAudioPostmixCallback((C.SDL_AudioDeviceID)(devid), C.uintptr_t(cgo.NewHandle(callback))) {
		return getError()
	}
	return nil
}

// Load the audio data of a WAVE file into memory.
//
// The entire data portion of the file is loaded into memory and decoded if
// necessary.
//
// Supported formats are RIFF WAVE files with the formats PCM (8, 16, 24, and
// 32 bits), IEEE Float (32 bits), Microsoft ADPCM and IMA ADPCM (4 bits), and
// A-law and mu-law (8 bits). Other formats are currently unsupported and
// cause an error.
//
// If this function succeeds, it returns the audio data and an [AudioSpec]
// describing the format of the data.
//
// Because of the underspecification of the .WAV format, there are many
// problematic files in the wild that cause issues with strict decoders. To
// provide compatibility with these files, this decoder is lenient in regards
// to the truncation of the file, the fact chunk, and the size of the RIFF
// chunk. The hints [HintWaveRiffChunkSize],
// [HintWaveTruncation], and [HintWaveFactChunk] can be used to
// tune the behavior of the loading process.
//
// Any file that is invalid (due to truncation, corruption, or wrong values in
// the headers), too big, or unsupported causes an error. Additionally, any
// critical I/O error from the data source will terminate the loading process
// with an error.
//
// It is required that the data source supports seeking.
//
// src: the data source for the WAVE data.
//
// closeio: if true, calls [IOStream.Close] on src before returning, even
// in the case of an error.
//
// Returns the audio data and format, or an error.
//
// This function returns an error if the .WAV file cannot be opened,
// uses an unknown data format, or is corrupt.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_LoadWAV_IO
func LoadWAV_IO(src *IOStream, closeio bool) ([]byte, AudioSpec, error) {
	var audio_buf *C.Uint8
	var audio_len C.Uint32
	var spec C.SDL_AudioSpec
	if !C.SDL_LoadWAV_IO((*C.SDL_IOStream)(src), (C.bool)(closeio), &spec, &audio_buf, &audio_len) {
		return nil, AudioSpec{}, getError()
	}
	out := C.GoBytes(unsafe.Pointer(audio_buf), C.int(audio_len))
	C.SDL_free(unsafe.Pointer(audio_buf))
	return out, AudioSpec{AudioFormat(spec.format), int(spec.channels), int(spec.freq)}, nil
}

// Loads a WAV from a file path.
//
// This is a convenience function that is effectively the same as:
//
//	LoadWAV_IO(IOFromFile(path, "rb"), true)
//
// path: the file path of the WAV file to open.
//
// Returns the audio data and format, or an error.
//
// This function returns an error if the .WAV file cannot be opened,
// uses an unknown data format, or is corrupt.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_LoadWAV
func LoadWAV(path string) ([]byte, AudioSpec, error) {
	var audio_buf *C.Uint8
	var audio_len C.Uint32
	var spec C.SDL_AudioSpec
	if !C.SDL_LoadWAV(tmpstring(path), &spec, &audio_buf, &audio_len) {
		return nil, AudioSpec{}, getError()
	}
	out := C.GoBytes(unsafe.Pointer(audio_buf), C.int(audio_len))
	C.SDL_free(unsafe.Pointer(audio_buf))
	return out, AudioSpec{AudioFormat(spec.format), int(spec.channels), int(spec.freq)}, nil
}

// Mix audio data in a specified format.
//
// This takes an audio buffer src of format data and mixes
// it into dst, performing addition, volume adjustment, and overflow
// clipping. The buffer pointed to by dst must be of the same length and
// format.
//
// This is provided for convenience -- you can mix your own audio data.
//
// Do not use this function for mixing together more than two streams of
// sample data. The output from repeated application of this function may be
// distorted by clipping, because there is no accumulator with greater range
// than the input (not to mention this being an inefficient way of doing it).
//
// It is a common misconception that this function is required to write audio
// data to an output stream in an audio callback. While you can do that,
// [MixAudio] is really only needed when you're mixing a single audio
// stream with a volume adjustment.
//
// dst: the destination for the mixed audio.
//
// src: the source audio buffer to be mixed.
//
// format: the [AudioFormat] structure representing the desired audio
// format.
//
// volume: ranges from 0.0 - 1.0, and should be set to 1.0 for full
// audio volume.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MixAudio
func MixAudio(dst []byte, src []byte, format AudioFormat, volume float32) error {
	if !C.SDL_MixAudio((*C.Uint8)(unsafe.SliceData(dst)), (*C.Uint8)(unsafe.SliceData(src)), (C.SDL_AudioFormat)(format), (C.Uint32)(min(len(dst), len(src))), (C.float)(volume)) {
		return getError()
	}
	return nil
}

// Convert some audio data of one format to another format.
//
// Please note that this function is for convenience, but should not be used
// to resample audio in blocks, as it will introduce audio artifacts on the
// boundaries. You should only use this function if you are converting audio
// data in its entirety in one call. If you want to convert audio in smaller
// chunks, use an [AudioStream], which is designed for this situation.
//
// Internally, this function creates and destroys an [AudioStream] on each
// use, so it's also less efficient than using one directly, if you need to
// convert multiple times.
//
// srcSpec: the format details of the input audio.
//
// src: the audio data to be converted.
//
// dstSpec: the format details of the output audio.
//
// Returns the converted audio data or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ConvertAudioSamples
func ConvertAudioSamples(srcSpec AudioSpec, src []byte, dstSpec AudioSpec) ([]byte, error) {
	var dstData *C.Uint8
	var dstLen C.int
	if !C.SDL_ConvertAudioSamples(
		&C.SDL_AudioSpec{C.SDL_AudioFormat(srcSpec.Format), C.int(srcSpec.Channels), C.int(srcSpec.Freq)},
		(*C.Uint8)(unsafe.SliceData(src)), (C.int)(len(src)),
		&C.SDL_AudioSpec{C.SDL_AudioFormat(dstSpec.Format), C.int(dstSpec.Channels), C.int(dstSpec.Freq)},
		&dstData, &dstLen) {
		return nil, getError()
	}
	dst := C.GoBytes(unsafe.Pointer(dstData), dstLen)
	C.SDL_free(unsafe.Pointer(dstData))
	return dst, nil
}

// Get the human readable name of an audio format.
//
// format: the audio format to query.
//
// Returns the human readable name of the specified audio format or
// "SDL_AUDIO_UNKNOWN" if the format isn't recognized.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAudioFormatName
func (format AudioFormat) Name() string {
	return C.GoString(C.SDL_GetAudioFormatName((C.SDL_AudioFormat)(format)))
}

// Get the appropriate memset value for silencing an audio format.
//
// A byte slice filled with the value returned by this function represents
// silence in the specified format.
//
// format: the audio data format to query.
//
// Returns a byte value that can be passed to memset.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSilenceValueForFormat
func GetSilenceValueForFormat(format AudioFormat) byte {
	return (byte)(C.SDL_GetSilenceValueForFormat((C.SDL_AudioFormat)(format)))
}
