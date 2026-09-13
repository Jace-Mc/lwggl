package utils

import (
	"os"
	"fmt"
)

type all any

func Write(__format ...all) {
	fmt.Print(__format)
}

func Writeln(__format ...all) {
	fmt.Println(__format)
}

func Argc() int {
	return len(os.Args)
}

func Argv() []string {
	return os.Args
}
