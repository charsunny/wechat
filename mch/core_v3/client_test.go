package core_v3

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("Not Equal. %d %d", a, b)
	}
}

// 新建帐号
func Test_NewClient(t *testing.T) {
	fmt.Println("Testing client creatation")
	var err error
	_, err = NewClient("1533391551", "72E6A550FCCC4AB90E1699D06989669221DF167A", "001rsrs001001rsrs001001rsrs001nb", "./apiclient_cert.pem", "./apiclient_key.pem")
	assertEqual(t, err, nil)
	return
}

func Test_DoGet(t *testing.T) {
	fmt.Println("Testing get request")
	var err error
	var cli *Client
	var resp []byte

	cli, _ = NewClient("1533391551", "72E6A550FCCC4AB90E1699D06989669221DF167A", "001rsrs001001rsrs001001rsrs001nb", "./apiclient_cert.pem", "./apiclient_key.pem")

	resp, err = cli.DoGet("/v3/certificates")
	assertEqual(t, err, nil)
	fmt.Println(string(resp))
}

// func Test_DoPost(t *testing.T) {
// 	fmt.Println("Testing post request")
// 	var err error
// 	var cli *Client
// 	var resp []byte

// 	cli, _ = NewClient("1533391551", "72E6A550FCCC4AB90E1699D06989669221DF167A", "001rsrs001001rsrs001001rsrs001nb", "./apiclient_cert.pem", "./apiclient_key.pem")

// 	resp, err = cli.DoPost("/v3/certificates", "")
// 	assertEqual(t, err, nil)
// 	fmt.Println(string(resp))
// }
