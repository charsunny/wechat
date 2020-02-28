package core_v3

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

const (
	GATEWAY = "https://api.mch.weixin.qq.com/"
)

type Client struct {
	GateWay           string           // 网关
	MerchantId        string           // 商户id
	SerialNumber      string           // 序列号
	AppSecret         string           // 商户密钥 微信支付在回调通知和平台证书下载接口中，对关键信息进行了AES-256-GCM加密
	PrivateKey        *rsa.PrivateKey  // 商户API私钥
	Certificate       tls.Certificate  // 商户证书
	WechatCertificate *tls.Certificate // 平台证书
}

// 实例化一个客户端
func NewClient(merchantId, serialNumber, appSecret, certFile, keyFile string) (cli *Client, err error) {

	var cert tls.Certificate
	var pk *rsa.PrivateKey

	cli = new(Client)

	cli.GateWay = GATEWAY
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
	cli.Certificate = cert

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
