package main
import "fmt"

func main() {
	s := []int{1,2,3,4,5,6,7,8,9,10}
	PrintSlice(s) 
	a := make([]int, 10)
	PrintSlice(a) 
	b := make([]int, 0, 10)
	PrintSlice(b) 
	c := s[0:4]
	PrintSlice(c) 
	d := c[2:5]
	PrintSlice(d) 
	}
	func PrintSlice(data []int) {
		fmt.Printf("%v :: len=%d cap=%d \n", data, len(data), cap(data))
		}