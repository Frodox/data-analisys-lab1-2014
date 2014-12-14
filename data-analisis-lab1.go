/*
 * Data analisys' Lab1
 * 
 * Teacher: Korotkov Alexander
 * Student: Vit Ry <developer@bitthinker.com> (c) 2014
 * Task: see --help option
 * ts: 4
 */

package main

import (
	"fmt"
	"flag"
	"path"
	"os"
	"code.google.com/p/gcfg" // config parser
	"math/big"
)

type Config struct {
	Data struct {
		Ncount		uint64;
		Dprovider	uint64;
		Dcustomer	uint64;
		Nmin		uint64
		Nmax		uint64
		Step		int;
		Alpha		float64
	}
}

// ------------------------------------------------------------------------- //

func main() {

	/* Setup cmd-line parser */

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"Usage: %s [options]\n",
			path.Base(os.Args[0]))
		fmt.Fprintf(os.Stderr, "Options:\n");
		flag.PrintDefaults()

		fmt.Fprintf(os.Stderr,
			"\nКакой-то типа текст про то, что нужно делать в лабе\n\n")
		os.Exit(0);
	}
	filenamePtr := flag.String(
		"filename",
		"lab1.data.var6.txt",
		"txt datafile for this lab")
	flag.Parse()

	if len(*filenamePtr) == 0 {
		flag.Usage();	// will Exit(0)
	}


	/* Read data from data-file in memory-struct */

	var cfg Config
	err := gcfg.ReadFileInto(&cfg, *filenamePtr)
	if err != nil {
		fmt.Fprintln(os.Stderr,
			"Error occured during parsing data file:\n",
			err.Error());
		os.Exit(1);
	}


	/* Init convinient vars */

	N := cfg.Data.Ncount
	n := cfg.Data.Nmin
	alpha := cfg.Data.Alpha
	D_provider := cfg.Data.Dprovider
	D_customer := cfg.Data.Dcustomer
	fmt.Println("N: \t", N);
	fmt.Println("n: \t", n);
	fmt.Println("a: \t", alpha);
	fmt.Println("D_п: \t", D_provider);
	fmt.Println("D_з: \t", D_customer);


	/* Minimal data checking */

	if D_provider > n {
		fmt.Fprintln(os.Stderr,
			"ERROR: D_provider must be <= n\n")
		os.Exit(1);
	}


	/* Calculate P( x=l | Ho ) */

	P_values := make([]float64, D_provider+1, D_provider+1)
	for l:= uint64(0); l <= D_provider; l++ {
		//fmt.Printf("l=%d , N=%d \n", l, N);
		P_values[l] = P1_sharp(l, N, n, D_provider)
	}


	/* Determine 'c' */

	c := 0
	for summ, i := 0.0, len(P_values)-1; i >= 0; i-- {
		summ += P_values[i]
		fmt.Printf("P(x = %2d) = %10.9f, summ: %f\n", i, P_values[i], summ);

		if summ > alpha {
			c = i
			break
		}
	}
	fmt.Printf("c = %d\n", c);


	/* Determine betta */
	betta := 0.0
	for i := 0; i <= c; i++ {
		//fmt.Printf("i: = %d\n", i);
		if (n-uint64(i)) > (N-D_customer) {
			//fmt.Println("skip")
			continue
		}
		betta += P1_sharp(uint64(i), N, n, D_customer)
	}
	fmt.Printf("betta = %f\n", betta);
}

// ------------------------------------------------------------------------- //

// P (x = l) = C (D_f, l) * C (N-D_f, n-l)   /   C (N, n)
func P1_sharp ( l uint64, N uint64, n uint64, D_f uint64 ) (res float64) {

	res = 0;
	C := CombinationBig // just alias

	//fmt.Println("1--------------");

	// C (D_f, l) * C (N-D_f, n-l)
	numerator_int := big.NewInt(1);
	//fmt.Printf("D: C(%d, %d) * C(%d, %d)\n", D_f, l, N-D_f, n-l);
	numerator_int.Mul( C(D_f, l), C(N-D_f, n-l) );
	//fmt.Println("2--------------");

	numerator_rat := big.NewRat(1, 1);
	numerator_rat.SetInt(numerator_int);


	// C (N, n)
	denominator_int := C(N, n);

	denominator_rat := big.NewRat(1, 1);
	denominator_rat.SetInt(denominator_int);

	// P = C (D_f, l) * C (N-D_f, n-l)   /   C (N, n)
	P := big.NewRat(1, 1);
	P.Quo(numerator_rat, denominator_rat);

	res, _ = P.Float64();
	return
}

