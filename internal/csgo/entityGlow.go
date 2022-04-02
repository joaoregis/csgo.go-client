package csgo

import (
	"bytes"
	"encoding/binary"
	"gosource/internal/global/configs"
	"gosource/internal/memory"
	"unsafe"
)

type EntityGlowStruct struct {
	Red              float32 //0x8
	Green            float32 //0xC
	Blue             float32 //0x10
	Alpha            float32 //0x14
	Pad_0            [8]uint8
	Unk_0            float32
	Pad_1            [4]uint8
	RenderOccluded   byte
	RenderUnoccluded byte
	FullBloom        byte
}

func GetEntityGlow(entity uintptr) *EntityGlowStruct {

	const structSize uint = uint(unsafe.Sizeof(EntityGlowStruct{}))
	dwGlowObjectManager, _ := memory.GameProcess.ReadIntPtr(memory.GameClient + configs.Offsets.Signatures.DwGlowObjectManager)
	iGlowIndex, _ := memory.GameProcess.ReadIntPtr(entity + configs.Offsets.Netvars.MIGlowIndex)
	v, _ := memory.GameProcess.ReadBytes(dwGlowObjectManager+(iGlowIndex*0x38)+0x8, structSize)
	return ToEntityGlowStruct(v)

}

func ToEntityGlowStruct(v []byte) *EntityGlowStruct {

	var entityGlowInstance EntityGlowStruct
	b := bytes.NewBuffer(v) // b is a disk block of  2048 bytes
	_ = binary.Read(b, binary.LittleEndian, &entityGlowInstance)
	return &entityGlowInstance

}

func (v *EntityGlowStruct) Bytes() []byte {

	const sz = int(unsafe.Sizeof(EntityGlowStruct{}))
	var asByteSlice []byte = (*(*[sz]byte)(unsafe.Pointer(v)))[:]
	return asByteSlice

}

func (v *EntityGlowStruct) SetColorRGBA(r float32, g float32, b float32, a float32) {
	v.Red = r
	v.Green = g
	v.Blue = b
	v.Alpha = a
}

func (v *EntityGlowStruct) Save(entity uintptr) {

	dwGlowObjectManager, _ := memory.GameProcess.ReadIntPtr(memory.GameClient + configs.Offsets.Signatures.DwGlowObjectManager)
	iGlowIndex, _ := memory.GameProcess.ReadIntPtr(entity + configs.Offsets.Netvars.MIGlowIndex)
	memory.GameProcess.WriteBytes(dwGlowObjectManager+(iGlowIndex*0x38)+0x8, v.Bytes())
	memory.GameProcess.WriteInt(dwGlowObjectManager+iGlowIndex*0x38+0x27, 1) // Enables Glow at Entity
	memory.GameProcess.WriteInt(dwGlowObjectManager+iGlowIndex*0x38+0x28, 1) // Enables Glow at Entity

}
