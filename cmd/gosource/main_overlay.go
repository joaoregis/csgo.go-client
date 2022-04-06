package main

import (
	"gosource/internal/global"
	"gosource/internal/global/utils"
	"gosource/internal/memory"
	"strings"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/lxn/win"
)

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
