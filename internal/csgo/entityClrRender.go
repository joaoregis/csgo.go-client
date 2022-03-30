package csgo

import "unsafe"

type CLRColorRender struct {
	R byte
	G byte
	B byte
	A byte
}

func (val CLRColorRender) Bytes() []byte {

	const sz = int(unsafe.Sizeof(CLRColorRender{}))
	var asByteSlice []byte = (*(*[sz]byte)(unsafe.Pointer(&val)))[:]
	return asByteSlice

}
