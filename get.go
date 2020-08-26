package gdic

import (
	"reflect"
)

// Get returns existing or creates new instance of specified type
func Get(i interface{}) (interface{}, error) {
	var svc = getService(i)
	if svc == nil {
		return nil, &ErrNoProviderFound{name: reflect.TypeOf(i).Elem().String()}
	}
	return svc.getInstance()
}

// GetOrPanic returns existing or creates new instance of specified type or calls panic on error
func GetOrPanic(i interface{}) interface{} {
	if instance, err := Get(i); err == nil {
		return instance
	} else {
		panic(err)
	}
}
