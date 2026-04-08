package cdp

import "context"

func (obj *WebSock) IndexedDBDeleteDatabase(preCtx context.Context, securityOrigin, dbName string) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "IndexedDB.deleteDatabase",
		Params: map[string]any{
			"securityOrigin": securityOrigin,
			"databaseName":   dbName,
		},
	})
}

func (obj *WebSock) IndexedDBRequestDatabaseNames(preCtx context.Context, securityOrigin string) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "IndexedDB.requestDatabaseNames",
		Params: map[string]any{
			"securityOrigin": securityOrigin,
		},
	})
}
