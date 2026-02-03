package cdp

import (
	"context"
)

func (obj *WebSock) PageEnable(ctx context.Context) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "Page.enable",
	})
}
func (obj *WebSock) PageGetFrameTree(ctx context.Context) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "Page.getFrameTree",
	})
}
func (obj *WebSock) PageAddScriptToEvaluateOnNewDocument(ctx context.Context, source string) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "Page.addScriptToEvaluateOnNewDocument",
		Params: map[string]any{
			"source":         source,
			"runImmediately": true, //立即运行
		},
	})
}

type ScreenshotOption struct {
	Format                string //图像压缩格式（默认为 webp）,允许的值：jpeg,png,webp
	Quality               int    //范围 [0..100] 的压缩质量（仅限 jpeg）。
	CaptureBeyondViewport bool   //捕获视口之外的屏幕截图。默认为 false。
}

func (obj *WebSock) PageCaptureScreenshot(ctx context.Context, rect Rect, options ...ScreenshotOption) (RecvData, error) {
	var option ScreenshotOption
	if len(options) > 0 {
		option = options[0]
	}
	if option.Format == "" {
		option.Format = "webp"
	}
	params := map[string]any{
		"format":                option.Format,
		"captureBeyondViewport": option.CaptureBeyondViewport,
	}
	if rect.Width > 0 && rect.Height > 0 {
		params["clip"] = map[string]float64{
			"x":      rect.X,
			"y":      rect.Y,
			"width":  rect.Width,
			"height": rect.Height,
			"scale":  1,
		}
	}
	if option.Quality > 0 {
		params["quality"] = option.Quality
	}
	return obj.send(ctx, commend{
		Method: "Page.captureScreenshot",
		Params: params,
	})
}
func (obj *WebSock) PageGetLayoutMetrics(ctx context.Context) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "Page.getLayoutMetrics",
		Params: map[string]any{},
	})
}
func (obj *WebSock) PageReload(ctx context.Context) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "Page.reload",
		Params: map[string]any{},
	})
}

type PageNavigateOption struct {
	Referrer string
}

func (obj *WebSock) PageNavigate(ctx context.Context, url string, options ...PageNavigateOption) (RecvData, error) {
	var option PageNavigateOption
	if len(options) > 0 {
		option = options[0]
	}
	params := map[string]any{
		"url": url,
	}
	if option.Referrer != "" {
		params["referrer"] = option.Referrer
	}
	return obj.send(ctx, commend{
		Method: "Page.navigate",
		Params: params,
	})
}

func (obj *WebSock) PageHandleJavaScriptDialog(ctx context.Context, accept bool, txts ...string) (RecvData, error) {
	params := map[string]any{
		"accept": accept,
	}
	if len(txts) > 0 {
		params["promptText"] = txts[0]
	}
	return obj.send(ctx, commend{
		Method: "Page.handleJavaScriptDialog",
		Params: params,
	})
}
func (obj *WebSock) PageBringToFront(ctx context.Context) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "Page.bringToFront",
	})
}
func (obj *WebSock) PageSetDocumentContent(ctx context.Context, frameId string, html string) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "Page.setDocumentContent",
		Params: map[string]any{
			"frameId": frameId,
			"html":    html,
		},
	})
}
