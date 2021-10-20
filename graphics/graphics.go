package graphics

import (
	"encoding/json"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type SourceFrame struct {
	X, Y, W, H int
}

type GFXFrame struct {
	Filename        string
	Frame           SourceFrame
	Anim_frames     int
	Ticks_per_frame int
}

type Data struct {
	Frames []GFXFrame
}

var spriteSheet *ebiten.Image
var spriteMap map[string]GFXFrame

func GetSpriteSheet() *ebiten.Image {
	return spriteSheet
}

func GetSpriteMap(k string) GFXFrame {
	return spriteMap[k]
}

func Load() {

	spriteMap = make(map[string]GFXFrame)
	spritedata := &Data{}

	dat, err := os.ReadFile("graphics/spritesheet.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(dat), spritedata)

	if err != nil {
		panic(err)
	}

	for _, v := range spritedata.Frames {
		spriteMap[v.Filename] = v
	}

	f, err := os.Open("graphics/spritesheet.png")
	if err != nil {
		panic(err)
	}

	sheetimg, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	spriteSheet = ebiten.NewImageFromImage(sheetimg)
}
