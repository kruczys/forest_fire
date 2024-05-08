package main

import (
	"fmt"
	"math/rand"
)

type forest struct {
	dimensions int
	trees      [][]int
}

func (f *forest) populateForest(treeProbability float32, forestSize int) {
	f.dimensions = forestSize
	f.trees = make([][]int, f.dimensions)

	for i := range f.trees {
		f.trees[i] = make([]int, f.dimensions)
		for j := range f.trees[i] {
			if rand.Float32() < treeProbability {
				f.trees[i][j] = 1
			}
		}
	}
}

func lightningStrikeCoords(forestSize int) (x, y int) {
	x = rand.Intn(forestSize)
	y = rand.Intn(forestSize)
	return
}

func (f *forest) didLigthningHitTree(x, y int) bool {
	return f.trees[x][y] == 1
}

func (f *forest) printForest() {
	for i := 0; i < len(f.trees); i++ {
		for j := 0; j < len(f.trees[i]); j++ {
			switch f.trees[i][j] {
			case 1:
				fmt.Print("\033[32m4\033[0m ")
			case 0:
				fmt.Print("\033[30m0\033[0m ")
			case -1:
				fmt.Print("\033[31m-4\033[0m ")
			default:
				fmt.Print(f.trees[i][j], " ")
			}
		}
		fmt.Println()
	}
}

func main() {
	var forest forest
	forest.populateForest(0.3, 10)
	forest.printForest()
}
