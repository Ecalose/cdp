package cdp

var userAgentMetadata = map[string]any{
	"brands": []map[string]string{
		{
			"brand":   "Chromium",
			"version": "143",
		},
		{
			"brand":   "Not A(Brand",
			"version": "24",
		},
	},
	"fullVersionList": []map[string]string{},
	"platform":        "macOS",
	"platformVersion": "",
	"architecture":    "",
	"bitness":         "",
	"uaFullVersion":   "",
	"model":           "",
	"mobile":          false,
	"wow64":           false,
	"formFactors":     []string{},
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
