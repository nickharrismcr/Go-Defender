package game

import (
	"Def/global"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	points []int
	img    *ebiten.Image
	ops    *ebiten.DrawImageOptions
}

func NewWorld() *World {

	w := &World{}
	w.points = make([]int, global.WorldWidth+1)
	var y int = 0
	var dy int = 1
	for i := 0; i <= global.WorldWidth; i++ {
		w.points[i] = y
		y += dy
		if y == 0 || y > global.ScreenHeight/4 || rand.Intn(10) == 1 {
			dy = -dy
		}
	}
	w.img = ebiten.NewImage(2, 2)
	w.ops = &ebiten.DrawImageOptions{}
	w.img.Fill(color.White)
	w.ops.ColorM.Scale(0.5, 0.3, 0, 1)
	return w
}

func (w *World) Draw(scr *ebiten.Image) {
	ww := global.WorldWidth
	i := int(global.CameraX)
	for x := 0; x < global.ScreenWidth; x++ {
		if i < 0 {
			i += ww
		} else if i > ww {
			i -= ww
		}
		h := w.points[i]
		w.ops.GeoM.Reset()
		w.ops.GeoM.Translate(float64(x), float64(global.ScreenHeight-h))
		scr.DrawImage(w.img, w.ops)
		i++
	}
	cx := global.CameraX - float64(ww/2)

	sw := float64(global.ScreenWidth)
	rs := sw / 4
	rw := sw / 2
	for j := 0; j < ww; j += 10 {
		ind := j + int(cx)
		if ind < 0 {
			ind += ww
		}
		if ind > ww-1 {
			ind -= ww
		}
		h := w.points[ind]
		sx := rs + rw*(float64(j)/float64(ww))
		w.ops.GeoM.Reset()
		w.ops.GeoM.Translate(sx, float64(global.ScreenHeight/10)-(float64(h)/10))
		scr.DrawImage(w.img, w.ops)
	}
}
