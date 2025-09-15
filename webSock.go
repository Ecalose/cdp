package cdp

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gospider007/gson"
	"github.com/gospider007/requests"
	"github.com/gospider007/websocket"
)

type commend struct {
	Id        int64          `json:"id"`
	Method    string         `json:"method"`
	Params    map[string]any `json:"params,omitempty"`
	SessionId string         `json:"sessionId,omitempty"`
}
type event struct {
	// Ctx      context.Context
	// Cnl      context.CancelFunc
	RecvData chan RecvData
}
type RecvData struct {
	Id     int64          `json:"id"`
	Method string         `json:"method"`
	Params map[string]any `json:"params"`
	Result map[string]any `json:"result"`
	Error  map[string]any `json:"error"`
}

type WebSock struct {
	ws       string
	err      error
	option   requests.RequestOption
	conn     *websocket.Conn
	ctx      context.Context
	cnl      context.CancelCauseFunc
	id       atomic.Int64
	reqCli   *requests.Client
	ids      sync.Map
	onEvents sync.Map
}

type RouteData struct {
	RequestId    string      `json:"requestId"`
	Request      RequestData `json:"request"`
	FrameId      string      `json:"frameId"`
	NetworkId    string      `json:"networkId"`
	ResourceType string      `json:"resourceType"`

	ResponseErrorReason string   `json:"responseErrorReason"`
	ResponseStatusCode  int      `json:"responseStatusCode"`
	ResponseStatusText  string   `json:"responseStatusText"`
	ResponseHeaders     []Header `json:"responseHeaders"`
}

func (obj *WebSock) Done() <-chan struct{} {
	return obj.ctx.Done()
}

func (obj *WebSock) Context() context.Context {
	return obj.ctx
}

func (obj *WebSock) recv(ctx context.Context, rd RecvData) error {
	defer recover()
	cmdDataAny, ok := obj.ids.LoadAndDelete(rd.Id)
	if ok {
		cmdData := cmdDataAny.(*event)
		select {
		case <-obj.Done():
			return errors.New("websocks closed")
		case <-ctx.Done():
			return context.Cause(ctx)
		case cmdData.RecvData <- rd:
		}
	}
	methodFuncAny, ok := obj.onEvents.Load(rd.Method)
	if ok && methodFuncAny != nil {
		if fun, funok := methodFuncAny.(func(ctx context.Context, rd RecvData)); funok {
			fun(ctx, rd)
		}
	}
	return nil
}
func (obj *WebSock) recvMain() (err error) {
	defer func() {
		obj.err = err
		obj.CloseWithError(err)
	}()
	for {
		select {
		case <-obj.ctx.Done():
			return context.Cause(obj.ctx)
		default:
			msgType, con, err := obj.conn.ReadMessage()
			if err != nil {
				return err
			}
			switch msgType {
			case websocket.TextMessage:
				rd := RecvData{}
				if _, err = gson.Decode(con, &rd); err == nil {
					if rd.Id == 0 {
						rd.Id = obj.id.Add(1)
					}
					go obj.recv(obj.ctx, rd)
				}
			case websocket.PingMessage:
				obj.conn.WriteMessage(websocket.PongMessage, con)
			case websocket.CloseMessage:
				return errors.New("websock closed")
			default:
				return errors.New("websock unknown message type")
			}
		}
	}
}

func NewWebSock(preCtx context.Context, globalReqCli *requests.Client, ws string, option requests.RequestOption) (*WebSock, error) {
	response, err := globalReqCli.Request(preCtx, "get", ws, requests.RequestOption{DisProxy: true})
	if err != nil {
		return nil, err
	}
	conn := response.WebSocket()
	if conn == nil {
		return nil, errors.New("new websock error")
	}
	cli := &WebSock{
		ws:     ws,
		conn:   response.WebSocket(),
		reqCli: globalReqCli,
		option: option,
	}
	cli.ctx, cli.cnl = context.WithCancelCause(preCtx)
	go cli.recvMain()
	return cli, err
}
func (obj *WebSock) AddEvent(method string, fun func(ctx context.Context, rd RecvData)) {
	obj.onEvents.Store(method, fun)
}
func (obj *WebSock) DelEvent(method string) {
	obj.onEvents.Delete(method)
}
func (obj *WebSock) CloseWithError(err error) {
	obj.cnl(err)
	obj.conn.Close()
}
func (obj *WebSock) Error() error {
	return obj.err
}

func (obj *WebSock) regId(preCtx context.Context, ids ...int64) *event {
	data := new(event)
	data.RecvData = make(chan RecvData)
	for _, id := range ids {
		obj.ids.Store(id, data)
	}
	return data
}
func (obj *WebSock) send(preCtx context.Context, cmd commend) (RecvData, error) {
	var cnl context.CancelFunc
	var ctx context.Context
	if preCtx == nil {
		ctx, cnl = context.WithTimeout(obj.ctx, time.Second*60)
	} else {
		ctx, cnl = context.WithTimeout(preCtx, time.Second*60)
	}
	defer cnl()
	cmd.Id = obj.id.Add(1)
	idEvent := obj.regId(ctx, cmd.Id)
	if err := obj.conn.WriteMessage(websocket.TextMessage, cmd); err != nil {
		obj.CloseWithError(err)
		return RecvData{}, err
	}
	select {
	case <-obj.Done():
		err := obj.Error()
		if err == nil {
			err = context.Cause(obj.ctx)
		}
		obj.CloseWithError(err)
		return RecvData{}, err
	case <-ctx.Done():
		err := obj.Error()
		if err == nil {
			err = context.Cause(ctx)
		}
		obj.CloseWithError(err)
		return RecvData{}, err
	case idRecvData := <-idEvent.RecvData:
		if idRecvData.Error != nil {
			return idRecvData, fmt.Errorf("websock error: %v", idRecvData.Error)
		}
		return idRecvData, nil
	}
}
