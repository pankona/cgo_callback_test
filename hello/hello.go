package main

import "fmt"
import "unsafe"
import "C"

var funcPointer unsafe.Pointer

//export Hello
func Hello(p unsafe.Pointer) {
	fmt.Println("Hello! This is from Golang!")
	funcPointer = p
}

//export InvokeCallback
func InvokeCallback() {
	go askForCallback(funcPointer)
}

func main() {}
