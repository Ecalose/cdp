package cdp

import "context"

func (obj *WebSock) TargetCreateTarget(ctx context.Context, url string) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "Target.createTarget",
		Params: map[string]any{
			"url": url,
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
	return obj.send(obj.ctx, commend{
		Method: "Target.setAutoAttach",
		Params: map[string]any{
			"autoAttach":             true,
			"flatten":                true,
			"waitForDebuggerOnStart": true,
			"filter": []any{
				map[string]any{"type": "iframe", "exclude": false},
			},
		},
	})
}
