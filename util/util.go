package util

import (
	"Def/cmp"
	"Def/global"
	"math"
	"math/rand"
)

func ScreenX(x float64) float64 {

	ww := float64(global.WorldWidth)
	sw := float64(global.ScreenWidth)
	cx := global.CameraX()
	over := sw - (ww - cx)
	if over > 0 && x < over {
		x += ww
	}
	sx := x - cx

	return sx
}

func OffScreen(x, y float64) bool {

	return (x < -100 || x > global.ScreenWidth+100 || y < 0 || y > global.ScreenHeight)
}

func RandChoiceF(lst []float64) float64 {
	return lst[rand.Intn(len(lst))]
}

func RandChoiceI(lst []int) int {
	return lst[rand.Intn(len(lst))]
}

func RandChoiceS(lst []string) string {
	return lst[rand.Intn(len(lst))]
}

func ComputeBullet(pos1, pos2 *cmp.Pos, ticks int) (float64, float64) {
	projected_x := pos2.X + (pos2.DX * float64(ticks))
	projected_y := pos2.Y + (pos2.DY * float64(ticks))
	dx := (projected_x - pos1.X) / float64(ticks)
	dy := (projected_y - pos1.Y) / float64(ticks)
	return dx, dy
}

func Collide(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {

	left := math.Max(x1, x2)
	top := math.Max(y1, y2)
	right := math.Min(x1+w1, x2+w2)
	bottom := math.Min(y1+h1, y2+h2)

	return left < right && top < bottom
}
