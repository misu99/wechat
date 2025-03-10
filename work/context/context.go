package context

import (
	"github.com/misu99/wechat/v2/credential"
	"github.com/misu99/wechat/v2/work/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
