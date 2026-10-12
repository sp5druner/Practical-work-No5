package main

import "fmt"

func main() {
	ch := make(chan int)
	select {
	case x := <-ch:
		fmt.Println("Получено:", x)
	default:
		fmt.Println("Данных нет")
	}
	ch2 := make(chan int, 1)
	ch2 <- 67
	select {
	case x := <-ch2:
		fmt.Println("Получено:", x)
	default:
		fmt.Println("Данных нет")
	}
}
