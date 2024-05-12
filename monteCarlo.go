package main

import "fmt"

type monteCarlo struct {
	forestSize int
	sampleSize int
}

func (m *monteCarlo) prepareExperiment(forestSize, sampleSize int) { // funkcja przygotowuje eksperyment
	m.forestSize = forestSize
	m.sampleSize = sampleSize
}

func (m *monteCarlo) conductExperiment() { // funkcja przeprowadza eksperyment losowy i wyswietla rezultaty za pomoca innej funkcji
	var currentWindSpeed int
	for k := 0; k < 21; k++ {
		var forestPercentage float32 = 0.1
		for i := 0; i < 10; i++ {
			var averageBurnRate float32
			for j := 0; j < m.sampleSize; j++ {
				var currentForest forest
				currentForest.populateForest(forestPercentage, m.forestSize, currentWindSpeed)
				currentForest.lightningStrike(false)
				averageBurnRate += currentForest.calculateBurnRatio()
			}
			averageBurnRate /= float32(m.sampleSize)
			m.printResults(averageBurnRate, forestPercentage, currentWindSpeed)
			forestPercentage += 0.1
		}
		currentWindSpeed += 50
	}
}

func (m *monteCarlo) printResults(burnRate, forestPercentage float32, windSpeed int) {
	fmt.Printf("Wind speed: %d, forest percentage: %.2f%%, average burn rate: %.2f%%\n", windSpeed, forestPercentage*100, burnRate*100)
}
