package main

import (
	"Swordplay/assets"
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"io"
	"log"
	"time"
)

func main() {
	ebiten.SetWindowSize(1280, 800)
	ebiten.SetWindowTitle("Swordplay")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	go func() {
		file, err := assets.Sounds.Open("music/menu.ogg")

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

		newPos, err := player.Seek(0, io.SeekStart)
		if err != nil {
			panic("player.Seek failed: " + err.Error())
		}
		println("Player is now at position:", newPos)
		player.Play()

		file.Close()
	}()

	game := NewGame()

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
