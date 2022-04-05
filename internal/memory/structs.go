package memory

import (
	"syscall"

	"github.com/Xustyx/w32"
	"golang.org/x/sys/windows"
)

const (
	MEM_COMMIT             = 0x1000
	PAGE_NOACCESS          = 0x01
	PAGE_EXECUTE_READWRITE = 0x40
)

var (
	kernel32              = windows.MustLoadDLL("kernel32.dll")
	virtualQueryEx        = kernel32.MustFindProc("VirtualQueryEx")
	virtualProtectEx      = kernel32.MustFindProc("VirtualProtectEx")
	procReadProcessMemory = kernel32.MustFindProc("ReadProcessMemory")
)

var (
	user32   = syscall.NewLazyDLL("user32.dll")
	ProcGAKS = user32.NewProc("GetAsyncKeyState")
)

type MemoryBasicInformation struct {
	BaseAddress       uintptr
	AllocationBase    uintptr
	AllocationProtect uint32
	RegionSize        uintptr
	State             uint32
	Protect           uint32
	Type              uint32
}

type Process struct {
	Name        string
	Handle      w32.HANDLE
	Pid         uint32
	ModBaseAddr uintptr
	ModBaseSize uint32
	Modules     map[string]Module
}

type Module struct {
	Name        string
	ModBaseAddr uintptr
	ModBaseSize uint32
}
