package option

import "fmt"

type Optionable interface {
	any | None
}

// Wrapper for optional value
type Option[T any] struct {
	value Optionable
}

type None struct{}

// Creates a new Option of type T
func Optional[T any](initial ...T) Option[T] {
	if len(initial) == 0 {
		return Option[T]{None{}}
	}

	return Option[T]{initial[0]}
}

// check if the current option value is None
func (o Option[T]) IsNone() bool {
	return o.value == None{}
}

// check if the current option value is Some
func (o Option[T]) IsSome() bool {
	return o.value != None{}
}

// get the current option value and if exists an error
func (o Option[T]) GetValue() (variable T, err error) {
	rawValue := o.value

	if o.IsNone() {
		v := getDefault[T]()
		return v, fmt.Errorf("option is none")
	} else {
		v := rawValue.(T)
		return v, nil
	}
}

// get the current option value or default value passed
func (o Option[T]) ValueOr(defaultVal T) T {
	if o.IsSome() {
		value := o.value.(T)
		return value 
	} else {
		return defaultVal
	}
}

// get the value or panic the program
func (o Option[T]) Unwrap() T {
	if o.IsNone() {
		panic("option is none")
	}

	return o.value.(T)
}

// set the current option value
func (o Option[T]) SetValue(value Optionable) Option[T] {
	o.value = value
	return o
}

func getDefault[T any]() T {
	var result T
	return result
}

