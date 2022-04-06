package sdk

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

type CLRColorRender struct {
	R byte
	G byte
	B byte
	A byte
}
