package features

import (
	c "gosource/internal/configs"
	"gosource/internal/csgo"
	kb "gosource/internal/hackFunctions/keyboard"
	"gosource/internal/memory"
	"time"
)

var triggerbotDelayCounter time.Time
var triggerbotDelayCounter_Registered bool = false

func Triggerbot() {

	if !c.G.Triggerbot.Enabled {
		return
	}

	hWId, _ := csgo.GetActiveWeapon()
	if !csgo.IsAimbotable(hWId) {
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

	isHotkeyPressed := kb.GetAsyncKeyState(kb.GetKey(c.G.Triggerbot.Key))

	if !isHotkeyPressed {
		triggerbotDelayCounter_Registered = false
		return
	}

	isEnemyTarget, _ := csgo.PlayerIsEnemy(targetEntityAtCrosshair)
	isCrosshairIdGreaterThanZero := crosshairTargetId > 0
	isCrosshairIdLowerThanMaxPlayers := int64(crosshairTargetId) <= int64(csgo.MaxPlayers)
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

		if c.G.Triggerbot.Delay > 0 {
			if !triggerbotDelayCounter_Registered {

				triggerbotDelayCounter_Registered = true
				triggerbotDelayCounter = time.Now()

				// return to check in the next frame if time elapsed had passed
				return

			} else {
				elapsed := time.Since(triggerbotDelayCounter).Milliseconds()
				if elapsed <= int64(c.G.Triggerbot.Delay) {
					// return to prevent trigger from firing before estabilished delay
					return
				} else {
					triggerbotDelayCounter_Registered = false
				}
			}
		}

		memory.GameProcess.Write(memory.GameClient+c.Offsets.Signatures.DwForceAttack, "int", 4)
		memory.GameProcess.Write(memory.GameClient+c.Offsets.Signatures.DwForceAttack, "int", 6)

	} else {

		triggerbotDelayCounter_Registered = false

	}

}
