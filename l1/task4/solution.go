package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Работник %d завершает работу\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Работник %d: канал закрыт, выход\n", id)
				return
			}
			fmt.Printf("Работник %d делает %d\n", id, job)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Нет данных по количеству требуемых работников\n")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Printf("Ошибка при чтении или некорректное количество работников")
	}

	jobs := make(chan int)

	// Контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i <= n; i++ {
		go worker(ctx, i, jobs)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	count := 1

loop:
	for {
		select {
		case <-stop:
			fmt.Printf("\nСигнал получен, завершение работы...\n\n")
			cancel()
			close(jobs)
			break loop
		default:
			jobs <- count
			count++
			time.Sleep(100 * time.Millisecond)
		}
	}
	time.Sleep(500 * time.Millisecond)
}
