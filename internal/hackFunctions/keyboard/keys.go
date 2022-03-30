package keyboard

type Key int64

var _KEYS map[string]uint = make(map[string]uint)

func GetKey(k string) uint {
	return _KEYS[k]
}

func InitKeys() {

	_KEYS["Numpad +"] = 0x6B
	_KEYS["Backspace"] = 0x08
	_KEYS["Pause Break"] = 0x03
	_KEYS["Numpad ."] = 0x6E
	_KEYS["Numpad /"] = 0x6F
	_KEYS["Esc"] = 0x1B
	_KEYS["0"] = 0x30
	_KEYS["1"] = 0x31
	_KEYS["2"] = 0x32
	_KEYS["3"] = 0x33
	_KEYS["4"] = 0x34
	_KEYS["5"] = 0x35
	_KEYS["6"] = 0x36
	_KEYS["7"] = 0x37
	_KEYS["8"] = 0x38
	_KEYS["9"] = 0x39
	_KEYS["A"] = 0x41
	_KEYS["B"] = 0x42
	_KEYS["C"] = 0x43
	_KEYS["D"] = 0x44
	_KEYS["E"] = 0x45
	_KEYS["F"] = 0x46
	_KEYS["G"] = 0x47
	_KEYS["H"] = 0x48
	_KEYS["I"] = 0x49
	_KEYS["J"] = 0x4A
	_KEYS["K"] = 0x4B
	_KEYS["L"] = 0x4C
	_KEYS["M"] = 0x4D
	_KEYS["N"] = 0x4E
	_KEYS["O"] = 0x4F
	_KEYS["P"] = 0x50
	_KEYS["Q"] = 0x51
	_KEYS["R"] = 0x52
	_KEYS["S"] = 0x53
	_KEYS["T"] = 0x54
	_KEYS["U"] = 0x55
	_KEYS["V"] = 0x56
	_KEYS["W"] = 0x57
	_KEYS["X"] = 0x58
	_KEYS["Y"] = 0x59
	_KEYS["Z"] = 0x5A
	_KEYS["Numpad *"] = 0x6A
	_KEYS["Numpad 0"] = 0x60
	_KEYS["Numpad 1"] = 0x61
	_KEYS["Numpad 2"] = 0x62
	_KEYS["Numpad 3"] = 0x63
	_KEYS["Numpad 4"] = 0x64
	_KEYS["Numpad 5"] = 0x65
	_KEYS["Numpad 6"] = 0x66
	_KEYS["Numpad 7"] = 0x67
	_KEYS["Numpad 8"] = 0x68
	_KEYS["Numpad 9"] = 0x69
	_KEYS[";"] = 0xBA
	_KEYS["<"] = 0xE2
	_KEYS["/"] = 0xBF
	_KEYS["`"] = 0xC0
	_KEYS["["] = 0xDB
	_KEYS["|"] = 0xDC
	_KEYS["]"] = 0xDD
	_KEYS["\""] = 0xDE
	_KEYS["!"] = 0xDF
	_KEYS["<"] = 0xBC
	_KEYS["-"] = 0xBD
	_KEYS[">"] = 0xBE
	_KEYS["="] = 0xBB
	_KEYS["Enter"] = 0x0D
	_KEYS["Space"] = 0x20
	_KEYS["Numpad -"] = 0x6D
	_KEYS["Tab"] = 0x09
	_KEYS["Context Menu"] = 0x5D
	_KEYS["Caps Lock"] = 0x14
	_KEYS["Delete"] = 0x2E
	_KEYS["Arrow Down"] = 0x28
	_KEYS["End"] = 0x23
	_KEYS["F1"] = 0x70
	_KEYS["F10"] = 0x79
	_KEYS["F11"] = 0x7A
	_KEYS["F12"] = 0x7B
	_KEYS["F13"] = 0x7C
	_KEYS["F14"] = 0x7D
	_KEYS["F15"] = 0x7E
	_KEYS["F16"] = 0x7F
	_KEYS["F17"] = 0x80
	_KEYS["F18"] = 0x81
	_KEYS["F19"] = 0x82
	_KEYS["F2"] = 0x71
	_KEYS["F20"] = 0x83
	_KEYS["F21"] = 0x84
	_KEYS["F22"] = 0x85
	_KEYS["F23"] = 0x86
	_KEYS["F24"] = 0x87
	_KEYS["F3"] = 0x72
	_KEYS["F4"] = 0x73
	_KEYS["F5"] = 0x74
	_KEYS["F6"] = 0x75
	_KEYS["F7"] = 0x76
	_KEYS["F8"] = 0x77
	_KEYS["F9"] = 0x78
	_KEYS["Home"] = 0x24
	_KEYS["Insert"] = 0x2D
	_KEYS["Mouse 1"] = 0x01
	_KEYS["Left Ctrl"] = 0xA2
	_KEYS["Arrow Left"] = 0x25
	_KEYS["Left Alt"] = 0xA4
	_KEYS["Left Shift"] = 0xA0
	_KEYS["Left Win"] = 0x5B
	_KEYS["Mouse 3"] = 0x04
	_KEYS["Page Down"] = 0x22
	_KEYS["Numpad Lock"] = 0x90
	_KEYS["Pause"] = 0x13
	_KEYS["Print"] = 0x2A
	_KEYS["Page Up"] = 0x21
	_KEYS["Mouse 2"] = 0x02
	_KEYS["Right Ctrl"] = 0xA3
	_KEYS["Arrow Right"] = 0x27
	_KEYS["Right Alt"] = 0xA5
	_KEYS["Right Shift"] = 0xA1
	_KEYS["Right Win"] = 0x5C
	_KEYS["Scroll Lock"] = 0x91
	_KEYS["Print Screen"] = 0x2C
	_KEYS["Arrow Up"] = 0x26
	_KEYS["Volume Down"] = 0xAE
	_KEYS["Volume Mute"] = 0xAD
	_KEYS["Volume Up"] = 0xAF
	_KEYS["Mouse 4"] = 0x05
	_KEYS["Mouse 5"] = 0x06

}
