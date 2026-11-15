package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements returns array of positive random integers, or nil for a non-positive size.
func generateRandomElements(size int) []int {
	if size <= 0 {
		fmt.Fprintf(os.Stderr, "generateRandomElements: размер слайса должен быть положительным, получено %d\n", size)
		return nil
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(math.MaxInt) + 1
	}
	return data
}

// maximum returns the largest value in data, or zero if data is empty.
func maximum(data []int) int {
	if len(data) == 0 {
		fmt.Fprintln(os.Stderr, "maximum: невозможно найти максимум в пустом слайсе")
		return 0
	}

	max := data[0]
	for _, value := range data[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

// maxChunks finds the maximum using 8 goroutines, or returns zero for empty data.
func maxChunks(data []int) int {
	if len(data) == 0 {
		fmt.Fprintln(os.Stderr, "maxChunks: невозможно найти максимум в пустом слайсе")
		return 0
	}

	chunkSize := len(data) / CHUNKS
	results := make(chan int, CHUNKS)
	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := range CHUNKS {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}

		go func(chunk []int) {
			defer wg.Done()
			if len(chunk) == 0 {
				return
			}
			results <- maximum(chunk)
		}(data[start:end])
	}

	wg.Wait()
	close(results)

	max := data[0]
	for value := range results {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
