//go:build !prod

package web

import (
	"io/fs"
	"os"
)

func GetFrontendAssets() fs.FS {
	return os.DirFS("web/assets")
}
