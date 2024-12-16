package cdp

import "context"

func (obj *WebSock) StorageClear(preCtx context.Context, href string) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Storage.clearDataForOrigin",
		Params: map[string]any{
			"origin":       href,
			"storageTypes": "all",
		},
	})
}
func (obj *WebSock) StorageEnable(preCtx context.Context, storageKey string) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Storage.setStorageBucketTracking",
		Params: map[string]any{
			"storageKey": storageKey,
			"enable":     true,
		},
	})
}

func (obj *WebSock) StorageGetStorageKeyForFrame(preCtx context.Context, frameId string) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Storage.getStorageKeyForFrame",
		Params: map[string]any{
			"frameId": frameId,
		},
	})
}
