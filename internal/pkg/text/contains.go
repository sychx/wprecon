package text

// ContainsInSlice :: This function will be used to check if a word/text already exists within a slice.
func ContainsInSlice(slice []string, value string) (int, bool) {
	for i, item := range slice {
		if value == item {
			return i, true
		}
	}

	return -1, false
}

// ContainsInSliceSlice :: This function will be used to check if a word/text already exists within a slice.
func ContainsInSliceSlice(slice [][]string, value string) (int, bool) {
	for i, item := range slice {
		for _, item2 := range item {
			if value == item2 {
				return i, true
			}
		}
	}

	return -1, false
}

// ContainsInSliceOfMap :: This function will be used to check if a word/text already exists within a Map.
func ContainsInSliceOfMap(slice []map[string]string, sub string, s string) (int, bool) {
	for i, item := range slice {
		if s == item[sub] {
			return i, true
		}
	}

	return -1, false
}
