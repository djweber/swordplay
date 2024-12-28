package main

import (
	swordplay "Swordplay"
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"log"
	"time"
)

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

	game := swordplay.NewGame()

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
