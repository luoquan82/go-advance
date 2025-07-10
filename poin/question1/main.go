package main

func Add10(x *int) *int {
	if x == nil {
		panic("nil pointer dereference")
	}
	*x += 10
	return x
}

func main() {
	// Example usage
	var num int = 5
	result := Add10(&num)
	println(*result) // Output: 15
}
