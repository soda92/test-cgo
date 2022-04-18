package main

// #include "myheader.h"
import "C"
import (
	"math"
	"unsafe"
)

func Abs(t int) int {
	return int(math.Abs(float64(t)))
}

func main() {
	cs := C.CString("Hello from stdio")
	C.myprint(cs)
	println(Abs(-22))
	C.free(unsafe.Pointer(cs))
}
