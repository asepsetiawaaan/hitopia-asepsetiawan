package weightedstrings

// ComputeWeights function to compute weights of all possible substrings
func ComputeWeights(s string) map[int]struct{} {
	weights := make(map[int]struct{})

	// Iterate over the string
	for i := 0; i < len(s); i++ {
		weight := int(s[i] - 'a' + 1)
		for j := i; j < len(s) && s[j] == s[i]; j++ {
			weights[weight*(j-i+1)] = struct{}{}
		}
	}

	return weights
}

// CheckQueries function to check the status of each query
func CheckQueries(weights map[int]struct{}, queries []int) []string {
	result := make([]string, len(queries))
	for i, q := range queries {
		if _, found := weights[q]; found {
			result[i] = "Yes"
		} else {
			result[i] = "No"
		}
	}
	return result
}
