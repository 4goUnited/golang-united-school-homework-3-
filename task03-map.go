package homework

import "sort"

func sortMapValues(input map[int]string) []string {
	//Place your code here
	result := make([]string, len(input))
	sl := make([]int, len(input))
	i := 0
	for key, _ := range input {
		sl[i] = key
		i++
	}
	sort.Ints(sl)
	for i, el := range sl {
		result[i] = input[el]
	}
	return result
}
