package config

import "go_learn/pkg/config"

func init() {
	config.Add("sms", func() map[string]interface{} {
		return map[string]interface{}{
			// 默认阿里云
			"aliyun": map[string]interface{}{
				"access_key_id":     config.Env("SMS_ALIYUN_ACCESS_ID"),
				"access_key_secret": config.Env("SMS_ALIYUN_ACCESS_SECRET"),
				"sign_name":         config.Env("SMS_ALIYUN_SIGN_NAME"),
				"template_code":     config.Env("SMS_ALIYUN_TEMPLATE_CODE"),
			},
		}
	})
}
