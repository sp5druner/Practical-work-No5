package main

import (
	"fmt"
	"time"
)

func request(result chan<- string) {
	fmt.Println("сервер обрабатывает запрос")
	time.Sleep(3 * time.Second)
	result <- "ответ сервера"
}

func main() {
	result := make(chan string)
	go request(result)
	select {
	case msg := <-result:
		fmt.Println("Получено:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Таймаут")
	}
}
