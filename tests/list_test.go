package tests

import (
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/mubu/mubu"
	"testing"
)

func init() {
	xerror.Panic(xenv.LoadFile("../.env"))
}

func TestLogin(t *testing.T) {
	mb := mubu.New()
	xerror.Debug(mb.API().Login(xenv.GetEnv("phone"), xenv.GetEnv("password")))
	_, err := mb.API().ListDoc("", "time", "", "")
	xerror.P(err, "")
}
