package mubu

import (
	"encoding/json"
	"github.com/pubgo/g/xerror"
)

var (
	Err                             = xerror.NewXErr("mubu")
	ErrParametersIncorrectOrMissing = Err.New("parameters incorrect or missing")
	ErrAuthIncorrect                = Err.New("auth incorrect")
	ErrPermissionsMissing           = Err.New("permissions missing")
	ErrNotFound                     = Err.New("not found")
	ErrServer                       = Err.New("server error")
	ErrRateLimitOver                = Err.New("too many requests")
	ErrUnknownCode                  = Err.New("unknown code")

	ErrJsonUnmarshal = Err.New("json unmarshal error")
)

func checkCode(code int) error {
	switch code {
	case 400:
		return ErrParametersIncorrectOrMissing
	case 401:
		return ErrAuthIncorrect
	case 403:
		return ErrPermissionsMissing
	case 404:
		return ErrNotFound
	case 500:
		return ErrServer
	case 429:
		return ErrRateLimitOver
	default:
		return ErrUnknownCode
	}
}

type errMsg struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func decode2Err(err xerror.IErr, data string) {
	var m errMsg
	xerror.Exit(json.Unmarshal([]byte(data), &m), ErrJsonUnmarshal)
	err.M("status", m.Status)
	err.M("code", m.Code)
	err.M("message", m.Message)
}
