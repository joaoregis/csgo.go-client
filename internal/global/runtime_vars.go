package global

import (
	"github.com/joaoregis/machineid"
)

var (
	HARDWARE_ID, _          = machineid.ProtectedID(APP_UUID_4_GO)
	CONFIG_NAME             = GetConfigName() + ".cfg"
	CONFIG_NAME_WITHOUT_EXT = GetConfigName()
	USER_HOME_PATH          = GetUserHomePath()
)