package models



// https://mubu.com/api/document/get
// docId: 7Ilk96_PD
type Doc struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}
type Data2 struct {
	Role        int    `json:"role"`
	BaseVersion int    `json:"baseVersion"`
	Name        string `json:"name"`
	Definition  string `json:"definition"`
}