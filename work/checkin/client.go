package checkin

import (
	"github.com/misu99/wechat/v2/work/context"
)

// Client 打卡接口实例
type Client struct {
	*context.Context
}

// NewClient 初始化实例
func NewClient(ctx *context.Context) *Client {
	return &Client{
		ctx,
	}
}
