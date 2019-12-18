package models


// https://mubu.com/api/document/create_link
// docId: g9-sXP3WD

type CreateShare struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data struct {
		Version    int64  `json:"version"`
		ShareID    string `json:"shareId"`
		SetVersion bool   `json:"setVersion"`
		SetShareID bool   `json:"setShareId"`
	} `json:"data"`
}


// https://mubu.com/api/document/close_link
// docId: g9-sXP3WD
type CloseShare struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data struct {
		ShareID interface{} `json:"shareId"`
		Version int64       `json:"version"`
	} `json:"data"`
}



// https://mubu.com/api/export/file
//
