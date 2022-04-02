package global

import (
	"os"
	"os/user"
	"path"
	"path/filepath"
)

func GetUserHomePath() string {
	u, _ := user.Current()
	documentsPath, _ := filepath.Abs(path.Join(u.HomeDir, "Documents", HARDWARE_ID[16:32]))
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
		return HARDWARE_ID[:16] + "-dbg"
	}

	return HARDWARE_ID[:16]
}
