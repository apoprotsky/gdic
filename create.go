package gdic

import (
	"reflect"
)

// Create creates new instance of specified type
func Create(i interface{}) (interface{}, error) {
	var svc = getService(i)
	if svc == nil {
		return nil, &ErrNoProviderFound{name: reflect.TypeOf(i).Elem().String()}
	}
	if svc.instance != nil {
		svc.instance = nil
	}
	return svc.getInstance()
}

// CreateOrPanic creates new instance of specified type or calls panic on error
func CreateOrPanic(i interface{}) interface{} {
	if instance, err := Create(i); err == nil {
		return instance
	} else {
		panic(err.Error())
	}
}
