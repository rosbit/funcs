package funcs

func Filter(p func(interface{})bool, arr ...interface{}) (it <-chan interface{}) {
	res := make(chan interface{})

	go func() {
		for val := range iter(arr...) {
			if p(val) {
				res <- val
			}
		}

		close(res)
	}()

	return res
}

