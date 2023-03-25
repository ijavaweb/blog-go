package handler

import (
	"blog-go/pkg/model"
	"blog-go/pkg/service"
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/go-xmlpath/xmlpath"
	"log"
	"net/http"
	"sort"
	"strings"
)

const (
	wechatToken     = "67_1ou0ZA_NLl_R8TunAGmB15VF3KWNkHusLRt4V8ZhRH9zRDQ_vmmK1F_Os2dqjVI8nO83NFE10EJUaLnpnIoA5cG0Ac_Nl_wupuf-RZbLdeXOqdmgPL9tVsJJSUwNKVeAIAHEP"
)
func  VerifyData(c *gin.Context) {
	req := c.Request
	signature := req.URL.Query().Get("signature")
	timestamp := req.URL.Query().Get("timestamp")
	nonce := req.URL.Query().Get("nonce")
	echostr := req.URL.Query().Get("echostr")
	log.Println(req)
	log.Println(signature)
	log.Println(nonce)
	log.Println(timestamp)
	log.Println(echostr)
	c.String(http.StatusOK,echostr)
	return
	//if checkSignature(wechatToken, signature, timestamp, nonce) {
	//	c.JSON(http.StatusOK,echostr)
	//	return
	//} else {
	//	c.JSON(http.StatusOK,"invalid signature")
	//}
}
//func MessageHandler (c *gin.Context) {
//	var receivedMessage model.TextMessage
//	err:=c.ShouldBindXML(&receivedMessage)
//	if err != nil {
//		log.Println(err.Error())
//		c.String(http.StatusBadRequest, "Invalid XML")
//		return
//	}
//	go service.GenerateGPTResponse(c,&receivedMessage)
//	return
//}
func MessageHandler (c *gin.Context) {
	var receivedMessage model.ReqMessage
	err:=c.ShouldBind(&receivedMessage)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, "Invalid XML")
		return
	}
	str := service.GenerateGPTTestResponse(receivedMessage.Content)
	c.JSON(http.StatusOK,str)
}
func checkSignature(token, signature, timestamp, nonce string) bool {
	values := []string{token, timestamp, nonce}
	sort.Strings(values)

	hash := sha1.New()
	hash.Write([]byte(strings.Join(values, "")))
	generatedSignature := hex.EncodeToString(hash.Sum(nil))

	return signature == generatedSignature
}
func parseXMLMessage(xmlData string) (string, string) {
	root, err := xmlpath.Parse(strings.NewReader(xmlData))
	if err != nil {
		return "", ""
	}

	contentPath := xmlpath.MustCompile("//xml/Content")
	content, _ := contentPath.String(root)

	toUserNamePath := xmlpath.MustCompile("//xml/FromUserName")
	toUserName, _ := toUserNamePath.String(root)

	return content, toUserName
}