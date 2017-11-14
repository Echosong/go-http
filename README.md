# go-http
自动记录cookie 模拟请求 http 封装

## 下载安装

go get github.com/echosong/go-http


···
package main

import (
	"github.com/Echosong/go-http"
	"fmt"
	"net/url"
)

func main()  {
	client := httpes.Http{"http://www.blog.com/"}
	
	rep,_ := client.PostForm("login", url.Values{"username":{"admin"}, "password":{"songfeiok"}}) //client.Post("login", "username=admin&password=songfeiok")

	fmt.Printf("%s \n", rep.Body);
	fmt.Println(rep.StatusCode);

	rep1, _ :=client.Get("login")

	fmt.Printf("%s \n", rep1.Body);
	fmt.Println(rep1.StatusCode);

}

···
