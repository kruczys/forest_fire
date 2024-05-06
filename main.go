package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createForest(treeProbability float32, forestSize int) [][]int {
	forest := make([][]int, forestSize)

	for i := range forest {
		forest[i] = make([]int, forestSize)
		for j := range forest[i] {
			if rand.Float32() < treeProbability {
				forest[i][j] = 1
			}
		}
	}

	return forest
}

func printForest(forest [][]int) {
	for i := 0; i < len(forest); i++ {
		fmt.Println(forest[i])
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	forest := createForest(0.3, 5)
	printForest(forest)
}
