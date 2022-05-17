package homework

func reverse(input []int64) []int64 {
	//Place your code here
	result := make([]int64, len(input))
	copy(result, input)
	for i, j := 0, len(input) - 1; i < len(input) / 2; i, j = i + 1, j - 1 {
		result[i], result[j] = result[j], result[i]
	}
	return result 
}
