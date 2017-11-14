package main

import (
	"github.com/Echosong/go-http"
	"fmt"
	"regexp"
	"strings"
	"net/url"
)

func main() {
	client := httpes.Http{"https://github.com/"}

	rep, _ := client.Get("login");
	var token = ""
	if rep.StatusCode == 200 {
		html := string(rep.Body[:]);
		reg := regexp.MustCompile("<input name=\"authenticity_token\" type=\"hidden\" value=\"(\\S*)\" />")
		regStr := reg.FindString(html);
		str := strings.Replace(regStr, "<input name=\"authenticity_token\" type=\"hidden\" value=\"", "", -1)
		token = strings.Replace(str, "\" />", "", -1)
	} else {
		fmt.Println(rep.StatusCode)
		return;
	}
	//login
	repSession, _ := client.PostForm("session", url.Values{"authenticity_token": {token}, "login": {"echosong"}, "password": {"Songfeiok123"}});
	if repSession.StatusCode == 200 {
		//be sure
		repIndex, _ := client.Get("settings/admin");
		fmt.Printf("%s", repIndex.Body)
	} else {
		fmt.Println("fail");
	}

}
