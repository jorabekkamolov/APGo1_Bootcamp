package main

import (
	"fmt"
	"sort"
)

func main() {
	hashMap := make(map[string]int)
	word := []string{"aa", "bb", "cc", "aa", "cc", "cc", "cc", "aa", "ab", "ac", "bb"}
	for _, i := range word {
		hashMap[i]++
	}
	answer := make([]string, 0, len(hashMap))
	for i, _ := range hashMap {
		answer = append(answer, i)
	}

	sort.Slice(answer, func(i, j int) bool {
		if hashMap[answer[i]] == hashMap[answer[j]] {
			return answer[i] < answer[j]
		}
		return hashMap[answer[i]] > hashMap[answer[j]]
	})
	fmt.Println(answer[:3])
}
