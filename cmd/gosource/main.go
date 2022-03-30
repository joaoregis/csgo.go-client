package main

import (
	"fmt"
	"gosource/internal/configs"
	"gosource/internal/csgo"
	"gosource/internal/features"
	"gosource/internal/hackFunctions/keyboard"
	"gosource/internal/memory"
	"gosource/internal/offsets"
	"time"
)

var ShouldContinue = true

func main() {

	fmt.Println("getting latest offsets for current game version")
	offsets.InitOffsets(&configs.Offsets)

	fmt.Println("checking client state: ", configs.Offsets.Signatures.DwClientState)
	if configs.Offsets.Signatures.DwClientState > 0 {
		fmt.Println("client state is ok. proceeding.")
	}

	fmt.Println("initializing key functions")
	keyboard.InitKeys()

	for !memory.Init() {
		fmt.Println("awaiting for game process to proceed ...")
		time.Sleep(1000 * time.Millisecond)
	}

	fmt.Println("reading user configs")
	configs.Read()

	// Mainloop
	fmt.Println("everything is fine. good hacking.")
	for ShouldContinue {
		if csgo.UpdatePlayerVars() != nil {
			continue
		}

		if keyboard.GetAsyncKeyStateOnce(keyboard.GetKey(configs.G.ReloadKey)) {
			configs.Read()
		}

		if keyboard.GetAsyncKeyStateOnce(keyboard.GetKey(configs.G.StopKey)) {
			ShouldContinue = false
		}

		features.Triggerbot()
		features.BunnyHop()
		features.Visuals()

		go features.Aimbot()

		if !ShouldContinue {
			break
		}

		time.Sleep(1 * time.Millisecond)
	}

	fmt.Println("good bye. cya!")
}
