package mubu

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"net/http"
)

func _post(c *resty.Request, url string, body map[string]string, dt interface{}) (err error) {
	defer xerror.RespErr(&err)

	c.SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	if body != nil {
		c.SetFormData(body)
	}

	resp := xerror.PanicErr(c.Post(url)).(*resty.Response)
	xerror.PanicTT(resp.StatusCode() != http.StatusOK, func(err xerror.IErr) {
		err.SetErr(checkCode(resp.StatusCode()))
		err.M("url", c.URL)
		decode2Err(err, resp.String())
		if body != nil {
			err.M("body", body)
		}
	})
	xerror.PanicM(json.Unmarshal(resp.Body(), dt), "%s resp decode error", url)
	return
}
