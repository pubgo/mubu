package models

// https://mubu.com/api/list/get
// folderId=i-LGHOgX7&sort=time&keywords=&source=
// content-type: application/x-www-form-urlencoded; charset=UTF-8



type FolderList struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data Data        `json:"data"`
}
type Folders struct {
	ID           string `json:"id"`
	FolderID     string `json:"folderId"`
	Name         string `json:"name"`
	UserID       int    `json:"userId"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
	Deleted      int    `json:"deleted"`
	DeleteTime   int    `json:"deleteTime"`
	Version      int64  `json:"version"`
	CommitClient string `json:"commitClient"`
	Encrypted    int    `json:"encrypted"`
	Stared       int    `json:"stared"`
	StarIndex    int    `json:"starIndex"`
	MetaVersion  int64  `json:"metaVersion"`
}
type Documents struct {
	Deleted    int    `json:"deleted"`
	Role       int    `json:"role"`
	Encrypted  int    `json:"encrypted"`
	Name       string `json:"name"`
	ShareID    string `json:"shareId,omitempty"`
	UpdateTime int64  `json:"updateTime"`
	ID         string `json:"id"`
	Stared     int    `json:"stared"`
	UserName   string `json:"userName"`
	UserID     int    `json:"userId"`
	FolderID   string `json:"folderId"`
}
type Dir struct {
	ID           string `json:"id"`
	FolderID     string `json:"folderId"`
	Name         string `json:"name"`
	UserID       int    `json:"userId"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
	Deleted      int    `json:"deleted"`
	DeleteTime   int    `json:"deleteTime"`
	Version      int64  `json:"version"`
	CommitClient string `json:"commitClient"`
	Encrypted    int    `json:"encrypted"`
	Stared       int    `json:"stared"`
	StarIndex    int    `json:"starIndex"`
	MetaVersion  int64  `json:"metaVersion"`
}
type Data struct {
	Folders   []Folders   `json:"folders"`
	Documents []Documents `json:"documents"`
	Dir       []Dir       `json:"dir"`
	FolderID  string      `json:"folderId"`
}


// /api/login/submit
//phone
// password
