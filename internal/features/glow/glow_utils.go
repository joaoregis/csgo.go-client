package glow

import (
	"gosource/internal/csgo"
	"gosource/internal/csgo/sdk"
)

func SetColorRGBA(v *sdk.EntityGlowStruct, r float32, g float32, b float32, a float32) {
	v.Red = r
	v.Green = g
	v.Blue = b
	v.Alpha = a
}

func ValidatePlayerGlow(entity uintptr) bool {

	if isValidPlayer, _ := csgo.PlayerIsValid(entity); !isValidPlayer {
		return false
	}

	return true
}
