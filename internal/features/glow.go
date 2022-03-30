package features

import (
	"gosource/internal/configs"
	"gosource/internal/csgo"
	"gosource/internal/hackFunctions/color"

	"github.com/google/gxui/math"
)

func Glow(entity uintptr) {

	isEnemy, _ := csgo.PlayerIsEnemy(entity)
	if !isEnemy {
		return
	}

	eGlow := csgo.GetEntityGlow(entity)
	if csgo.IsDefusing(entity) {

		eGlow.SetColorRGBA(1, 1, 1, 1) // Set glow to white when enemy is defusing!

	} else {

		if configs.G.Glow.IsHealthBased {

			playerHealth := csgo.GetPlayerHealth(entity)
			c := float32((math.Lerpf(0, 1, playerHealth/100)))
			eGlow.SetColorRGBA(1-c, c, 0, configs.G.Glow.Alpha)

		} else {

			rgba := color.HexToRGBA(color.Hex(configs.G.Glow.BaseColor))
			eGlow.SetColorRGBA(rgba.Red, rgba.Green, rgba.Blue, configs.G.Glow.Alpha)

		}

	}

	eGlow.Save(entity)
}
