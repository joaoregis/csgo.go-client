package csgo

import (
	"gosource/internal/csgo/sdk"
	"gosource/internal/global"
	"gosource/internal/global/configs"
	"gosource/internal/hackFunctions/vector"
	"gosource/internal/memory"
	"math"
)

func GetDistance(entity uintptr, other vector.Vector3) float64 {
	entPosition, _ := memory.GameProcess.ReadVec3(entity + configs.Offsets.Netvars.MVecOrigin)
	delta := other.CalcVector3WithOtherVector3(*entPosition, "-")

	return math.Sqrt(delta.X*delta.X + delta.Y*delta.Y + delta.Z*delta.Z)
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ConvertToRange(Point *vector.Vector2, w, h float64) {

	Point.X /= w
	Point.X *= 2.0
	Point.X -= 1.0

	Point.Y /= h
	Point.Y *= 2.0
	Point.Y -= 1.0
}

func WorldToScreen(VecOrigin *vector.Vector3, VecScreen *vector.Vector2) bool {

	sdk.UpdateViewMatrix()
	Matrix := sdk.ViewMatrix

	var clipCoords vector.Vector4
	clipCoords.X = VecOrigin.X*float64(Matrix[0]) + VecOrigin.Y*float64(Matrix[1]) + VecOrigin.Z*float64(Matrix[2]) + float64(Matrix[3])
	clipCoords.Y = VecOrigin.X*float64(Matrix[4]) + VecOrigin.Y*float64(Matrix[5]) + VecOrigin.Z*float64(Matrix[6]) + float64(Matrix[7])
	clipCoords.Z = VecOrigin.X*float64(Matrix[8]) + VecOrigin.Y*float64(Matrix[8]) + VecOrigin.Z*float64(Matrix[10]) + float64(Matrix[11])
	clipCoords.W = VecOrigin.X*float64(Matrix[12]) + VecOrigin.Y*float64(Matrix[13]) + VecOrigin.Z*float64(Matrix[14]) + float64(Matrix[15])

	if clipCoords.W < 0.001 {
		return false
	}

	var NDC vector.Vector3
	NDC.X = clipCoords.X / clipCoords.W
	NDC.Y = clipCoords.Y / clipCoords.W
	NDC.Z = clipCoords.Z / clipCoords.W

	w, h := global.WINDOW_OVERLAY.GetSize()
	VecScreen.X = (float64(w) / 2 * NDC.X) + (NDC.X + float64(w)/2)
	VecScreen.Y = (float64(h) / 2 * NDC.Y) + (NDC.Y + float64(h)/2)

	ConvertToRange(VecScreen, float64(w), float64(h))

	return true
}
