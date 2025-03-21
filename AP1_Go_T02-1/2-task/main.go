package main

import (
	"fmt"
	"os"
	"strconv"
)

func generator(K, N int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := K; i <= N; i++ {
			ch <- i
		}
	}()
	return ch
}

func squarer(ch1 <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := range ch1 {
			ch <- i * i
		}
	}()
	return ch
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Dasturni ishga tushurish uchun ---> go run main.go <K> <N> <--- K va N bu yerda butun musbat sonlar")
		return
	}

	K, err1 := strconv.Atoi(os.Args[1])
	N, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil || K < 0 || K > N || N <= 0 {
		fmt.Println("K va N musbat butun sonlar bo'lishi kerak --- N katta bolsin K dan --> N < K <--")
		return
	}

	ch1 := generator(K, N)
	ch2 := squarer(ch1)

	for i := range ch2 {
		fmt.Println(i)
	}
}
