package features

import (
	c "gosource/internal/configs"
	"gosource/internal/csgo"
	kb "gosource/internal/hackFunctions/keyboard"
	"gosource/internal/memory"
)

func Triggerbot() {

	if !c.G.Triggerbot.Enabled {
		return
	}

	if csgo.MaxPlayers <= 0 || csgo.LocalEntity == 0 {
		return
	}

	crosshairTargetId, _ := memory.GameProcess.ReadInt(csgo.LocalEntity + c.Offsets.Netvars.MICrosshairID)
	targetIndex := int(crosshairTargetId) - 1

	targetEntityAtCrosshair, err := csgo.GetPlayerByIndex(targetIndex)
	if err != nil {
		return
	}

	isEnemyTarget, _ := csgo.PlayerIsEnemy(targetEntityAtCrosshair)
	isCrosshairIdGreaterThanZero := crosshairTargetId > 0
	isCrosshairIdLowerThanMaxPlayers := int64(crosshairTargetId) <= int64(csgo.MaxPlayers)
	isHotkeyPressed := kb.GetAsyncKeyState(kb.GetKey(c.G.Triggerbot.Key))
	isValidPlayer, _ := csgo.PlayerIsValid(targetEntityAtCrosshair)
	isImmune, _ := csgo.IsEntityImmune(targetEntityAtCrosshair)

	shouldShot :=
		isEnemyTarget &&
			isCrosshairIdGreaterThanZero &&
			isCrosshairIdLowerThanMaxPlayers &&
			isHotkeyPressed &&
			isValidPlayer &&
			!isImmune

	if shouldShot {

		memory.GameProcess.Write(memory.GameClient+c.Offsets.Signatures.DwForceAttack, "int", 6)
	}

}
