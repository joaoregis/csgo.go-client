package main

import (
	"gosource/internal/global"
	"gosource/internal/global/logs"
	"gosource/internal/global/utils"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type VoidFunc func()
type VoidRetErrorFunc func()

var (
	PROJECT_NAME      = "CSGO.GO"
	BUILD_TIMESTAMP   = "development"
	IS_OVERLAY_ACTIVE = false
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()

}

func main() {

	initGameHwnd()

	initOpenGL()
	defer glfw.Terminate()
	defer global.WINDOW_OVERLAY.Destroy()

	initClientHeader()
	postInitOpenGL()

	if b := checkHwidAuth(); !b {
		logs.Info("your hwid is not registered yet. please, submit it to a admin:")
		logs.Info(utils.GetProtectHwid())
		return
	}

	go clientVMatrixLoop()
	clientMainLoop()
	gracefulExit()

}
