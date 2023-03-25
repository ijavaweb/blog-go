package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/go-xmlpath/xmlpath"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

const (
	wechatToken     = "your_wechat_token"
	openAIAPIClient = "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"
)
type openAIResponse struct {
	Choices []Choice `json:"choices"`
}
type Choice struct {
	Message M `json:"message"`
	FinishReason string `json:"finish_reason"`
	Index int `json:"index"`
}
type M struct {
	Role string `json:"role"`
	Content string `json:"content"`
}
type openAIRequest struct {
	Model string `json:"model"`
	Messages []Message `json:"messages"`
}
type  Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}
type TextMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
}
type GptRequest struct {
	Content      string   `json:"content"`
}
func main() {
	http.HandleFunc("/wechat", wechatHandler)
	http.ListenAndServe(":8080", nil)
}

func wechatHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		signature := r.URL.Query().Get("signature")
		timestamp := r.URL.Query().Get("timestamp")
		nonce := r.URL.Query().Get("nonce")
		echostr := r.URL.Query().Get("echostr")

		if checkSignature(wechatToken, signature, timestamp, nonce) {
			fmt.Fprint(w, echostr)
		} else {
			fmt.Fprint(w, "Invalid request")
		}

	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		//content, _ := parseXMLMessage(string(body))
		req := &GptRequest{}
		json.Unmarshal(body, req)

		reply := generateGPTResponse(req.Content)
		//
		//response := TextMessage{
		//	ToUserName:   toUserName,
		//	FromUserName: "gh_abcdefgh12345678",
		//	CreateTime:   time.Now().Unix(),
		//	MsgType:      "text",
		//	Content:      reply,
		//}
		//
		//responseXML, err := xml.MarshalIndent(response, "", "  ")
		//if err != nil {
		//	http.Error(w, "Error generating response", http.StatusInternalServerError)
		//	return
		//}
		//
		//w.Header().Set("Content-Type", "application/xml")
		//w.Write(responseXML)
		w.Write([]byte(reply))

	default:
		fmt.Fprint(w, "Invalid request")
	}
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

func generateGPTResponse(prompt string) string {
	fmt.Print(prompt)
	apiURL := "https://api.openai.com/v1/chat/completions"
	messages := make([]Message,0)
	messages = append(messages,Message{
		Role:    "user",
		Content: prompt,
	})
	data := &openAIRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "Error generating response"
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "Error generating response"
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", openAIAPIClient))

	resp, err := client.Do(req)
	if err != nil {
		return "Error generating response"
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error generating response"
	}

	fmt.Println(string(body))
	var result openAIResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "Error generating response"
	}
	for _, t := range result.Choices {
		fmt.Print(t)
	}

	return strings.TrimSpace(result.Choices[0].Message.Content)
}