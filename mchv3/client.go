package mchv3

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	GATEWAY   = "https://api.mch.weixin.qq.com"
	UserAgent = "User-Agent:Mozilla/5.0 (Windows NT 10.0; WOW64; rv:38.0) Gecko/20100101 Firefox/38.0"
)
const DEBUG bool = true

var Cli *Client // 默认客户端

type Client struct {
	GateWay            string            // 网关
	Isv                bool              // 是否是服务商client
	MerchantId         string            // 商户id
	SerialNumber       string            // 商户序列号
	AppSecret          string            // 商户密钥 微信支付在回调通知和平台证书下载接口中，对关键信息进行了AES-256-GCM加密
	PrivateKey         *rsa.PrivateKey   // 商户API私钥
	WechatCertificate  *x509.Certificate // 平台证书
	WechatSerialNumber string            // 微信平台序列号
	httpClient         *http.Client      // http客户端
}

// 实例化一个客户端
func NewClient(isv bool, merchantId, serialNumber, appSecret, keyFile string) (cli *Client, err error) {

	var pk *rsa.PrivateKey

	cli = new(Client)
	cli.Isv = isv
	cli.GateWay = GATEWAY
	cli.SerialNumber = serialNumber
	cli.MerchantId = merchantId
	cli.AppSecret = appSecret

	// cert, err = tls.LoadX509KeyPair(certFile, keyFile)
	// if err != nil {
	// 	return
	// }

	pk, err = loadPrivateKey(keyFile)
	if err != nil {
		return
	}

	cli.PrivateKey = pk
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
	req.Header.Add("User-Agent", UserAgent)

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
func (cli *Client) DoPost(url, body string, needVerify ...bool) (resp []byte, httpCode int, err error) {
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
	req.Header.Add("User-Agent", UserAgent)

	respser, err := cli.httpClient.Do(req)
	if err != nil {
		return
	}
	defer respser.Body.Close()

	httpCode = respser.StatusCode
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

// 敏感信息加密
// https://wechatpay-api.gitbook.io/wechatpay-api-v3/qian-ming-zhi-nan-1/min-gan-xin-xi-jia-mi
func (cli *Client) SecretFieldEncrypt(plaintext []byte) (ciphertext string, err error) {
	var cipherdata []byte
	cipherdata, err = rsa.EncryptOAEP(sha1.New(), rand.Reader, (cli.WechatCertificate.PublicKey).(*rsa.PublicKey), plaintext, nil)
	if err != nil {
		return
	}
	ciphertext = base64.StdEncoding.EncodeToString(cipherdata)
	return
}

// 图片上传
// https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/tool/chapter3_1.shtml
func (cli *Client) ImageUpload(content []byte, filename string) (mediaId string, err error) {
	var format, meta, string2sign, sign, boundary, nonce, body, auth string
	var hashed, resp []byte
	var digest [32]byte
	var req *http.Request

	format = strings.Split(filename, ".")[1]
	if format == "jpg" || format == "bmp" || format == "png" {
		// pass
	} else {
		err = fmt.Errorf("format not allowed")
		return
	}

	sum := sha256.Sum256(content)
	meta = fmt.Sprintf("{\"filename\":\"%s\",\"sha256\":\"%x\"}", filename, sum)
	nonce = genRandomString(12)
	string2sign = "POST\n" + "/v3/merchant/media/upload\n" + timestamp() + "\n" + nonce + "\n" + meta + "\n"

	fmt.Println(string2sign)

	digest = sha256.Sum256([]byte(string2sign))
	hashed, err = rsa.SignPKCS1v15(nil, cli.PrivateKey, crypto.SHA256, digest[:])
	if err != nil {
		return
	}

	sign = base64.StdEncoding.EncodeToString(hashed)

	boundary = "------------" + genRandomString(12)

	body = "--" + boundary + "\r\n" + "Content-Disposition: form-data; name=\"meta\";" + "\r\n" + "Content-Type: application/json" + "\r\n\r\n" + meta + "\r\n--" + boundary + "\r\n" + fmt.Sprintf("Content-Disposition: form-data; name=\"file\"; filename=\"%s\";", filename) + "\r\n" + fmt.Sprintf("Content-Type: image/%s", format) + "\r\n\r\n" + string(content) + "\r\n--" + boundary + "--\r\n"

	req, err = http.NewRequest("POST", GATEWAY+"/v3/merchant/media/upload", bytes.NewBuffer([]byte(body)))
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data.image/%s", format))
	auth = fmt.Sprintf("mchid=\"%s\",nonce_str=\"%s\",signature=\"%s\",timestamp=\"%s\",serial_no=\"%s\"", cli.MerchantId, nonce, sign, timestamp(), cli.SerialNumber)
	req.Header.Add("Authorization", "WECHATPAY2-SHA256-RSA2048 "+auth)
	req.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data;boundary=\"%s\"", boundary))
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Accept", "application/json")

	respser, err := cli.httpClient.Do(req)
	if err != nil {
		return
	}

	defer respser.Body.Close()
	resp, err = ioutil.ReadAll(respser.Body)
	if err != nil {
		return
	}

	type UploadResult struct {
		MediaId string `json:"media_id"` // 媒体文件标识 Id
	}
	var res *UploadResult

	fmt.Println(string(resp))
	res = new(UploadResult)
	err = json.Unmarshal(resp, res)
	if err != nil {
		return
	}

	mediaId = res.MediaId
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
