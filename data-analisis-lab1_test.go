package main

import (
	//"os"
	"fmt"
	//"math/big"
	"testing"
)

func TestKnownP1_sharp(t *testing.T) {
	// Test for known value
	 fmt.Printf("%10.9f ?= %10.9f\n", 0.073684211, P1_sharp(3, 20, 9, 3));
	 fmt.Printf("%10.9f ?= %10.9f\n", 0.347368421, P1_sharp(2, 20, 9, 3));
	 fmt.Printf("%10.9f ?= %10.9f\n", 0.434210526, P1_sharp(1, 20, 9, 3));
}
