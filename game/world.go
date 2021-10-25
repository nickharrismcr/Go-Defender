package game

import (
	"Def/global"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var scrh = float64(global.ScreenHeight)
var scrtop = float64(global.ScreenTop)

type World struct {
	points []float64
	img    *ebiten.Image
	ops    *ebiten.DrawImageOptions
	engine *Engine
}

func NewWorld(engine *Engine) *World {

	w := &World{
		engine: engine,
	}
	w.points = make([]float64, global.WorldWidth+1)
	var y float64 = 0
	var dy float64 = 1
	for i := 0; i <= global.WorldWidth; i++ {
		w.points[i] = y
		y += dy
		if i > 100 && (y < 100 || y > global.ScreenHeight/4 || rand.Intn(10) == 1) {
			dy = -dy
		}
	}
	y = 0
	dy = 1
	for i := global.WorldWidth; i > 0; i-- {
		if y == w.points[i] {
			break
		}
		w.points[i] = y
		y += dy

	}
	w.img = ebiten.NewImage(2, 2)
	w.ops = &ebiten.DrawImageOptions{}
	w.img.Fill(color.White)
	w.ops.ColorM.Scale(0.5, 0.3, 0, 1)
	return w
}

func (w *World) At(wx float64) float64 {

	if wx < 0 {
		wx = 0
	}
	if wx > global.WorldWidth {
		wx = global.WorldWidth
	}
	return w.points[int(wx)]
}

func (w *World) Draw(scr *ebiten.Image) {
	ww := global.WorldWidth
	i := int(w.engine.CameraX)
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

	sw := float64(global.ScreenWidth)
	cx := w.engine.CameraX - float64(ww/2) + sw/2

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
		w.ops.GeoM.Scale(0.5, 0.5)
		w.ops.GeoM.Translate(sx, float64(scrtop)-(float64(h*(scrtop/scrh))))

		scr.DrawImage(w.img, w.ops)
	}
}
