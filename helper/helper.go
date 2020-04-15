package helper

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
)

func CalculateSliceAverage(c1 chan float64, c2 chan float64, sli primitive.A){
	fnSum := 0.0
	weekSum := 0.0
	sliLen := len(sli)
	for i := range sli{
		fnSum += sli[i].(float64)
	}

	for i := 7; i < sliLen-1; i++ {
		weekSum += sli[i].(float64)
	}
	c1 <- fnSum/float64(sliLen)
	c2 <- weekSum/float64(sliLen/2)
}

func OneWeekProjection(c3 chan float64, current, dropRate float64, weeks int32){
	rate := current * dropRate
	newCurrent := current - rate
	c3 <- newCurrent
}

func GetDailyDifference(sli []float64)([]float64) {
	var diffSlice []float64
  for i := 0; i < len(sli)-1; i++ {
    val := math.Abs(sli[i] - sli[i+1])
		diffSlice = append(diffSlice, val)
	}
  return diffSlice
}