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
