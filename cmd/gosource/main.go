package main

import (
	"gosource/internal/global"
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
	go clientVMatrixLoop()
	clientMainLoop()
	gracefulExit()

}
