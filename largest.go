package main

func Largest(array[]int)int{
	largest := array[0]
	for _,value := range  array{
		if value >largest{
			largest=value
		}
	}
	return largest
}
func Smallest(array[]int)int{
	smallest := array[0]
	for _,value :=range array{
		if value<smallest{
			smallest=value
		}
	}
	return smallest
}
func Average(arr []int) float64 {
	sum := 0

	for _, value := range arr {
		sum += value
	}

	average := float64(sum) / float64(len(arr))
	return average
}
