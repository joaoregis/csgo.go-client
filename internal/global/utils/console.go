package utils

import (
	"os/exec"
)

//Set windows console title
func SetConsoleTitle(title string) error {

	return exec.Command("cmd", "/C", "title", title).Run()

}
