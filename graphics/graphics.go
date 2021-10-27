package graphics

import (
	"Def/global"
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// loads a spritesheet and JSON sprite map from embedded files

// define structs mapping to JSON layout
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
	rv, ok := spriteMap[k]
	if !ok {
		panic(fmt.Sprintf("No graphics for %s", k))
	}
	return rv
}

//go:embed spritesheet.json
var spritedataJSON []byte

//go:embed spritesheet.png
var spritesheetPNG []byte

func Load() {

	spriteMap = make(map[string]GFXFrame)
	spritedata := &Data{}

	err := json.Unmarshal(spritedataJSON, spritedata)

	if err != nil {
		panic(err)
	}

	for _, v := range spritedata.Frames {
		v.Ticks_per_frame /= 60 / global.MaxTPS
		spriteMap[v.Filename] = v
	}

	sheetimg, _, err := image.Decode(bytes.NewReader(spritesheetPNG))
	if err != nil {
		panic(err)
	}

	spriteSheet = ebiten.NewImageFromImage(sheetimg)
}
