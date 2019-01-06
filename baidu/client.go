package baidu

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/magichuihui/sms-go/baidu/api"
)

const (
	DEFAULT_SERVICE_DOMAIN = "https://sms." + bce.DEFAULT_REGION + ".baidubce.com"
)

type Client struct {
	*bce.BceClient
}

// NewClient make the sms service client with default configuration.
// Use `cli.Config.xxx` to access the config or change it to non-default value.
func NewClient(ak, sk, endpoint string) (*Client, error) {
	var credentials *auth.BceCredentials
	var err error
	if len(ak) == 0 && len(sk) == 0 { // to support public-read-write request
		credentials, err = nil, nil
	} else {
		credentials, err = auth.NewBceCredentials(ak, sk)
		if err != nil {
			return nil, err
		}
	}
	if len(endpoint) == 0 {
		endpoint = DEFAULT_SERVICE_DOMAIN
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:    endpoint,
		Region:      bce.DEFAULT_REGION,
		UserAgent:   bce.DEFAULT_USER_AGENT,
		Credentials: credentials,
		SignOption:  defaultSignOptions,
		Retry:       bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	v1Signer := &auth.BceV1Signer{}

	client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
}


// GetMessage - 查询短信的流水信息
//
// PARAMS:
//		- messageId: sms message id
// RETURNS:
//		- *api.MessageResult: Detail of sms mesage
//     	- error: the return error if any occurs
func (c *Client) GetMessage(messageId string) (*api.MessageResult, error) {
	return api.GetMessage(c, messageId)
}

// SendMessage - send sms message
//
// PARAMS:
//		- content: message content
// RETURNS:
//		- *api.SendReuslt:
//		- error: the return error if any occurs
func (c *Client) SendMessage(content string) (*api.SendResult, error) {
	body, err := bce.NewBodyFromString(content)
	if err != nil {
		return nil, err
	}
	return api.SendMessage(c, body)
}