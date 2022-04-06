package utils

import (
	"log"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

var (
	mod                 = windows.NewLazyDLL("user32.dll")
	procMapWindowPoints = mod.NewProc("MapWindowPoints")
)

func FindWindow(wndName string) win.HWND {

	wnd, err := syscall.UTF16PtrFromString(wndName)
	if err != nil {
		log.Fatal(err)
	}

	return win.FindWindow(nil, wnd)

}

func GetClientRect(hwnd win.HWND) win.RECT {

	var rect win.RECT
	win.GetClientRect(hwnd, &rect)
	return rect

}

func GetLocalCoordinates(hWnd win.HWND) win.RECT {

	var rect win.RECT
	win.GetWindowRect(hWnd, &rect)

	procMapWindowPoints.Call(
		uintptr(win.HWND_DESKTOP),
		uintptr(win.GetParent(hWnd)),
		uintptr(unsafe.Pointer(&rect)),
		uintptr(2),
	)

	return rect
}
