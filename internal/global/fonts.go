package global

import (
	"os"

	"github.com/flopp/go-findfont"
	"github.com/go-gl/gltext"
)

var DefaultFontPath string
var Font12 *gltext.Font

func InitFonts() {

	DefaultFontPath, _ = findfont.Find("arial.ttf")
	Font12, _ = loadFont(DefaultFontPath, 12)

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
