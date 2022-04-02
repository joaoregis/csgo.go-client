package global

import (
	"crypto/sha256"
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
)

func GetUserHomePath() string {
	u, _ := user.Current()
	clientHwidHash := fmt.Sprintf("%x", sha256.Sum256([]byte(HARDWARE_ID)))
	documentsPath, _ := filepath.Abs(path.Join(u.HomeDir, "Documents", clientHwidHash[16:32]))
	if _, err := os.Stat(documentsPath); os.IsNotExist(err) {
		// creating default client directory if not exists yet
		os.Mkdir(documentsPath, os.ModeDir)
	}
	return documentsPath
}

func IsConfigExists() bool {
	documentsPath, _ := filepath.Abs(path.Join(USER_HOME_PATH, CONFIG_NAME))
	if _, err := os.Stat(documentsPath); os.IsNotExist(err) {
		return false
	}
	return true
}

func GetConfigName() string {

	if DEBUG_MODE {
		clientHwidHash := fmt.Sprintf("%x", sha256.Sum256([]byte(HARDWARE_ID+"-dbg")))
		return clientHwidHash[:16]
	}

	clientHwidHash := fmt.Sprintf("%x", sha256.Sum256([]byte(HARDWARE_ID)))
	return clientHwidHash[:16]
}
