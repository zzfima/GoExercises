package main

import (
	"fmt"
	"time"
)

func do_parallel() {
	go hello_print()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("finished wait")
}

func hello_print() {
	fmt.Println("Hello World")
}
