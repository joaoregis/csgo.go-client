package color

import (
	"fmt"
	"strconv"
	str "strings"
)

func ConvertHexToRGB(hex Hex) (*RGB, error) {
	var rgb RGB

	if str.Contains(string(hex), "#") {
		hex = Hex(str.Split(string(hex), "#")[1])
	}

	values, err := strconv.ParseUint(string(hex), 16, 32)

	if err != nil {
		return &RGB{}, err
	}

	rgb = RGB{
		Red:   int(uint8(values >> 16)),
		Green: int(uint8((values >> 8) & 0xFF)),
		Blue:  int(uint8(values & 0xFF)),
	}

	return &rgb, nil
}

func ConvertRGBToFloat(rgb RGB, alpha *float32) *RGBA {
	result := &RGBA{Red: float32(rgb.Red) / 255.0, Green: float32(rgb.Green) / 255.0, Blue: float32(rgb.Blue) / 255.0, Alpha: 1}
	if alpha != nil {
		result.Alpha = *alpha
	}

	return result
}

func HexToRGBA(hex Hex, alpha *float32) *RGBA {
	rgb, err := ConvertHexToRGB(hex)

	if err != nil {
		fmt.Println("Error on convert hex to rgb")
		return &RGBA{}
	}

	return ConvertRGBToFloat(*rgb, alpha)
}
