package api

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func GetMessage(cli bce.Client, messageId string) (*MessageResult, error) {
	req := &bce.BceRequest{}
	req.SetUri("/v1/message/" + messageId)
	req.SetMethod(http.GET)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &MessageResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

func SendMessage(cli bce.Client, body *bce.Body) (*SendResult, error) {
	req := &bce.BceRequest{}
	req.SetUri("/api/v3/sendSms?clientToken=123123412")
	req.SetMethod(http.POST)

	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	req.SetHeader(http.BCE_DATE, "2019-12-04T06:53:12Z")
	req.SetBody(body)

	fmt.Println(req.Request)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &SendResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}