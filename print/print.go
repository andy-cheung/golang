package print

/*
#cgo CFLAGS: -Qunused-arguments
#include <stdlib.h>
#include <stdio.h>
void print(char *str) {
    printf("%s\n", str);
}
*/
import "C"

import "unsafe"

func Print(s string) {
	cs := C.CString(s)
	//	C.fputs(cs, (*C.FILE)(C.stdout))
	C.print(cs)
	C.free(unsafe.Pointer(cs))
}
