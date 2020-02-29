package core_v3

import (
	"fmt"
	"testing"
)

// 签名
func Test_Signer(t *testing.T) {
	fmt.Println("----------------------------------")
	fmt.Println("Testing sign")
	var err error
	var cli *Client

	cli, err = NewClient("1533391551", "", "001rsrs001001rsrs001001rsrs001nb", "./apiclient_cert.pem", "./apiclient_key.pem")
	assertEqual(t, err, nil)

	_, err = cli.Sign("GET", "/v3/certificates", "")
	assertEqual(t, err, nil)
	return
}
