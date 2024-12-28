package swordplay

import (
	"embed"
	"io/fs"
)

//go:embed assets
var assets embed.FS

var Sounds, _ = fs.Sub(assets, "assets/sounds")
