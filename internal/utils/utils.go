package utils

import (
	"log"
	"os"
)

func print(out *os.File, format string, a ...any) {
	log.SetOutput(out)
	log.Printf(format+"\n", a...)
}

// PrintError prints a message to stderr. Arguments are handled in the
// manner of fmt.Printf.
func PrintError(format string, a ...any) {
	print(os.Stderr, format, a...)
}

// PrintInfo prints a message to stdout. Arguments are handled in the
// manner of fmt.Printf.
func PrintInfo(format string, a ...any) {
	print(os.Stdout, format, a...)
}

// PrintVerbose prints a message to stdout only if verbose output is
// enabled. Arguments are handled in the manner of fmt.Printf.
func PrintVerbose(verbose bool, format string, a ...any) {
	if verbose {
		print(os.Stdout, format, a...)
	}
}

// PrintFatal is a shorthand for a call to PrintError followed by
// os.Exit(1).
func PrintFatal(format string, a ...any) {
	PrintError(format, a...)
	os.Exit(1)
}
