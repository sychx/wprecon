package text

// ContainsSliceString :: This function will be used to check if there is a certain string within a slice/array.
func ContainsSliceString(slice []string, value string) (int, bool) {
	for i, item := range slice {
		if item == value {

			return i, true
		}
	}

	return -1, false
}
