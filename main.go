package main

import (
	"fmt"
	"math/rand"
)

type tree struct {
	value    int
	isOnFire bool
	age      int // P spalenia drzewa jest odwrotnie proporcjalne do jego wieku, gdy drzewo ma 300 lat to P = 1%, gdy 1 rok to P = 100%
}

type forest struct {
	dimensions int
	trees      [][]tree
	windDir    int // 0 - NW, 1 - N, 2 - NE, 3 - E, 4 - SE, 5 - S, 6 - SW, 7 - W, to podwaja P spalenia drzew na kierunkach wiatru
	windSpeed  int // windSpeed / 10 okresla ile pol dalej drzewa beda narazone na spalenie
}

func (f *forest) populateForest(treeProbability float32, forestSize, windDir, windSpeed int) {
	f.dimensions = forestSize
	f.trees = make([][]tree, f.dimensions)
	f.windDir = windDir
	f.windSpeed = windSpeed

	for i := range f.trees {
		f.trees[i] = make([]tree, f.dimensions)
		for j := range f.trees[i] {
			if rand.Float32() < treeProbability {
				f.trees[i][j].value = 1
				f.trees[i][j].age = rand.Intn(300)
			}
		}
	}
}

func (t *tree) setOnFire() {
	t.value = -1
	t.isOnFire = true
}

func (f *forest) lightningStrike() {
	x := rand.Intn(f.dimensions)
	y := rand.Intn(f.dimensions)
	if f.didLigthningHitTree(x, y) {
		f.trees[x][y].value = -2
		f.burnForest(x, y)
	}
}

func (f *forest) didLigthningHitTree(x, y int) bool {
	return f.trees[x][y].value == 1
}

func (f *forest) burnForest(x, y int) {
	windSpeedIndex := f.windSpeed / 10
💀
}

func (f *forest) printForest() {
	for i := 0; i < len(f.trees); i++ {
		for j := 0; j < len(f.trees[i]); j++ {
			switch f.trees[i][j].value {
			case 1:
				age := f.trees[i][j].age
				if age >= 200 {
					fmt.Print("🌲 ")
				} else if age >= 100 {
					fmt.Print("🌳 ")
				} else {
					fmt.Print("🌿 ")
				}
			case 0:
				fmt.Print("🟫 ")
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
}

func main() {
	var forest forest
	forest.populateForest(0.89, 10, 0, 30)
	forest.printForest()
	fmt.Println()
	forest.lightningStrike()
	forest.printForest()
}
