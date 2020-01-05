package models

import (
	"encoding/json"
	"fmt"
	"github.com/pubgo/g/xerror"
	"regexp"
	"strings"
)

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
	nodes    Nodes
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

func (t *GetDoc) Markdown() string {
	t.markdown = fmt.Sprintf("# %s\n\n", t.Title())
	for _, node := range xerror.PanicErr(t.Nodes()).(*Nodes).Nodes {
		t.markdown += t._Markdown(0, node) + "\n\n"
	}
	return t.markdown
}

var trim = strings.TrimSpace
var replaceAll = strings.ReplaceAll
var classRex = regexp.MustCompile(`class="(.+)"`)
var valueRex = regexp.MustCompile(`.*>(.+)<`)

func (t *GetDoc) _Markdown(level int, node Node) string {
	node.Text = replaceAll(trim(node.Text), "\\", "")
	if node.Text == "" {
		return ""
	}

	_markdown := ""

	if classRex.MatchString(node.Text) {
		_text1 := valueRex.FindStringSubmatch(node.Text)
		_text := _text1[len(_text1)-1]
		for _, c := range strings.Split(classRex.FindStringSubmatch(node.Text)[1], " ") {

			switch c {
			case "bold":
				node.Text = fmt.Sprintf("**%s**", _text)
			case "italic":
				node.Text = fmt.Sprintf("*%s*", _text)
			case "underline":
				node.Text = fmt.Sprintf("_%s_", _text)
			}
		}
	}

	if node.Finish {
		node.Text = fmt.Sprintf("~~%s~~", node.Text)
	}

	if node.Heading == 0 {
		level += 1
		_markdown = fmt.Sprintf("%s1. %s",strings.Repeat("  ", level), node.Text)
	} else {
		level = 0
		_markdown = fmt.Sprintf("%s %s\n", strings.Repeat("#", node.Heading), node.Text)
	}

	if node.Images != nil && len(node.Images) != 0 {
		_markdown += "\n\n"
		_markdown += fmt.Sprintf("![image](https://mubu.com/%s)\n\n", node.Images[0].URI)
	}

	for _, n := range node.Children {
		_a := t._Markdown(level, n)
		if _a == "" {
			continue
		}

		_markdown += "\n"
		_markdown += _a
	}

	return _markdown
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
