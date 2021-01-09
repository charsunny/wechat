package mchv3

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func assertEqual(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("Not Equal. %v %v", a, b)
	}
}

func init() {
	Cli, _ = NewClient(false, "1533391551", "72E6A550FCCC4AB90E1699D06989669221DF167A", "001rsrs001001rsrs001001rsrs001nb", "./apiclient_cert.pem")
}

func Test_DoGet(t *testing.T) {
	fmt.Println("----------------------------------")
	fmt.Println("Testing get request")
	var err error
	var resp []byte

	err = Cli.GetWechatCertificate()
	assertEqual(t, err, nil)

	resp, err = Cli.DoGet("/v3/certificates", true)
	assertEqual(t, err, nil)
	fmt.Println(string(resp))
}

func Test_SecretMsgEncrypt(t *testing.T) {
	fmt.Println("----------------------------------")
	fmt.Println("Testing secret msg encrypt")

	var err error
	var ciphertext string

	ciphertext, err = Cli.SecretFieldEncrypt([]byte("secret msg"))
	assertEqual(t, err, nil)
	fmt.Println(ciphertext)
}

func Test_uploadImage(t *testing.T) {
	fmt.Println("----------------------------------")
	fmt.Println("Testing upload image")

	var content []byte
	var err error
	var mediaId string

	content, _ = ioutil.ReadFile("./5W7A1549.jpg")
	mediaId, err = Cli.ImageUpload(content, "babe.jpg")
	assertEqual(t, err, nil)
	fmt.Println("media_id:", mediaId)
}
