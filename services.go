package gdic

import "reflect"

var services = map[interface{}]*service{}

func getService(i interface{}) *service {
	var svc *service
	var ok bool
	if svc, ok = services[reflect.TypeOf(i).Elem()]; ok {
		return svc
	}
	return nil
}
