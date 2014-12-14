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
)

type Config struct {
	Data struct {
		Ncount		int;
		Dprovider	int;
		Dcustomer	int;
		Nmin		int
		Nmax		int
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

	fmt.Println("7:, ", FactorialBig(7))
	hi();

	// debug

}

