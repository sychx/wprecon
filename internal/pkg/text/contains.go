package text

// ContainsInSlice :: This function will be used to check if a word/text already exists within a slice.
func ContainsInSlice(slice []string, value string) (int, bool) {
	for i, item := range slice {
		if item == value {
			return i, true
		}
	}

	return 0, false
}

// ContainsInSliceSlice :: This function will be used to check if a word/text already exists within a slice.
func ContainsInSliceSlice(slice [][]string, value string) (int, bool) {
	for i, item := range slice {
		if item[0] == value {
			return i, true
		}
	}

	return 0, false
}

// ContainsInSliceOfMap :: This function will be used to check if a word/text already exists within a Map.
func ContainsInSliceOfMap(slice []map[string]string, sub string, s string) (int, bool) {
	for i, item := range slice {
		if item[sub] == s {
			return i, true
		}
	}

	return 0, false
}
