package models

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/pubgo/g/xerror"
	"gopkg.in/russross/blackfriday.v2"
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
	Msg   string `json:"msg"`
	nodes Nodes
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
	markdown := fmt.Sprintf("# %s\n", t.Title())
	for _, node := range xerror.PanicErr(t.Nodes()).(*Nodes).Nodes {
		_md := trim(t._Markdown(0, node))
		if _md == "" {
			continue
		}

		markdown += _md + "\n\n"
	}
	return trim(markdown)
}

var trim = strings.TrimSpace
var printf = fmt.Sprintf
var replaceAll = strings.ReplaceAll
var classRex = regexp.MustCompile(`class="(.+)"`)

func (t *GetDoc) _Markdown(level int, node Node) string {
	_markdown := ""
	_text := replaceAll(trim(node.Text), "\\", "")
	if _text == "" {
		return _markdown
	}

	if classRex.MatchString(_text) {
		_text1 := xerror.PanicErr(goquery.NewDocumentFromReader(strings.NewReader(_text))).(*goquery.Document).Text()
		for _, c := range strings.Split(classRex.FindStringSubmatch(_text)[1], " ") {
			switch c {
			case "bold":
				_text1 = printf("**%s**", _text1)
			case "italic":
				_text1 = printf("*%s*", _text1)
			case "underline":
				_text1 = printf("_%s_", _text1)
			}
		}
		_text = _text1
	}

	if node.Finish {
		_text = printf("~~%s~~", _text)
	}

	if node.Heading == 0 {
		level += 1
		_markdown = printf("%s1. %s\n", strings.Repeat("  ", level), _text)

		if node.Images != nil && len(node.Images) != 0 {
			_markdown += printf("%s![image](https://mubu.com/%s)\n", strings.Repeat("  ", level+1), node.Images[0].URI)
		}
	} else {
		level = 0
		_markdown = printf("%s %s\n\n", strings.Repeat("#", node.Heading), _text)
		if node.Images != nil && len(node.Images) != 0 {
			_markdown += printf("%s![image](https://mubu.com/%s)\n\n", strings.Repeat("  ", level+1), node.Images[0].URI)
		}
	}

	if len(node.Children) > 0 {
		for _, n := range node.Children {
			_a := t._Markdown(level, n)
			if _a == "" {
				continue
			}
			_markdown += _a
		}
	}

	return _markdown
}

func (t *GetDoc) Html() string {
	_m1 := blackfriday.Run([]byte(t.Markdown()))
	return string(_m1)
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
