package features

import (
	"fmt"
	c "gosource/internal/configs"
	"gosource/internal/csgo"
	kb "gosource/internal/hackFunctions/keyboard"
	"gosource/internal/memory"
)

func BunnyHop() {

	if c.G.Bunnyhop.Enabled && kb.GetAsyncKeyState(kb.GetKey(c.G.Bunnyhop.Key)) {

		localPlayer, err := csgo.GetLocalPlayer()
		if err != nil {
			return
		}

		fFlags, err := memory.GameProcess.ReadInt(localPlayer + c.Offsets.Netvars.MFFlags)

		if err != nil {
			fmt.Println("Error on found fFlags", err)
			return
		}

		if fFlags == 256 {
			memory.GameProcess.WriteInt(memory.GameClient+c.Offsets.Signatures.DwForceJump, 4)
		} else {
			memory.GameProcess.WriteInt(memory.GameClient+c.Offsets.Signatures.DwForceJump, 5)
		}

	}

}
