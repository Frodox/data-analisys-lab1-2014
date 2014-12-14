package main

import (
	"os"
	"fmt"
	"math/big"
)

// C (n, k) = n! / k! (n-k)!
func CombintationBig(n uint64, k uint64) (r *big.Int) {

	if k <= 0 || n <= 0 {
		fmt.Fprintf(os.Stderr,
			"ERROR: n (%d) and k (%d) must be positive\n", n, k);
		return nil
	}

	if k > n {
		fmt.Fprintf(os.Stderr,
			"ERROR: k (%d) must be <= n (%d)\n", k, n);
		return nil
	}


	//fmt.Println("Came: n ", n, " k: ", k);
	n_factorial := FactorialBig(n);
	//fmt.Printf("%d! \t: %v\n", n, n_factorial);
	k_factorial := FactorialBig(k);
	//fmt.Printf("%d! \t: %v\n", k, k_factorial);
	n_minus_k_f := FactorialBig(n-k);
	//fmt.Printf("%d! \t: %v\n", n-k, n_minus_k_f);
	denominator := big.NewInt(1);
	denominator.Mul(k_factorial, n_minus_k_f)
	//fmt.Printf("%d*%d \t: %v\n", k_factorial, n_minus_k_f, denominator);

	r = big.NewInt(1)
	r = r.Div( n_factorial, denominator );

	return
}
