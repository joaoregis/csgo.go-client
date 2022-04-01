package main

import (
	"fmt"
	"time"
)

func Say(s string, done chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	done <- "Terminei"
}

func RunChannels() {
	done := make(chan string)
	go Say("world", done)
	fmt.Println(<-done)
}
