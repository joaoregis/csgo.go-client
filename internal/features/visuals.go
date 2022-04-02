package features

import (
	"gosource/internal/csgo"
	"gosource/internal/global/configs"
)

func Visuals() {

	if csgo.MaxPlayers <= 0 {
		return
	}

	// entity list loop
	for i := 0; i < csgo.MaxPlayers; i++ {

		entity, _ := csgo.GetPlayerByIndex(i)

		if configs.G.D.Glow.Enabled {
			Glow(entity)
		}

		if configs.G.D.Radar {
			Radar(entity)
		}

		EngineChams(entity)

	}
}
