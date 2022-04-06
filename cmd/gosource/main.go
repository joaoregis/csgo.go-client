package main

import (
	"gosource/internal/csgo/offsets"
	"gosource/internal/global"
	"gosource/internal/global/configs"
	"gosource/internal/global/logs"
	"gosource/internal/global/utils"
	"gosource/internal/hackFunctions/keyboard"
	"gosource/internal/memory"
	"log"
	"runtime"
	"syscall"
	"time"
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/lxn/win"
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

	for global.HWND_GAME == 0 {
		global.HWND_GAME = utils.FindWindow("Counter-Strike: Global Offensive - Direct3D 9")
		logs.Info("awaiting for csgo ...")
		time.Sleep(1000 * time.Millisecond)
	}

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

	/* After initialization of opengl */
	utils.SetConsoleTitle(PROJECT_NAME + " - build: " + BUILD_TIMESTAMP)

	if global.DEBUG_MODE {

		logs.Warn("NOTE: you are using a work-in-progress version of our client")
		logs.Warn("NOTE: keep-in-mind that this can have several bugs")
		logs.Warn("NOTE: this version isn't the true experience that we want to delivery")

	}

	logs.Info("getting latest offsets for current game version")
	offsets.InitOffsets(&configs.Offsets)

	logs.Info("initializing key functions")
	keyboard.InitKeys()

	for !memory.Init() {
		logs.Info("trying to initialize the memory module ...")
		time.Sleep(1000 * time.Millisecond)
	}

	logs.Info("initializing user configs")
	configs.Init()

	logs.Info("scanning for new game patterns")
	updateOffsetsByPatterns()

	// Overlay hit test handling
	global.HWND_OVERLAY = win.HWND(uintptr(unsafe.Pointer(global.WINDOW_OVERLAY.GetWin32Window())))
	wproc := syscall.NewCallback(wndProc)
	win.SetWindowLongPtr(global.HWND_OVERLAY, win.GWLP_WNDPROC, wproc)
	extendedStyle := win.GetWindowLong(global.HWND_OVERLAY, win.GWL_EXSTYLE)
	win.SetWindowLong(global.HWND_OVERLAY, win.GWL_EXSTYLE, extendedStyle|win.WS_EX_TRANSPARENT|win.WS_EX_LAYERED|win.WS_EX_NOACTIVATE)

	// create bitmaps for the device context font's first 256 glyphs
	win.WglUseFontBitmaps(win.HDC(global.HWND_OVERLAY), 0, 256, 1000)

	go clientVMatrixLoop()
	clientMainLoop()

	endCheat()
	logs.Info("good bye. cya!")
}
