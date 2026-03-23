package cdp

type Brand struct {
	Brand   string `json:"brand"`
	Version string `json:"version"`
}
type UserAgentData struct {
	Brands          []Brand  `json:"brands"`
	FullVersionList []Brand  `json:"fullVersionList"`
	Platform        string   `json:"platform"`
	PlatformVersion string   `json:"platformVersion"`
	Architecture    string   `json:"architecture"`
	Bitness         string   `json:"bitness"`
	UaFullVersion   string   `json:"uaFullVersion"`
	Model           string   `json:"model"`
	Mobile          bool     `json:"mobile"`
	Wow64           bool     `json:"wow64"`
	FormFactors     []string `json:"formFactors"`
}
