//go:build prod

package web

import (
	"embed"
	"io/fs"
)

//go:embed assets
var embedFrontend embed.FS

//go:embed templates
var embedTemplates embed.FS

func GetFrontendAssets() fs.FS {
	f, err := fs.Sub(embedFrontend, "assets")
	if err != nil {
		panic(err)
	}
	return f
}

func GetTemplates() fs.FS {
	f, err := fs.Sub(embedTemplates, "templates")
	if err != nil {
		panic(err)
	}
	return f
}
