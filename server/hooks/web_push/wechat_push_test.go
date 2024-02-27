package web_push

import (
	"pmail/config"
	"pmail/dto/parsemail"
	"testing"
)

func testInit() {

	config.Init()

}
func TestWebPushHook_ReceiveParseAfter(t *testing.T) {
	testInit()

	w := NewWebPushHook()
	w.ReceiveParseAfter(&parsemail.Email{Subject: "标题", Text: []byte("文本内容")})
}
