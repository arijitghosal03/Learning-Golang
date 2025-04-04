package main
import (
	"fmt"
)
func SumInts(m map[string]int64) int64 {
	var sum int64
	for _, v := range m {
		sum += v
	}
	return sum
}
func SumFloats(m map[string]float64) float64 {
	var sum float64
	for _, v := range m {
		sum += v
	}
	return sum
}
func SumIntsORFloats[K comparable, V int64|float64](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
	
}
func main() {
	ints := map[string]int64{"a": 1, "b": 2, "c": 3}
	floats := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
	fmt.Println(SumInts(ints))
	fmt.Println(SumFloats(floats))
	fmt.Println(SumIntsORFloats(ints))
	fmt.Println(SumIntsORFloats(floats))
}