package main

import "fmt"

// a function is a block of code which is used to perform a specific task

func recProps(len, wid float64) (float64, float64) { //return mutiple values
	var area = len * wid
	var par = (len + wid) * 2
	return area, par
}

func main() {

	area, par := recProps(10.45, 12.6)
	area1, _ := recProps(10.45, 12.6)

	// here we'r discarding parameter using _(blank Identifier)

	fmt.Printf("\n normal return is:%f  \t %f \n blak Identifier of per :%f", area, par, area1)
}
