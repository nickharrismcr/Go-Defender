package game

import (
	"Def/gl"
	"Def/graphics"
	"Def/types"
	"image"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

var charlist = "0123456789:?ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Char struct {
	offset int
	x, y   float64
}

type Characters struct {
	chars  [][]Char
	sIdx   int
	ops    *ebiten.DrawImageOptions
	col    types.ColorF
	colIdx int
	colCtr float64
	speed  float64
	img    *ebiten.Image
	smap   graphics.GFXFrame
}

func NewCharacters() *Characters {

	w := &Characters{

		chars:  [][]Char{},
		ops:    &ebiten.DrawImageOptions{},
		colIdx: 0,
		colCtr: 0,
		speed:  1,
		sIdx:   0,
		img:    graphics.GetSpriteSheet(),
		smap:   graphics.GetSpriteMap("font.png"),
	}

	w.ops.ColorM.Scale(0.5, 0.3, 0, 1)
	return w
}

func (w *Characters) SetSpeed(s float64) {
	w.speed = s
}

func (w *Characters) Clear() {
	w.chars = [][]Char{}
}

func (w *Characters) Remove(idx int) {
	w.chars[idx] = []Char{}
}

func (w *Characters) Add(s string, x, y float64) int {

	idx := w.sIdx
	chs := []Char{}
	for i, c := range s {
		ch := Char{
			offset: w.getOffset(string(c)),
			x:      x + float64(i*31),
			y:      y,
		}
		chs = append(chs, ch)
	}
	w.chars = append(w.chars, chs)
	w.sIdx++
	return idx
}

func (w *Characters) Change(idx int, s string) {
	chs := w.chars[idx]
	for i, c := range s {
		if i >= len(chs) {
			break
		}
		chs[i].offset = w.getOffset(string(c))
	}
}

func (w *Characters) getOffset(c string) int {

	if c == " " {
		return -1
	}
	return strings.Index(charlist, c) * w.smap.Frame.H
}

func (w *Characters) Update() {

	w.colCtr += w.speed
	if w.colCtr == gl.MaxTPS {
		w.colCtr = 0
		w.colIdx++
	}
	currCol := gl.Cols[w.colIdx%5]
	nextCol := gl.Cols[(w.colIdx+1)%5]
	dcol := nextCol.Subtract(currCol)
	dcol = dcol.Multiply(1.0 / float64(gl.MaxTPS))
	ncol := dcol.Multiply(w.colCtr)
	ncol = ncol.Add(currCol)
	w.col = ncol

}

func (w *Characters) Draw(scr *ebiten.Image) {
	for _, chs := range w.chars {
		for _, c := range chs {
			if c.offset == -1 {
				continue
			}
			sx := w.smap.Frame.X
			sy := w.smap.Frame.Y
			sh := w.smap.Frame.H
			si := w.img.SubImage(image.Rect(sx+c.offset, sy, sx+c.offset+sh, sy+sh)).(*ebiten.Image)
			w.ops.GeoM.Reset()
			w.ops.GeoM.Translate(c.x, c.y)
			w.ops.ColorM.Reset()
			w.ops.ColorM.Scale(w.col.R, w.col.G, w.col.B, w.col.A)
			scr.DrawImage(si, w.ops)
		}
	}
}
