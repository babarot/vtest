package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mcuadros/go-version"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Fprintln(os.Stderr, "requires the following arguments. '0.12.37 < 0.12.40'")
		os.Exit(1)
	}

	parts := strings.Split(os.Args[1], " ")
	if len(parts) != 3 {
		fmt.Fprintln(os.Stderr, "requires the following arguments. '0.12.37 < 0.12.40'")
		os.Exit(1)
	}

	var (
		v1       = parts[0]
		v2       = parts[2]
		operater = parts[1]
	)

	if !version.Compare(v1, v2, operater) {
		os.Exit(1)
	}
}
