package main

import (
	"gosource/internal/global"
	"gosource/internal/global/configs"
	"gosource/internal/hackFunctions/keyboard"
)

func handleKeyboardEvents() {

	if keyboard.GetAsyncKeyStateOnce(keyboard.GetKey(configs.G.D.ReloadKey)) {
		configs.Reload()
	}

	if keyboard.GetAsyncKeyStateOnce(keyboard.GetKey(configs.G.D.StopKey)) {
		global.WINDOW_OVERLAY.SetShouldClose(true)
	}

}
