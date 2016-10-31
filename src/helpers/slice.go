package helpers

// StringInSlice returns whether a string is within a slice or not.
func StringInSlice(s string, list []string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}

	return false
}
