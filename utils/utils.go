package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func DashRow(count int, doNewLine rune) {
	for i := 0; i < count-1; i++ {
		fmt.Print("-")
	}

	switch doNewLine {
	case 'y':
		fmt.Printf("-\n")
	case 'n':
		fmt.Printf("-")
	}
}

func AskWithMsgs(params ...interface{}) {
	messages := params[len(params)-1].([]string) // interface{} => slice
	vars := make([]string, len(params)-1)        // Length is params length minus the messages
	for i := 0; i < len(params)-1; i++ {
		fmt.Printf("%s", messages[i])
		fmt.Scanln(&vars[i])                 // Save to a index in vars-slice
		vars[i] = strings.TrimSpace(vars[i]) // Trim whitespace which may raise errors later

		// Attempt to convert to float64
		if varAsFloat, err := TryFloatConvert(vars[i]); err != nil {
			fmt.Print("Tuo arvo ei ollut numero! Anna numero. (Käytä pisteitä desimaaleina)")
			i--      // Forces loop to stay in current iteration
			continue // Kicks back loop to rewind
		} else {
			*params[i].(*float64) = varAsFloat // Appending dereferenced elements of interface{} which are converted to pointers of type float64 before dereferencing. (that's really mouthful)
		}
	}
}

func TryFloatConvert(str string) (float64, error) {
	if v, err := strconv.ParseFloat(str, 64); err != nil {
		return 0.0, errors.New("Couldn't convert to float64")
	} else {
		return v, nil
	}
}

func IsOfType(val interface{}, wantedType interface{}) bool {
	return fmt.Sprintf("%T", val) == fmt.Sprintf("%T", wantedType)
}
