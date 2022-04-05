package main

import (
	"fmt"
	"gosource/internal/csgo"
	"gosource/internal/csgo/offsets"
	"gosource/internal/csgo/sdk"
	"gosource/internal/features"
	"gosource/internal/global"
	"gosource/internal/global/configs"
	"gosource/internal/global/utils"
	"gosource/internal/hackFunctions/keyboard"
	"gosource/internal/memory"
	"log"
	"runtime"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/lxn/win"
)

var (
	PROJECT_NAME    = "CSGO.GO"
	BUILD_TIMESTAMP = "development"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()

}

func WndProc(hWnd win.HWND, Msg uint32, wParam uintptr, lParam uintptr) uintptr {

	if Msg == win.WM_NCHITTEST {
		return win.HTNOWHERE
	}

	return win.DefWindowProc(hWnd, Msg, wParam, lParam)

}

var AlreadyShowed bool = false

func main() {

	for global.HWND == 0 {
		global.HWND = utils.FindWindow("Counter-Strike: Global Offensive - Direct3D 9")
		fmt.Println("awaiting for csgo ...")
		time.Sleep(1000 * time.Millisecond)
	}

	/**** START: OPEN-GL INITIALIZATION ****/
	fmt.Println("initializing resources ...")

	err := glfw.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Floating, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)

	global.WINDOW, err = glfw.CreateWindow(1, 1, global.HARDWARE_ID, nil, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer global.WINDOW.Destroy()

	global.WINDOW.SetAttrib(glfw.Decorated, glfw.False)

	global.WINDOW.MakeContextCurrent()

	glfw.SwapInterval(1)

	if err := gl.Init(); err != nil {
		log.Fatal(err)
	}

	global.InitFonts()

	/**** END: OPEN-GL INITIALIZATION ****/

	/* After initialization of opengl */
	utils.SetConsoleTitle(PROJECT_NAME + " - build: " + BUILD_TIMESTAMP)

	if global.DEBUG_MODE {

		fmt.Println("=====================================================================================")
		fmt.Println("hardware id: ", global.HARDWARE_ID)
		fmt.Println("default dir: ", global.USER_HOME_PATH)
		fmt.Println("config file: ", global.CONFIG_NAME)
		fmt.Println("=====================================================================================")

	}

	fmt.Println("getting latest offsets for current game version")
	offsets.InitOffsets(&configs.Offsets)

	fmt.Println("initializing key functions")
	keyboard.InitKeys()

	for !memory.Init() {
		fmt.Println("trying to initialize the memory module ...")
		time.Sleep(1000 * time.Millisecond)
	}

	fmt.Println("initializing user configs")
	configs.Init()

	fmt.Println("scanning for new game patterns")
	updateOffsetsByPatterns()

	// Overlay hit test handling
	global.OVERLAY_HWND = win.HWND(uintptr(unsafe.Pointer(global.WINDOW.GetWin32Window())))
	wproc := syscall.NewCallback(WndProc)
	win.SetWindowLongPtr(global.OVERLAY_HWND, win.GWLP_WNDPROC, wproc)
	extendedStyle := win.GetWindowLong(global.OVERLAY_HWND, win.GWL_EXSTYLE)
	win.SetWindowLong(global.OVERLAY_HWND, win.GWL_EXSTYLE, extendedStyle|win.WS_EX_TRANSPARENT|win.WS_EX_LAYERED|win.WS_EX_NOACTIVATE)

	// View Matrix Update Loop
	go func() {
		for !global.WINDOW.ShouldClose() && global.HWND != 0 {
			sdk.UpdateViewMatrix()
		}
	}()

	// Mainloop
	fmt.Println("everything is fine. good hacking.")
	for !global.WINDOW.ShouldClose() && global.HWND != 0 {

		glfw.PollEvents()
		gl.Clear(gl.COLOR_BUFFER_BIT)

		global.HWND_RECT = utils.GetClientRect(global.HWND)
		global.WINDOW.SetSize(int(global.HWND_RECT.Right), int(global.HWND_RECT.Bottom))

		hCsPositionOnScreen := utils.GetLocalCoordinates(global.HWND)
		global.WINDOW.SetPos(int(hCsPositionOnScreen.Left), int(hCsPositionOnScreen.Top))

		// only will perform client actions when counter-strike is focused
		if hwnd := win.GetForegroundWindow(); hwnd != 0 {

			hwndText := memory.GetWindowText(memory.HWND(hwnd))
			if !strings.Contains(hwndText, "Counter-Strike: Global Offensive") {
				if AlreadyShowed {
					global.WINDOW.Hide()
					AlreadyShowed = false
				}
				continue
			}

			if !AlreadyShowed {

				AlreadyShowed = true
				global.WINDOW.Show()

			}

		}

		win.BringWindowToTop(global.OVERLAY_HWND)
		win.SetFocus(global.HWND)

		//
		if csgo.UpdatePlayerVars() != nil {
			continue
		}

		if keyboard.GetAsyncKeyStateOnce(keyboard.GetKey(configs.G.D.ReloadKey)) {
			configs.Reload()
		}

		if keyboard.GetAsyncKeyStateOnce(keyboard.GetKey(configs.G.D.StopKey)) {
			global.WINDOW.SetShouldClose(true)
		}

		// skip these features when cursor is enabled
		if !csgo.IsCursorEnabled() {

			features.Visuals()
			features.AutoWeapons()
			features.Triggerbot()
			features.BunnyHop()
			features.Aimbot()

		}

		if global.WINDOW.ShouldClose() {
			break
		}

		global.HWND = utils.FindWindow("Counter-Strike: Global Offensive - Direct3D 9")

		// Rendering
		display_w, display_h := glfw.GetCurrentContext().GetFramebufferSize()
		gl.Viewport(0, 0, int32(display_w), int32(display_h))
		global.WINDOW.SwapBuffers()

	}

	endCheat()
	fmt.Println("good bye. cya!")
}

func endCheat() {
	fmt.Println("clearing client residues ...")
}

func updateOffsetsByPatterns() {

	var patterns = map[string]string{
		"DwEntityList":        "BB ? ? ? ? 83 FF 01 0F 8C ? ? ? ? 3B F8",
		"DwForceAttack":       "89 0D ? ? ? ? 8B 0D ? ? ? ? 8B F2 8B C1 83 CE 04",
		"DwForceJump":         "8B 0D ? ? ? ? 8B D6 8B C1 83 CA 02",
		"DwLocalPlayer":       "8D 34 85 ? ? ? ? 89 15 ? ? ? ? 8B 41 08 8B 48 04 83 F9 FF",
		"DwGlowObjectManager": "A1 ? ? ? ? A8 01 75 4B",
		"DwClientState":       "A1 ? ? ? ? 33 D2 6A 00 6A 00 33 C9 89 B0",
	}

	configs.Offsets.Signatures.DwEntityList, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwEntityList"], true, 0x1, 0x0)
	configs.Offsets.Signatures.DwForceAttack, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwForceAttack"], true, 0x2, 0x0)
	configs.Offsets.Signatures.DwForceJump, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwForceJump"], true, 0x2, 0x0)
	configs.Offsets.Signatures.DwLocalPlayer, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwLocalPlayer"], true, 0x3, 0x4)
	configs.Offsets.Signatures.DwGlowObjectManager, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwGlowObjectManager"], true, 0x1, 0x4)
	configs.Offsets.Signatures.DwClientState, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Engine], patterns["DwClientState"], true, 0x1, 0x0)

}
