package features

import (
	"bytes"
	"encoding/binary"
	"gosource/internal/configs"
	"gosource/internal/hackFunctions/vector"
	"gosource/internal/memory"
	"unsafe"

	"golang.org/x/sys/windows"
)

type Radar_player_t struct {
	Origin     vector.Vector3 //0x0000
	Viewangles vector.Vector3 //0x000C
	Pad_0018   [56]byte       //0x0018
	Health     uint32         //0x0050
	Name       [32]byte       //0x0054
	Pad_00D4   [117]byte      //0x00D4
	Visible    uint32         //0x00E9
} //Size: 0x0B32

func Radar(ientity int) {

	radar_base, _ := memory.GameProcess.ReadIntPtr(memory.GameClient + configs.Offsets.Signatures.DwRadarBase + 0x74)

	radar_player_addr := radar_base + (0x174 * (uintptr(ientity) + 1)) - 0x3C
	radar_player_b, _ := memory.GameProcess.ReadBytes(radar_player_addr, 0x0B32)

	var oldprotect uint32
	windows.VirtualProtect(radar_player_addr, 0x0B32, windows.PAGE_EXECUTE_READWRITE, &oldprotect)

	var radar_player_s Radar_player_t
	b := bytes.NewBuffer(radar_player_b)
	_ = binary.Read(b, binary.LittleEndian, &radar_player_s)

	radar_player_s.Visible = 0x1 // change visibility to true

	const sz = int(unsafe.Sizeof(Radar_player_t{}))
	var asByteSlice []byte = (*(*[sz]byte)(unsafe.Pointer(&radar_player_s)))[:]

	memory.GameProcess.WriteBytes(radar_player_addr, asByteSlice)

	windows.VirtualProtect(radar_player_addr, 0x0B32, oldprotect, &oldprotect)

	// memory.GameProcess.WriteInt(entity+configs.Offsets.Netvars.MBSpotted, 1)

}
