// Package workspace_lib provides essential (mock) features to work related apps.
package workspace_lib

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// AddNums returns sum of two finite numbers.
// https://www.mathsisfun.com/numbers/addition.html
func AddNums[N Number](a, b N) N {
	return a + b
}

// SubNums subtracts two finite numbers and returns result.
func SubNums(a, b int) int {
	return a - b
}
