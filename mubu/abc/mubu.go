package abc

import (
	"github.com/pubgo/mubu/mubu/models"
	"net/http"
)

type IMubuAPI interface {
	// 登陆获取Cookie
	Login(phone, password string) ([]*http.Cookie, error)
	// 文件列表 , sort=time
	ListDoc(folderId, sort, keywords, source string) (*models.ListDoc, error)
	// 列出所有的文件和目录
	Walk(folderId, sort, keywords, source string, fn func(*models.ListDoc, error) error) error
	// 获取文件
	GetDoc(docId string) (*models.GetDoc, error)
	// 创建分享链接
	CreateLink(docId string) (*models.CreateLink, error)
	// 关闭分享链接
	CloseLink(docId string) (*models.CloseLink, error)
	// 刷新链接
	RefreshLink(docId string) (*models.RefreshLink, error)
}

type IMubu interface {
	API() IMubuAPI
}
