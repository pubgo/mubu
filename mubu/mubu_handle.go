package mubu


import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"net/http"
)

const BaseUrlPrefix = "/api"

func yqPut(c *resty.Request, url string, body interface{}, dt interface{}) (err error) {
	defer xerror.RespErr(&err)

	url = BaseUrlPrefix + url

	c.SetHeader("Content-Type", "application/json")

	if body != nil {
		c.SetBody(body)
	}

	resp := xerror.PanicErr(c.Put(url)).(*resty.Response)
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

func yqDelete(c *resty.Request, url string, queryParams map[string]string, dt interface{}) (err error) {
	defer xerror.RespErr(&err)

	url = BaseUrlPrefix + url

	if queryParams != nil {
		c.SetQueryParams(queryParams)
	}

	resp := xerror.PanicErr(c.Delete(url)).(*resty.Response)
	xerror.PanicTT(resp.StatusCode() != http.StatusOK, func(err xerror.IErr) {
		err.SetErr(checkCode(resp.StatusCode()))
		err.M("url", c.URL)
		decode2Err(err, resp.String())
		if queryParams != nil {
			err.M("query_params", queryParams)
		}
	})
	xerror.PanicM(json.Unmarshal(resp.Body(), dt), "%s resp decode error", url)
	return
}

func yqGet(c *resty.Request, url string, queryParams map[string]string, dt interface{}) (err error) {
	defer xerror.RespErr(&err)

	url = BaseUrlPrefix + url

	if queryParams != nil {
		c.SetQueryParams(queryParams)
	}

	resp := xerror.PanicErr(c.Get(url)).(*resty.Response)
	xerror.PanicTT(resp.StatusCode() != http.StatusOK, func(err xerror.IErr) {
		err.SetErr(checkCode(resp.StatusCode()))
		err.M("url", c.URL)
		decode2Err(err, resp.String())
		if queryParams != nil {
			err.M("query_params", queryParams)
		}

	})
	xerror.PanicM(json.Unmarshal(resp.Body(), dt), "%s resp decode error", url)

	return
}

func yqPost(c *resty.Request, url string, body interface{}, dt interface{}) (err error) {
	defer xerror.RespErr(&err)

	url = BaseUrlPrefix + url

	c.SetHeader("Content-Type", "application/json")

	if body != nil {
		c.SetBody(body)
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
