// Code generated by hertz generator.

package Message

import (
	"github.com/cloudwego/hertz/pkg/app"
	"offer_tiktok/biz/mw"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _messageMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _message_ctionMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _chatMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _messagechatMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}
