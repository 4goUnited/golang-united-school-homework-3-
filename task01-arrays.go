package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	for _, el := range input {
		result += el
	}
	result /= float32(len(input))
	return 
}
