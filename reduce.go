package funcs

func Reduce(initVal interface{}, f func(acc interface{}, a interface{})(res interface{}), arr ...interface{}) (total interface{}) {
	total = initVal
	for val := range iter(arr...) {
		total = f(total, val)
	}

	return
}

