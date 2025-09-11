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
		fmt.Printf("Нет данных по лимиту таймера\n")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Printf("Ошибка при чтении или некорректное значение времени (количество секунд)")
		return
	}

	jobs := make(chan int)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs)
	}

	timer := time.After(time.Duration(n) * time.Second)
	count := 1

	for {
		select {
		case <-timer:
			fmt.Printf("Время вышло")
			close(jobs)
			return
		default:
			jobs <- count
			count++
			time.Sleep(500 * time.Millisecond)
		}
	}
}
