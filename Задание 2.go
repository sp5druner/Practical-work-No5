package main

import (
	"fmt"
	"time"
)

var done = make(chan struct{})

func printer() {
	ticker := time.NewTicker(1 * time.Second) // 7.2, стр. 4
	defer ticker.Stop()                       // Остановка go-подпрограммы ticker

	for {
		select {
		case <-done:
			fmt.Println("stop")
			return
		case <-ticker.C:
			fmt.Println("12312211")
		}
	}
}

func main() {
	go printer()
	time.Sleep(3500 * time.Millisecond)
	fmt.Println("done")
	close(done)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("close")
}
