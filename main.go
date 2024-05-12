package main

func main() {
	// var forest forest
	var dataAnalysis monteCarlo
	// forest.populateForest(0.01, 40, 0)
	// forest.lightningStrike(true)
	dataAnalysis.prepareExperiment(50, 100)
	dataAnalysis.conductExperiment()
}
