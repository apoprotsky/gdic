package gdic

import (
	"reflect"
)

// Get used for creating or retrieving instance of specified type
func Get(i interface{}) (interface{}, error) {
	var svc = getService(i)
	if svc == nil {
		return nil, &ErrNoProviderFound{name: reflect.TypeOf(i).Elem().String()}
	}
	return svc.getInstance()
}
