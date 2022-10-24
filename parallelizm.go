package main

import (
	"fmt"
	"sync"
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
	time.Sleep(100 * time.Millisecond)
	fmt.Println("finished wait")
}

func hello_print_wait() {
	fmt.Println("somebody wait for me...")
}

func go_routine_race_condition() {
	var print_line = "hello in closure1"
	go func() {
		fmt.Println(print_line)
	}()
	print_line = "goodbye in closure1"
	time.Sleep(100 * time.Millisecond)
}

func go_routine_no_race_condition_pass_by_value() {
	var print_line = "hello in closure2"
	go func(msg string) {
		fmt.Println(msg)
	}(print_line)
	print_line = "goodbye in closure2"
}

func go_routine_no_race_condition_with_sync() {
	var wg = sync.WaitGroup{}
	wg.Add(1)
	var print_line = "hello in closure3"
	go func() {
		fmt.Println(print_line)
		wg.Done()
	}()
	wg.Wait()

	print_line = "goodbye in closure3"
}
