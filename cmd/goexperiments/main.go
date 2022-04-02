package main

import (
	"fmt"
	"gosource/internal/global"
	"os/user"
	"path"
	"path/filepath"
)

func main() {

	u, _ := user.Current()
	documentsPath, _ := filepath.Abs(path.Join(u.HomeDir, "Documents", global.CONFIG_NAME))
	fmt.Println(documentsPath)

}
