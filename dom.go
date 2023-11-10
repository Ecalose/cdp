package cdp

import (
	"context"

	"github.com/gospider007/gson"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type NodeType int64

var (
	NodeTypeElement               NodeType = 1
	NodeTypeAttribute             NodeType = 2
	NodeTypeText                  NodeType = 3
	NodeTypeCDATA                 NodeType = 4
	NodeTypeEntityReference       NodeType = 5
	NodeTypeEntity                NodeType = 6
	NodeTypeProcessingInstruction NodeType = 7
	NodeTypeComment               NodeType = 8
	NodeTypeDocument              NodeType = 9
	NodeTypeDocumentType          NodeType = 10
	NodeTypeDocumentFragment      NodeType = 11
	NodeTypeNotation              NodeType = 12
)

func (obj NodeType) HtmlNodeType() html.NodeType {
	switch obj {
	case NodeTypeElement:
		return html.ElementNode
	case NodeTypeText:
		return html.TextNode
	case NodeTypeComment:
		return html.CommentNode
	case NodeTypeDocument:
		return html.DocumentNode
	case NodeTypeDocumentType:
		return html.DoctypeNode
	default:
		return html.RawNode
	}
}
func ParseJsonDom(data *gson.Client) *html.Node {
	attrs := []html.Attribute{}
	attributes := data.Get("attributes").Array()
	for i := 0; i < len(attributes)/2; i++ {
		attrs = append(attrs, html.Attribute{
			Key: attributes[i*2].String(),
			Val: attributes[i*2+1].String(),
		})
	}
	attrs = append(attrs, html.Attribute{Key: "gospiderNodeId", Val: data.Get("nodeId").String()})
	nodeType := NodeType(data.Get("nodeType").Int())
	curNode := &html.Node{Type: nodeType.HtmlNodeType(), Attr: attrs}
	curNode.DataAtom = atom.Lookup(data.Get("localName").Bytes())
	switch nodeType {
	case NodeTypeText:
		curNode.Data = data.Get("nodeValue").String()
	case NodeTypeElement:
		curNode.Data = data.Get("localName").String()
	default:
		if curNode.Data = data.Get("nodeValue").String(); curNode.Data == "" {
			curNode.Data = data.Get("localName").String()
		}
	}
	for _, children := range data.Get("children").Array() {
		if node := ParseJsonDom(children); node != nil {
			curNode.AppendChild(node)
		}
	}
	for _, children := range data.Get("contentDocument.children").Array() {
		if node := ParseJsonDom(children); node != nil {
			curNode.AppendChild(node)
		}
	}
	return curNode
}
func (obj *WebSock) DOMEnable(ctx context.Context) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "DOM.enable",
	})
}
func (obj *WebSock) DOMDescribeNode(ctx context.Context, nodeId, backendNodeId int64) (RecvData, error) {
	params := map[string]any{
		"depth": 0,
	}
	if backendNodeId != 0 {
		params["backendNodeId"] = backendNodeId
	} else {
		params["nodeId"] = nodeId
	}
	return obj.send(ctx, commend{
		Method: "DOM.describeNode",
		Params: params,
	})
}
func (obj *WebSock) DOMResolveNode(ctx context.Context, backendNodeId int64) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "DOM.resolveNode",
		Params: map[string]any{
			"backendNodeId": backendNodeId,
		},
	})
}
func (obj *WebSock) DOMGetFrameOwner(ctx context.Context, frameId string) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "DOM.getFrameOwner",
		Params: map[string]any{
			"frameId": frameId,
		},
	})
}

func (obj *WebSock) DOMRequestNode(ctx context.Context, objectId string) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "DOM.requestNode",
		Params: map[string]any{
			"objectId": objectId,
		},
	})
}

func (obj *WebSock) DOMSetOuterHTML(ctx context.Context, nodeId int64, outerHTML string) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "DOM.setOuterHTML",
		Params: map[string]any{
			"nodeId":    nodeId,
			"outerHTML": outerHTML,
		},
	})
}
func (obj *WebSock) DOMGetOuterHTML(ctx context.Context, nodeId int64, backendNodeId int64) (RecvData, error) {
	params := map[string]any{}
	if backendNodeId != 0 {
		params["backendNodeId"] = backendNodeId
	} else {
		params["nodeId"] = nodeId
	}
	return obj.send(ctx, commend{
		Method: "DOM.getOuterHTML",
		Params: params,
	})
}
func (obj *WebSock) DOMFocus(ctx context.Context, nodeId int64) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "DOM.focus",
		Params: map[string]any{
			"nodeId": nodeId,
		},
	})
}
func (obj *WebSock) DOMQuerySelector(ctx context.Context, nodeId int64, selector string) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "DOM.querySelector",
		Params: map[string]any{
			"nodeId":   nodeId,
			"selector": selector,
		},
	})
}
func (obj *WebSock) DOMQuerySelectorAll(ctx context.Context, nodeId int64, selector string) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "DOM.querySelectorAll",
		Params: map[string]any{
			"nodeId":   nodeId,
			"selector": selector,
		},
	})
}
func (obj *WebSock) DOMGetBoxModel(ctx context.Context, nodeId int64) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "DOM.getBoxModel",
		Params: map[string]any{
			"nodeId": nodeId,
		},
	})
}
func (obj *WebSock) DOMGetDocument(ctx context.Context) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "DOM.getDocument",
		Params: map[string]any{
			"depth": 0,
		},
	})
}
func (obj *WebSock) DOMGetDocuments(ctx context.Context) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "DOM.getDocument",
		Params: map[string]any{
			"depth":  -1,
			"pierce": true,
		},
	})
}
func (obj *WebSock) DOMScrollIntoViewIfNeeded(ctx context.Context, nodeId int64) (RecvData, error) {
	return obj.send(ctx, commend{
		Method: "DOM.scrollIntoViewIfNeeded",
		Params: map[string]any{
			"nodeId": nodeId,
		},
	})
}
