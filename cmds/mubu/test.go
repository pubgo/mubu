package mubu

import (
	"github.com/pubgo/g/xcmd"
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/mubu/mubu"
	"github.com/pubgo/mubu/mubu/models"
)

func Test() *xcmd.Command {
	return &xcmd.Command{
		Use:   "test",
		Short: "test login and list first page",
		RunE: func(cmd *xcmd.Command, args []string) (err error) {
			defer xerror.RespErr(&err)

			xerror.Panic(xenv.LoadFile(".env"))

			mb := mubu.New()
			xerror.PanicErr(mb.API().Login(xenv.GetEnv("phone"), xenv.GetEnv("password")))
			xerror.Debug(xerror.PanicErr(mb.API().ListDoc("", "time", "", "")).(*models.ListDoc))
			return
		},
	}
}
