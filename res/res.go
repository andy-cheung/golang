package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "icon_zhoujijin.png"
	fin, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	stat, err := fin.Stat()
	buf := make([]byte, stat.Size())

	for {
		n, _ := fin.Read(buf)
		if 0 == n {
			break
		}
	}
	fin.Close()

	const HeadSize = 8
	count := stat.Size() - HeadSize
	for index := int64(0); index < count; index++ {
		buf1 := buf[index:]
		ok := buf1[5] == buf1[7]

		for ci := 1; ci < HeadSize; ci++ {
			if buf1[ci] == buf1[ci-1] {
				ok = false
				break
			}
		}
		if !ok {
			continue
		}

		fmt.Println(index)
		for i := 0; i < i; i++ {
			fmt.Printf("%d ", buf1[i])
		}
		fmt.Println()
	}
	// if buf[0] == 0xef {
	// 	fout, _ := os.Create("tmp.png")
	// 	defer fout.Close()
	// 	fout.Write(buf[3:])
	// }
}
