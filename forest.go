package main

import (
	"fmt"
	"math/rand"
	"time"
)

type forest struct {
	dimensions      int
	trees           [][]tree
	windSpeed       int // windSpeed / 10 okresla ile bedzie prob spalenia pojedynczego drzewa
	treesTotal      int
	burntTreesTotal int
}

func (f *forest) populateForest(treeProbability float32, forestSize, windSpeed int) {
	f.dimensions = forestSize
	f.trees = make([][]tree, f.dimensions)
	f.windSpeed = windSpeed

	for i := range f.trees {
		f.trees[i] = make([]tree, f.dimensions)
		for j := range f.trees[i] {
			if rand.Float32() < treeProbability {
				f.trees[i][j].value = 1
				f.trees[i][j].age = rand.Intn(101)
				f.treesTotal += 1
			}
		}
	}
}

func (f *forest) lightningStrike(animate bool) {
	var x, y int
	for {
		x = rand.Intn(f.dimensions)
		y = rand.Intn(f.dimensions)
		if f.didLigthningHitTree(x, y) {
			break
		}
	}
	f.trees[x][y].value = -2
	f.burnForest(x, y, animate)
}

func (f *forest) didLigthningHitTree(x, y int) bool {
	return f.trees[x][y].value == 1
}

func (f *forest) burnForest(x, y int, animate bool) {
	if animate {
		f.printForest()
	}

	dx := []int{-1, 0, 1}
	dy := []int{-1, 0, 1}

	for i := 0; i < f.windSpeed/10+1; i++ {
		newX := x + dx[rand.Intn(3)]
		newY := y + dy[rand.Intn(3)]

		if newX < 0 || newX >= f.dimensions || newY < 0 || newY >= f.dimensions {
			continue
		}

		if f.trees[newX][newY].tryToSetOnFire() {
			f.burntTreesTotal += 1
			f.burnForest(newX, newY, animate)
		}
	}
}

func (f *forest) calculateBurnRatio() float32 {
	return float32(f.burntTreesTotal) / float32(f.treesTotal)
}

func (f *forest) printForest() {
	time.Sleep(10 * time.Millisecond)
	clearScreen()

	for i := 0; i < len(f.trees); i++ {
		for j := 0; j < len(f.trees[i]); j++ {
			switch f.trees[i][j].value {
			case 1:
				fmt.Print("🌲 ")
			case 0:
				fmt.Print("   ")
			case -1:
				fmt.Print("🔥 ")
			case -2:
				fmt.Print("🌩️  ")
			default:
				fmt.Print(f.trees[i][j], "  ")
			}
		}
		fmt.Println()
	}

	burnRatio := f.calculateBurnRatio()
	fmt.Printf("\033[31mBurn ratio: %.2f%% \033[0m", burnRatio*100)
	fmt.Print("\n\n")
}
