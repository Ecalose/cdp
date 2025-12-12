package cdp

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"net/http"

	"github.com/gospider007/gson"
	"github.com/gospider007/netx"
	"github.com/gospider007/re"
	"github.com/gospider007/requests"
	"github.com/gospider007/tools"
)

type RequestOption struct {
	Url      string      `json:"url"`
	Method   string      `json:"method"`
	PostData string      `json:"postData"`
	Headers  http.Header `json:"headers"`
	Proxy    string
}
type RequestData struct {
	Url              string            `json:"url"`
	UrlFragment      string            `json:"urlFragment"`
	Method           string            `json:"method"`
	Headers          map[string]string `json:"headers"`
	PostData         string            `json:"postData"`
	HasPostData      bool              `json:"hasPostData"`
	MixedContentType string            `json:"mixedContentType"`
	InitialPriority  string            `json:"initialPriority"` //初始优先级
	ReferrerPolicy   string            `json:"referrerPolicy"`
	IsLinkPreload    bool              `json:"isLinkPreload"`   //是否通过链路预加载加载。
	PostDataEntries  []DataEntrie      `json:"postDataEntries"` //是否通过链路预加载加载。
}
type DataEntrie struct {
	Bytes string `json:"bytes"`
}
type Route struct {
	webSock  *WebSock
	recvData RouteData
	used     bool
}

func NewRoute(webSock *WebSock, recvData RouteData) *Route {
	return &Route{webSock: webSock, recvData: recvData}
}
func (obj *Route) Used() bool {
	return obj.used
}
func (obj *Route) IsResponse() bool {
	if obj.recvData.ResponseErrorReason != "" ||
		obj.recvData.ResponseStatusCode != 0 || obj.recvData.ResponseStatusText != "" ||
		obj.recvData.ResponseHeaders != nil {
		return true
	}
	return false
}
func (obj *Route) Error() error {
	if obj.recvData.ResponseErrorReason != "" {
		return errors.New(obj.recvData.ResponseErrorReason)
	}
	return nil
}
func (obj *Route) StatusCode() int {
	return obj.recvData.ResponseStatusCode
}
func (obj *Route) StatusText() string {
	return obj.recvData.ResponseStatusText
}
func (obj *Route) ResponseHeaders() http.Header {
	head := http.Header{}
	for _, hd := range obj.recvData.ResponseHeaders {
		head.Add(hd.Name, hd.Value)
	}
	return head
}

func (obj *Route) NewRequestOption() RequestOption {
	return RequestOption{
		Url:      obj.Url(),
		Method:   obj.Method(),
		PostData: obj.PostData(),
		Headers:  obj.Headers(),
	}
}
func (obj *Route) NewFulData(ctx context.Context) (fulData FulData, err error) {
	if !obj.IsResponse() {
		err = errors.New("not response route")
		return
	}
	if err = obj.Error(); err != nil {
		return
	}
	fulData.Body, err = obj.ResponseBody(ctx)
	fulData.StatusCode = obj.StatusCode()
	fulData.Headers = obj.ResponseHeaders()
	fulData.ResponsePhrase = obj.StatusText()
	return
}

type ResourceType string

const (
	ResourceTypeDocument           ResourceType = "Document"
	ResourceTypeStylesheet         ResourceType = "Stylesheet"
	ResourceTypeImage              ResourceType = "Image"
	ResourceTypeMedia              ResourceType = "Media"
	ResourceTypeFont               ResourceType = "Font"
	ResourceTypeScript             ResourceType = "Script"
	ResourceTypeTextTrack          ResourceType = "TextTrack"
	ResourceTypeXHR                ResourceType = "XHR"
	ResourceTypeFetch              ResourceType = "Fetch"
	ResourceTypePrefetch           ResourceType = "Prefetch"
	ResourceTypeEventSource        ResourceType = "EventSource"
	ResourceTypeWebSocket          ResourceType = "WebSocket"
	ResourceTypeManifest           ResourceType = "Manifest"
	ResourceTypeSignedExchange     ResourceType = "SignedExchange"
	ResourceTypePing               ResourceType = "Ping"
	ResourceTypeCSPViolationReport ResourceType = "CSPViolationReport"
	ResourceTypePreflight          ResourceType = "Preflight"
	ResourceTypeOther              ResourceType = "Other"
)

// Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, Prefetch, EventSource, WebSocket, Manifest, SignedExchange, Ping, CSPViolationReport, Preflight, Other
func (obj *Route) ResourceType() ResourceType {
	return obj.recvData.ResourceType
}
func (obj *Route) Url() string {
	return obj.recvData.Request.Url
}
func (obj *Route) Method() string {
	return obj.recvData.Request.Method
}
func (obj *Route) PostData() string {
	return obj.recvData.Request.PostData
}
func (obj *Route) Headers() http.Header {
	delete(obj.recvData.Request.Headers, "If-Modified-Since")
	head := http.Header{}
	for kk, vv := range obj.recvData.Request.Headers {
		head.Add(kk, vv)
	}
	if useragent := head.Get("user-agent"); strings.Contains(useragent, "HeadlessChrome") {
		head.Set("user-agent", re.Sub("HeadlessChrome", "Chrome", useragent))
	}
	return head
}
func (obj *Route) SetHeader(key, val string) {
	obj.recvData.Request.Headers[key] = val
}
func (obj *Route) Cookies() (requests.Cookies, error) {
	return requests.ReadCookies(obj.Headers())
}

func (obj *Route) GetCacheKey(routeOption RequestOption) string {
	keyStr := routeOption.Url
	nt := strconv.Itoa(int(time.Now().Unix() / 1000))
	keyStr = re.Sub(fmt.Sprintf(`=%s\d*?&`, nt), "=&", keyStr)
	keyStr = re.Sub(fmt.Sprintf(`=%s\d*?$`, nt), "=", keyStr)
	keyStr = re.Sub(fmt.Sprintf(`=%s\d*?\.\d+?&`, nt), "=&", keyStr)
	keyStr = re.Sub(fmt.Sprintf(`=%s\d*?\.\d+?$`, nt), "=", keyStr)
	keyStr = re.Sub(`=0\.\d{10,}&`, "=&", keyStr)
	keyStr = re.Sub(`=0\.\d{10,}$`, "=", keyStr)
	md5Str := tools.Md5(fmt.Sprintf("%s,%s,%s", routeOption.Method, keyStr, routeOption.PostData))
	return tools.Hex(md5Str)
}
func (obj *Route) request(ctx context.Context, enableBandwidthStats bool, routeOptions ...RequestOption) (fulData FulData, r, w int64, err error) {
	var routeOption RequestOption
	if len(routeOptions) > 0 {
		routeOption = routeOptions[0]
	} else {
		routeOption = obj.NewRequestOption()
	}
	option := obj.webSock.option
	if routeOption.PostData != "" {
		option.Body = routeOption.PostData
	}
	option.Headers = routeOption.Headers
	if routeOption.Proxy != "" {
		option.Proxy = routeOption.Proxy
	}
	option.DialOption = &netx.DialOption{
		EnableBandwidthStats: enableBandwidthStats,
	}
	rs, err := obj.webSock.reqCli.Request(ctx, routeOption.Method, routeOption.Url, option)
	if err != nil {
		return fulData, 0, 0, err
	}
	bwStatus := rs.BWStatus()
	if enableBandwidthStats {
		r, w = bwStatus.R(), bwStatus.W()
	}
	fulData.StatusCode = rs.StatusCode()
	fulData.Body = rs.Text()
	fulData.Headers = rs.Headers()
	fulData.ResponsePhrase = rs.Status()
	if rs.WebSocket() != nil {
		rs.WebSocket().Close()
	}
	if rs.SSE() != nil {
		rs.SSE().Close()
	}
	return
}
func (obj *Route) RawRequest() *requests.Client {
	return obj.webSock.reqCli
}
func (obj *Route) FulFill(ctx context.Context, fulDatas ...FulData) error {
	obj.used = true
	var fulData FulData
	if len(fulDatas) > 0 {
		fulData = fulDatas[0]
	}
	_, err := obj.webSock.FetchFulfillRequest(ctx, obj.recvData.RequestId, fulData)
	if err != nil {
		obj.Fail(ctx)
	}
	return err
}
func (obj *Route) Request(ctx context.Context, routeOptions ...RequestOption) (fulData FulData, err error) {
	f, _, _, err := obj.request(ctx, false, routeOptions...)
	return f, err
}
func (obj *Route) RequestWithBandwidth(ctx context.Context, routeOptions ...RequestOption) (fulData FulData, r, w int64, err error) {
	return obj.request(ctx, true, routeOptions...)
}
func (obj *Route) RequestContinueWithBandwidth(ctx context.Context, options ...RequestOption) (FulData, int64, int64, error) {
	obj.used = true
	fulData, r, w, err := obj.RequestWithBandwidth(ctx, options...)
	if err != nil {
		obj.Fail(ctx)
	} else {
		err = obj.FulFill(ctx, fulData)
	}
	return fulData, r, w, err
}
func (obj *Route) RequestContinue(ctx context.Context, options ...RequestOption) (FulData, error) {
	obj.used = true
	fulData, err := obj.Request(ctx, options...)
	if err != nil {
		obj.Fail(ctx)
	} else {
		err = obj.FulFill(ctx, fulData)
	}
	return fulData, err
}

