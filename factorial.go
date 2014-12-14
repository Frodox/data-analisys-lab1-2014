package main

import (
	"math/big"
)

func FactorialBig(n uint64) (r *big.Int) {

	r = big.NewInt(1)
	one, bn := big.NewInt(1), new(big.Int).SetUint64(n)

	if bn.Cmp(one) <= 0 {
		return
	}

	for i := big.NewInt(2); i.Cmp(bn) <= 0; i.Add(i, one) {
		r.Mul(r, i)
	}

	return
}
