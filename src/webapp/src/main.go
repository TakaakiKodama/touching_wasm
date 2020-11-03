package main

import (
	"golang.org/x/net/html"
	"syscall/js"

	"github.com/TakaakiKodama/touching_wasm/src/webapp/utils"
)

var document = js.Global().Get("document")

func main() {
	tRootDOM := document.Call("getElementById", "app")
	tControl := utils.CreateHTMLTag(utils.HTMLTag{
		Tag:   "div",
		Class: []string{"control"},
		Attrs: []html.Attribute{
			{Key: "id", Val: "control"},
		},
	})

	tMsgField := CreateHTMLTag(HTMLTag{
		Tag:   "input",
		Class: []string{"input"},
		Attrs: []html.Attribute{
			{Key: "id", Val: "sending_msg"},
			{Key: "value", Val: "default"},
		},
	})
	appendChild(tControl, tMsgField)

	tSender := CreateHTMLTag(HTMLTag{
		Tag:   "input",
		Class: []string{"input"},
		Attrs: []html.Attribute{
			{Key: "id", Val: "sender"},
			{Key: "value", Val: "default"},
		},
	})
	appendChild(tControl, tSender)

	tButton := CreateHTMLTag(HTMLTag{
		Tag:   "input",
		Class: []string{"button"},
		Attrs: []html.Attribute{
			{Key: "type", Val: "button"},
			{Key: "value", Val: "送信"},
		},
	})
	tButton.Call("addEventListener", "click", js.FuncOf(SendMsg))
	appendChild(tControl, tButton)

	appendChild(tRootDOM, tControl)

	createHTMLTree(
		[]HTMLTag{
			{
				Tag: "div", Attrs: []html.Attribute{{Key: "id", Val: "chat"}},
				Children: []HTMLTag{
					{Tag: "div", Attrs: []html.Attribute{{Key: "id", Val: "sender_column"}},
						Children: nil,
					},
					{Tag: "div", Attrs: []html.Attribute{{Key: "id", Val: "msg_column"}},
						Children: nil,
					},
				}},
		}, tRootDOM)

	select {}
}
