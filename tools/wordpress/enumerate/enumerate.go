package enumerate


func findStringInSliceSlice(slice [][5]string, key int, value string) (int, bool) {
	for i, item := range slice {
		if value == item[key] {
			return i, true
		}
	}

	return -1, false
}

func findByValueInIndex(slice [][5]string, sub string) int {
	for i, item := range slice {
		for _, item2 := range item {
			if sub == item2 {
				return i
			}
		}
	}

	return -1
}
