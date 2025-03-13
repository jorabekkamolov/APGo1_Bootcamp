package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("1 - Sonni kiriting: ")
	var first float64
	if _, err := fmt.Scanf("%f", &first); err != nil {
		log.Println("Invalid input")
		return
	}
	fmt.Printf("Kiritilgan son: %.3f\n", first)
	fmt.Print("Operation kiriting (+, -, *, /): ")
	op, _, err := reader.ReadRune()
	if err != nil {
		log.Println("Xatolik: noto'g'ri operator!")
		return
	}

	fmt.Println("Tanlangan operator:", string(op))

	fmt.Print("2 - Sonni kiriting: ")
	var second float64
	if _, err := fmt.Scanf("%f", &second); err != nil {
		log.Println("Invalid input")
		return
	}

	fmt.Printf("Kiritilgan son: %.3f\n", second)
	var result float64
	var validOp bool = true

	switch op {
	case '+':
		result = first + second
	case '-':
		result = first - second
	case '*':
		result = first * second
	case '/':
		if second == 0 {
			log.Println("Invalid input")
			return
		}
		result = first / second
	default:
		validOp = false
		log.Println("Invalid input")
	}

	if validOp {
		fmt.Printf("Natija:\n%.3f %c %.3f = %.3f\n", first, op, second, result)
	}
}
