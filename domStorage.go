package cdp

import (
	"context"
	_ "embed"
	"log"
)

func (obj *WebSock) SetDOMStorageItem(preCtx context.Context, storageKey, key, value string, isLocalStorage bool) (RecvData, error) {
	log.Print(storageKey)
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
