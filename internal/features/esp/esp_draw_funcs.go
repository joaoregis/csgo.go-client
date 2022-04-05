package esp

import (
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

func DrawFilledOutlinedRect(pos1 vector.Vector2, pos2 vector.Vector2, thickness float32, c *color.RGBA, cOuline *color.RGBA) {

	// gl.Color4f(color.Red, color.Green, color.Blue, color.Alpha)
	// gl.Rectf(float32(pos1.X), float32(pos1.Y), float32(pos2.X), float32(pos2.Y))

}

func DrawOutlinedRect(pos1 vector.Vector2, pos2 vector.Vector2, c *color.RGBA) {

	// gl.LineWidth(thickness)
	// gl.Begin(gl.LINES)
	// gl.Color4f(color.Red, color.Green, color.Blue, color.Alpha)
	// // DrawLine(topLeft, topRight, thickness, c)
	// // DrawLine(bottomLeft, bottomRight, thickness, c)
	// // DrawLine(topLeft, bottomLeft, thickness, c)
	// // DrawLine(topRight, bottomRight, thickness, c)

	// gl.Rect(float32(pos1.X), float32(pos1.Y), float32(pos2.X), float32(pos2.Y))

}

func DrawStringf(pos vector.Vector2, color *color.RGBA, format string, argv ...interface{}) {

	// gl.Color4f(color.Red, color.Blue, color.Green, color.Alpha)
	// gl.RasterPos2f(float32(pos.X), float32(pos.Y))

	// text := fmt.Sprintf(format, argv...)

	// gl.PushAttrib(gl.LIST_BIT)
	// gl.ListBase(32)
	// gl.CallLists(int32(len(text)), gl.UNSIGNED_BYTE, unsafe.Pointer(&text))
	// gl.PopAttrib()

}
