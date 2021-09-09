package funcs

import (
	"testing"
	"fmt"
)

func test(it <-chan interface{}, prompt string) {
	for v := range it {
		fmt.Printf("%v\n", v)
	}
	fmt.Printf("-- done for %s --\n", prompt)
}

func TestIter(t *testing.T) {
	test(iter([]int{1, 2, 3, 4}), "[]int")
	test(iter(1, 2, 3), "many int")
}

func TestFilter(t *testing.T) {
	a := []int{1, 2, 3, 4}
	sq := Filter(func(i interface{})bool {
		return i.(int) > 2
	}, a)

	test(sq, "filtering")
}

func TestMap(t *testing.T) {
	a := []int{1, 2, 3, 4}
	sq := Map(func(i interface{})(interface{}) {
		return i.(int) * 2
	}, a)

	test(sq, "mapping")
}

func TestReduce(t *testing.T) {
	a := []int{1, 2, 3, 4}
	res := Reduce(0, func(acc interface{}, a interface{})(interface{}) {
		return acc.(int) + a.(int)
	}, a)

	fmt.Printf("reduce(+) %v => %v\n", a, res)

	res = Reduce(1, func(acc interface{}, a interface{})(interface{}) {
		return acc.(int) * a.(int)
	}, a)
	fmt.Printf("reduce(*) %v => %v\n", a, res)
}
