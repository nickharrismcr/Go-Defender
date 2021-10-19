package util

import "Def/constants"

func OffScreen(x, y float64) bool {
	return (x < 0 || x > constants.ScreenWidth || y < 0 || y > constants.ScreenHeight)
}
