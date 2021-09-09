package funcs

func IIF(cond bool, ifVal interface{}, elseVal interface{}) (retVal interface{}) {
	if cond {
		retVal = ifVal
	} else {
		retVal = elseVal
	}
	return
}
