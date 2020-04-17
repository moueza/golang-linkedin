package main

import (
	"fmt"
	//"strings"
	"math"
	"math/big"
)

func main() {
	// bitwise AND OR XOR
	// bit clear &^
	//left shift <<
	//right >>

	//numeric types do not implicitly convert
	//no add  int to float
	//var indep int = 5
	//var int1 int = 5 //undefined: intl lbl2
	var int2 int = 5
	var floatl float64 = 42
	//  sum := intl+floatl //lbl2

	sum2 := float64(int2) + floatl //float( KO
	//fmt.Printf("Sum: %v, Type: %T",sum,sum)
	fmt.Printf("Sum2: %v, Type: %T", sum2, sum2)

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum:", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	fSum := f1 + f2 + f3
	fmt.Println("float sum:", fSum)

	var b1, b2, b3, b4, b5, b6, bigSum, bigSum2 big.Float
	b1.SetFloat64(23.5)
	b2.SetFloat64(65.1)
	b3.SetFloat64(76.3)

	b4.SetFloat64(23.5)
	b5.SetFloat64(65.1)
	b6.SetFloat64(26.3)

	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Println("BigSumInternal= %.10g\n", bigSum) //internal cplx form
	fmt.Println("BigSum = %.10g\n", &bigSum)

	bigSum2.Add(&b4, &b5).Add(&bigSum, &b6)
	fmt.Println("BigSumInternal2= %.10g\n", bigSum2) //internal cplx form
	fmt.Println("BigSum2 = %.10g\n", &bigSum2)       //still bad val

	circleRadius := 15.5
	circumference := circleRadius * math.Pi
	fmt.Printf("cirumference: %.2f\n",circumference)

}
