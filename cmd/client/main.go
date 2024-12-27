package main

import (
	swordplay "Swordplay"
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"time"
)

type Game struct{}

func (g Game) Update() error {
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hi")
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(1280, 800)
	ebiten.SetWindowTitle("Swordplay")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	go func() {
		file, err := swordplay.Sounds.Open("music/menu.ogg")

		if err != nil {
			panic(err)
		}

		decodedOgg, err := vorbis.DecodeF32(file)

		if err != nil {
			panic("could not decode file")
		}

		op := &oto.NewContextOptions{}

		op.SampleRate = 44100

		op.ChannelCount = 2

		op.Format = oto.FormatFloat32LE

		otoCtx, readyChan, err := oto.NewContext(op)

		<-readyChan

		player := otoCtx.NewPlayer(decodedOgg)

		player.Play()

		for player.IsPlaying() {
			time.Sleep(time.Millisecond)
		}

		file.Close()
	}()

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
