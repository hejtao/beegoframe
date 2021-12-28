package msg

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

const (
	// 阿里云Short Message Service
	regionId     = "cn-hangzhou"
	accessKeyId  = "xxxx"
	accessSecret = "xxxx"
	scheme       = "https"

	SignName                  = "xxxx"
	DomesticTemplateCode      = "xxxx"
	InternationalTemplateCode = "xxxx"
)

func Send(phone, signName, templateCode, templateParam string) error {
	client, err := dysmsapi.NewClientWithAccessKey(regionId, accessKeyId, accessSecret)
	if err != nil {
		return err
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = scheme
	request.PhoneNumbers = phone
	request.SignName = signName
	request.TemplateCode = templateCode
	request.TemplateParam = templateParam

	_, err = client.SendSms(request)
	return err
}
