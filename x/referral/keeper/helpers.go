package keeper

// isExistsInSlice checks if value exists in slice.
func isExistsInSlice(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}
