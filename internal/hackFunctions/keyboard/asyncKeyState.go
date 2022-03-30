package keyboard

import (
	"gosource/internal/memory"
)

func GetAsyncKeyState(key uint) bool {

	r, _, _ := memory.ProcGAKS.Call(uintptr(key))
	return int(r) != 0

}

func GetAsyncKeyStateOnce(key uint) bool {

	r, _, _ := memory.ProcGAKS.Call(uintptr(key))

	if int(r) != 0 {
		for {

			r, _, _ = memory.ProcGAKS.Call(uintptr(key))
			if r == 0 {
				break
			}

		}

		return true
	}

	return false

}
