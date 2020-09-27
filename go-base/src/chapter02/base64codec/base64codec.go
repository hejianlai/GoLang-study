package main

import (
	`encoding/base64`
	`fmt`
)

func main()  {
	// 需要处理的字符串
	// message := "Away from keyboard.https://golang.org/"
	message := "123"
	// 编码消息
	encodeMessage := base64.StdEncoding.EncodeToString([]byte(message))
	// 输出编码完成的消息
	fmt.Println(encodeMessage)
	// 解码消息
	data,err := base64.StdEncoding.DecodeString(encodeMessage)
	// 出错处理
	if err != nil{
		fmt.Println(err)

	} else {
		// 打印解码完成的数据
		fmt.Println(string(data))
	}
}
