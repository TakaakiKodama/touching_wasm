package utils

import (
	"golang.org/x/net/html"
	"syscall/js"
)

type HTMLTag struct {
	Tag      string
	Class    []string
	Attrs    []html.Attribute
	Children []HTMLTag
}

var document = js.Global().Get("document")

func createHTMLTree(aHTMLTag []HTMLTag, aParentTag js.Value) {
	for tChildIndex := range aHTMLTag {
		tNewTag := CreateHTMLTag(aHTMLTag[tChildIndex])
		createHTMLTree(aHTMLTag[tChildIndex].Children, tNewTag)
		class := tNewTag.Get("classList")
		for _, tClassName := range aHTMLTag[tChildIndex].Class {
			class.Call("add", tClassName)
		}
		appendChild(aParentTag, tNewTag)
	}
}

func CreateHTMLTag(aHTMLTag HTMLTag) js.Value {
	tNewTag := document.Call("createElement", aHTMLTag.Tag)
	class := tNewTag.Get("classList")
	for _, tClassName := range aHTMLTag.Class {
		class.Call("add", tClassName)
	}
	AddAttrs(tNewTag, aHTMLTag.Attrs)
	return tNewTag
}

func appendChild(aDOM js.Value, aElm js.Value) js.Value {
	return aDOM.Call("appendChild", aElm)
}

func AddAttrs(aDOM js.Value, aAttrs []html.Attribute) {
	for tAttrsIndex := range aAttrs {
		aDOM.Set(aAttrs[tAttrsIndex].Key, aAttrs[tAttrsIndex].Val)
	}
}

func SendMsg(this js.Value, args []js.Value) interface{} {
	tMsgText := document.Call("getElementById", "sending_msg").Get("value").String()
	tMsgParagraph := CreateHTMLTag(HTMLTag{
		Tag: "p",
		Attrs: []html.Attribute{
			{Key: "class", Val: "msg"},
			{Key: "textContent", Val: tMsgText},
		},
	})
	appendChild(document.Call("getElementById", "msg_column"), tMsgParagraph)

	tSenderNameText := document.Call("getElementById", "sender").Get("value").String()
	tSender := CreateHTMLTag(HTMLTag{
		Tag: "p",
		Attrs: []html.Attribute{
			{Key: "class", Val: "sender"},
			{Key: "textContent", Val: tSenderNameText},
		},
	})
	appendChild(document.Call("getElementById", "sender_column"), tSender)

	return nil
}
