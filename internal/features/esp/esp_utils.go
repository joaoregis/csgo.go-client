package esp

import (
	"gosource/internal/csgo"
	"gosource/internal/hackFunctions/color"
	"gosource/internal/hackFunctions/vector"
)

var lineOrigin vector.Vector2 = vector.Vector2{
	X: 0,
	Y: -1,
}
var colorA color.RGBA = color.RGBA{
	Red:   1,
	Green: 0,
	Blue:  0,
	Alpha: 1,
}

func ValidatePlayerESP(entity uintptr) bool {

	if isValidPlayer, _ := csgo.PlayerIsValid(entity); !isValidPlayer {
		return false
	}

	if bDormant, _ := csgo.PlayerIsDormant(entity); bDormant {
		return false
	}

	if csgo.PlayerIsLocalEntity(entity) {
		return false
	}

	return true
}
