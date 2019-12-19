package mubu

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/mubu/mubu/abc"
	"github.com/pubgo/mubu/mubu/models"
	"net/http"
)

var (
	_            abc.IMubu = (*mubuImpl)(nil)
	_login                 = _url("/api/login/submit")
	_getDoc                = _url("/api/document/get")
	_getList               = _url("/api/list/get")
	_createLink            = _url("/api/document/create_link")
	_closeLink             = _url("/api/document/close_link")
	_refreshLink           = _url("/api/document/refresh_link")
)

type mubuImpl struct {
	c *resty.Client
}

func (t *mubuImpl) Login(phone, password string) (cookie []*http.Cookie, err error) {
	defer xerror.RespErr(&err)

	_req := t.c.R()

	_req.SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	_req.SetFormData(map[string]string{"phone": phone, "password": password})

	resp := xerror.PanicErr(_req.Post(_login())).(*resty.Response)
	xerror.PanicTT(resp.StatusCode() != http.StatusOK, func(err xerror.IErr) {
		err.SetErr(checkCode(resp.StatusCode()))
		err.M("url", _req.URL)
		decode2Err(err, resp.String())
	})

	return resp.Cookies(), nil
}

func (t *mubuImpl) ListDoc(folderId, sort, keywords, source string) (*models.ListDoc, error) {
	panic("implement me")
}

func (t *mubuImpl) GetDoc(docId string) (*models.GetDoc, error) {
	panic("implement me")
}

func (t *mubuImpl) CreateLink(docId string) (*models.CreateLink, error) {
	panic("implement me")
}

func (t *mubuImpl) CloseLink(docId string) (*models.CloseLink, error) {
	panic("implement me")
}

func (t *mubuImpl) RefreshLink(docId string) (*models.RefreshLink, error) {
	panic("implement me")
}
