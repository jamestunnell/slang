package examples

import (
	"embed"
	"io/fs"
)

//go:embed garage/**/*.sl
var garage embed.FS

//go:embed calculator/*.sl
var calculator embed.FS

func Garage() fs.FS {
	return garage
}

func Calculator() fs.FS {
	return calculator
}
