package main

import (
	"fmt"
	"math/rand"
)

func createForest(treeProbability float32, forestSize int) (forest [][]int, err error) {
	forest = make([][]int, forestSize)

	for i := range forest {
		forest[i] = make([]int, forestSize)
		for j := range forest[i] {
			if rand.Float32() < treeProbability {
				forest[i][j] = 1
			}
		}
	}

	return
}

func lightningStrike(forestSize int) (x, y int) {
	x = forestSize
	return
}

func printForest(forest [][]int) {
	for i := 0; i < len(forest); i++ {
		fmt.Println(forest[i])
	}
}

func main() {
	forest, err := createForest(0.3, 5)
	if err == nil {
		printForest(forest)
	}
}
