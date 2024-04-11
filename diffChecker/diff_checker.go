package diffchecker

import (
	"math/big"
)

func CheckDifference(current, next *big.Int) bool {
	// Avoid division by zero
	if current.Sign() == 0 {
		return false
	}

	// Calculate the difference: diff = (next - current)
	diff := new(big.Int).Sub(next, current)

	// Calculate the absolute value of the difference
	diffAbs := new(big.Int).Abs(diff)

	// Calculate the threshold for 0.5% of the current value: threshold = current * 0.005
	threshold := new(big.Int).Div(new(big.Int).Mul(current, big.NewInt(5)), big.NewInt(1000))

	// Check if the absolute difference is at least the threshold
	return diffAbs.Cmp(threshold) >= 0
}
