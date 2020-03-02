package core_v3

import (
	"bytes"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	GATEWAY = "https://api.mch.weixin.qq.com"
)
const DEBUG bool = true

type Client struct {
	GateWay            string            // 网关
	MerchantId         string            // 商户id
	SerialNumber       string            // 商户序列号
	AppSecret          string            // 商户密钥 微信支付在回调通知和平台证书下载接口中，对关键信息进行了AES-256-GCM加密
	PrivateKey         *rsa.PrivateKey   // 商户API私钥
	Certificate        *tls.Certificate  // 商户证书
	WechatCertificate  *x509.Certificate // 平台证书
	WechatSerialNumber string            // 微信平台序列号
	httpClient         *http.Client      // http客户端
}

// 实例化一个客户端
func NewClient(merchantId, serialNumber, appSecret, certFile, keyFile string) (cli *Client, err error) {

	var cert tls.Certificate
	var pk *rsa.PrivateKey

	cli = new(Client)

	cli.GateWay = GATEWAY
	cli.SerialNumber = serialNumber
	cli.MerchantId = merchantId
	cli.AppSecret = appSecret

	cert, err = tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return
	}

	pk, err = loadPrivateKey(keyFile)
	if err != nil {
		return
	}

	cli.PrivateKey = pk
	cli.Certificate = &cert
	cli.httpClient = http.DefaultClient

	return
}

// GET一个API
func (cli *Client) DoGet(url string, needVerify ...bool) (resp []byte, err error) {

	var req *http.Request
	var sign, auth, nonce string

	req, err = http.NewRequest("GET", GATEWAY+url, nil)
	if err != nil {
		return
	}
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Accept", "application/json")

	// sign
	nonce = genRandomString(12)
	sign, err = cli.Sign("GET", url, "", nonce)
	if err != nil {
		return
	}

	auth = fmt.Sprintf("mchid=\"%s\",nonce_str=\"%s\",signature=\"%s\",timestamp=\"%s\",serial_no=\"%s\"", cli.MerchantId, nonce, sign, timestamp(), cli.SerialNumber)
	req.Header.Add("Authorization", "WECHATPAY2-SHA256-RSA2048 "+auth)

	respser, err := cli.httpClient.Do(req)
	if err != nil {
		return
	}

	defer respser.Body.Close()
	resp, err = ioutil.ReadAll(respser.Body)
	if err != nil {
		return
	}

	if len(needVerify) > 0 && needVerify[0] {
		if !cli.Verify(respser, resp) {
			err = fmt.Errorf("Verify error")
		}
	}
	return
}

// POST一个API
func (cli *Client) DoPost(url, body string, needVerify ...bool) (resp []byte, err error) {
	var req *http.Request
	var sign, auth, nonce string

	req, err = http.NewRequest("POST", GATEWAY+url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return
	}
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Accept", "application/json")

	// sign
	nonce = genRandomString(12)
	sign, err = cli.Sign("POST", url, body, nonce)
	if err != nil {
		return
	}
	auth = fmt.Sprintf("mchid=\"%s\",nonce_str=\"%s\",signature=\"%s\",timestamp=\"%s\",serial_no=\"%s\"", cli.MerchantId, nonce, sign, timestamp(), cli.SerialNumber)
	req.Header.Add("Authorization", "WECHATPAY2-SHA256-RSA2048 "+auth)
	respser, err := cli.httpClient.Do(req)
	if err != nil {
		return
	}

	defer respser.Body.Close()
	resp, err = ioutil.ReadAll(respser.Body)
	if err != nil {
		return
	}

	if len(needVerify) > 0 && needVerify[0] {
		if !cli.Verify(respser, resp) {
			err = fmt.Errorf("Verify error")
		}
	}

	return
}

// 加载一个私钥
func loadPrivateKey(path string) (*rsa.PrivateKey, error) {

	priv, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	privPem, _ := pem.Decode(priv)

	var privPemBytes []byte

	privPemBytes = privPem.Bytes

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(privPemBytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(privPemBytes); err != nil { // note this returns type `interface{}`
			return nil, err
		}
	}

	return parsedKey.(*rsa.PrivateKey), nil
}
