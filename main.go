package main

func main() {
	var forest forest
	var dataAnalysis monteCarlo
	dataAnalysis.prepareExperiment(50, 20)
	dataAnalysis.conductExperiment()
	forest.populateForest(0.85, 40, 200)
	forest.lightningStrike(true)
}
