package functional
// Memo wrapper for arbitrary functions
// Requires Go 1.1+

import (
	"reflect"
)

func Memo(fn interface{}) interface{} {
	cache := make(map[interface{}]reflect.Value)
	fnVal := reflect.ValueOf(fn)
	memo := func(in []reflect.Value) []reflect.Value {
		arg := in[0].Interface()
		result, ok := cache[arg]
		if ok {
			return []reflect.Value{result}
		}

		result = fnVal.Call(in)[0]
		cache[arg] = result
		return []reflect.Value{result}
	}
	return reflect.MakeFunc(reflect.TypeOf(fn), memo).Interface()
}
