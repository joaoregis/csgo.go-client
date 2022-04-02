package features

import (
	"gosource/internal/csgo"
	"gosource/internal/global/configs"
	"gosource/internal/memory"
)

func ClearEngineChams() {

	// entity list loop
	for i := 0; i < csgo.MaxPlayers; i++ {

		entity, _ := csgo.GetPlayerByIndex(i)

		memory.GameProcess.WriteBytes(entity+configs.Offsets.Netvars.MClrRender, csgo.CLRColorRender{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		}.Bytes())

	}
}

func EngineChams(entity uintptr) {

	if !configs.G.D.EngineChams {

		memory.GameProcess.WriteBytes(entity+configs.Offsets.Netvars.MClrRender, csgo.CLRColorRender{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		}.Bytes())

		return

	}

	isEnemy, _ := csgo.PlayerIsEnemy(entity)
	if !isEnemy {
		return
	}

	eGlow := csgo.GetEntityGlow(entity)
	clrColorStruct := csgo.CLRColorRender{
		R: byte(eGlow.Red * 255),
		G: byte(eGlow.Green * 255),
		B: byte(eGlow.Blue * 255),
		A: 255,
	}

	memory.GameProcess.WriteBytes(entity+configs.Offsets.Netvars.MClrRender, clrColorStruct.Bytes())

}
