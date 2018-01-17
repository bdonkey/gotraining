// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show the basic concept of using a pointer
// to share data.
package main

func main() {

	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")

	// Pass the "address of" count.
	// 2018-01-16 ss passing address of count - but still considered passing by value
	increment(&count)

	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")
}

// increment declares count as a pointer variable whose value is
// always an address and points to values of type int.
// ss 2018-01-16 *int is a type, not an operator. All types get their pointer type automatically
//go:noinline
func increment(inc *int) {

	// Increment the "value of" count that the "pointer points to".
	// ss 2018-01-16 here the * is an deferencing operator
	// allows this function to change memory in another function stack frame
	*inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}
