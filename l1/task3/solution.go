package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Работник %d делает %d\n", id, job)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Нет данных\n")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Printf("Ошибка при чтении или некорректное количество работников")
		return
	}

	jobs := make(chan int)

	for i := 1; i <= n; i++ {
		go worker(i, jobs)
	}

	count := 1

	for {
		jobs <- count
		count++
		time.Sleep(100 * time.Millisecond)
	}
}
