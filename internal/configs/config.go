package configs

import (
	"encoding/json"
	"fmt"
	"gosource/internal/offsets"
	"log"
	"os"
)

type Config struct {
	ReloadKey string `json:"reloadKey"`
	StopKey   string `json:"stopKey"`
	Radar     struct {
		Enabled bool `json:"enabled"`
	} `json:"radar"`
	EngineChams struct {
		Enabled bool `json:"enabled"`
	} `json:"engineChams"`
	Glow struct {
		Enabled       bool    `json:"enabled"`
		BaseColor     string  `json:"glowBaseColor"`
		Alpha         float32 `json:"glowAlpha"`
		IsHealthBased bool    `json:"healthBased"`
	} `json:"glow"`
	Bunnyhop struct {
		Enabled bool   `json:"enabled"`
		Key     string `json:"key"`
	} `json:"bhop"`
	Triggerbot struct {
		Enabled bool   `json:"enabled"`
		Key     string `json:"key"`
	} `json:"trigger"`
	Aimbot struct {
		Enabled bool    `json:"enabled"`
		Key     string  `json:"key"`
		Fov     float64 `json:"fov"`
		Smooth  float64 `json:"smooth"`
	} `json:"aimbot"`
}

func Read() error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return err
	}

	file, err := os.Open(dir + "/configs/config.json")
	if err != nil {
		log.Fatal(err)
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&G)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("config loaded successfully.")
	return nil

}

var G Config
var Offsets offsets.GameOffset
