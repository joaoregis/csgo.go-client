package csgo

import (
	"gosource/internal/global/configs"
	"gosource/internal/memory"
)

func IsCursorEnabled() bool {

	hDwMouseEnable, _ := memory.GameProcess.ReadInt(memory.GameClient + configs.Offsets.Signatures.DwMouseEnable)
	hDwMouseEnablePtr, _ := memory.GameProcess.ReadInt(memory.GameClient + configs.Offsets.Signatures.DwMouseEnablePtr)
	hDwMouseEnabledValue, _ := memory.GameProcess.ReadInt(uintptr(hDwMouseEnable))
	IsMouseCursorEnabled := (hDwMouseEnabledValue == hDwMouseEnablePtr)
	return IsMouseCursorEnabled

}

func GetMaxPlayers() (int, error) {
	moduleBase, err := memory.GameProcess.ReadIntPtr(memory.GameEngine + configs.Offsets.Signatures.DwClientState)

	if err != nil {
		return 0, err
	}

	return memory.GameProcess.ReadInt(moduleBase + configs.Offsets.Signatures.DwClientStateMaxPlayer)
}
