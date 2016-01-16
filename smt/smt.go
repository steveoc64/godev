package smt

import (
	"runtime"
)

func Init() int {

	numCores := runtime.NumCPU()

	useCores := numCores
	// If we are on a REAL machine, then dont get too greedy - leave one core alone
	if numCores > 2 {
		useCores--
	}
	runtime.GOMAXPROCS(useCores)
	return useCores
}
