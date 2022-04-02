package features

import (
	"gosource/internal/csgo"
	"gosource/internal/global/configs"
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

		if configs.G.D.Glow.IsHealthBased {

			playerHealth := csgo.GetPlayerHealth(entity)
			c := float32((math.Lerpf(0, 1, playerHealth/100)))
			eGlow.SetColorRGBA(1-c, c, 0, configs.G.D.Glow.Alpha)

		} else {

			rgba := color.HexToRGBA(color.Hex(configs.G.D.Glow.BaseColor))
			eGlow.SetColorRGBA(rgba.Red, rgba.Green, rgba.Blue, configs.G.D.Glow.Alpha)

		}

	}

	eGlow.Save(entity)
}
