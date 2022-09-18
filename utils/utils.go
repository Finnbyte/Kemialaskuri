package utils

import (
	"fmt"
)

func DashRow(count int, doNewLine rune) {
	for i := 0; i < count-1; i++ {
		fmt.Print("-")
	}

	if doNewLine == 'y' {
		fmt.Printf("-\n")
	} else if doNewLine == 'n' {
		fmt.Printf("-")
	}
}
