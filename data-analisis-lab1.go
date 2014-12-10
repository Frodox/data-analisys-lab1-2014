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
	"log"
	// config-file parser
	"code.google.com/p/gcfg"
	//"os/exec"
	//"strconv"
	//"strings"
	//"math/rand"
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

// --------------------------------------------------------------------------- //

func main() {

	/* Init cmd-line parser */
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", path.Base(os.Args[0]))
		fmt.Fprintf(os.Stderr, "Options:\n");
		flag.PrintDefaults()

		fmt.Fprintf(os.Stderr, "\nКакой-то типа текст прото, что нужно делать в лабе\n\n")
		os.Exit(0);
	}
	filenamePtr := flag.String("filename", "", "txt datafile for this lab")
	flag.Parse()

	if len(*filenamePtr) == 0 {
		flag.Usage();
	}


	/* Read data from data-file in memory */

	var cfg Config
	err := gcfg.ReadFileInto(&cfg, *filenamePtr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error occured during parsing data file:\n", err.Error());
		os.Exit(1);
	}

	log.Fatal("text");

	fmt.Printf("Name: %d\n", cfg.Data.Ncount);
}
