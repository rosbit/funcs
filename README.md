# funcs: some useful FP functions

## usage

```go
import (
	"github.com/rosbit/funcs"
)

func main() {
	a := []int{1,2,3,4}
	d := funcs.Map(func(i interface{})interface{} {
		return i.(int)*2
	}, a)

	acc := funcs.Reduce(0, func(ac interface{}, i interface{})interface{} {
		return ac.(int) + i.(int)
	}, a)

	r := funcs.Filter(func(i interface{})bool {
		return i.(int) > 2
	}, a)

	res := funcs.IFF(len(a)>0, "has items", "no items")
```
