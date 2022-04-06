package main

import (
	"fmt"
	"gosource/internal/csgo/offsets"
	"gosource/internal/global/configs"
	"gosource/internal/memory"
)

func endCheat() {
	fmt.Println("clearing client residues ...")
}

func updateOffsetsByPatterns() {

	var patterns = map[string]string{
		"DwEntityList":        "BB ? ? ? ? 83 FF 01 0F 8C ? ? ? ? 3B F8",
		"DwForceAttack":       "89 0D ? ? ? ? 8B 0D ? ? ? ? 8B F2 8B C1 83 CE 04",
		"DwForceJump":         "8B 0D ? ? ? ? 8B D6 8B C1 83 CA 02",
		"DwLocalPlayer":       "8D 34 85 ? ? ? ? 89 15 ? ? ? ? 8B 41 08 8B 48 04 83 F9 FF",
		"DwGlowObjectManager": "A1 ? ? ? ? A8 01 75 4B",
		"DwClientState":       "A1 ? ? ? ? 33 D2 6A 00 6A 00 33 C9 89 B0",
	}

	configs.Offsets.Signatures.DwEntityList, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwEntityList"], true, 0x1, 0x0)
	configs.Offsets.Signatures.DwForceAttack, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwForceAttack"], true, 0x2, 0x0)
	configs.Offsets.Signatures.DwForceJump, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwForceJump"], true, 0x2, 0x0)
	configs.Offsets.Signatures.DwLocalPlayer, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwLocalPlayer"], true, 0x3, 0x4)
	configs.Offsets.Signatures.DwGlowObjectManager, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwGlowObjectManager"], true, 0x1, 0x4)
	configs.Offsets.Signatures.DwClientState, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Engine], patterns["DwClientState"], true, 0x1, 0x0)

}
