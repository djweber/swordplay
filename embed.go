package swordplay

import (
	"embed"
	"io/fs"
)

//go:embed assets/sounds
var soundFs embed.FS
var Sounds, _ = fs.Sub(soundFs, "assets/sounds")
