package paymentchecker

import "math/big"

func SufficientFunds(availableFunds *big.Int, paymentAmount *big.Int, oracleCount uint8) bool {
	min := big.NewInt(int64(oracleCount))
	min = min.Mul(min, paymentAmount)

	return availableFunds.Cmp(min) >= 0
}

// SufficientPayment checks if the available payment is enough to submit an
// answer. It compares the payment amount on chain with the min payment amount
// listed in the job / ENV var.
// func SufficientPayment(payment *big.Int) bool {
// 	payment.Cmp(c.MinContractPayment.ToInt()) >= 0
// }
