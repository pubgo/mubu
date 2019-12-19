package mubu

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/mubu/mubu/abc"
	"github.com/pubgo/mubu/mubu/models"
	"net/http"
)

var (
	_            abc.IMubu = (*mubuImpl)(nil)
	_login                 = _url("/login/submit")
	_getDoc                = _url("/document/get")
	_getList               = _url("/list/get")
	_createLink            = _url("/document/create_link")
	_closeLink             = _url("/document/close_link")
	_refreshLink           = _url("/document/refresh_link")
)

type mubuImpl struct {
	c *resty.Client
}

func (t *mubuImpl) Login(phone, password string) []*http.Cookie {
	panic("implement me")
}

func (t *mubuImpl) ListDoc(folderId, sort, keywords, source string) *models.ListDoc {
	panic("implement me")
}

func (t *mubuImpl) GetDoc(docId string) *models.GetDoc {
	panic("implement me")
}

func (t *mubuImpl) CreateLink(docId string) *models.CreateLink {
	panic("implement me")
}

func (t *mubuImpl) CloseLink(docId string) *models.CloseLink {
	panic("implement me")
}

func (t *mubuImpl) RefreshLink(docId string) *models.RefreshLink {
	panic("implement me")
}

func (t *mubuImpl) Init() {
}
