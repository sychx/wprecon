package text

import (
	"fmt"
)

func ContainsSliceString(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if fmt.Sprintf("%s", item) == val {
			return i, true
		}
	}

	return -1, false
}