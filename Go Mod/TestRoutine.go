package main

import (
    "fmt"
    "runtime"
)

func print(text string) {
    for i := 0; i < 10; i++ {
        runtime.Gosched()
        fmt.Println(text)
    }
}

func main() {
	runtime.GOMAXPROCS(2)
	go print("Hello") // create new goroutine
	print("World")    // existing goroutine	
	fmt.Println(runtime.NumCPU())
}