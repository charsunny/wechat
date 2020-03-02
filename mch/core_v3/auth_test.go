package core_v3

import (
	"fmt"
	"testing"
)

func Test_GetWechatCertificate(t *testing.T) {
	fmt.Println("----------------------------------")
	fmt.Println("Testing get wechat certificate")

	var err error
	var cli *Client

	cli, _ = NewClient("1533391551", "72E6A550FCCC4AB90E1699D06989669221DF167A", "001rsrs001001rsrs001001rsrs001nb", "./apiclient_cert.pem", "./apiclient_key.pem")

	err = cli.GetWechatCertificate()
	assertEqual(t, err, nil)
	return
}
