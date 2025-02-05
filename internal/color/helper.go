package color

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrFlagCollision = errors.New("flags cannot be set at the sametime")
)

var colorOptions = []string{"help", "set"}

// Check if the option is valid
func isValidOptions(option string) bool {
	for _, opt := range colorOptions {
		if option == opt {
			return true
		}
	}
	return false
}

// Convert HEX to RGB
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

// Generate ANSI escape sequence for 24-bit True Color
func rgbToAnsiTrueColor(r, g, b int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%d", r, g, b)
}

// Validate HEX color
func isValidHexColor(hex string) bool {
	match, _ := regexp.MatchString(`^#([a-fA-F0-9]{6})$`, hex)
	return match
}

// Process and set a new color from HEX
func setColor(name, hex string) error {
	if len(name) > 32 {
		return fmt.Errorf("color name cannot exceed 32 characters: %v", len(name))
	}

	if !isValidHexColor(hex) {
		return fmt.Errorf("%v is not a valid hex color", hex)
	}

	r, g, b, err := hexToRGB(hex)
	if err != nil {
		return err
	}

	Colors[name] = rgbToAnsiTrueColor(r, g, b)
	return nil
}

// Get color escape code by name
func getColor(name string) (string, error) {
	color, exists := Colors[name]
	if !exists {
		return "", fmt.Errorf("color %v does not exist", name)
	}
	return color, nil
}
