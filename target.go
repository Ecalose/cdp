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

func (obj *WebSock) TargetCreateBrowserContext(ctx context.Context) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "Target.createBrowserContext",
		Params: map[string]any{
			"disposeOnDetach": true,
		},
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
