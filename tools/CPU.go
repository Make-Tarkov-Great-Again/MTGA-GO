package tools

import "runtime"

// Calculate max number of CPU cores we have available for parallel execution
func CalculateWorkers() int {
	numCPU := runtime.NumCPU()

	// Adjust the parallelism factor based on your workload and requirements
	// For CPU-bound tasks, you can use numCPU
	// For I/O-bound tasks, you can use a higher factor (e.g., numCPU * 2)

	// Ensure a minimum number of workers (e.g., 1)
	minWorkers := 1

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	return max(numCPU, minWorkers)
}
