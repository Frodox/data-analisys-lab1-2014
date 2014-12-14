package main

import (
	//"os"
	//"fmt"
	"math/big"
	"testing"
)

func TestAllCasesCombinationBig (t *testing.T) {
	var bres 	*big.Int
	var n, k 	uint64
	var res 	int64
	

	n, k = 0, 0; res = 1
	bres = CombinationBig(n, k)
	if 0 != bres.Cmp(big.NewInt(res)) {
		t.Errorf("expect %v but result is %v", res, bres)
	}

	n, k = 0, 1; res = -1
	bres = CombinationBig(n, k)
	if nil != bres {
		t.Errorf("expect 'nil' but result is %v", bres)
	}

	n, k = 1, 0; res = 1
	bres = CombinationBig(n, k)
	if 0 != bres.Cmp(big.NewInt(res)) {
		t.Errorf("expect %v but result is %v", res, bres)
	}

	n, k = 1, 50; res = -1
	bres = CombinationBig(n, k)
	if nil != bres {
		t.Errorf("expect 'nil' but result is %v", res, bres)
	}

	n, k = 1, 1; res = 1
	bres = CombinationBig(n, k)
	if 0 != bres.Cmp(big.NewInt(res)) {
		t.Errorf("expect %v but result is %v", res, bres)
	}

	n, k = 2, 1; res = 2
	bres = CombinationBig(n, k)
	if 0 != bres.Cmp(big.NewInt(res)) {
		t.Errorf("expect %v but result is %v", res, bres)
	}

	n, k = 10, 5; res = 252
	bres = CombinationBig(n, k)
	if 0 != bres.Cmp(big.NewInt(res)) {
		t.Errorf("expect %v but result is %v", res, bres)
	}

	n, k = 100, 3; res = 1.617e5
	bres = CombinationBig(n, k)
	if 0 != bres.Cmp(big.NewInt(res)) {
		t.Errorf("expect %v but result is %v", res, bres)
	}

	n, k = 50, 20; res = 4.712921224396e13
	bres = CombinationBig(n, k)
	if 0 != bres.Cmp(big.NewInt(res)) {
		t.Errorf("expect %v but result is %v", res, bres)
	}
}
