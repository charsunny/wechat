package core_v3

import (
	"fmt"
	"testing"
)

// 新建帐号
func Test_NewClient(t *testing.T) {
	var err error
	_, err = NewClient("1533391551", "", "001rsrs001001rsrs001001rsrs001nb", "./apiclient_cert.pem", "./apiclient_key.pem")
	if err != nil {
		fmt.Println("error: ", err)
	}
	return
}
