package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	r1 := bufio.NewReader(os.Stdin)
	input1, _ := r1.ReadString('\n')
	input1 = strings.TrimSpace(input1)

	array1 := strings.Split(input1, " ")
	hashMap := make(map[int]bool)
	for _, j := range array1 {
		num, err := strconv.Atoi(j)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		hashMap[num] = true
	}

	r2 := bufio.NewReader(os.Stdin)
	input2, _ := r2.ReadString('\n')
	input2 = strings.TrimSpace(input2)

	array2 := strings.Split(input2, " ")
	answer := make([]int, 0)

	for _, j := range array2 {
		num, err := strconv.Atoi(j)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		if hashMap[num] {
			answer = append(answer, num)
			hashMap[num] = false
		}
	}
	sort.Ints(answer)
	fmt.Println(answer)
}
