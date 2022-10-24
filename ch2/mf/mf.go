package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/lengthconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mf: %v\n", err)
			os.Exit(1)
		}

		f := lengthconv.Foot(t)
		m := lengthconv.Meter(t)

		fmt.Printf("%s = %s, %s = %s\n", f,
			lengthconv.FToM(f), m, lengthconv.MToF(m))
	}
	if len(os.Args) < 2 {
		var t float64
		for {
			fmt.Scanf("%f", &t)
			f := lengthconv.Foot(t)
			m := lengthconv.Meter(t)

			fmt.Printf("%s = %s, %s = %s\n", f,
				lengthconv.FToM(f), m, lengthconv.MToF(m))
		}
	}
}
