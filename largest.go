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