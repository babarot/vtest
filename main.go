package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/mcuadros/go-version"
)

// error message template
var (
	ErrInvalidUsage = errors.New("requires the following arguments. '0.12.37 < 0.12.40'")
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Fprintln(os.Stderr, ErrInvalidUsage)
		os.Exit(1)
	}

	parts := strings.Split(os.Args[1], " ")
	if len(parts) != 3 {
		fmt.Fprintln(os.Stderr, ErrInvalidUsage)
		os.Exit(1)
	}

	var (
		lhs      = parts[0]
		rhs      = parts[2]
		operater = parts[1]
	)

	if !version.Compare(lhs, rhs, operater) {
		os.Exit(1)
	}
}
