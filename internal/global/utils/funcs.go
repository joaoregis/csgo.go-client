package utils

import (
	"crypto/sha256"
	"fmt"
	"gosource/internal/global"
)

type cointainable interface {
	string | int | int32 | int64 | float64 | float32
}

func Contains[T cointainable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetProtectHwid() string {

	bb := sha256.Sum256([]byte(global.HARDWARE_ID))
	return fmt.Sprintf("%x", bb[:])

}
