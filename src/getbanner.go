package src

import "fmt"

// GetBannerFile returns the banner file path based on the style argument
func GetBannerFile(style string) (string, error) {
	switch style {
	case "standard":
		return "banners/standard.txt", nil
	case "shadow":
		return "banners/shadow.txt", nil
	case "thinkertoy":
		return "banners/thinkertoy.txt", nil
	case "train":
		return "banners/train.txt", nil
	default:
		return "", fmt.Errorf("unknown style: %s. Available styles: standard, shadow, thinkertoy, train", style)
	}
}
