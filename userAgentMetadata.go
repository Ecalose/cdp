package cdp

import (
	"strings"

	"github.com/mileusna/useragent"
)

func autoBuildUAParams(userAgent string) map[string]any {
	ua := useragent.Parse(userAgent)
	platform := detectPlatform(ua.OS)
	return map[string]any{
		"userAgent": userAgent,
		"platform":  platform,
		"userAgentMetadata": UserAgentMetadata{
			Platform:        ua.OS,
			PlatformVersion: strings.ReplaceAll(ua.OSVersion, "_", "."),
			Architecture:    "",
			Model:           "",
			Mobile:          ua.Mobile,
		},
	}
}

func detectPlatform(ua string) string {
	ua = strings.ToLower(ua)
	switch {
	case strings.Contains(ua, "mac"):
		return "MacIntel"
	case strings.Contains(ua, "windows"):
		return "Win32"
	case strings.Contains(ua, "linux"):
		return "Linux x86_64"
	case strings.Contains(ua, "iphone"):
		return "iPhone"
	case strings.Contains(ua, "ipad"):
		return "iPad"
	case strings.Contains(ua, "android"):
		return "Linux armv8l"
	default:
		return "Unknown"
	}
}
