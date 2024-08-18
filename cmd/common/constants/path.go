package constants

import (
	"os"
	"os/user"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/arch"
)

var ApplicationPath = ""

var (
	TempPath      = ""
	TempAudioPath = ""
	TempSavesPath = ""
)

var (
	ResourcesPath = ""
	IconsPath     = ""
	TexturesPath  = ""
	FontsPath     = ""
	DataPath      = ""
	SoundsPath    = ""
)

var (
	ApplicationProfileCpu  = ""
	ApplicationProfileHeap = ""
	ApplicationIcon        = ""
	ApplicationFontDefault = ""
	ApplicationFontPixeled = ""
	ApplicationFontOutline = ""
)

func LoadConstants() error {
	ex, err := os.Executable()
	if err != nil {
		return err
	}

	if arch.KurinDebug {
		ApplicationPath = path.Join(path.Dir(ex), "..", "..")
	} else {
		u, err := user.Current()
		if err != nil {
			return err
		}

		ApplicationPath = path.Join(u.HomeDir, "Documents", "kurin")
	}

	TempPath = path.Join(ApplicationPath, "temp")
	os.MkdirAll(TempPath, 777)
	TempAudioPath = path.Join(TempPath, "audio")
	os.MkdirAll(TempAudioPath, 777)
	TempSavesPath = path.Join(TempPath, "saves")
	os.MkdirAll(TempSavesPath, 777)

	ResourcesPath = path.Join(ApplicationPath, "resources")
	IconsPath = path.Join(ResourcesPath, "icons")
	TexturesPath = path.Join(ResourcesPath, "textures")
	FontsPath = path.Join(ResourcesPath, "fonts")
	DataPath = path.Join(ResourcesPath, "data")
	SoundsPath = path.Join(ResourcesPath, "sounds")

	ApplicationProfileCpu = path.Join(TempPath, "cpu.prof")
	ApplicationProfileHeap = path.Join(TempPath, "heap.out")
	ApplicationIcon = path.Join(IconsPath, "icon.png")
	ApplicationFontDefault = path.Join(FontsPath, "Roboto-Regular.ttf")
	ApplicationFontPixeled = path.Join(FontsPath, "Pixeled.ttf")
	ApplicationFontOutline = path.Join(FontsPath, "Pixeled-Outline.ttf")

	return nil
}
