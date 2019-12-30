package mubu

import (
	"github.com/pubgo/g/xcmd"
	"github.com/pubgo/g/xerror"
	"io/ioutil"
	"strings"
)

func Env() *xcmd.Command {
	_name := ".env"
	_data := `
env_prefix=project name 
phone=!{Encrypted phone}
password=!{Encrypted password}
`
	return &xcmd.Command{
		Use:   "env",
		Short: "create env file",
		RunE: func(cmd *xcmd.Command, args []string) (err error) {
			defer xerror.RespErr(&err)

			if len(args) > 0 {
				_name = args[0] + _name
			}

			xerror.Panic(ioutil.WriteFile(_name, []byte(strings.TrimSpace(_data)), 0644))
			return
		},
	}
}
