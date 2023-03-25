package service

import (
	"blog-go/pkg/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func GenerateGPTResponse(c *gin.Context,receivedMessage *model.TextMessage)  {
	log.Println(*receivedMessage)
	start := time.Now().Unix()
	apiURL := "https://api.openai.com/v1/chat/completions"
	messages := make([]model.Message,0)
	messages = append(messages,model.Message{
		Role:    "user",
		Content: receivedMessage.Content,
	})
	data := &model.OpenAIRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}
	secret := os.Getenv("OPENAI_API_KEY")
	if secret == "" {
		log.Println(" empty secret")
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", secret))

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	end := time.Now().Unix()
	log.Printf("time cost : %v",end-start)
	log.Println()
	log.Println()
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	var result model.OpenAIResponse
	err = json.Unmarshal(body, &result)
	log.Println(result)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if len(result.Choices) == 0 {
		return
	}
	reply := strings.TrimSpace(result.Choices[0].Message.Content)
	//response := model.TextMessage{
	//	ToUserName:   receivedMessage.ToUserName,
	//	FromUserName: receivedMessage.FromUserName,
	//	CreateTime:   time.Now().Unix(),
	//	MsgType:      receivedMessage.MsgType,
	//	Content:       reply,
	//}
	t := `<xml>
<ToUserName><![CDATA[%s]]></ToUserName>
<FromUserName><![CDATA[%s]]></FromUserName>
<CreateTime>%d</CreateTime>
<MsgType><![CDATA[%s]]></MsgType>
<Content><![CDATA[%s]]></Content>
</xml>`
	content := fmt.Sprintf(t,receivedMessage.ToUserName, receivedMessage.FromUserName,time.Now().Unix(),receivedMessage.MsgType,reply)
	c.Data(200, "application/xml; charset=utf-8", []byte(content))
	return
}
func GenerateGPTTestResponse(content string) string {
	start := time.Now().Unix()
	apiURL := "https://api.openai.com/v1/chat/completions"
	messages := make([]model.Message,0)
	messages = append(messages,model.Message{
		Role:    "user",
		Content: content,
	})
	data := &model.OpenAIRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return " "
	}
	secret ,_:= os.LookupEnv("OPENAI_API_KEY")
	log.Println(secret)
	log.Println(len(secret))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", secret))

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	end := time.Now().Unix()
	log.Printf("time cost : %v",end-start)
	log.Println()
	log.Println()
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	var result model.OpenAIResponse
	err = json.Unmarshal(body, &result)
	log.Println(result)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	if len(result.Choices) == 0 {
		return ""
	}
	return strings.TrimSpace(result.Choices[0].Message.Content)

}