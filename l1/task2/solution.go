package main

import (
	"fmt"
)

func square(workerID int, jobs <-chan int, res chan<- int) {
	for num := range jobs {
		fmt.Printf("Горутина №%d начала обрабатывать число %d\n", workerID, num)
		res <- num * num
	}
}

func main() {
	arr := []int{2, 4, 6, 8, 10}

	jobs := make(chan int, len(arr))
	results := make(chan int, len(arr))

	for w := 1; w <= 4; w++ {
		go square(w, jobs, results)
	}

	for _, num := range arr {
		jobs <- num
	}

	close(jobs)

	for i := 0; i < len(arr); i++ {
		fmt.Println("Результат:", <-results)
	}
}
