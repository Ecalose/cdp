package cdp

import (
	"context"
)

func (obj *WebSock) RuntimeEvaluate(ctx context.Context, expression string) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "Runtime.evaluate",
		Params: map[string]any{
			// "disableBreaks":               true,       //执行期间禁用断点
			"disableBreaks":               true,       //执行期间禁用断点
			"expression":                  expression, //表达式
			"awaitPromise":                true,       //异步函数
			"returnByValue":               true,
			"silent":                      true,
			"allowUnsafeEvalBlockedByCSP": true,
			// "includeCommandLineAPI":       true,
			"includeCommandLineAPI": false,
		},
	})
}
func (obj *WebSock) RuntimeEnable(ctx context.Context) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "Runtime.enable",
	})
}
