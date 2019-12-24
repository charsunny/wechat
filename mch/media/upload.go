package media

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/xml"
	"github.com/charsunny/wechat/mch/core"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)


var cert tls.Certificate
var caCert []byte

// UploadMediaRequest 上传图片数据结构
type UploadMediaRequest struct {
	// XMLName xml根节点标识
	XMLName xml.Name `xml:"xml"`
	// Media 媒体文件
	Media string `xml:"media"`
	// MediaHash 媒体文件hash
	MediaHash string `xml:"media_hash"`
}

// Submit 申请入驻接口提交你的小微商户资料，申请后一般5分钟左右可以查询到具体的申请结果
//  NOTE: 请求需要双向证书.
func Upload(clt *core.Client, file string) (resp map[string]string, err error) {
	m1 := make(map[string]string, 16)

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	// Add your image file
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("open img file error:", err)
	}

	// calculate hash first
	md5h := md5.New()
	io.Copy(md5h, f)
	f.Seek(0, os.SEEK_SET)

	defer f.Close()

	//file part
	fw, err := w.CreateFormFile("media", file)
	if err != nil {
		log.Fatal(err)
		return
	}
	if _, err = io.Copy(fw, f); err != nil {
		log.Fatal(err)
		return
	}

	var uploadMediaRequest UploadMediaRequest

	uploadMediaRequest.MediaHash = hex.EncodeToString(md5h.Sum(nil))

	data, err :=  xml.Marshal(uploadMediaRequest)
	if err != nil {
		return
	}
	err = xml.Unmarshal(data, &m1)
	if err != nil {
		return
	}

	return clt.PostXML(core.APIBaseURL()+"/applyment/micro/submit", m1)
}

//func main() {
//	log.Println("===start===")
//	var uploadMediaRequest UploadMediaRequest
//	uploadMediaRequest.MchID = config.MchID
//	//multipart post form
//
//	var b bytes.Buffer
//	w := multipart.NewWriter(&b)
//	// Add your image file
//	file := config.ImgFile
//	f, err := os.Open(file)
//	if err != nil {
//		log.Fatal("open img file error:", err)
//	}
//
//	// calculate hash first
//	md5h := md5.New()
//	io.Copy(md5h, f)
//	f.Seek(0, os.SEEK_SET)
//
//	defer f.Close()
//
//	//file part
//	fw, err := w.CreateFormFile("media", file)
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//	if _, err = io.Copy(fw, f); err != nil {
//		log.Fatal(err)
//		return
//	}
//
//	uploadMediaRequest.MediaHash = hex.EncodeToString(md5h.Sum(nil))
//	// calculate sign
//	var hashString = "mch_id=" + uploadMediaRequest.MchID + "&media_hash=" + uploadMediaRequest.MediaHash + "&key=" + config.Key
//	log.Println("to hash string: ",hashString)
//	hasher := md5.New()
//	hasher.Write([]byte(hashString))
//	uploadMediaRequest.Sign = strings.ToUpper(hex.EncodeToString(hasher.Sum(nil)))
//
//	output, err := xml.MarshalIndent(uploadMediaRequest, "  ", "    ")
//	if err != nil {
//		log.Printf("error: %v\n", err)
//	}
//
//	log.Println("Request in XML:\n", string(output))
//
//	// Add the other fields
//	// mch_id
//	fw, err = w.CreateFormField("mch_id")
//	if  err != nil {
//		return
//	}
//	if _, err = fw.Write([]byte(uploadMediaRequest.MchID)); err != nil {
//		return
//	}
//
//	// media_hash
//	if fw, err = w.CreateFormField("media_hash"); err != nil {
//		return
//	}
//	if _, err = fw.Write([]byte(uploadMediaRequest.MediaHash)); err != nil {
//		return
//	}
//
//	// sign
//	if fw, err = w.CreateFormField("sign"); err != nil {
//		return
//	}
//	if _, err = fw.Write([]byte(uploadMediaRequest.Sign)); err != nil {
//		return
//	}
//
//	// Don't forget to close the multipart writer.
//	// If you don't close it, your request will be missing the terminating boundary.
//	w.Close()
//
//	newReq, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/secapi/mch/uploadmedia", &b)
//	newReq.Header.Set("Content-Type", w.FormDataContentType())
//
//	var client *http.Client
//	caCertPool := x509.NewCertPool()
//	caCertPool.AppendCertsFromPEM(caCert)
//	tlsConfig := &tls.Config{
//		Certificates: []tls.Certificate{cert},
//		RootCAs:      caCertPool,
//	}
//	tlsConfig.BuildNameToCertificate()
//	transport := &http.Transport{TLSClientConfig: tlsConfig}
//	client = &http.Client{Transport: transport}
//
//	resp, err := client.Do(newReq)
//
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//
//	log.Println(resp.StatusCode)
//
//	respBody, err := ioutil.ReadAll(resp.Body)
//	log.Printf("[INFO][IN] Respond content:%v", string(respBody))
//	defer resp.Body.Close()
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//
//	return
//}