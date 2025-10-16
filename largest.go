package main
import "fmt"

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
func main (){
	fmt.Println(Largest([]int{1,2,3,4}))
	fmt.Println(Smallest([]int{1,2,3,4}))
	fmt.Printf("%.2f\n",Average([]int{1,2,3}))

}
