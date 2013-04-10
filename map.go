package functional
// Implements map and reduce operations

import (
	"reflect"
)

func Map(fn interface{}, arg interface{}) interface{} {
	fnVal := reflect.ValueOf(fn)
	sVal := reflect.ValueOf(arg)

	styp := reflect.SliceOf(fnVal.Type().Out(0))
	ns := reflect.MakeSlice(styp, sVal.Len(), sVal.Len())

	for j := 0; j < sVal.Len(); j++ {
		val := fnVal.Call([]reflect.Value{sVal.Index(j)})[0]
		ns.Index(j).Set(val)
	}
	return ns.Interface()
}


