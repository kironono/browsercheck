//go:build prod

package web

import (
	"embed"
	"io/fs"
)

//go:embed dist
var embedFrontend embed.FS

func GetFrontendAssets() fs.FS {
	f, err := fs.Sub(embedFrontend, "assets")
	if err != nil {
		panic(err)
	}
	return f
}
