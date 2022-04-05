package glow

import (
	"gosource/internal/csgo"
	"gosource/internal/csgo/sdk"
	"gosource/internal/global/configs"
	"gosource/internal/memory"
)

func ClearEngineChams() {

	// entity list loop
	for i := 0; i < csgo.MaxPlayers; i++ {

		entity, _ := csgo.GetPlayerByIndex(i)

		memory.Write(&memory.GameProcess, entity+configs.Offsets.Netvars.MClrRender, &sdk.CLRColorRender{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		})

	}
}

func EngineChams(entity uintptr) {

	if !ValidatePlayerGlow(entity) {
		return
	}

	if !configs.G.D.EngineChams {

		memory.Write(&memory.GameProcess, entity+configs.Offsets.Netvars.MClrRender, &sdk.CLRColorRender{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		})

		return

	}

	isEnemy, _ := csgo.PlayerIsEnemy(entity)
	if !isEnemy {
		return
	}

	eGlow := csgo.GetEntityGlow(entity)
	clrColorStruct := sdk.CLRColorRender{
		R: byte(eGlow.Red * 255),
		G: byte(eGlow.Green * 255),
		B: byte(eGlow.Blue * 255),
		A: 255,
	}

	memory.Write(&memory.GameProcess, entity+configs.Offsets.Netvars.MClrRender, &clrColorStruct)

}
