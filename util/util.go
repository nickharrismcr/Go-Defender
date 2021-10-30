package util

import (
	"Def/cmp"
	"Def/gl"
	"math"
	"math/rand"
)

func ScreenX(x float64) float64 {

	ww := float64(gl.WorldWidth)
	sw := float64(gl.ScreenWidth)
	cx := gl.CameraX()
	over := sw - (ww - cx)
	if over > 0 && x < over {
		x += ww
	}
	sx := x - cx

	return sx
}

func OffScreen(x, y float64) bool {

	return (x < -100 || x > gl.ScreenWidth+100 || y < 0 || y > gl.ScreenHeight)
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

func ComputeBullet(firepos, playpos *cmp.Pos, time float64) (float64, float64) {

	tt := gl.MaxTPS * time
	projected_x := playpos.X + (playpos.DX * tt)
	projected_y := playpos.Y
	dx := (projected_x - firepos.X) / tt
	dy := (projected_y - firepos.Y) / tt
	return dx, dy
}

func Collide(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {

	left := math.Max(x1, x2)
	top := math.Max(y1, y2)
	right := math.Min(x1+w1, x2+w2)
	bottom := math.Min(y1+h1, y2+h2)

	return left < right && top < bottom
}
