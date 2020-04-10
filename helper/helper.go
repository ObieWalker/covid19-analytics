package helper

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CalculateSliceAverage(c chan float64, sli primitive.A, n int){
	fmt.Print("it got here")
	sum := 0.0
	for i := range sli{
		sum += sli[i].(float64)
	}
	fmt.Print("sum.. ", sum)
	fmt.Print("n.. ", n)

	c <- sum/float64(n)
}