func (obj *Route) Continue(ctx context.Context, options ...RequestOption) error {
	obj.used = true
	_, err := obj.webSock.FetchContinueRequest(ctx, obj.recvData.RequestId, options...)
	if err != nil {
		obj.Fail(ctx)
	}
	return err
}
func (obj *Route) ResponseBody(ctx context.Context) (string, error) {
	if err := obj.Error(); err != nil {
		obj.Continue(ctx)
		return "", err
	}
	rs, err := obj.webSock.FetchGetResponseBody(ctx, obj.recvData.RequestId)
	if err != nil {
		return "", err
	}
	jsonData, err := gson.Decode(rs.Result)
	if err != nil {
		return "", err
	}
	body := jsonData.Get("body").String()
	if body == "" {
		return body, nil
	}
	if jsonData.Get("base64Encoded").Bool() {
		bodyByte, err := tools.Base64Decode(body)
		if err != nil {
			return body, err
		}
		body = tools.BytesToString(bodyByte)
	}
	return body, nil
}

// Failed, Aborted, TimedOut, AccessDenied, ConnectionClosed, ConnectionReset, ConnectionRefused, ConnectionAborted, ConnectionFailed, NameNotResolved, InternetDisconnected, AddressUnreachable, BlockedByClient, BlockedByResponse
func (obj *Route) Fail(ctx context.Context, errorReasons ...string) error {
	obj.used = true
	var errorReason string
	if len(errorReasons) > 0 {
		errorReason = errorReasons[0]
	}
	if errorReason == "" {
		errorReason = "Failed"
	}
	_, err := obj.webSock.FetchFailRequest(ctx, obj.recvData.RequestId, errorReason)
	return err
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type FulData struct {
	StatusCode     int         `json:"responseCode"`
	Headers        http.Header `json:"responseHeaders"`
	Body           string      `json:"body"`
	ResponsePhrase string      `json:"responsePhrase"`
}

func (obj FulData) Cookies() (requests.Cookies, error) {
	cookies := []*http.Cookie{}
	for _, cook := range obj.Headers.Values("Set-Cookie") {
		result, err := http.ParseSetCookie(cook)
		if err != nil {
			return nil, err
		}
		cookies = append(cookies, result)
	}
	return cookies, nil
}
