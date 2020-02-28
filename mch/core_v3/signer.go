package core_v3

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

func (cli *Client) Sign(method, url, body string) (sign string, err error) {
	var string2sign string
	var digest [32]byte
	var hashed []byte

	string2sign = method + "\n" + url + "\n" + timestamp() + "\n" + genRandomString(12) + "\n" + body + "\n"
	if DEBUG {
		fmt.Println("string to sign: ", string2sign)
	}

	digest = sha256.Sum256([]byte(string2sign))
	hashed, err = rsa.SignPKCS1v15(nil, cli.PrivateKey, crypto.SHA256, digest[:])
	if err != nil {
		return
	}

	sign = base64.StdEncoding.EncodeToString(hashed)

	return
}

func genRandomString(size int) string {

	if DEBUG {
		return "123456"
	}

	var a, b []byte
	var i int

	a = []byte("1234567890abcdefghijklmnopqrstuvwxyz")
	b = make([]byte, size)

	for i = 0; i < size; i++ {
		b[i] = a[rand.Intn(36)]
	}

	return string(b)
}

func timestamp() string {
	// if DEBUG {
	// 	return "1582883827"
	// }
	return fmt.Sprintf("%d", time.Now().Unix())
}
