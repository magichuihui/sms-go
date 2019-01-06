package api


type MessageResult struct {
	MessageId string `json:"messageId"`
	Content string `json:"content"`
	Receiver []string `json:"receiver"`
	SendTime string `json:"sendTime"`
}

type SendResult struct {
	RequestId string `json:"requestId"`
	Code string `json:"code"`
	Message string `json:"message"`
}