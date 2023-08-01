package main

import (
	"fmt"
	"poc-interfaces/pk1"
)

func main() {
	var impl pk1.Interface1

	impl = pk1.Impl1{Value: "Impl1"}
	fmt.Println(impl.SayFoo())

	impl = &pk1.Impl2{OtherValue: "Impl2"}
	fmt.Println(impl.SayFoo())
}
