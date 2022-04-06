package global

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/joaoregis/machineid"
	"github.com/lxn/win"
)

var (
	HARDWARE_ID, _          = machineid.ProtectedID(APP_UUID_4_GO)
	CONFIG_NAME             = GetConfigName() + ".cfg"
	CONFIG_NAME_WITHOUT_EXT = GetConfigName()
	USER_HOME_PATH          = GetUserHomePath()
	SPRINT_F                = fmt.Sprintf
)

var (
	HWND_GAME      win.HWND
	HWND_GAME_POS  win.RECT
	HWND_GAME_RECT win.RECT
	HWND_GAME_SIZE win.RECT
	HWND_OVERLAY   win.HWND
	WINDOW_OVERLAY *glfw.Window
)
