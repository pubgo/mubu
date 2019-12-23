package models

import (
	"encoding/json"
	"fmt"
	"github.com/pubgo/g/xerror"
	"regexp"
	"strings"
)

var classRex = regexp.MustCompile(`<span class="(.+)">`)

type Node struct {
	Collapsed bool   `json:"collapsed,omitempty"`
	Finish    bool   `json:"finish,omitempty"`
	ID        string `json:"id,omitempty"`
	Color     string `json:"color,omitempty"`
	Heading   int    `json:"heading,omitempty"`
	Modified  int    `json:"modified,omitempty"`
	Text      string `json:"text,omitempty"`
	Children  []Node `json:"children,omitempty"`
	// bold italic underline
	// font-weight: bold; font-style: italic; text-decoration: underline; text-decoration: line-through;
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
	Msg      string `json:"msg"`
	markdown string
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
	t.markdown = fmt.Sprintf("# %s\n\n", t.Title())
	for _, node := range xerror.PanicErr(t.Nodes()).(*Nodes).Nodes {
		t.markdown += t._Markdown(0, node)
	}
	return t.markdown
}

/*
	Text      string `json:"text,omitempty"`
	Children  []Node `json:"children,omitempty"`
	Images []struct {
		ID  string `json:"id"`
		Oh  int    `json:"oh"`
		Ow  int    `json:"ow"`
		URI string `json:"uri"`
		W   int    `json:"w"`
	} `json:"images,omitempty"`
*/

func (t *GetDoc) _Markdown(level int, node Node) string {
	node.Text = strings.TrimSpace(node.Text)
	node.Text = strings.ReplaceAll(node.Text, "\\", "")
	if node.Text == "" {
		return "\n"
	}

	_markdown := ""

	if classRex.MatchString(node.Text) {
		for _, c := range classRex.FindStringSubmatch(node.Text)[1:] {
			switch c {
			case "bold":
				node.Text = fmt.Sprintf("**%s**", node.Text)
			case "italic":
				node.Text = fmt.Sprintf("_%s_", node.Text)
			case "underline":
				node.Text = fmt.Sprintf("__%s__", node.Text)
			}
		}
	}

	if node.Finish {
		node.Text = fmt.Sprintf("~~%s~~", node.Text)
	}

	if level < 4 {
		node.Text = strings.TrimSpace(fmt.Sprintf("%s %s", strings.Repeat("#", node.Heading), node.Text))
	} else {

	}

	if node.Images != nil && len(node.Images) != 0 {
		_markdown += fmt.Sprintf("![image](%s)", node.Images[0].URI)
	}

	return ""
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
