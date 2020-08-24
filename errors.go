package gdic

// ErrNoProviderFound error
type ErrNoProviderFound struct {
	name string
}

func (e *ErrNoProviderFound) Error() string {
	return "provider for " + e.name + " not found"
}

// ErrProviderNotFunction error
type ErrProviderNotFunction struct {
	name string
}

func (e *ErrProviderNotFunction) Error() string {
	return "provider '" + e.name + "' is not a function"
}

// ErrProviderReturnsTooManyValues error
type ErrProviderReturnsTooManyValues struct {
	name string
}

func (e *ErrProviderReturnsTooManyValues) Error() string {
	return "provider '" + e.name + "' returns more then two values"
}

// ErrProviderReturnsNonErrorSecondValue error
type ErrProviderReturnsNonErrorSecondValue struct {
	name string
}

func (e *ErrProviderReturnsNonErrorSecondValue) Error() string {
	return "provider '" + e.name + "' must return error as second value"
}
