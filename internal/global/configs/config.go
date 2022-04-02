package configs

import (
	"encoding/json"
	"errors"
	"fmt"
	"gosource/internal/global"
	"gosource/internal/offsets"
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

func write() error {

	path := path.Join(global.USER_HOME_PATH, global.CONFIG_NAME)
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

	file, err = os.Open(path)
	if err != nil {
		fmt.Println("write err 1")
		return err
	}

	if global.DEBUG_MODE {
		j, _ = json.MarshalIndent(G, "", "	")
	} else {
		j, _ = json.Marshal(G)
	}

	if !global.DEBUG_MODE {
		j = []byte(global.CONFIG_NAME_WITHOUT_EXT + ":" + global.Encrypt(string(j), global.APP_HASH_ENC_KEY))
	}

	file.Write(j)

	return nil

}

func read() error {

	var err error

	path := path.Join(global.USER_HOME_PATH, global.CONFIG_NAME)
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

	if dummy["version"] != global.CONFIG_VERSION {
		// version has changed, need to be updated
		err = json.Unmarshal(j, &G)
		if err != nil {
			// cannot recover data :( [possibly too old config]
			G = defaultConfig()
			write()
		} else {
			// read successfully
			G.Version = global.CONFIG_VERSION
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
				Enabled:       false,
				BaseColor:     "#F000FF",
				Alpha:         0.7,
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
		},
	}
}

var G Config
var Offsets offsets.GameOffset
