package cdp

var userAgentMetadata = map[string]any{
	"brands": []map[string]string{
		{
			"brand":   "Google Chrome",
			"version": "143",
		},
		{
			"brand":   "Chromium",
			"version": "143",
		},
		{
			"brand":   "Not A(Brand",
			"version": "24",
		},
	},
	"fullVersionList": []map[string]string{
		{
			"brand":   "Google Chrome",
			"version": "143.0.7499.193",
		},
		{
			"brand":   "Chromium",
			"version": "143.0.7499.193",
		},
		{
			"brand":   "Not A(Brand",
			"version": "24.0.0.0",
		},
	},
	"platform":        "macOS",
	"platformVersion": "26.1.0",
	"architecture":    "arm",
	"bitness":         "64",
	"uaFullVersion":   "143.0.7499.193",
	"model":           "",
	"mobile":          false,
	"wow64":           false,
	"formFactors":     []string{"Desktop"},
}

// await navigator.userAgentData.getHighEntropyValues([
//
//	'brands',
//	'fullVersionList',
//	'platform',
//	'platformVersion',
//	'architecture',
//	'bitness',
//	'uaFullVersion',
//	'model',
//	'mobile',
//	'wow64',
//	'formFactors',
//
// ]);
func autoBuildUAParams(userAgent string, acceptLanguage string) map[string]any {
	params := map[string]any{
		"userAgent":         userAgent,
		"platform":          "MacIntel",
		"userAgentMetadata": userAgentMetadata,
	}
	if acceptLanguage != "" {
		params["acceptLanguage"] = acceptLanguage
	}
	return params
}
