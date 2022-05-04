package goctory


import "reflect"

type AttrFunc[T any] func(*T)

type AttrFuncs[T any] []AttrFunc[T]

func WithAttr[T any, V any](key string, value V) AttrFunc[T] {
	return func(obj *T) {
		item := reflect.ValueOf(obj).Elem().FieldByName(key)
		if item.CanSet() {
			item.Set(reflect.ValueOf(value))
		}
	}
}

func (afs AttrFuncs[T]) Apply(data *T) {
	for _, af := range afs {
		af(data)
	}
}

func NewStruct[T any](opts ...AttrFunc[T]) *T {
	data := new(T)
	AttrFuncs[T](opts).Apply(data)
	return data
}