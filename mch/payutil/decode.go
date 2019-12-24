package payutil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"github.com/romain-jacotin/aesgcm"
)

func AesGcmDecrypt(ciphertext, key, associated_data, nonce string) (result string, err error)  {
	aead, err := aesgcm.NewAES_256_GCM([]byte(key))
	if err != nil {
		return
	}
	content, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return
	}
	plaintext := make([]byte, len(content))
	tag := make([]byte, 16)
	success := aead.AuthenticateThenDecrypt(tag, plaintext, []byte(associated_data), content, []byte(nonce))
	if !success {
		err = errors.New("decode failed")
	} else {
		result = string(plaintext)
	}
	return
}

func ParsePKCS1PublicKey(data []byte) (key *rsa.PublicKey, err error) {
	var block *pem.Block
	block, _ = pem.Decode(data)
	if block == nil {
		return nil, errors.New("public key error")
	}

	var pubInterface interface{}
	pubInterface, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	key, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("public key error")
	}

	return key, err
}

func packageData(originalData []byte, packageSize int) (r [][]byte) {
	var src = make([]byte, len(originalData))
	copy(src, originalData)

	r = make([][]byte, 0)
	if len(src) <= packageSize {
		return append(r, src)
	}
	for len(src) > 0 {
		var p = src[:packageSize]
		r = append(r, p)
		src = src[packageSize:]
		if len(src) <= packageSize {
			r = append(r, src)
			break
		}
	}
	return r
}

func RSAEncryptPKCS1(plaintext, key []byte) (string, error) {
	pub, err := ParsePKCS1PublicKey(key)
	if err != nil {
		return "", err
	}

	var data = packageData(plaintext, pub.N.BitLen()/8-11)
	var cipherData = make([]byte, 0, 0)

	for _, d := range data {
		var c, e = rsa.EncryptPKCS1v15(rand.Reader, pub, d)
		if e != nil {
			return "", e
		}
		cipherData = append(cipherData, c...)
	}
	str := base64.StdEncoding.EncodeToString(cipherData)
	return str, nil
}
