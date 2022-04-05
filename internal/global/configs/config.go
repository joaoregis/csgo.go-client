package configs

import (
	"encoding/json"
	"errors"
	"fmt"
	"gosource/internal/csgo/offsets"
	"gosource/internal/global"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type Config struct {
	Version string     `json:"version"`
	D       ConfigData `json:"data"`
}

func Init() {

	if !global.IsConfigExists() {
		write()
	}

	read()
	fmt.Println("config initialized successfully.")
}

func Reload() {
	read()
	fmt.Println("config reloaded successfully.")
}

func getDirPath() string {
	return global.USER_HOME_PATH
}

func getFilePath() string {
	return path.Join(getDirPath(), global.CONFIG_NAME)
}

func write() error {

	path := getFilePath()
	var err error
	var file *os.File
	var j []byte
	if !global.IsConfigExists() {

		G = defaultConfig()

		if global.DEBUG_MODE {
			j, _ = json.MarshalIndent(G, "", "	")
		} else {
			j, _ = json.Marshal(G)
		}

		if !global.DEBUG_MODE {
			j = []byte(global.CONFIG_NAME_WITHOUT_EXT + ":" + global.Encrypt(string(j), global.APP_HASH_ENC_KEY))
		}

		err = os.WriteFile(path, j, os.ModeAppend)
		if err != nil {
			fmt.Println("write err 0")
			return err
		}

		return nil

	}

	file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println("write err 1")
		return err
	}

	defer file.Close()

	if global.DEBUG_MODE {
		j, _ = json.MarshalIndent(G, "", "	")
	} else {
		j, _ = json.Marshal(G)
	}

	if !global.DEBUG_MODE {
		j = []byte(global.CONFIG_NAME_WITHOUT_EXT + ":" + global.Encrypt(string(j), global.APP_HASH_ENC_KEY))
	}

	content := string(j)
	if global.DEBUG_MODE {
		fmt.Printf("%s\n", content)
	}

	file.WriteString(content)

	return nil

}

func read() error {

	var err error

	path := getFilePath()
	file, _ := os.Open(path)
	j, _ := ioutil.ReadAll(file)

	if !global.DEBUG_MODE {

		enc_check := strings.Split(string(j), ":")
		if len(enc_check) != 2 || enc_check[0] != global.CONFIG_NAME_WITHOUT_EXT {
			fmt.Println("cfg is not properly encrypted. delete it and regenerate a new one.")
			return errors.New("cfg encryption issue")
		}

		j = []byte(global.Decrypt(string(enc_check[1]), global.APP_HASH_ENC_KEY))

	}

	var dummy map[string]interface{}
	if err = json.Unmarshal(j, &dummy); err != nil {
		fmt.Println("read err 1")
		log.Fatal(err)
		return err
	}

	fmt.Printf("detected config version: %s | current config version: %s \n", dummy["version"], global.CONFIG_VERSION)
	if dummy["version"] != global.CONFIG_VERSION {
		// version has changed, need to be updated
		err = json.Unmarshal(j, &G)
		if err != nil {
			fmt.Printf("cannot recover data. config will be regenerated.\n")
			G = defaultConfig()
			write()
		} else {
			// read successfully
			fmt.Printf("config updated successfully.\n")
			G.Version = global.CONFIG_VERSION

			/* New features need to be defined here. Theres nothing to do about that */
			G.D.ESP = newConfigEntry_defaultValues_ESP()

			write()
		}
	}

	err = json.Unmarshal(j, &G)
	if err != nil {
		fmt.Println("read err 2")
		log.Fatal(err)
		return err
	}

	return nil

}

func defaultConfig() Config {
	return Config{
		Version: global.CONFIG_VERSION,
		D: ConfigData{
			ReloadKey:   "Insert",
			StopKey:     "Delete",
			Radar:       false,
			EngineChams: false,
			Bunnyhop:    false,
			Glow: ConfigDataGlow{
				Enabled:       true,
				BaseColor:     "#6821a6",
				Alpha:         0.6,
				IsHealthBased: false,
			},
			Triggerbot: ConfigDataTrigger{
				Enabled: false,
				Key:     "Mouse 5",
				Delay:   50,
			},
			AutoWeapons: ConfigDataAutoWeapons{
				Enabled: false,
				Delay:   25,
			},
			Aimbot: ConfigDataAimbot{
				Enabled: false,
				Key:     "Mouse 4",
				Fov:     5.0,
				Smooth:  10.0,
			},
			ESP: newConfigEntry_defaultValues_ESP(),
		},
	}
}

func newConfigEntry_defaultValues_ESP() ConfigDataESP {
	return ConfigDataESP{
		Enabled:          true,
		AllyBoundingBox:  newConfigEntry_defaultValues_ESP_BoundingBox(false),
		EnemyBoundingBox: newConfigEntry_defaultValues_ESP_BoundingBox(true),
		DrawSnapLines:    false,
	}
}

func newConfigEntry_defaultValues_ESP_BoundingBox(e bool) ConfigDataESPBoundingBox {
	return ConfigDataESPBoundingBox{
		Enabled:               e,
		DrawBox:               true,
		Layout:                0,
		Outline:               false,
		OutlineColor:          "#000000",
		Color:                 "#6821a6",
		ColorAlpha:            .5,
		IsColorHealthBasedBox: false,
		FullfillBox:           true,
		FullfillBoxColor:      "#333333",
		FullfillBoxColorAlpha: 0.4,
		DrawName:              true,
		DrawHealth:            true,
		Thickness:             4,
	}
}

var G Config
var Offsets offsets.GameOffset
