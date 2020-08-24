package gdic

import (
	"reflect"
)

type service struct {
	provider interface{}
	instance interface{}
}

func (svc *service) getInstance() (interface{}, error) {
	if svc.instance == nil {
		var providerType = reflect.TypeOf(svc.provider)
		var argumentsCount = providerType.NumIn()
		var arguments = make([]reflect.Value, argumentsCount)
		for i := 0; i < argumentsCount; i++ {
			var s = getService(reflect.New(providerType.In(i).Elem()).Interface())
			if s == nil {
				return nil, &ErrNoProviderFound{name: providerType.In(i).String()}
			}
			if instance, err := s.getInstance(); err == nil {
				arguments[i] = reflect.ValueOf(instance)
			} else {
				return nil, err
			}
		}
		var results = reflect.ValueOf(svc.provider).Call(arguments)
		svc.instance = results[0].Interface()
	}
	return svc.instance, nil
}
