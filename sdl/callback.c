#include <SDL3/SDL.h>

/* audio */
extern void cb_AudioStreamCallback(void *userdata, SDL_AudioStream *stream, int additional_amount, int requested_amount);
bool wrap_SDL_SetAudioStreamPutCallback(SDL_AudioStream *stream, uintptr_t userdata) {
	return SDL_SetAudioStreamPutCallback(stream, cb_AudioStreamCallback, (void *)userdata);
}
bool wrap_SDL_SetAudioStreamGetCallback(SDL_AudioStream *stream, uintptr_t userdata) {
	return SDL_SetAudioStreamGetCallback(stream, cb_AudioStreamCallback, (void *)userdata);
}
SDL_AudioStream *wrap_SDL_OpenAudioDeviceStream(SDL_AudioDeviceID devid, const SDL_AudioSpec *spec, uintptr_t userdata) {
	return SDL_OpenAudioDeviceStream(devid, spec, userdata ? cb_AudioStreamCallback : NULL, (void *)userdata);
}
extern void cb_AudioPostmixCallback(void *userdata, const SDL_AudioSpec *spec, float *buffer, int buflen);
bool wrap_SDL_SetAudioPostmixCallback(SDL_AudioDeviceID devid, uintptr_t userdata) {
	return SDL_SetAudioPostmixCallback(devid, cb_AudioPostmixCallback, (void *)userdata);
}

/* clipboard */
extern const void *cb_ClipboardDataCallback(void *userdata, const char *mime_type, size_t *size);
extern void cb_ClipboardCleanupCallback(void *userdata);
bool wrap_SDL_SetClipboardData(uintptr_t userdata, const char **mime_types, size_t num_mime_types) {
	return SDL_SetClipboardData(cb_ClipboardDataCallback, cb_ClipboardCleanupCallback, (void *)userdata, mime_types, num_mime_types);
}

/* dialog */
extern void cb_DialogFileCallback(void *userdata, const char *const *filelist, int filter);
void wrap_SDL_ShowOpenFileDialog(uintptr_t userdata, SDL_Window *window, const SDL_DialogFileFilter *filters, int nfilters, const char *default_location, bool allow_many) {
	SDL_ShowOpenFileDialog(cb_DialogFileCallback, (void *)userdata, window, filters, nfilters, default_location, allow_many);
}
void wrap_SDL_ShowSaveFileDialog(uintptr_t userdata, SDL_Window *window, const SDL_DialogFileFilter *filters, int nfilters, const char *default_location) {
	SDL_ShowSaveFileDialog(cb_DialogFileCallback, (void *)userdata, window, filters, nfilters, default_location);
}
void wrap_SDL_ShowOpenFolderDialog(uintptr_t userdata, SDL_Window *window, const char *default_location, bool allow_many) {
	SDL_ShowOpenFolderDialog(cb_DialogFileCallback, (void *)userdata, window, default_location, allow_many);
}
void wrap_SDL_ShowFileDialogWithProperties(SDL_FileDialogType type, uintptr_t userdata, SDL_PropertiesID props) {
	SDL_ShowFileDialogWithProperties(type, cb_DialogFileCallback, (void *)userdata, props);
}

/* events */
extern bool cb_EventFilter(void *userdata, SDL_Event *event);
void wrap_SDL_SetEventFilter(uintptr_t userdata) {
	SDL_SetEventFilter(cb_EventFilter, (void *)userdata);
}
void wrap_SDL_FilterEvents(uintptr_t userdata) {
	SDL_FilterEvents(cb_EventFilter, (void *)userdata);
}
bool wrap_SDL_AddEventWatch(uintptr_t userdata) {
	return SDL_AddEventWatch(cb_EventFilter, (void *)userdata);
}
void wrap_SDL_RemoveEventWatch(uintptr_t userdata) {
	SDL_RemoveEventWatch(cb_EventFilter, (void *)userdata);
}

/* hints */
extern void cb_HintCallback(void *userdata, const char *name, const char *oldValue, const char *newValue);
bool wrap_SDL_AddHintCallback(const char *name, uintptr_t userdata) {
	return SDL_AddHintCallback(name, cb_HintCallback, (void *)userdata);
}
void wrap_SDL_RemoveHintCallback(const char *name, uintptr_t userdata) {
	SDL_RemoveHintCallback(name, cb_HintCallback, (void *)userdata);
}

/* init */
extern void cb_MainThreadCallback(void *userdata);
bool wrap_SDL_RunOnMainThread(uintptr_t userdata, bool wait) {
	return SDL_RunOnMainThread(cb_MainThreadCallback, (void *)userdata, wait);
}

