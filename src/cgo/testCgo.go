package main

import (
	"print"
)

func main() {
	print.Print("hello")
	//	C.fputs(cs, (*C.FILE)(C.stdout))
	C.print(cs)
	C.free(unsafe.Pointer(cs))
}
