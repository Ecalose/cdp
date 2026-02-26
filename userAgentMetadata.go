package cdp

import (
	"fmt"
	"strconv"
)

func getUserAgentMetadata(major int, fullVersion, osVersion string) map[string]any {
	majorStr := strconv.Itoa(major)
	versioin := fmt.Sprintf("%d.%s", major, fullVersion)
	return map[string]any{
		"brands": []map[string]string{
			{
				"brand":   "Google Chrome",
				"version": majorStr,
			},
			{
				"brand":   "Chromium",
				"version": majorStr,
			},
			{
				"brand":   "Not A(Brand",
				"version": "24",
			},
		},
		"fullVersionList": []map[string]string{
			{
				"brand":   "Google Chrome",
				"version": versioin,
			},
			{
				"brand":   "Chromium",
				"version": versioin,
			},
			{
				"brand":   "Not A(Brand",
				"version": "24.0.0.0",
			},
		},
		"platform":        "macOS",
		"platformVersion": osVersion,
		"architecture":    "arm",
		"bitness":         "64",
		"uaFullVersion":   versioin,
		"model":           "",
		"mobile":          false,
		"wow64":           false,
		"formFactors":     []string{"Desktop"},
	}
}
func autoBuildUAParams(userAgent string, major int, acceptLanguage string, fullVersion string, osVersion string) map[string]any {
	params := map[string]any{
		"userAgent":         userAgent,
		"platform":          "MacIntel",
		"userAgentMetadata": getUserAgentMetadata(major, fullVersion, osVersion),
	}
	if acceptLanguage != "" {
		params["acceptLanguage"] = acceptLanguage
	}
	return params
}
