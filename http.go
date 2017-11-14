package http

import (
	"net/http"
	"net/url"
	"strings"
	"io/ioutil"
)

type Jar struct {
	cookies []*http.Cookie
}

func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.cookies = cookies
}
func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
	return jar.cookies
}

type Http struct {
	Host string
}

type ResponseBody struct {
	Body []byte
	StatusCode int
}

var client = http.Client{nil, nil, new(Jar), 99999999999992}

func (http *Http) Post(url string, body string) (resp *ResponseBody, err error)   {

	rep, err := client.Post(http.Host+url, "application/x-www-form-urlencoded", strings.NewReader(body));
	if err == nil{
		resp.Body ,_= ioutil.ReadAll(rep.Body)
		resp.StatusCode = rep.StatusCode
		defer rep.Body.Close()
		return  resp, nil
	}else{
		return nil, err
	}
}

func (http *Http) Get(url string)(resp *ResponseBody, err error)  {
	rep, err := client.Get(http.Host + url);
	if err == nil{
		resp.Body ,_= ioutil.ReadAll(rep.Body)
		resp.StatusCode = rep.StatusCode
		defer rep.Body.Close()
		return  resp, nil
	}else{
		return nil, err
	}
}

func  (http *Http) PostForm(url string, values url.Values)(resp *ResponseBody, err error)  {
	rep, err := client.PostForm(http.Host+ url,values)
	if err == nil{
		resp.Body ,_= ioutil.ReadAll(rep.Body)
		resp.StatusCode = rep.StatusCode
		defer rep.Body.Close()
		return  resp, nil
	}else{
		return nil, err
	}
}
