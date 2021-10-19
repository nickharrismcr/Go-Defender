package util

import (
	"Def/constants"
	"math/rand"
)

func OffScreen(x, y float64) bool {
	return (x < 0 || x > constants.ScreenWidth || y < 0 || y > constants.ScreenHeight)
}

func RandChoiceF(lst []float64) float64 {
	return lst[rand.Intn(len(lst))]
}
