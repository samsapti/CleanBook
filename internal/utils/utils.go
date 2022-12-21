package utils

import (
	"fmt"
	"os"
)

func print(out *os.File, format string, a ...any) (int, error) {
	return fmt.Fprintf(out, format+"\n", a...)
}

// PrintError prints a message to stderr. Arguments are handled in the
// manner of fmt.Printf.
func PrintError(format string, a ...any) (int, error) {
	return print(os.Stderr, format, a...)
}

// PrintInfo prints a message to stdout. Arguments are handled in the
// manner of fmt.Printf.
func PrintInfo(format string, a ...any) (int, error) {
	return print(os.Stdout, format, a...)
}

// PrintFatal is a shorthand for a call to PrintError followed by
// os.Exit(1).
func PrintFatal(format string, a ...any) {
	PrintError(format, a...)
	os.Exit(1)
}
