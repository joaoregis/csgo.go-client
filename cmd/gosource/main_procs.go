package main

/* Main package private members */

import (
	"gosource/internal/csgo"
	"gosource/internal/csgo/sdk"
	"gosource/internal/features"
	"gosource/internal/global"
	"gosource/internal/global/logs"

	"github.com/lxn/win"
)

func clientVMatrixLoop() {
	for !global.WINDOW_OVERLAY.ShouldClose() && global.HWND_GAME != 0 {
		sdk.UpdateViewMatrix()
	}
}

func clientMainLoop() {

	// Mainloop
	logs.Info("everything is fine. good hacking.")
	for !global.WINDOW_OVERLAY.ShouldClose() && global.HWND_GAME != 0 {

		if b := beginFrame(); !b {
			// Should go to the next frame
			finishFrame()
			continue
		}

		//
		if csgo.UpdatePlayerVars() != nil {
			continue
		}

		handleKeyboardEvents()

		// prevent client from working when cursor is enabled
		if !csgo.IsCursorEnabled() {

			features.Visuals()
			features.AutoWeapons()
			features.Triggerbot()
			features.BunnyHop()
			features.Aimbot()

		}

		if b := finishFrame(); !b {
			// Should break main loop to end the process
			break
		}
	}

}

func wndProc(hWnd win.HWND, Msg uint32, wParam uintptr, lParam uintptr) uintptr {

	if Msg == win.WM_NCHITTEST {
		return win.HTNOWHERE
	}

	return win.DefWindowProc(hWnd, Msg, wParam, lParam)

}
