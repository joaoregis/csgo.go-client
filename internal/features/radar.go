package features

import (
	"gosource/internal/configs"
	"gosource/internal/memory"
)

func Radar(entity uintptr) {

	memory.GameProcess.WriteInt(entity+configs.Offsets.Netvars.MBSpotted, 1)

}
