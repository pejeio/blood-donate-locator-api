package utils

func ContainsAllString(slice1, slice2 []string) bool {
	elements := make(map[string]bool)

	for _, value := range slice1 {
		elements[value] = true
	}

	for _, value := range slice2 {
		if !elements[value] {
			return false
		}
	}

	return true
}
