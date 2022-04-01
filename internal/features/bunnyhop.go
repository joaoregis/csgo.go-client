package features

import (
	c "gosource/internal/configs"
	"gosource/internal/csgo"
	kb "gosource/internal/hackFunctions/keyboard"
	"gosource/internal/memory"
)

var isBhopping bool = false

const FL_ONGROUND = 0x100

func BunnyHop() {

	if jumping, _ := memory.GameProcess.ReadInt(memory.GameClient + c.Offsets.Signatures.DwForceJump); jumping == 5 {

		if c.G.Bunnyhop.Enabled {

			if localPlayer, err := csgo.GetLocalPlayer(); err == nil {

				if isBhopping {
					return
				}

				// async processing of bhop
				isBhopping = true
				go func() {

					for kb.GetAsyncKeyState(kb.GetKey(c.G.Bunnyhop.Key)) {

						fFlags, _ := memory.GameProcess.ReadInt(localPlayer + c.Offsets.Netvars.MFFlags)
						if fFlags == FL_ONGROUND {
							memory.GameProcess.WriteInt(memory.GameClient+c.Offsets.Signatures.DwForceJump, 4)
						} else {
							memory.GameProcess.WriteInt(memory.GameClient+c.Offsets.Signatures.DwForceJump, 5)
						}

					}

					isBhopping = false
				}()

			}

		}

	}

}
