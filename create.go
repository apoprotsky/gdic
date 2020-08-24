package gdic

import (
	"reflect"
)

// Create used to create new instance of specified type
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
