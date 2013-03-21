package main

import "fmt"
import "time"

func pump(ch chan int) {
	for i := 0; i < 10; i++ {
		//time.Sleep(2 * 1e9)	// sleep 2 sec
		println(i)
		ch <- i
	}
	
}
func suck(ch chan int) {
	for { fmt.Println(<-ch)}
}

func main() {
    ch := make(chan int)
    go pump(ch)
    go suck(ch)
    time.Sleep(10 * 1e9)
}