package autoreply

import (
	"github.com/charsunny/wechat/mp/core"
)

// 查询自动回复规则
func Get(clt *core.Client) (ar *AutoReply, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/get_current_autoreply_info?access_token="

	var result struct {
		core.Error
		AutoReply
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	ar = &result.AutoReply
	return
}
