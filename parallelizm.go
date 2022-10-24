package main

import (
	"fmt"
	"time"
)

func go_routine_no_wait() {
	go hello_print_no_wait()
}

func hello_print_no_wait() {
	fmt.Println("no body wait for me...")
}

func go_routine_wait() {
	go hello_print_wait()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("finished wait")
}

func hello_print_wait() {
	fmt.Println("somebody wait for me...")
}

func go_routine_wait_with_closure() {
	var print_line = "hello in closure"
	go func() {
		fmt.Println(print_line)
	}()
	print_line = "goodbye in closure"
	time.Sleep(1000 * time.Millisecond)
}
