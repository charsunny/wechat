package mmpaymkttransfers
import (
	"github.com/charsunny/wechat/mch/core"
)

// 请求单次分账.
// https://pay.weixin.qq.com/wiki/doc/api/allocation_sl.php?chapter=25_6&index=1
//  NOTE: 请求需要双向证书
func ProfitSharing(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/secapi/pay/profitsharing", req)
}

// 请求多次分账.
// https://pay.weixin.qq.com/wiki/doc/api/allocation_sl.php?chapter=25_6&index=2
//  NOTE: 请求需要双向证书
func MultiProfitSharing(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/secapi/pay/multiprofitsharing", req)
}

// 查询分账结果
// https://pay.weixin.qq.com/wiki/doc/api/allocation_sl.php?chapter=25_2&index=3
func ProfitSharingQuery(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/pay/profitsharingquery", req)
}

// 回退结果查询
// https://pay.weixin.qq.com/wiki/doc/api/allocation_sl.php?chapter=25_2&index=3
func ProfitSharingReturnQuery(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/pay/profitsharingreturnquery", req)
}

// 添加分账接收方
func ProfitSharingAddReceiver(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/pay/profitsharingaddreceiver", req)
}

// 删除分账接收方
func ProfitSharingRemoveReceiver(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/pay/profitsharingremovereceiver", req)
}

// 完结分账
//  NOTE: 请求需要双向证书
func ProfitSharingFinish(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/secapi/pay/profitsharingfinish", req)
}

// 分账回退
//  NOTE: 请求需要双向证书
func ProfitSharingReturn(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/secapi/pay/profitsharingreturn", req)
}