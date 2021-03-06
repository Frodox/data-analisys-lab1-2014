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
	"math"
	"math/big"
	hm "./mathfuncs"
)

type Config struct {
	Data struct {
		Ncount		uint64;
		Dprovider	uint64;
		Dcustomer	uint64;
		Nmin		uint64
		Nmax		uint64
		Step		uint64;
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
		"data/lab1.data.var6.txt",
		"txt datafile for this lab")
	plotablePtr := flag.Bool(
		"plotable",
		false,
		"print output in format for gnuplot")
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


	/* Init convinient vars aliases */

	N		:= cfg.Data.Ncount
	n_min		:= cfg.Data.Nmin
	n_max		:= cfg.Data.Nmax
	alpha		:= cfg.Data.Alpha
	D_provider	:= cfg.Data.Dprovider
	D_customer	:= cfg.Data.Dcustomer
	step		:= cfg.Data.Step
	fmt.Println("# N: \t\t", N);
	fmt.Println("# n_min: \t", n_min);
	fmt.Println("# n_max: \t", n_max);
	fmt.Println("# alpha: \t", alpha);
	fmt.Println("# D_п: \t", D_provider);
	fmt.Println("# D_з: \t", D_customer);
	fmt.Println("# =========================")
	pb := *plotablePtr
	if pb {
		fmt.Println("# \t n\tc1\tbetta1\t\tc2\tbetta2\t\tc3\tbetta3")
	}


	/* Minimal data checking */

	if D_provider > n_min {
		fmt.Fprintln(os.Stderr,
			"ERROR: D_provider must be <= n_min\n")
		os.Exit(1);
	}


	for n:= n_min; n <= n_max; n+= step {
		if pb {
			fmt.Printf("\t%d", n)
		} else {
			fmt.Printf("n = %d\n", n)
			fmt.Println("---------")
		}

		/* Calculate P(x=l | Ho) */

		slice_size := D_provider+1
		P1_sharp_values		:= make([]float64, slice_size, slice_size)
		P2_binom_aprox_values	:= make([]float64, slice_size, slice_size)
		P3_aprox_values		:= make([]float64, slice_size, slice_size)

		for l:= uint64(0); l <= D_provider; l++ {
			//fmt.Printf("l=%d , N=%d \n", l, N);
			P1_sharp_values[l] = P1_sharp(l, N, n, D_provider)
			P2_binom_aprox_values[l] = P2_binom(l, N, n, D_provider)
			P3_aprox_values[l] = P3_aprox(l, N, n, D_provider)

			if !pb {
				fmt.Printf("P1(l) = %10.9f, P2(l) = %10.9f, P3(l) = %10.9f\n",
					P1_sharp_values[l],
					P2_binom_aprox_values[l],
					P3_aprox_values[l] );
			}
		}


		/* Determine 'c' and 'betta' */

		c1 := find_c(P1_sharp_values, alpha);
		c2 := find_c(P2_binom_aprox_values, alpha);
		c3 := find_c(P3_aprox_values, alpha);

		betta1 := find_betta(N, n, D_customer, c1, P1_sharp);
		betta2 := find_betta(N, n, D_customer, c2, P2_binom);
		betta3 := find_betta(N, n, D_customer, c3, P3_aprox);

		if pb {
			fmt.Printf("\t%d\t%10.9f\t%d\t%10.9f\t%d\t%10.9f\n",
					c1, betta1,
					c2, betta2,
					c3, betta3);
		} else {
			fmt.Printf("c1 = %d, betta1 = %f\n", c1, betta1);
			fmt.Printf("c2 = %d, betta2 = %f\n", c2, betta2);
			fmt.Printf("c3 = %d, betta3 = %f\n", c3, betta3);
			fmt.Println("=========================")
		}
	}
}

// ------------------------------------------------------------------------- //
func find_betta(N, n, D_customer uint64, c int, P func(_, _, _, _ uint64) float64) (betta float64) {
	betta = 0.0

	for i := 0; i <= c; i++ {
		if (n-uint64(i)) > (N-D_customer) {
			continue
		}
		betta += P(uint64(i), N, n, D_customer)
	}
	return
}

// ------------------------------------------------------------------------- //
func find_c(data_values[] float64, alpha float64) (c int) {
	c = 0

	for summ, i := 0.0, len(data_values)-1; i >= 0; i-- {
		summ += data_values[i]
		//fmt.Printf("P(x = %2d) = %10.9f, summ: %f\n", i, data_values[i], summ);

		if summ > alpha {
			c = i
			break
		}
	}
	return
}

// ------------------------------------------------------------------------- //

// P (x = l) = C (D_f, l) * C (N-D_f, n-l)   /   C (N, n)
func P1_sharp ( l, N, n, D_f uint64 ) (res float64) {

	res = 0;
	C := hm.CombinationBig // just alias

	// C (D_f, l) * C (N-D_f, n-l)
	//fmt.Printf("D: C(%d, %d) * C(%d, %d)\n", D_f, l, N-D_f, n-l);
	numerator_int := big.NewInt(1);
	numerator_int.Mul( C(D_f, l), C(N-D_f, n-l) );

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

// ------------------------------------------------------------------------- //

// P(x=l) = C(n, l) * po^l * (1 - po)^(n-l); po = D_f / N
func P2_binom ( l, N, n, D_f uint64 ) (res float64) {

	res = 0;
	C := hm.CombinationBig // just alias
	po := float64(D_f) / float64(N)

	first_arg := C(n, l); // big.Int
	first_arg_rat := big.NewRat(1, 1);
	first_arg_rat.SetInt(first_arg);

	second_arg := math.Pow(po, float64(l)) * math.Pow(1-po, float64(n-l));
	second_arg_rat := big.NewRat(1, 1);
	second_arg_rat.SetFloat64(second_arg);

	P := big.NewRat(1, 1);
	P.Mul( first_arg_rat, second_arg_rat );

	res, _ = P.Float64();
	return
}

// ------------------------------------------------------------------------- //

// P(x=l) = lambda^l / (l!)  * exp(-lambda)
func P3_aprox ( l, N, n, D_f uint64 ) (res float64) {
	res = 0;
	po := float64(D_f) / float64(N)

	lambda := float64(n) * po
	lambda_l_rat := big.NewRat(1, 1);
	lambda_l_rat.SetFloat64( math.Pow(lambda, float64(l)) );

	exponenta_rat := big.NewRat(1, 1);
	exponenta_rat.SetFloat64( math.Exp(-1 * lambda) );

	l_factorial_rat := big.NewRat(1, 1);
	l_factorial_rat.SetInt( hm.FactorialBig(l) )

	P := big.NewRat(1, 1);
	P.Quo(lambda_l_rat, l_factorial_rat);
	P.Mul(P, exponenta_rat);

	res, _ = P.Float64();
	return
}
