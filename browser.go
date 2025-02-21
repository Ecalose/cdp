package cdp

import (
	"context"
)

func (obj *WebSock) BrowserClose() error {
	_, err := obj.send(obj.ctx, commend{
		Method: "Browser.close",
	})
	return err
}
func (obj *WebSock) Cdp(ctx context.Context, sessid string, method string, params ...map[string]any) (RecvData, error) {
	comd := commend{
		SessionId: sessid,
		Method:    method,
	}
	if len(params) > 0 {
		comd.Params = params[0]
	}
	return obj.send(ctx, comd)
}

// Browser.PermissionType
// Allowed Values: ar, audioCapture, automaticFullscreen, backgroundFetch, backgroundSync, cameraPanTiltZoom, capturedSurfaceControl, clipboardReadWrite, clipboardSanitizedWrite, displayCapture, durableStorage, geolocation, handTracking, idleDetection, keyboardLock, localFonts, midi, midiSysex, nfc, notifications, paymentHandler, periodicBackgroundSync, pointerLock, protectedMediaIdentifier, sensors, smartCard, speakerSelection, storageAccess, topLevelStorageAccess, videoCapture, vr, wakeLockScreen, wakeLockSystem, webAppInstallation, webPrinting, windowManagement

func (obj *WebSock) BrowserSetPermission(ctx context.Context, name string, setting string, origins ...string) error {
	params := map[string]any{
		"permission": map[string]string{
			"name": name,
		},
		"setting": setting,
	}
	if len(origins) > 0 {
		params["origin"] = origins[0]
	}
	_, err := obj.send(ctx, commend{
		Method: "Browser.setPermission",
		Params: params,
	})
	return err
}

func (obj *WebSock) BrowserGrantPermissions(ctx context.Context, names []string, origins ...string) error {
	permissions := make([]map[string]string, len(names))
	for i, name := range names {
		permissions[i] = map[string]string{
			"name": name,
		}
	}
	params := map[string]any{
		"permissions": permissions,
	}
	if len(origins) > 0 {
		params["origin"] = origins[0]
	}

	_, err := obj.send(ctx, commend{
		Method: "Browser.grantPermissions",
		Params: params,
	})
	return err
}
