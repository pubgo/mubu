package abc

import (
	"github.com/pubgo/mubu/mubu/models"
	"net/http"
)

type IMubu interface {
	// 登陆获取Cookie
	Login(phone, password string) []*http.Cookie
	// 文件列表 , sort=time
	ListDoc(folderId, sort, keywords, source string) *models.ListDoc
	// 获取文件
	GetDoc(docId string) *models.GetDoc
	// 创建分享链接
	CreateLink(docId string) *models.CreateLink
	// 关闭分享链接
	CloseLink(docId string) *models.CloseLink
	// 刷新链接
	RefreshLink(docId string) *models.RefreshLink
}
