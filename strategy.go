package main

type Result struct {
	winProbability float64
	tieProbability float64
}

type Strategy interface {
	Run() Result
}
