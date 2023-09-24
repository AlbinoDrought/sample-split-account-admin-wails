//go:build !windows

package main

import (
	"os"
)

func runWithFixedToken() {
	println("Not supported")
	os.Exit(1)
}
