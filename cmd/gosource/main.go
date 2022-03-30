package main

import (
	"fmt"
	"gosource/internal/configs"
	"gosource/internal/csgo"
	"gosource/internal/features"
	"gosource/internal/hackFunctions/keyboard"
	"gosource/internal/memory"
	"gosource/internal/offsets"
	"os"
	"os/signal"
	"time"
)

var ShouldContinue = true

func main() {

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	go func() {
		select {
		case sig := <-c:
			fmt.Println("operation aborted by user. closing client.", sig)
			endCheat()
			ShouldContinue = false
			os.Exit(1)
		}
	}()

	fmt.Println("getting latest offsets for current game version")
	offsets.InitOffsets(&configs.Offsets)

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

		go checkCsgoProcess(&ShouldContinue)

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

func endCheat() {
	fmt.Println("clearing client residues ...")
	features.ClearEngineChams()
}

func checkCsgoProcess(sc *bool) {
	if !memory.Init() {
		fmt.Println("game has been closed.")
		endCheat()
		*sc = false
	}
}
