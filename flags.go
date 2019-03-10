package main

import (
	"flag"
)

var (
	srcAddress string
)

/*
/	parseFlags parses all command line arguments.
*/
func parseFlags() {
	//Get flags
	flagSrcAddress := flag.String("S", "", "Address to block.")

	//Get Values
	flag.Parse()
	srcAddress = *flagSrcAddress
}

/*	flagsComplete
/	Returns true if all flags are satisfied and valid
*/
func flagsComplete() (allValid bool, err string) {
	allValid = true
	if srcAddress == "" {
		err = err + "Invalid source address.\n"
		allValid = false
	}

	return allValid, err
}
