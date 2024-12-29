package assets

import (
	"embed"
	"io/fs"
)

//go:embed sounds
var soundFs embed.FS
var Sounds, _ = fs.Sub(soundFs, "sounds")
