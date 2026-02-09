package cdp

import (
	"context"
)

func (obj *WebSock) TargetCreateTarget(ctx context.Context, browserContextId string, url string) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "Target.createTarget",
		Params: map[string]any{
			"browserContextId": browserContextId,
			"url":              url,
			"transitionType":   "address_bar",
		},
	})
}
func (obj *WebSock) TargetCloseTarget(targetId string) (RecvData, error) {
	return obj.send(obj.ctx, commend{
		Method: "Target.closeTarget",
		Params: map[string]any{
			"targetId": targetId,
		},
	})
}
func (obj *WebSock) TargetDetachFromTarget(sessionId string) (RecvData, error) {
	return obj.send(obj.ctx, commend{
		Method: "Target.detachFromTarget",
		Params: map[string]any{
			"sessionId": sessionId,
		},
	})
}
func (obj *WebSock) TargetSetAutoAttach(ctx context.Context) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "Target.setAutoAttach",
		Params: map[string]any{
			"autoAttach":             true,
			"flatten":                true,
			"waitForDebuggerOnStart": true,
		},
	})
}

type TargetFilter struct {
	Type    string `json:"type"`
	Exclude bool   `json:"exclude"`
}

func (obj *WebSock) TargetSetDiscoverTargets(ctx context.Context, discover bool, filters ...TargetFilter) (RecvData, error) {
	params := map[string]any{
		"discover": discover,
	}
	if len(filters) > 0 {
		params["filter"] = filters
	}
	return obj.send(ctx, commend{
		Method: "Target.setDiscoverTargets",
		Params: params,
	})
}

func (obj *WebSock) TargetCreateBrowserContext(ctx context.Context, proxyServer string) (RecvData, error) {
	params := map[string]any{
		"disposeOnDetach": true,
	}
	if proxyServer != "" {
		params["proxyServer"] = proxyServer
	}
	return obj.send(ctx, commend{
		Method: "Target.createBrowserContext",
		Params: params,
	})
}
func (obj *WebSock) TargetDisposeBrowserContext(browserContextId string) (RecvData, error) {
	return obj.send(obj.ctx, commend{
		Method: "Target.disposeBrowserContext",
		Params: map[string]any{
			"browserContextId": browserContextId,
		},
	})
}
func (obj *WebSock) TargetAttachToTarget(targetId string) (RecvData, error) {
	return obj.send(obj.ctx, commend{
		Method: "Target.attachToTarget",
		Params: map[string]any{
			"targetId": targetId,
		},
	})
}
