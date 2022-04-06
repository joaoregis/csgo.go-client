package main

import (
	"gosource/internal/global"
	"gosource/internal/global/logs"
	"gosource/internal/global/utils"
	"gosource/internal/memory"
	"log"
	"strings"
	"syscall"
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/lxn/win"
)

func postInitOpenGL() {

	// Overlay hit test handling
	global.HWND_OVERLAY = win.HWND(uintptr(unsafe.Pointer(global.WINDOW_OVERLAY.GetWin32Window())))
	wproc := syscall.NewCallback(wndProc)
	win.SetWindowLongPtr(global.HWND_OVERLAY, win.GWLP_WNDPROC, wproc)
	extendedStyle := win.GetWindowLong(global.HWND_OVERLAY, win.GWL_EXSTYLE)
	win.SetWindowLong(global.HWND_OVERLAY, win.GWL_EXSTYLE, extendedStyle|win.WS_EX_TRANSPARENT|win.WS_EX_LAYERED|win.WS_EX_NOACTIVATE)

	// create bitmaps for the device context font's first 256 glyphs
	win.WglUseFontBitmaps(win.HDC(global.HWND_OVERLAY), 0, 256, 1000)

}

func initOpenGL() {

	/**** START: OPEN-GL INITIALIZATION ****/
	logs.Info("initializing resources ...")

	err := glfw.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Floating, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)

	global.WINDOW_OVERLAY, err = glfw.CreateWindow(1, 1, global.HARDWARE_ID, nil, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer global.WINDOW_OVERLAY.Destroy()
	global.WINDOW_OVERLAY.SetAttrib(glfw.Decorated, glfw.False)
	global.WINDOW_OVERLAY.MakeContextCurrent()

	glfw.SwapInterval(1)

	if err := gl.Init(); err != nil {
		log.Fatal(err)
	}

	global.InitFonts()
	/**** END: OPEN-GL INITIALIZATION ****/

}

func beginFrame() bool {

	display_w, display_h := glfw.GetCurrentContext().GetFramebufferSize()
	gl.Viewport(0, 0, int32(display_w), int32(display_h))
	global.WINDOW_OVERLAY.SwapBuffers()
	glfw.PollEvents()
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// only will perform client actions when counter-strike is focused
	if hwnd := win.GetForegroundWindow(); hwnd != 0 {

		hwndText := memory.GetWindowText(memory.HWND(hwnd))
		if !strings.Contains(hwndText, "Counter-Strike: Global Offensive") {
			return false
		}

	}

	global.HWND_GAME_RECT = utils.GetClientRect(global.HWND_GAME)
	global.HWND_GAME_POS = utils.GetLocalCoordinates(global.HWND_GAME)

	win.SetWindowPos(
		global.HWND_OVERLAY,
		win.HWND_TOPMOST,
		int32(global.HWND_GAME_POS.Left),
		int32(global.HWND_GAME_POS.Top),
		int32(global.HWND_GAME_RECT.Right),
		int32(global.HWND_GAME_RECT.Bottom),
		0,
	)

	return true
}

func finishFrame() bool {

	if global.WINDOW_OVERLAY.ShouldClose() {
		return false
	}

	global.HWND_GAME = utils.FindWindow("Counter-Strike: Global Offensive - Direct3D 9")
	gl.Flush()

	return true

}
