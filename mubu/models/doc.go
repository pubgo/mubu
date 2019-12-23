package models

import "encoding/json"

type Node struct {
	Collapsed bool   `json:"collapsed,omitempty"`
	Finish    bool   `json:"finish,omitempty"`
	ID        string `json:"id,omitempty"`
	Color     string `json:"color,omitempty"`
	Heading   string `json:"heading,omitempty"`
	Modified  int    `json:"modified,omitempty"`
	Text      string `json:"text,omitempty"`
	Children  []Node `json:"children,omitempty"`
	// bold italic underline
	// font-weight: bold; text-decoration: underline; font-style: italic;
	Class  []string               `json:"-"`
	Style  map[string]interface{} `json:"-"`
	Images []struct {
		ID  string `json:"id"`
		Oh  int    `json:"oh"`
		Ow  int    `json:"ow"`
		URI string `json:"uri"`
		W   int    `json:"w"`
	} `json:"images,omitempty"`
}

type Nodes struct {
	Nodes []Node `json:"nodes,omitempty"`
}

type GetDoc struct {
	Code int `json:"code"`
	Data struct {
		BaseVersion int    `json:"baseVersion"`
		Definition  string `json:"definition"`
		Name        string `json:"name"`
		Role        int    `json:"role"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func (t *GetDoc) Nodes() (nodes *Nodes, _ error) {
	if t.Data.Definition == "" {
		return nil, nil
	}
	nodes = new(Nodes)
	return nodes, json.Unmarshal([]byte(t.Data.Definition), nodes)
}

func (t *GetDoc) Title() string {
	return t.Data.Name
}

func (t *GetDoc) Version() int {
	return t.Data.BaseVersion
}

func (t *GetDoc) Role() int {
	return t.Data.Role
}

func (t *GetDoc) Markdown(level int) string {
	panic("not implemented")
}

func (t *GetDoc) Html() string {
	panic("not implemented")
}

type Dir struct {
	CommitClient string `json:"commitClient"`
	CreateTime   int    `json:"createTime"`
	DeleteTime   int    `json:"deleteTime"`
	Deleted      int    `json:"deleted"`
	Encrypted    int    `json:"encrypted"`
	FolderID     string `json:"folderId"`
	ID           string `json:"id"`
	MetaVersion  int    `json:"metaVersion"`
	Name         string `json:"name"`
	StarIndex    int    `json:"starIndex"`
	Stared       int    `json:"stared"`
	UpdateTime   int    `json:"updateTime"`
	UserID       int    `json:"userId"`
	Version      int    `json:"version"`
}

type Document struct {
	Deleted    int    `json:"deleted"`
	Encrypted  int    `json:"encrypted"`
	FolderID   string `json:"folderId"`
	ID         string `json:"id"`
	Name       string `json:"name"`
	Role       int    `json:"role"`
	ShareID    string `json:"shareId"`
	Stared     int    `json:"stared"`
	UpdateTime int    `json:"updateTime"`
	UserID     int    `json:"userId"`
	UserName   string `json:"userName"`
}

type Folder struct {
	CommitClient string `json:"commitClient"`
	CreateTime   int    `json:"createTime"`
	DeleteTime   int    `json:"deleteTime"`
	Deleted      int    `json:"deleted"`
	Encrypted    int    `json:"encrypted"`
	FolderID     string `json:"folderId"`
	ID           string `json:"id"`
	MetaVersion  int    `json:"metaVersion"`
	Name         string `json:"name"`
	StarIndex    int    `json:"starIndex"`
	Stared       int    `json:"stared"`
	UpdateTime   int    `json:"updateTime"`
	UserID       int    `json:"userId"`
	Version      int    `json:"version"`
}

type ListDoc struct {
	Code int `json:"code"`
	Data struct {
		Dir       []Dir      `json:"dir"`
		Documents []Document `json:"documents"`
		FolderID  string     `json:"folderId"`
		Folders   []Folder   `json:"folders"`
	} `json:"data"`
	Msg interface{} `json:"msg"`
}
