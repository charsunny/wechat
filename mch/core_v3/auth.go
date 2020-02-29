package core_v3

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
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

	if DEBUG {
		fmt.Printf("%s\n", string(plaintext))
	}

	// 加载证书
	pblock, _ = pem.Decode([]byte(plaintext))

	cert, err = x509.ParseCertificate(pblock.Bytes)
	if err != nil {
		return
	}

	cli.WechatCertificate = cert
	return
}

func main() {}
