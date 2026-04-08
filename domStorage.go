package cdp

import (
	"context"
	_ "embed"
)

func (obj *WebSock) SetDOMStorageItem(preCtx context.Context, storageKey, key, value string, isLocalStorage bool) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "DOMStorage.setDOMStorageItem",
		Params: map[string]any{
			"storageId": map[string]any{
				"storageKey":     storageKey,
				"isLocalStorage": isLocalStorage,
			},
			"key":   key,
			"value": value,
		},
	})
}

func (obj *WebSock) DOMStorageClear(preCtx context.Context, storageKey string, isLocalStorage bool) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "DOMStorage.clear",
		Params: map[string]any{
			"storageId": map[string]any{
				"storageKey":     storageKey,
				"isLocalStorage": isLocalStorage,
			},
		},
	})
}
