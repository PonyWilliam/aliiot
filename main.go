// This file is auto-generated, don't edit it. Thanks.
package main

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	iot20180120 "github.com/alibabacloud-go/iot-20180120/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"encoding/base64"
	"github.com/gin-gonic/gin"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient (accessKeyId *string, accessKeySecret *string) (_result *iot20180120.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("iot.cn-shanghai.aliyuncs.com")
	_result = &iot20180120.Client{}
	_result, _err = iot20180120.NewClient(config)
	return _result, _err
}

func _main (content string) (_err error) {
	client, _err := CreateClient(tea.String("LTAI5tGF78Bm2ybxV3YEiyB8"), tea.String("s1Ic1TOyklUIS9FKuYIWMuP39bCAf0"))
	if _err != nil {
		return _err
	}

	pubRequest := &iot20180120.PubRequest{
		TopicFullName: tea.String("/a1rdpdDHgQI/window1/user/get"),
		MessageContent: tea.String(base64.StdEncoding.EncodeToString([]byte(content))),
		ProductKey: tea.String("a1rdpdDHgQI"),
		Qos: tea.Int32(0),
	}
	// 复制代码运行请自行打印 API 的返回值
	_, _err = client.Pub(pubRequest)
	if _err != nil {
		return _err
	}
	return _err
}


func main() {
	r := gin.Default()
	r.GET("/open",func(c *gin.Context){
		err := _main("open")
		if err != nil{
			c.JSON(200,gin.H{
				"message":"error",
			})
		}
		c.JSON(200,gin.H{
			"message":"open",
		})
	})
	r.GET("/close",func(c *gin.Context){
		err := _main("close")
		if err != nil{
			c.JSON(200,gin.H{
				"message":"error",
			})
		}
		c.JSON(200,gin.H{
			"message":"close",
		})
	})
	r.Run(":8080");
}
