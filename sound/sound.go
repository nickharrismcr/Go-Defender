package sound

import (
	"Def/gl"
	"bytes"
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type SoundType string

const (
	Background  SoundType = "background"
	Baiterdie   SoundType = "baiterdie"
	Bomberdie   SoundType = "bomberdie"
	Bullet      SoundType = "bullet"
	Caughthuman SoundType = "caughthuman"
	Die         SoundType = "die"
	Dropping    SoundType = "dropping"
	Grabbed     SoundType = "grabbed"
	Humandie    SoundType = "humandie"
	Landerdie   SoundType = "landerdie"
	Laser       SoundType = "laser"
	Levelstart  SoundType = "levelstart"
	Materialise SoundType = "materialise"
	Mutant      SoundType = "mutant"
	Placehuman  SoundType = "placehuman"
	Poddie      SoundType = "poddie"
	Start       SoundType = "start"
	Swarmer     SoundType = "swarmer"
	Thruster    SoundType = "thruster"
)
const sampleRate = 48000

var audioContext *audio.Context
var players map[SoundType]*audio.Player

//go:embed background.wav
var backgroundWAV []byte

//go:embed baiterdie.wav
var baiterdieWAV []byte

//go:embed bomberdie.wav
var bomberdieWAV []byte

//go:embed bullet.wav
var bulletWAV []byte

//go:embed caughthuman.wav
var caughthumanWAV []byte

//go:embed die.wav
var dieWAV []byte

//go:embed dropping.wav
var droppingWAV []byte

//go:embed grabbed.wav
var grabbedWAV []byte

//go:embed humandie.wav
var humandieWAV []byte

//go:embed landerdie.wav
var landerdieWAV []byte

//go:embed laser.wav
var laserWAV []byte

//go:embed levelstart.wav
var levelstartWAV []byte

//go:embed materialise.wav
var materialiseWAV []byte

//go:embed mutant.wav
var mutantWAV []byte

//go:embed placehuman.wav
var placehumanWAV []byte

//go:embed poddie.wav
var poddieWAV []byte

//go:embed start.wav
var startWAV []byte

//go:embed swarmer.wav
var swarmerWAV []byte

//go:embed thruster.wav
var thrusterWAV []byte

func init() {

	players = make(map[SoundType]*audio.Player)
	audioContext = audio.NewContext(sampleRate)
	addSample(Background, backgroundWAV, true)
	addSample(Baiterdie, baiterdieWAV, false)
	addSample(Bomberdie, bomberdieWAV, false)
	addSample(Bullet, bulletWAV, false)
	addSample(Caughthuman, caughthumanWAV, false)
	addSample(Die, dieWAV, false)
	addSample(Dropping, droppingWAV, false)
	addSample(Grabbed, grabbedWAV, false)
	addSample(Humandie, humandieWAV, false)
	addSample(Landerdie, landerdieWAV, false)
	addSample(Laser, laserWAV, false)
	addSample(Levelstart, levelstartWAV, false)
	addSample(Materialise, materialiseWAV, false)
	addSample(Mutant, mutantWAV, false)
	addSample(Placehuman, placehumanWAV, false)
	addSample(Poddie, poddieWAV, false)
	addSample(Start, startWAV, false)
	addSample(Swarmer, swarmerWAV, false)
	addSample(Thruster, thrusterWAV, true)

}
func addSample(name SoundType, data []byte, loop bool) {

	var player *audio.Player
	decoded, err := wav.Decode(audioContext, bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	if loop {
		looped := audio.NewInfiniteLoop(decoded, decoded.Length()*sampleRate)
		player, err = audioContext.NewPlayer(looped)
		if err != nil {
			panic(err)
		}
	} else {
		player, err = audioContext.NewPlayer(decoded)
		if err != nil {
			panic(err)
		}
	}

	players[name] = player
}

func Play(name SoundType) {
	if gl.Mute {
		return
	}
	players[name].Rewind()
	players[name].Play()
}

func PlayIfNot(name SoundType) {
	if gl.Mute {
		return
	}
	if players[name].IsPlaying() {
		return
	}
	Play(name)
}

func Stop(name SoundType) {
	players[name].Pause()
}
