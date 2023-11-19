package cdp

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

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
	Ctx      context.Context
	Cnl      context.CancelFunc
	RecvData chan RecvData
}
type RecvData struct {
	Id     int64          `json:"id"`
	Method string         `json:"method"`
	Params map[string]any `json:"params"`
	Result map[string]any `json:"result"`
}

type WebSock struct {
	option   WebSockOption
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

func (obj *WebSock) recv(ctx context.Context, rd RecvData) error {
	// log.Print(gson.Decode(rd))
	defer recover()
	cmdDataAny, ok := obj.ids.LoadAndDelete(rd.Id)
	if ok {
		cmdData := cmdDataAny.(*event)
		select {
		case <-obj.Done():
			return errors.New("websocks closed")
		case <-ctx.Done():
			return ctx.Err()
		case <-cmdData.Ctx.Done():
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
	defer obj.Close()
	for {
		select {
		case <-obj.ctx.Done():
			return obj.ctx.Err()
		default:
			rd := RecvData{}
			if err := obj.conn.RecvJson(obj.ctx, &rd); err != nil {
				return err
			}
			if rd.Id == 0 {
				rd.Id = obj.id.Add(1)
			}
			go obj.recv(obj.ctx, rd)
		}
	}
}

type WebSockOption struct {
	Proxy string
}

func NewWebSock(preCtx context.Context, globalReqCli *requests.Client, ws string, option WebSockOption) (*WebSock, error) {
	response, err := globalReqCli.Request(preCtx, "get", ws, requests.RequestOption{DisProxy: true})
	if err != nil {
		return nil, err
	}
	conn := response.WebSocket()
	if conn == nil {
		return nil, errors.New("new websock error")
	}
	conn.SetReadLimit(1024 * 1024 * 1024) //1G
	cli := &WebSock{
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
func (obj *WebSock) Close() {
	obj.conn.Close()
	obj.cnl(nil)
}

func (obj *WebSock) regId(preCtx context.Context, ids ...int64) *event {
	data := new(event)
	data.Ctx, data.Cnl = context.WithCancel(preCtx)
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
	select {
	case <-obj.Done():
		return RecvData{}, context.Cause(obj.ctx)
	case <-ctx.Done():
		return RecvData{}, obj.ctx.Err()
	default:
		cmd.Id = obj.id.Add(1)
		idEvent := obj.regId(ctx, cmd.Id)
		defer idEvent.Cnl()
		if err := obj.conn.SendJson(ctx, cmd); err != nil {
			return RecvData{}, err
		}
		select {
		case <-obj.Done():
			return RecvData{}, context.Cause(obj.ctx)
		case <-ctx.Done():
			return RecvData{}, ctx.Err()
		case idRecvData := <-idEvent.RecvData:
			return idRecvData, nil
		}
	}
}
