package main

import "fmt"

func test(fn func() int) int {
	return fn()
}

type FormatFunc func(s string, x, y int) string

func format(formatFunc FormatFunc, s string, x, y int) string {
	return formatFunc(s, x, y)
}

func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
func main() {
	t1 := test(func() int {
		return 100
	})
	println(t1)
	t2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d--%d", 10, 20)
	println(t2)
	x, y := 100, 200
	swap(&x, &y)
}
