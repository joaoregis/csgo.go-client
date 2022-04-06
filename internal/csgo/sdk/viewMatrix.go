package sdk

import (
	"gosource/internal/global/configs"
	"gosource/internal/memory"
)

type VMatrix [16]float32

var ViewMatrix VMatrix

func UpdateViewMatrix() {

	results, _ := memory.GameProcess.ReadFloats32(memory.GameClient+configs.Offsets.Signatures.DwViewMatrix, 0x10*0x4)
	copy(ViewMatrix[:], results)

}
