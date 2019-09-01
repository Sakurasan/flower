package main

/*
#include <stdio.h>

void hello(char *str) {
    // insert code here...
	// printf("嘿! Hello, World!\n");
	printf("%s \n",str);
}
*/
import (
	"C"
)

func main() {
	cs := C.CString("嘿Cgo!")
	C.hello(cs)
}
