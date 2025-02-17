package sdl

import "runtime/cgo"

func (props PropertiesID) SetGoValue(name string, value any) error {
	return props.SetPointerWithCleanup(name, uintptr(cgo.NewHandle(value)), propHandleCleanup)
}

func (props PropertiesID) GoValue(name string) (value any) {
	p := props.Pointer(name, nil)
	if p == nil {
		return nil
	}
	defer func() {
		if r := recover(); r != nil {
			value = nil
		}
	}()
	return cgo.Handle(p).Value()
}
