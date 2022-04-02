package main

import (
	"fmt"
	"gosource/internal/csgo"
	"gosource/internal/features"
	"gosource/internal/global"
	"gosource/internal/global/configs"
	"gosource/internal/global/utils"
	"gosource/internal/hackFunctions/keyboard"
	"gosource/internal/memory"
	"gosource/internal/offsets"
	"os"
	"os/signal"
	"strings"
	"time"
)

var (
	PROJECT_NAME    = "CSGO.GO"
	BUILD_TIMESTAMP = "development"
)

var ShouldContinue = true
var found = false

func main() {

	utils.SetConsoleTitle(PROJECT_NAME + " - build: " + BUILD_TIMESTAMP)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	go func() {

		sig := <-c
		if sig == os.Interrupt {
			fmt.Println("operation aborted by user. closing client.", sig)
			endCheat()
			ShouldContinue = false
			os.Exit(1)
		}

	}()

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
		fmt.Println("awaiting for game process to proceed ...")
		time.Sleep(1000 * time.Millisecond)
	}

	fmt.Println("initializing user configs")
	configs.Init()

	fmt.Println("scanning for new game patterns")
	updateOffsetsByPatterns()

	// Mainloop
	tries := 0
	fmt.Println("everything is fine. good hacking.")
	for ShouldContinue {

		// guarantee that offsets will be loaded correctly
		if !found {

			for configs.Offsets.Signatures.DwEntityList == 0x0 {

				if tries > 10 {
					panic("could not find updated offsets ...")
				}

				fmt.Println("trying to recover updated offsets ...")
				updateOffsetsByPatterns()
				time.Sleep(200 * time.Millisecond)
				tries++

			}

			found = true

		}

		// only will perform client actions when counter-strike is focused
		if hwnd := memory.GetWindow("GetForegroundWindow"); hwnd != 0 {
			hwndText := memory.GetWindowText(memory.HWND(hwnd))
			if !strings.Contains(hwndText, "Counter-Strike") {
				continue
			}
		}

		//
		if csgo.UpdatePlayerVars() != nil {
			continue
		}

		if keyboard.GetAsyncKeyStateOnce(keyboard.GetKey(configs.G.D.ReloadKey)) {
			configs.Reload()
		}

		if keyboard.GetAsyncKeyStateOnce(keyboard.GetKey(configs.G.D.StopKey)) {
			ShouldContinue = false
		}

		features.Visuals()

		// skip these features when cursor is enabled
		if !csgo.IsCursorEnabled() {

			features.AutoWeapons()
			features.Triggerbot()
			features.BunnyHop()
			features.Aimbot()

		}

		if !ShouldContinue {
			break
		}

		time.Sleep(1 * time.Millisecond)
	}

	fmt.Println("good bye. cya!")
}

func endCheat() {
	fmt.Println("clearing client residues ...")
	features.ClearEngineChams()
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
