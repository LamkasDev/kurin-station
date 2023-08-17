package constants

import (
	"os"
	"os/user"
	"path"

	"github.com/LamkasDev/kitsune/cmd/common/arch"
)

var ApplicationPath = ""
var ResourcesPath = ""

var ApplicationIcon = ""
var ApplicationIconClose = ""

var ApplicationFontRegular = ""
var ApplicationFontBold = ""

func LoadConstants() *error {
	ex, err := os.Executable()
	if err != nil {
		return &err
	}

	if arch.KitsuneDebug {
		ApplicationPath = path.Join(path.Dir(ex), "..", "..")
	} else {
		u, err := user.Current()
		if err != nil {
			return &err
		}

		ApplicationPath = path.Join(u.HomeDir, "Documents", "kitsune")
	}

	ResourcesPath = path.Join(ApplicationPath, "resources")
	ApplicationIcon = path.Join(ResourcesPath, "icons", "icon.png")
	ApplicationIconClose = path.Join(ResourcesPath, "icons", "close.png")
	ApplicationFontRegular = path.Join(ResourcesPath, "fonts", "Roboto-Regular.ttf")
	ApplicationFontBold = path.Join(ResourcesPath, "fonts", "Roboto-Bold.ttf")

	return nil
}
