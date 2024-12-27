package main

import (
	"embed"
	_ "embed"
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"time"
)

//go:embed "assets"
var f embed.FS

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
		file, err := f.Open("assets/sound/music/main.wav")

		if err != nil {
			panic(err)
		}

		decodedWav, err := wav.DecodeWithoutResampling(file)

		if err != nil {
			panic("could not decode file")
		}

		op := &oto.NewContextOptions{}

		op.SampleRate = 44100

		op.ChannelCount = 2

		op.Format = oto.FormatSignedInt16LE

		otoCtx, readyChan, err := oto.NewContext(op)

		<-readyChan

		player := otoCtx.NewPlayer(decodedWav)

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
