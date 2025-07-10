package main

import "fmt"

func doubleSliceEle(d []int) {
	if d == nil {
		panic("nil slice dereference")
	}
	for i := range d {
		d[i] *= 2
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	doubleSliceEle(arr)
	fmt.Println(arr) // Output: [2 4 6 8 10]
}
