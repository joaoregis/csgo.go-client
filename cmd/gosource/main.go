package main

import (
	"runtime"
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
	initClientHeader()
	postInitOpenGL()
	go clientVMatrixLoop()
	clientMainLoop()
	gracefulExit()

}
