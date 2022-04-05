package features

import (
	"gosource/internal/csgo"
	"gosource/internal/features/esp"
	"gosource/internal/features/glow"
	"gosource/internal/global/configs"
)

func Visuals() {

	if csgo.MaxPlayers <= 0 {
		return
	}

	csgo.LoadUserInfoTable()

	// entity list loop
	for i := 0; i < csgo.MaxPlayers; i++ {

		/* TODO: Convert entity pointer to struct object that represents the BasePlayer (similar to SDK) */
		entity, _ := csgo.GetPlayerByIndex(i)

		if configs.G.D.ESP.Enabled {
			esp.Esp(entity, i)
		}

		if configs.G.D.Radar {
			Radar(entity)
		}

		if configs.G.D.Glow.Enabled {
			glow.Glow(entity)
		}

		glow.EngineChams(entity)

	}
}
