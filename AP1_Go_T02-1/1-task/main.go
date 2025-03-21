// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"sort"
// 	"strconv"
// 	"sync"
// 	"time"
// )

// func main() {
// 	if len(os.Args) < 3 {
// 		fmt.Println("Foydalnish uchun ---> go run main.go <N> <M> <--- N va M bu butun musbat son")
// 		return
// 	}

// 	N, err1 := strconv.Atoi(os.Args[1])
// 	M, err2 := strconv.Atoi(os.Args[2])

// 	if err1 != nil || err2 != nil || N <= 0 || M <= 0 {
// 		fmt.Println("N va M musbat butun sonlar bo'lishi kerak !!!")
// 		return
// 	}

// 	result := make([]struct {
// 		index int
// 		sleep int
// 	}, N)

// 	var wsg sync.WaitGroup
// 	wsg.Add(N)
// 	for i := range N {
// 		go func(i int) {
// 			defer wsg.Done()
// 			sleepTime := rand.Intn(M + 1)
// 			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
// 			result[i] = struct {
// 				index int
// 				sleep int
// 			}{i, sleepTime}
// 		}(i)
// 	}
// 	wsg.Wait()

// 	sort.Slice(result, func(i, j int) bool {
// 		return result[i].sleep > result[j].sleep
// 	})

// 	for _, res := range result {
// 		fmt.Printf("Gorutina: %d, Uyqu vaqti: %d ms\n", res.index, res.sleep)
// 	}
// }

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

func main() {
	N := flag.Int("N", 0, "Gorutinalar soni")
	M := flag.Int("M", 0, "Maksimal uyqu vaqti (ms)")
	flag.Parse()

	if *N <= 0 || *M <= 0 {
		fmt.Println("N va M musbat butun son bo'lishi kerak!")
		return
	}
	results := make([]struct {
		index int
		sleep int
	}, *N)
	var wg sync.WaitGroup
	wg.Add(*N)

	for i := range *N {
		go func(i int) {
			defer wg.Done()
			sleepTime := rand.Intn(*M + 1)
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
			results[i] = struct {
				index int
				sleep int
			}{i, sleepTime}
		}(i)
	}

	wg.Wait()
	sort.Slice(results, func(i, j int) bool {
		return results[i].sleep > results[j].sleep
	})

	fmt.Println("Natijalar (kamayish tartibida):")
	for _, res := range results {
		fmt.Printf("Gorutina: %d, Uyqu vaqti: %d ms\n", res.index, res.sleep)
	}
}
