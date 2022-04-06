package esp

import (
	"fmt"
	"gosource/internal/global"
	"gosource/internal/hackFunctions/color"
	"gosource/internal/hackFunctions/vector"

	"github.com/go-gl/gl/v2.1/gl"
)

func DrawLine(Start vector.Vector2, End vector.Vector2, thickness float32, color *color.RGBA) {

	gl.LineWidth(thickness)
	gl.Begin(gl.LINES)
	gl.Color4f(color.Red, color.Green, color.Blue, color.Alpha)
	gl.Vertex2f(float32(Start.X), float32(Start.Y))
	gl.Vertex2f(float32(End.X), float32(End.Y))
	gl.End()

}

func DrawFilledRect(pos1 vector.Vector2, pos2 vector.Vector2, c *color.RGBA) {

	gl.Color4f(c.Red, c.Green, c.Blue, c.Alpha)
	gl.Rectf(float32(pos1.X), float32(pos1.Y), float32(pos2.X), float32(pos2.Y))

}

func DrawStringf(pos vector.Vector2, c *color.RGBA, format string, argv ...interface{}) {

	if !global.Initialization_Font_Was_Done {
		return
	}

	f := global.Font14

	// converting coordinate system
	var viewport [4]int32
	gl.GetIntegerv(gl.VIEWPORT, &viewport[0])
	x := ((pos.X + 1.0) * float64(viewport[2]) / 2)
	y := ((1.0 - pos.Y) * float64(viewport[3]) / 2)

	text := fmt.Sprintf(format, argv...)

	// background drawing font
	gl.PushAttrib(gl.CURRENT_BIT)
	gl.Color3f(float32(0), float32(0), float32(0))
	f.Printf(float32(x+1), float32(y+1), "%s\n", text)

	// actual text
	gl.Color3f(c.Red, c.Green, c.Blue)
	f.Printf(float32(x), float32(y), "%s\n", text)
	gl.PopAttrib()

}
