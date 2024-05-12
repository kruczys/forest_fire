package main

import (
	"math/rand"
)

type tree struct {
	value int
	age   int // Prawdopodobienstwo spalenia drzewa jest odwrotnie proporcjalne do jego wieku, gdy drzewo ma 100 lat to P = 1%, gdy 1 rok to P = 100%
}

func (t *tree) tryToSetOnFire() (success bool) { // funkcja przeprowadzajaca probe zapalenia drzewa
	chanceByAge := t.age <= rand.Intn(101)
	if chanceByAge && t.value == 1 {
		t.value = -1
		success = true
	}
	return
}
