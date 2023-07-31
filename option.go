package option

import "fmt"

type Optionable interface {
	any | None
}

type Option[T any] struct {
	value Optionable
}

type None struct{}

func Optional[T any](initial ...T) Option[T] {
	if len(initial) == 0 {
		return Option[T]{None{}}
	}

	return Option[T]{initial[0]}
}

func (o Option[T]) IsNone() bool {
	return o.value == None{}
}

func (o Option[T]) IsSome() bool {
	return o.value != None{}
}

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

func (o Option[T]) ValueOr(defaultVal T) T {
	if o.IsSome() {
		value, _ := o.GetValue()
		return value 
	} else {
		return defaultVal
	}
}

func (o Option[T]) SetValue(value Optionable) Option[T] {
	o.value = value
	return o
}

func getDefault[T any]() T {
	var result T
	return result
}

