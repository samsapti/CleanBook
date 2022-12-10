package utils

import (
	"fmt"
	"os"
)

func print(out *os.File, format string, a ...any) (int, error) {
	return fmt.Fprintf(out, format+"\n", a...)
}

func PrintError(format string, a ...any) (int, error) {
	return print(os.Stderr, format, a...)
}

func PrintInfo(format string, a ...any) (int, error) {
	return print(os.Stdout, format, a...)
}

func PrintPanic(format string, a ...any) {
	panic(fmt.Sprintf(format, a...))
}
