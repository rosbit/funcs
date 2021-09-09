package funcs

func Map(m func(a interface{})(b interface{}), arr ...interface{}) (it <-chan interface{}) {
	res := make(chan interface{})

	go func() {
		for val := range iter(arr...) {
			res <- m(val)
		}

		close(res)
	}()

	return res
}

