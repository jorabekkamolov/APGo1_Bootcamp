package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Dasturni ishga tushurish uchun ---> go run main.go <K> <--- K bu yerda butun musbat sonlar")
		os.Exit(1)
	}

	K, err := strconv.Atoi(os.Args[1])
	if err != nil || K <= 0 {
		fmt.Println("K Musbat son bo'lishi kerak")
		os.Exit(1)
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	start := time.Now()
	i := 1

	go func() {
		for {
			time.Sleep(time.Duration(K) * time.Second)
			fmt.Printf("Tick %d since %d\n", i, int(time.Since(start).Seconds()))
			i++
		}
	}()

	<-sigChan
	fmt.Println("Termination")
}
