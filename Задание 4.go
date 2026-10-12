package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID       string
	Name     string
	Priority int
}

func worker(name string, in <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range in {
		fmt.Printf("[%s] выполняю задачу %s (%s)\n", name, t.ID, t.Name)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Printf("[%s] канал закрыт, воркер завершает работу\n", name)
}

func main() {
	urgent := make(chan Task, 5)
	normal := make(chan Task, 10)

	var wg sync.WaitGroup
	wg.Add(2)
	go worker("СРОЧНЫЙ", urgent, &wg)
	go worker("ОБЫЧНЫЙ", normal, &wg)

	queue := []Task{
		{"#1", "обычный заказ", 2},
		{"#2", "срочный заказ", 1},
		{"#3", "обычный заказ", 2},
		{"#4", "срочный заказ", 1},
		{"#5", "обычный заказ", 2},
	}

	for _, t := range queue {
		if t.Priority == 1 {
			urgent <- t
		} else {
			normal <- t
		}
	}

	time.Sleep(3 * time.Second)
	close(urgent)
	close(normal)

	wg.Wait()
	fmt.Println("Все воркеры корректно завершены.")
}
