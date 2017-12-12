package main

// #include <stdio.h>
// #include <stdlib.h>
/*
void print(char *str){
	printf("%s\n", str);
}
 */
import "C"
import (
	"unsafe"
	"time"
)

func main() {
	s := "hello world"
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.print(cs)
	time.Sleep(5 * time.Second)
}
