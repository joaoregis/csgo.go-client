package features

import (
	"gosource/internal/configs"
	"gosource/internal/csgo"
)

func Visuals() {

	if csgo.MaxPlayers <= 0 {
		return
	}

	// entity list loop
	for i := 0; i < csgo.MaxPlayers; i++ {

		entity, _ := csgo.GetPlayerByIndex(i)

		if configs.G.Glow.Enabled {
			Glow(entity)
		}

		if configs.G.Radar {
			Radar(entity)
		}

		EngineChams(entity)

	}
}
