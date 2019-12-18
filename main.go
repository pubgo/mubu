package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func main() {

	ss := resty.New().
		SetContentLength(true).
		SetHostURL("http://mubu.com")
	resp, _ := ss.R().SetFormData(map[string]string{
		"phone":    "",
		"password": "",
	}).Post("/api/login/submit")
	fmt.Println(resp.Header())
	fmt.Println(resp.String())
	//cmds.Execute()
}
