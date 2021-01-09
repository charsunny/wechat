package mchv3

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type CertResponse struct {
	Data []struct {
		SerialNo           string `json:"serial_no"`      // 证书序列号
		EffectiveTime      string `json:"effective_time"` // 生效时间
		ExpireTime         string `json:"expire_time"`    // 过期时间
		EncryptCertificate struct {
			Algorithm      string `json:"algorithm"`       // 算法
			Nonce          string `json:"nonce"`           // 随机串
			AssociatedData string `json:"associated_data"` // 附加数据包
			Ciphertext     string `json:"ciphertext"`      // 密文
		} `json:"encrypt_certificate"` // 证书内容
	} `json:"data"`
}

// 获取微信支付平台证书
// 平台证书会周期性更换。建议商户定时通过API下载新的证书
// 可以弄个Cronjob线程周期执行这个方法
func (cli *Client) GetWechatCertificate() (err error) {
	var certResp *CertResponse
	var resp, plaintext, ciphertext, nonce, ad []byte
	var block cipher.Block
	var pblock *pem.Block
	var aesgcm cipher.AEAD
	var cert *x509.Certificate

	resp, err = cli.DoGet("/v3/certificates")
	if err != nil {
		return
	}
	certResp = new(CertResponse)
	err = json.Unmarshal(resp, certResp)
	if err != nil {
		return
	}

	nonce = []byte(certResp.Data[0].EncryptCertificate.Nonce)
	ciphertext, err = base64.StdEncoding.DecodeString(certResp.Data[0].EncryptCertificate.Ciphertext)
	if err != nil {
		return
	}
	ad = []byte(certResp.Data[0].EncryptCertificate.AssociatedData)

	// 解密过程 https://wechatpay-api.gitbook.io/wechatpay-api-v3/qian-ming-zhi-nan-1/zheng-shu-he-hui-tiao-bao-wen-jie-mi
	block, err = aes.NewCipher([]byte(cli.AppSecret))
	if err != nil {
		return
	}
	aesgcm, err = cipher.NewGCM(block)
	if err != nil {
		return
	}

	plaintext, err = aesgcm.Open(nil, nonce, ciphertext, ad)
	if err != nil {
		return
	}

	// if DEBUG {
	// 	fmt.Printf("%s\n", string(plaintext))
	// }

	// 加载证书
	pblock, _ = pem.Decode([]byte(plaintext))

	cert, err = x509.ParseCertificate(pblock.Bytes)
	if err != nil {
		return
	}

	cli.WechatCertificate = cert
	cli.WechatSerialNumber = strings.ToUpper(cert.SerialNumber.Text(16))
	return
}

// 加签
// https://wechatpay-api.gitbook.io/wechatpay-api-v3/qian-ming-zhi-nan-1/qian-ming-sheng-cheng
func (cli *Client) Sign(method, url, body, nonce string) (sign string, err error) {
	string2sign := method + "\n" + url + "\n" + timestamp() + "\n" + nonce + "\n" + body + "\n"

	return cli.SignStr(string2sign)
}

func (cli *Client) SignStr(str string) (sign string, err error) {
	digest := sha256.Sum256([]byte(str))
	hashed, err := rsa.SignPKCS1v15(nil, cli.PrivateKey, crypto.SHA256, digest[:])
	if err != nil {
		return
	}

	sign = base64.StdEncoding.EncodeToString(hashed)
	return
}

// 验签
// https://wechatpay-api.gitbook.io/wechatpay-api-v3/qian-ming-zhi-nan-1/qian-ming-yan-zheng
func (cli *Client) Verify(resp *http.Response, body []byte) bool {
	var wechatpaySerial, wechatpayTimestamp, wechatpayNonce, wechatpaySignature string
	var sig, digest []byte
	var err error
	var string2verify string

	wechatpaySerial = resp.Header.Get("Wechatpay-Serial")
	wechatpayTimestamp = resp.Header.Get("Wechatpay-Timestamp")
	wechatpayNonce = resp.Header.Get("Wechatpay-Nonce")
	wechatpaySignature = resp.Header.Get("Wechatpay-Signature")

	if wechatpaySerial != cli.WechatSerialNumber {
		fmt.Println(wechatpaySerial, cli.WechatSerialNumber)
		return false
	}

	string2verify = fmt.Sprintf("%s\n%s\n%s\n", wechatpayTimestamp, wechatpayNonce, string(body))
	// if DEBUG {
	// 	fmt.Println(digest)
	// }
	h := sha256.New()
	h.Write([]byte(string2verify))
	digest = h.Sum(nil)

	sig, _ = base64.StdEncoding.DecodeString(wechatpaySignature)

	err = rsa.VerifyPKCS1v15((cli.WechatCertificate.PublicKey).(*rsa.PublicKey), crypto.SHA256, digest, sig)
	if err != nil {
		fmt.Println("verify error:", err)
		return false
	}
	return true

}

func genRandomString(size int) string {

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
