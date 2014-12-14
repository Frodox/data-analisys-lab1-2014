package main

import (
	"testing"
	"math/big"
	"fmt"
)


func TestAllCasesFactorialBig (t *testing.T) {
	var bres *big.Int
	var val uint64
	var res int64

	val = 0; res = 1
	bres = FactorialBig(val)
	if 0 != bres.Cmp(big.NewInt(res)) {
		t.Errorf("expect %v but result is %v", res, bres)
	}

	val = 1; res = 1
	bres = FactorialBig(val)
	if 0 != bres.Cmp(big.NewInt(res)) {
		t.Errorf("expect %v but result is %v", res, bres)
	}

	val = 7; res = 5040
	bres = FactorialBig(val)
	if 0 != bres.Cmp(big.NewInt(res)) {
		t.Errorf("expect %v but result is %v", res, bres)
	}

	val = 10; res = 3628800
	bres = FactorialBig(val)
	if 0 != bres.Cmp(big.NewInt(res)) {
		t.Errorf("expect %v but result is %v", res, bres)
	}

	fmt.Printf("Extra testing...\n");

	val = 100
	bres = FactorialBig(val)
	fmt.Printf("%d! \t: %v\n", val, bres);

	val = 1000
	bres = FactorialBig(val)
	fmt.Printf("%d! \t: %v\n", val, bres);
}
