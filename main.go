package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	CHUNKS = 8
	SIZE   = 100_000_000
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)

	for i := range data {
		data[i] = rand.Intn(1_000_000) + 1
	}

	return data
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]

	for _, value := range data {
		if value > max {
			max = value
		}
	}

	return max
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	var wg sync.WaitGroup

	maxima := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * len(data) / CHUNKS
		end := (i + 1) * len(data) / CHUNKS

		wg.Add(1)

		go func(chunk []int, index int) {
			defer wg.Done()

			maxima[index] = maximum(chunk)
		}(data[start:end], i)
	}

	wg.Wait()

	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)

	data := generateRandomElements(SIZE)

	fmt.Println("Поиск максимума в один поток")

	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение: %d\n", max)
	fmt.Printf("Время выполнения: %d мкс\n\n", elapsed)

	fmt.Printf("Поиск максимума в %d потоков\n", CHUNKS)

	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение: %d\n", max)
	fmt.Printf("Время выполнения: %d мкс\n", elapsed)
}
