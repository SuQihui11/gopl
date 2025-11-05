package main

import (
	"fmt"
	"gopl/meterconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		m := meterconv.Meter(t)
		km := meterconv.Kilometer(t)
		fmt.Printf("%s=%s,%s=%s\n\n", m, meterconv.MToKm(m), km, meterconv.KmToM(km))
	}
}
