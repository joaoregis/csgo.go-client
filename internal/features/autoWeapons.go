package features

import (
	"gosource/internal/csgo"
	"gosource/internal/global/configs"
	"gosource/internal/hackFunctions/keyboard"
	"gosource/internal/memory"
	"math/rand"
	"time"
)

var t time.Time

func AutoWeapons() {

	if hWIdx, _ := csgo.GetActiveWeapon(); configs.G.D.AutoWeapons.Enabled && (csgo.IsPistol(hWIdx) || csgo.IsDeagle(hWIdx)) {
		elapsed := time.Since(t).Milliseconds()
		// 15ms is the minimum value and 25ms is the "maximum" minimum value to make it random
		if elapsed > int64(rand.Intn(25-15)+15)+int64(configs.G.D.AutoWeapons.Delay) {
			if keyboard.GetAsyncKeyState(keyboard.GetKey("Mouse 1")) {
				t = time.Now()
				memory.GameProcess.Write(memory.GameClient+configs.Offsets.Signatures.DwForceAttack, "int", 4)
				memory.GameProcess.Write(memory.GameClient+configs.Offsets.Signatures.DwForceAttack, "int", 6)
			}
		}
	}

}
