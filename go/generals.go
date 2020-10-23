package main

import "fmt"

func main() {
	/*
	   %T--->returs the type of data
	   %v---> retuns values
	   %d---->retuns the size
	   INT: %o(base 8),%b(base 2),%d(base 10),%x(base 16)
	   Float: %e (scientific notation),%f|%F(without decimal),%g(large exponents)
	   String: %q(with quotes),%s(normal string|char)
	*/
	var a = 54
	fmt.Printf("My data type is: %T", a)

	fmt.Printf("\n MY value is %v", a)

	fmt.Printf("\n Octal of a  is :%o ", a)

	fmt.Printf("\n Base of a is:%b", a)

	fmt.Printf("\n Decimal of a is:%d", a)

	fmt.Printf("\n Hexadecimal of a is:%X ", a)
	var f1 float64 = 56.871298479812
	fmt.Printf("\n Sceintic note of f1 is: %e \n Without decimal of f1 is:%F \n Actual Exponent of f1 is:%g", f1, f1, f1)
	var str1 = "Upputuri"
	fmt.Printf("\n Actual String of str1:%s \n With quotes:%q", str1, str1)
	fmt.Println("\nActual String of str1:%s \n With quotes:%q", str1, str1) // Note: here %q is not wrk bcz we used "Println"

}
