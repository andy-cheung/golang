package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse() // Scans the arg list and sets up flags
	// for i := 0; i < flag.NArg(); i++ {
	// 	fmt.Println(flag.Arg(i))
	// }
	userFile := flag.Arg(0)
	if userFile == "" {
		return
	}

	fin, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	stat, err := fin.Stat()
	buf := make([]byte, stat.Size())
	total := 0
	for {
		n, _ := fin.Read(buf)
		total += n
		if 0 == n {
			break
		}
	}
	fin.Close()

	fout, err := os.Create(userFile)
	defer fout.Close()
	if buf[0] == 0xef {
		fout.Write(buf[3:])
	} else {
		fout.Write(buf[:])
	}
}
