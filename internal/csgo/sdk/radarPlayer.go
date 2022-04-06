package sdk

import "gosource/internal/hackFunctions/vector"

type RadarPlayer_T struct {
	Origin     vector.Vector3 //0x0000
	ViewAngles vector.Vector3 //0x000C
	PAD_0018   [56]byte       //0x0018
	Health     uint32         //0x0050
	Name       [128]rune      //0x0054
	PAD_00D4   [117]byte      //0x00D4
	Visible    uint8          //0x00E9
} //Size: 0x0B32
