package features

import (
	"gosource/internal/csgo"
	c "gosource/internal/global/configs"
	kb "gosource/internal/hackFunctions/keyboard"
	"gosource/internal/memory"
)

var isBhopping bool = false

const (
	FL_ONGROUND = 0x100
	VK_SPACE    = 0x20
)

func BunnyHop() {

	if c.G.D.Bunnyhop {

		if jumping, _ := memory.GameProcess.ReadInt(memory.GameClient + c.Offsets.Signatures.DwForceJump); jumping == 5 {

			if localPlayer, err := csgo.GetLocalPlayer(); err == nil {

				if isBhopping {
					return
				}

				// async processing of bhop
				isBhopping = true
				go func() {

					for kb.GetAsyncKeyState(VK_SPACE) {

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
