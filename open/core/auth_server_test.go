package core

import (
	"encoding/xml"
	"testing"
)

func TestXmlUnmarshal(t *testing.T) {
	data := []byte(`<xml>
    <AppId><![CDATA[wxfa6cb88385ba59d8]]></AppId>
    <Encrypt><![CDATA[wBXNqCMv40htfPEaxpiYKU03wtnotwA4tLnJIzliKDTQAVqGUZd4PUbYoEnxFO026H12f13fPTOJcUrxGlsHi0paCM1LJb0PiATXhKYPAq0KbBoqRjg8MNqdx2A+HnbaK5YRD34BcBNfZDv59qGQ0q/mUHIu4NSHPOKwThEFxl4etWNIQOawgdQVIq6F5EWl6+yO2/tQB1FIY2DIcSfXgugDPapNicD3koxpLqYZN3YVyEWC2QVrd1YgYUTPHuxDIXenjCx6ZnuJB8k9k6Y1sHjAA2GHkJLzu+ZJiBOnQGtEQoGw7zjpCgLmEKm8HU9MTpekRGankdJxmvJ7Nj36fNpUTec28U8ykpsedBcvFnUrLRnJlo5fv2W+vZ+8enJivH9K5ejJzRutlIp8l8p402vicH5d6v3YwX6S3k96e0qpvJuhGpWwTCPzSVV/gFjsZ17AfNd8Hh01lV09pExZJg==]]></Encrypt>
</xml>`)

	var x cipherRequestHttpBody
	if err := xmlUnmarshal(data, &x); err != nil {
		t.Error(err)
		return
	}
	if x.ToUserName != "gh_b1eb3f8bd6c6" {
		t.Errorf("ToUserName mismatch,\nhave: %s\nwant: %s\n", x.ToUserName, "gh_b1eb3f8bd6c6")
		return
	}
	wantEncrypt := `DlCGq+lWQuyjNNK+vDaO0zUltpdUW3u4V00WCzsdNzmZGEhrU7TPxG52viOKCWYPwTMbCzgbCtakZHyNxr5hjoZJ7ORAUYoIAGQy/LDWtAnYgDO+ppKLp0rDq+67Dv3yt+vatMQTh99NII6x9SEGpY3O2h8RpG99+NYevQiOLVKqiQYzan21sX/jE4Y3wZaeudsb4QVjqzRAPaCJ5nS3T31uIR9fjSRgHTDRDOzjQ1cHchge+t6faUhniN5VQVTE+wIYtmnejc55BmHYPfBnTkYah9+cTYnI3diUPJRRiyVocJyHlb+XOZN22dsx9yzKHBAyagaoDIV8Yyb/PahcUbsqGv5wziOgLJQIa6z93/VY7d2Kq2C2oBS+Qb+FI9jLhgc3RvCi+Yno2X3cWoqbsRwoovYdyg6jme/H7nMZn77PSxOGRt/dYiWx2NuBAF7fNFigmbRiive3DyOumNCMvA==`
	if string(x.Base64EncryptedMsg) != wantEncrypt {
		t.Errorf("Encrypt mismatch,\nhave: %s\nwant: %s\n", x.Base64EncryptedMsg, wantEncrypt)
		return
	}
}

func BenchmarkXmlUnmarshal(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	data := []byte(`<xml>
    <ToUserName><![CDATA[gh_b1eb3f8bd6c6]]></ToUserName>
    <Encrypt><![CDATA[DlCGq+lWQuyjNNK+vDaO0zUltpdUW3u4V00WCzsdNzmZGEhrU7TPxG52viOKCWYPwTMbCzgbCtakZHyNxr5hjoZJ7ORAUYoIAGQy/LDWtAnYgDO+ppKLp0rDq+67Dv3yt+vatMQTh99NII6x9SEGpY3O2h8RpG99+NYevQiOLVKqiQYzan21sX/jE4Y3wZaeudsb4QVjqzRAPaCJ5nS3T31uIR9fjSRgHTDRDOzjQ1cHchge+t6faUhniN5VQVTE+wIYtmnejc55BmHYPfBnTkYah9+cTYnI3diUPJRRiyVocJyHlb+XOZN22dsx9yzKHBAyagaoDIV8Yyb/PahcUbsqGv5wziOgLJQIa6z93/VY7d2Kq2C2oBS+Qb+FI9jLhgc3RvCi+Yno2X3cWoqbsRwoovYdyg6jme/H7nMZn77PSxOGRt/dYiWx2NuBAF7fNFigmbRiive3DyOumNCMvA==]]></Encrypt>
</xml>`)
	var x cipherRequestHttpBody
	for i := 0; i < b.N; i++ {
		xmlUnmarshal(data, &x)
	}
}

func BenchmarkStdXmlUnmarshal(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	data := []byte(`<xml>
    <ToUserName><![CDATA[gh_b1eb3f8bd6c6]]></ToUserName>
    <Encrypt><![CDATA[DlCGq+lWQuyjNNK+vDaO0zUltpdUW3u4V00WCzsdNzmZGEhrU7TPxG52viOKCWYPwTMbCzgbCtakZHyNxr5hjoZJ7ORAUYoIAGQy/LDWtAnYgDO+ppKLp0rDq+67Dv3yt+vatMQTh99NII6x9SEGpY3O2h8RpG99+NYevQiOLVKqiQYzan21sX/jE4Y3wZaeudsb4QVjqzRAPaCJ5nS3T31uIR9fjSRgHTDRDOzjQ1cHchge+t6faUhniN5VQVTE+wIYtmnejc55BmHYPfBnTkYah9+cTYnI3diUPJRRiyVocJyHlb+XOZN22dsx9yzKHBAyagaoDIV8Yyb/PahcUbsqGv5wziOgLJQIa6z93/VY7d2Kq2C2oBS+Qb+FI9jLhgc3RvCi+Yno2X3cWoqbsRwoovYdyg6jme/H7nMZn77PSxOGRt/dYiWx2NuBAF7fNFigmbRiive3DyOumNCMvA==]]></Encrypt>
</xml>`)
	var x cipherRequestHttpBody
	for i := 0; i < b.N; i++ {
		xml.Unmarshal(data, &x)
	}
}
