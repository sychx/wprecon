package text

// ContainsSliceString :: This function will be used to check if there is a certain string within a slice/array.
func ContainsSliceString(slice []string, value string) (int, bool) {
	for i, item := range slice {
		if item == value {
			return i, true
		}
	}

	return 0, false
}

func ContainsSliceSliceString(slice [][]string, value string) (int, bool) {
	for i, item := range slice {
		if item[0] == value {
			return i, true
		}
	}

	return 0, false
}

func ContainsSliceOfMap(slice []map[string]string, sub string, s string) (int, bool) {
	for i, item := range slice {
		if item[sub] == s {
			return i, true
		}
	}

	return 0, false
}
