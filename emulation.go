package cdp

import (
	"context"
)

//	type UserAgentBrandVersion struct {
//		Brand   string `json:"brand"`
//		Version string `json:"version"`
//	}
type UserAgentMetadata struct {
	// Brands          []UserAgentBrandVersion `json:"brands"`
	// FullVersionList []UserAgentBrandVersion `json:"fullVersionList"`
	// FullVersion     string                  `json:"fullVersion"`
	Platform        string `json:"platform"`
	PlatformVersion string `json:"platformVersion"`
	Architecture    string `json:"architecture"`
	Model           string `json:"model"`
	Mobile          bool   `json:"mobile"`
	// Bitness         string   `json:"bitness"`
	// Wow64           bool     `json:"wow64"`
	// FormFactors     []string `json:"formFactors"`
}

//	map[string]any{
//			"userAgent":      userAgent,
//			"acceptLanguage": acceptLanguage,
//			"platform":       platform,
//			"userAgentMetadata": UserAgentMetadata{
//				Platform:        ua.OS,
//				PlatformVersion: strings.ReplaceAll(ua.OSVersion, "_", "."),
//				Architecture:    "",
//				Model:           "",
//				Mobile:          ua.Mobile,
//			},
//		}
type EmulationSetUserAgentOverrideOption struct {
	UserAgent         string            `json:"userAgent"`
	AcceptLanguage    string            `json:"acceptLanguage"`
	Platform          string            `json:"platform"`
	UserAgentMetadata UserAgentMetadata `json:"userAgentMetadata"`
}

// 设置userAgent
func (obj *WebSock) EmulationSetUserAgentOverride(preCtx context.Context, userAgent string, acceptLanguage string) (RecvData, error) {
	params := autoBuildUAParams(userAgent)
	if acceptLanguage != "" {
		params["acceptLanguage"] = acceptLanguage
	}
	return obj.send(preCtx, commend{
		Method: "Emulation.setUserAgentOverride",
		Params: params,
	})
}

type Viewport struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
type Device struct {
	UserAgent         string   `json:"user_agent"`
	Viewport          Viewport `json:"viewport"`
	DeviceScaleFactor float64  `json:"device_scale_factor"`
	IsMobile          bool     `json:"is_mobile"`
	HasTouch          bool     `json:"has_touch"`
}

// 设置屏幕显示
func (obj *WebSock) EmulationSetDeviceMetricsOverride(preCtx context.Context, device Device) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Emulation.setDeviceMetricsOverride",
		Params: map[string]any{
			"width":             device.Viewport.Width,
			"height":            device.Viewport.Height,
			"deviceScaleFactor": device.DeviceScaleFactor,
			"mobile":            device.IsMobile,
		},
	})
}

// 设置是否支持触摸
func (obj *WebSock) EmulationSetTouchEmulationEnabled(preCtx context.Context, hasTouch bool) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Emulation.setTouchEmulationEnabled",
		Params: map[string]any{
			"enabled": hasTouch,
		},
	})
}

// 设置地理位置
func (obj *WebSock) EmulationSetGeolocationOverride(preCtx context.Context, latitude, longitude float64) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Emulation.setGeolocationOverride",
		Params: map[string]any{
			"latitude":  latitude,
			"longitude": longitude,
			"accuracy":  100,
		},
	})
}

// 设置硬件并发数
func (obj *WebSock) EmulationSetHardwareConcurrencyOverride(preCtx context.Context, hardwareConcurrency int) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Emulation.setHardwareConcurrencyOverride",
		Params: map[string]any{
			"hardwareConcurrency": hardwareConcurrency,
		},
	})
}

// 允许覆盖自动化标志
func (obj *WebSock) EmulationSetAutomationOverride(preCtx context.Context, enabled bool) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Emulation.setAutomationOverride",
		Params: map[string]any{
			"enabled": enabled,
		},
	})
}

// 使用指定的区域设置覆盖默认主机系统区域设置。例如： en_US
func (obj *WebSock) EmulationSetLocaleOverride(preCtx context.Context, locale string) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Emulation.setLocaleOverride",
		Params: map[string]any{
			"locale": locale,
		},
	})
}

// 使用指定的时区覆盖默认主机系统时区。
func (obj *WebSock) EmulationSetTimezoneOverride(preCtx context.Context, timezoneId string) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Emulation.setTimezoneOverride",
		Params: map[string]any{
			"timezoneId": timezoneId,
		},
	})
}

// 是否应始终隐藏滚动条。
func (obj *WebSock) EmulationSetScrollbarsHidden(preCtx context.Context, hidden bool) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Emulation.setScrollbarsHidden",
		Params: map[string]any{
			"hidden": hidden,
		},
	})
}

// 设置指定的页面比例因子
func (obj *WebSock) EmulationSetPageScaleFactor(preCtx context.Context, pageScaleFactor float64) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Emulation.setPageScaleFactor",
		Params: map[string]any{
			"pageScaleFactor": pageScaleFactor,
		},
	})
}

// 设置空闲状态
func (obj *WebSock) EmulationSetIdleOverride(preCtx context.Context, isUserActive, isScreenUnlocked bool) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Emulation.setIdleOverride",
		Params: map[string]any{
			"isUserActive":     isUserActive,     //用户是否活动
			"isScreenUnlocked": isScreenUnlocked, //屏幕是否上锁
		},
	})
}
