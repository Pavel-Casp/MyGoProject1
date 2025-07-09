package service

func CalculateSum(numbers []float64) float64 {
	var sum float64
	for _, n := range numbers {
		sum += n
	}
	return sum
}
