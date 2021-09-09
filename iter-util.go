package funcs

import (
	"reflect"
)

func iter(vals ...interface{}) (<-chan interface{}) {
	it := make(chan interface{})

	go func() {
		switch len(vals) {
		case 0:
			// no args
			break
		case 1:
			// only 1 arg
			if vals[0] == nil {
				// the arg is nil
				break
			}

			// whether the arg is array or slice
			exit := false
			val := reflect.ValueOf(vals[0])
			switch val.Kind() {
			case reflect.Array, reflect.Slice:
			default:
				// not array/slice, wrap it as a slice
				it <- vals[0]
				exit = true
			}
			if exit {
				break
			}

			// convert array/slice into []interface{}
			arrLen := val.Len()
			if arrLen == 0 {
				break
			}
			for i:=0; i<arrLen; i++ {
				it <- val.Index(i).Interface()
			}
		default:
			// args is []interface{}, return it directly
			for i, _ := range vals {
				it <- vals[i]
			}
		}

		close(it)
	}()

	return it
}
