package color

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	FlagCollisionError = errors.New("flags Cannot be set at the same time.")
)

var colorOptions = []string{"help", "set"}

func isValidOptions(option string) bool {
	for _, opt := range colorOptions {
		if option == opt {
			return true
		}
	}
	return false
}

// Converts hex to RGB
func hexToRGB(hex string) (int, int, int, error) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return 0, 0, 0, fmt.Errorf("invalid hex color")
	}
	r, err := strconv.ParseInt(hex[0:2], 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}
	g, err := strconv.ParseInt(hex[2:4], 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}
	b, err := strconv.ParseInt(hex[4:6], 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}
	return int(r), int(g), int(b), nil
}

// Maps RGB to the closest ANSI 256 color
func rgbToAnsi256(r, g, b int) int {
	// Grayscale (232-255)
	if r>>4 == g>>4 && g>>4 == b>>4 {
		grayscale := (r + g + b) / 3
		if grayscale < 8 {
			return 16
		}
		if grayscale > 238 {
			return 231
		}
		return 232 + (grayscale-8)/10
	}

	// 6x6x6 Cube (16-231)
	rIndex := (r * 5) / 255
	gIndex := (g * 5) / 255
	bIndex := (b * 5) / 255
	return 16 + (rIndex*36 + gIndex*6 + bIndex)
}

// Validate hex color using regex
func isValidHexColor(hex string) bool {
	// Matches #RRGGBB or #RGB
	match, _ := regexp.MatchString(`^#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$`, hex)
	return match
}
