package main

/*
typedef void (*callbackFunc)(int in_arg);

void
askForCallback(void *func, int in_arg) {
    callbackFunc cb = (callbackFunc)func;
    return cb(in_arg);
}
*/
import "C"
import "unsafe"

func askForCallback(p unsafe.Pointer) {
	C.askForCallback(p, 5)
}
