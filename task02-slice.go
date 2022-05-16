package homework

import "sort"

func reverse(input []int64) []int64 {
	//Place your code here
	result := make([]int64, len(input))
	copy(result, input)
	sort.Slice(result, func(i, j int) bool {return int(result[i]) > int(result[j])})
	return result
}
