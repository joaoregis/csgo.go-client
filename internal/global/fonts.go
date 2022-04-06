package global

import (
	"os"

	"github.com/flopp/go-findfont"
	"github.com/go-gl/gltext"
)

var DefaultFontPath string
var Font12 *gltext.Font
var Font14 *gltext.Font
var Font16 *gltext.Font
var Initialization_Font_Was_Done bool = false

func InitFonts() {

	DefaultFontPath, _ = findfont.Find("arial.ttf")
	Font12, _ = loadFont(DefaultFontPath, 12)
	Font14, _ = loadFont(DefaultFontPath, 14)
	Font16, _ = loadFont(DefaultFontPath, 16)

	//
	Initialization_Font_Was_Done = true

}

// loadFont loads the specified font at the given scale.
func loadFont(file string, scale int32) (*gltext.Font, error) {
	fd, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	return gltext.LoadTruetype(fd, scale, 32, 127, gltext.LeftToRight)
}
