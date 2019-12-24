package mubu

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/mubu/mubu/abc"
	"time"
)

func WithRetry(retryCount, retryWaitTime, retryMaxWaitTime int) func(*mubu) {
	return func(m *mubu) {
		m.retryCount = retryCount
		m.retryWaitTime = retryWaitTime
		m.retryMaxWaitTime = retryMaxWaitTime
	}
}

func WithTimeout(timeout int) func(*mubu) {
	return func(m *mubu) {
		m.timeout = timeout
	}
}

type mubu struct {
	retryCount       int  // default 3
	retryWaitTime    int  // default 5 second
	retryMaxWaitTime int  // default 20 second
	timeout          int  // default 60 second
	debug            bool // default true
	client           *resty.Client
}

func (t *mubu) API() abc.IMubuAPI {
	return &mubuImpl{c: t.client}
}

func (t *mubu) _init() {
	if t.retryCount < 1 {
		t.retryCount = 3
	}

	if t.retryWaitTime < 1 {
		t.retryWaitTime = 5
	}

	if t.retryMaxWaitTime < 1 {
		t.retryMaxWaitTime = 20
	}

	if t.timeout < 1 {
		t.timeout = 60
	}

	t.debug = xenv.IsDebug()
	t.client = resty.New().
		SetDebug(t.debug).
		SetContentLength(true).
		SetHostURL("http://mubu.com").
		SetRetryCount(t.retryCount).
		SetRetryWaitTime(time.Second * time.Duration(t.retryWaitTime)).
		SetRetryMaxWaitTime(time.Second * time.Duration(t.retryMaxWaitTime)).
		SetTimeout(time.Second * time.Duration(t.timeout))
}

func New(fn ...func(*mubu)) abc.IMubu {
	_mb := &mubu{}
	for _, f := range fn {
		f(_mb)
	}
	_mb._init()
	return _mb
}

var _url = func(url string) func(...interface{}) string {
	return func(params ...interface{}) string {
		return fmt.Sprintf(url, params...)
	}
}
