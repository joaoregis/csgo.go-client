package global

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/joaoregis/machineid"
	"github.com/lxn/win"
)

var (
	HARDWARE_ID, _          = machineid.ProtectedID(APP_UUID_4_GO)
	CONFIG_NAME             = GetConfigName() + ".cfg"
	CONFIG_NAME_WITHOUT_EXT = GetConfigName()
	USER_HOME_PATH          = GetUserHomePath()
)

var (
	HWND         win.HWND
	HWND_RECT    win.RECT
	WINDOW       *glfw.Window
	OVERLAY_HWND win.HWND
)
