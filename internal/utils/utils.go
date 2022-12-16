package utils

import (
	"fmt"
	"os"

	"golang.org/x/text/encoding/charmap"
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

// FixEncoding converts strings from UTF-8 to ISO 8859-1 (Latin-1).
// Facebook incorrectly exports data to Latin-1 instead of UTF-8,
// causing problems with non-English characters, emojis, etc.
func FixEncoding(s *string) {
	if ret, err := charmap.ISO8859_1.NewEncoder().String(*s); err == nil {
		*s = ret
	}
}
