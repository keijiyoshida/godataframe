package dataframe

import "runtime"

var numConcurrency = runtime.GOMAXPROCS(0)

// SetNumConcurrency sets the number of concurrency of the data frame processing.
// The default value of this parameter is the maximum number of CPUs
// that can be executing simultaneously.
func SetNumConcurrency(n int) {
	numConcurrency = n
}