/* iostream */
extern Sint64 cb_IOStreamSize(void *userdata);
extern Sint64 cb_IOStreamSeek(void *userdata, Sint64 offset, SDL_IOWhence whence);
extern size_t cb_IOStreamRead(void *userdata, void *ptr, size_t size, SDL_IOStatus *status);
extern size_t cb_IOStreamWrite(void *userdata, const void *ptr, size_t size, SDL_IOStatus *status);
extern bool cb_IOStreamFlush(void *userdata, SDL_IOStatus *status);
extern bool cb_IOStreamClose(void *userdata);
SDL_IOStream *wrap_SDL_OpenIO(uintptr_t userdata) {
	SDL_IOStreamInterface iface;
	SDL_INIT_INTERFACE(&iface);
	iface.size = cb_IOStreamSize;
	iface.seek = cb_IOStreamSeek;
	iface.read = cb_IOStreamRead;
	iface.write = cb_IOStreamWrite;
	iface.flush = cb_IOStreamFlush;
	iface.close = cb_IOStreamClose;
	return SDL_OpenIO(&iface, (void *)userdata);
}

/* joystick */
extern void cb_VirtualJoystickUpdate(void *userdata);
extern void cb_VirtualJoystickSetPlayerIndex(void *userdata, int player_index);
extern bool cb_VirtualJoystickRumble(void *userdata, Uint16 low_frequency_rumble, Uint16 high_frequency_rumble);
extern bool cb_VirtualJoystickRumbleTriggers(void *userdata, Uint16 left_rumble, Uint16 right_rumble);
extern bool cb_VirtualJoystickSetLED(void *userdata, Uint8 red, Uint8 green, Uint8 blue);
extern bool cb_VirtualJoystickSendEffect(void *userdata, const void *data, int size);
extern bool cb_VirtualJoystickSetSensorsEnabled(void *userdata, bool enabled);
extern void cb_VirtualJoystickCleanup(void *userdata);
SDL_JoystickID wrap_SDL_AttachVirtualJoystick(SDL_VirtualJoystickDesc *desc, uintptr_t userdata) {
	desc->version = sizeof(*desc);
	desc->userdata = (void *)userdata;
	desc->Update = cb_VirtualJoystickUpdate;
	desc->SetPlayerIndex = cb_VirtualJoystickSetPlayerIndex;
	desc->Rumble = cb_VirtualJoystickRumble;
	desc->RumbleTriggers = cb_VirtualJoystickRumbleTriggers;
	desc->SetLED = cb_VirtualJoystickSetLED;
	desc->SendEffect = cb_VirtualJoystickSendEffect;
	desc->SetSensorsEnabled = cb_VirtualJoystickSetSensorsEnabled;
	desc->Cleanup = cb_VirtualJoystickCleanup;
	return SDL_AttachVirtualJoystick(desc);
}

/* properties */
extern void cb_CleanupPropertyCallback(void *userdata, void* value);
bool wrap_SDL_SetPointerPropertyWithCleanup(SDL_PropertiesID props, const char *name, uintptr_t value, uintptr_t userdata) {
	return SDL_SetPointerPropertyWithCleanup(props, name, (void *)value, cb_CleanupPropertyCallback, (void *)userdata);
}
extern void cb_EnumeratePropertiesCallback(void *userdata, SDL_PropertiesID props, const char *name);
bool wrap_SDL_EnumerateProperties(SDL_PropertiesID props, uintptr_t userdata) {
	return SDL_EnumerateProperties(props, cb_EnumeratePropertiesCallback, (void *)userdata);
}

/* tray */
extern void cb_TrayCallback(void *userdata, SDL_TrayEntry *entry);
void wrap_SDL_SetTrayEntryCallback(SDL_TrayEntry *entry, uintptr_t userdata) {
	SDL_SetTrayEntryCallback(entry, cb_TrayCallback, (void *)userdata);
}

/* video */
extern SDL_EGLAttrib *cb_EGLAttribArrayCallback(void *userdata);
extern SDL_EGLint *cb_EGLSurfaceArrayCallback(void *userdata, SDL_EGLDisplay display, SDL_EGLConfig config);
extern SDL_EGLint *cb_EGLContextArrayCallback(void *userdata, SDL_EGLDisplay display, SDL_EGLConfig config);
void wrap_SDL_EGL_SetAttributeCallbacks(uintptr_t userdata) {
	SDL_EGL_SetAttributeCallbacks(cb_EGLAttribArrayCallback, cb_EGLSurfaceArrayCallback, cb_EGLContextArrayCallback, (void *)userdata);
}
