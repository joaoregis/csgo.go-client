package main

/* Main package private members */

import (
	"encoding/json"
	"gosource/internal/csgo"
	"gosource/internal/csgo/sdk"
	"gosource/internal/features"
	"gosource/internal/global"
	"gosource/internal/global/logs"
	"gosource/internal/global/utils"
	"io/ioutil"
	"net/http"

	"github.com/lxn/win"
	"github.com/pkg/errors"
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

func checkHwidAuth() bool {

	var uri = "https://raw.githubusercontent.com/joaoregis/csgo.go-client-control/main/user-list.dat"

	logs.Info("checking hwid ...")
	resp, err := http.Get(uri)
	if err != nil {
		logs.Fatal(errors.Wrap(err, "cannot check hwid. aborting ...").Error())
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Fatal(errors.Wrap(err, "cannot read hwid list. aborting ...").Error())
	}

	var o []string
	err = json.Unmarshal(b, &o)
	if err != nil {
		logs.Fatal(errors.Wrap(err, "could not decode hwid. aborting ...").Error())
	}

	if utils.Contains(o, utils.GetProtectHwid()) {
		return true
	}

	return false

}
