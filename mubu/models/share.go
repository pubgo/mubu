package models

type CreateLink struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data struct {
		Version    int64  `json:"version"`
		ShareID    string `json:"shareId"`
		SetVersion bool   `json:"setVersion"`
		SetShareID bool   `json:"setShareId"`
	} `json:"data"`
}

type CloseLink struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data struct {
		ShareID interface{} `json:"shareId"`
		Version int64       `json:"version"`
	} `json:"data"`
}

type RefreshLink struct {
	Code int `json:"code"`
	Data struct {
		SetShareID bool   `json:"setShareId"`
		SetVersion bool   `json:"setVersion"`
		ShareID    string `json:"shareId"`
		Version    int    `json:"version"`
	} `json:"data"`
	Msg interface{} `json:"msg"`
}
