package mubu

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xenv"
	"time"
)

type mubu struct {
	RetryCount       int  // default 3
	RetryWaitTime    int  // default 5 second
	RetryMaxWaitTime int  // default 20 second
	Timeout          int  // default 60 second
	Debug            bool // default true
	client           *resty.Client
}

func (t *mubu) _init() {
	if t.RetryCount < 1 {
		t.RetryCount = 3
	}

	if t.RetryWaitTime < 1 {
		t.RetryWaitTime = 5
	}

	if t.RetryMaxWaitTime < 1 {
		t.RetryMaxWaitTime = 20
	}

	if t.Timeout < 1 {
		t.Timeout = 60
	}

	t.Debug = xenv.IsDebug()

	t.client = resty.New().
		SetDebug(t.Debug).
		SetContentLength(true).
		SetHostURL("http://mubu.com").
		SetRetryCount(t.RetryCount).
		SetRetryWaitTime(time.Second * time.Duration(t.RetryWaitTime)).
		SetRetryMaxWaitTime(time.Second * time.Duration(t.RetryMaxWaitTime)).
		SetTimeout(time.Second * time.Duration(t.Timeout))
}

func New() *mubu {
	_mb := &mubu{}
	_mb._init()
	return _mb
}
