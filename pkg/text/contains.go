package text

import (
	"fmt"
)

// ContainsSliceString :: This function will be used to check if there is a certain string within a slice/array.
func ContainsSliceString(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if fmt.Sprintf("%s", item) == val {

			return i, true
		}
	}

	return -1, false
}
