package main

import (
	"gosource/internal/csgo/offsets"
	"gosource/internal/global"
	"gosource/internal/global/configs"
	"gosource/internal/global/logs"
	"gosource/internal/global/utils"
	"gosource/internal/hackFunctions/keyboard"
	"gosource/internal/memory"
	"time"
)

func initClientHeader() {

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

	initMemory()

	logs.Info("initializing user configs")
	configs.Init()

	logs.Info("scanning for new game patterns")
	updateOffsetsByPatterns()

}

func initMemory() {

	for !memory.Init() {
		logs.Info("initializing memory module ...")
		time.Sleep(1000 * time.Millisecond)
	}

}

func initGameHwnd() {

	for global.HWND_GAME == 0 {
		global.HWND_GAME = utils.FindWindow("Counter-Strike: Global Offensive - Direct3D 9")
		logs.Info("awaiting for csgo ...")
		time.Sleep(1000 * time.Millisecond)
	}

}

func gracefulExit() {
	logs.Info("clearing client residues ...")
	logs.Info("good bye. cya!")
}

func updateOffsetsByPatterns() {

	var patterns = map[string]string{
		//
		"DwEntityList":        "BB ? ? ? ? 83 FF 01 0F 8C ? ? ? ? 3B F8",
		"DwForceAttack":       "89 0D ? ? ? ? 8B 0D ? ? ? ? 8B F2 8B C1 83 CE 04",
		"DwForceJump":         "8B 0D ? ? ? ? 8B D6 8B C1 83 CA 02",
		"DwLocalPlayer":       "8D 34 85 ? ? ? ? 89 15 ? ? ? ? 8B 41 08 8B 48 04 83 F9 FF",
		"DwGlowObjectManager": "A1 ? ? ? ? A8 01 75 4B",
		//
		"DwViewMatrix":            "0F 10 05 ? ? ? ? 8D 85 ? ? ? ? B9",
		"DwMouseEnable":           "B9 ? ? ? ? FF 50 34 85 C0 75 10",
		"DwMouseEnablePtr":        "B9 ? ? ? ? FF 50 34 85 C0 75 10",
		"DwRadarBase":             "A1 ? ? ? ? 8B 0C B0 8B 01 FF 50 ? 46 3B 35 ? ? ? ? 7C EA 8B 0D",
		"DwClientState":           "A1 ? ? ? ? 33 D2 6A 00 6A 00 33 C9 89 B0",
		"DwClientStateMaxPlayer":  "A1 ? ? ? ? 8B 80 ? ? ? ? C3 CC CC CC CC 55 8B EC 8A 45 08",
		"DwClientStatePlayerInfo": "8B 89 ? ? ? ? 85 C9 0F 84 ? ? ? ? 8B 01",
		"DwClientStateViewAngles": "F3 0F 11 86 ? ? ? ? F3 0F 10 44 24 ? F3 0F 11 86",
		"MBDormant":               "8A 81 ? ? ? ? C3 32 C0",
	}

	//
	configs.Offsets.Signatures.DwEntityList, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwEntityList"], true, 0x1, 0x0)
	configs.Offsets.Signatures.DwForceAttack, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwForceAttack"], true, 0x2, 0x0)
	configs.Offsets.Signatures.DwForceJump, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwForceJump"], true, 0x2, 0x0)
	configs.Offsets.Signatures.DwLocalPlayer, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwLocalPlayer"], true, 0x3, 0x4)
	configs.Offsets.Signatures.DwGlowObjectManager, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwGlowObjectManager"], true, 0x1, 0x4)
	//
	configs.Offsets.Signatures.DwViewMatrix, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwViewMatrix"], true, 0x3, 0xB0)
	configs.Offsets.Signatures.DwMouseEnable, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwMouseEnable"], true, 0x1, 0x30)
	configs.Offsets.Signatures.DwMouseEnablePtr, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwMouseEnablePtr"], true, 0x1, 0x0)
	configs.Offsets.Signatures.DwRadarBase, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["DwRadarBase"], true, 0x1, 0x0)
	configs.Offsets.Signatures.DwClientState, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Engine], patterns["DwClientState"], true, 0x1, 0x0)
	configs.Offsets.Signatures.DwClientStateMaxPlayer, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Engine], patterns["DwClientStateMaxPlayer"], false, 0x7, 0x0)
	configs.Offsets.Signatures.DwClientStatePlayerInfo, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Engine], patterns["DwClientStatePlayerInfo"], false, 0x2, 0x0)
	configs.Offsets.Signatures.DwClientStateViewAngles, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Engine], patterns["DwClientStateViewAngles"], false, 0x4, 0x0)
	configs.Offsets.Signatures.MBDormant, _ = memory.GameProcess.FindOffset(memory.GameProcess.Modules[offsets.Client], patterns["MBDormant"], false, 0x2, 0x8)

}
