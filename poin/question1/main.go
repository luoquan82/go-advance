package main

func Add10(x *int) {
	if x == nil {
		panic("nil pointer dereference")
	}
	*x += 10
}

func main() {
	// Example usage
	var num int = 5
	Add10(&num)
	println(num) // Output: 15
}
