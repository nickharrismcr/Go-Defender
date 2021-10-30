package game

import (
	"Def/gl"
	"Def/util"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var scrh = float64(gl.ScreenHeight)
var scrtop = float64(gl.ScreenTop)

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
	w.points = make([]float64, gl.WorldWidth+1)
	var y float64 = 50
	var dy float64 = 1
	for i := 0; i <= gl.WorldWidth; i++ {
		w.points[i] = y
		y += dy
		if i > 50 {
			if y < 50 || y > gl.ScreenHeight/4 || rand.Intn(10) == 1 {
				dy = -dy
			}
		} else {
			dy = 1
		}

	}
	y = 50
	dy = 1
	for i := gl.WorldWidth; i > 0; i-- {
		if y >= w.points[i] {
			break
		}
		w.points[i] = y
		dy := util.RandChoiceF([]float64{0, 1, 1})
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
	if wx > gl.WorldWidth {
		wx = gl.WorldWidth
	}
	return w.points[int(wx)]
}

func (w *World) Draw(scr *ebiten.Image) {
	ww := gl.WorldWidth
	i := int(gl.CameraX())
	for x := 0; x < gl.ScreenWidth; x++ {
		if i < 0 {
			i += ww
		} else if i > ww {
			i -= ww
		}
		h := w.points[i]
		w.ops.GeoM.Reset()
		w.ops.GeoM.Translate(float64(x), float64(gl.ScreenHeight-h))
		scr.DrawImage(w.img, w.ops)
		i++
	}

	sw := float64(gl.ScreenWidth)
	cx := gl.CameraX() - float64(ww/2) + sw/2

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
