package gdic

import (
	"reflect"
)

// RegisterProvider register function than provides functionality to create instance of specified type
func RegisterProvider(provider interface{}) error {
	providerType := reflect.TypeOf(provider)
	if providerType.Kind() != reflect.Func {
		return &ErrProviderNotFunction{name: providerType.String()}
	}
	if providerType.NumOut() > 2 {
		return &ErrProviderReturnsTooManyValues{name: providerType.String()}
	}
	var instance reflect.Type
	if providerType.NumOut() > 1 {
		instance = providerType.Out(1)
		if !instance.Implements(reflect.TypeOf((*error)(nil)).Elem()) {
			return &ErrProviderReturnsNonErrorSecondValue{name: providerType.String()}
		}
	}
	instance = providerType.Out(0).Elem()
	if _, ok := services[instance]; !ok {
		services[instance] = &service{provider: provider}
	}
	return nil
}
