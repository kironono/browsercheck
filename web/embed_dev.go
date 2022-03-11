//go:build !prod

package web

import (
	"io/fs"
	"os"
)

func GetFrontendAssets() fs.FS {
	return os.DirFS("web/assets")
}

func GetTemplates() fs.FS {
	return os.DirFS("web/templates")
}
