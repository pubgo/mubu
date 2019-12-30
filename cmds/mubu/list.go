package mubu

import (
	"github.com/c-bata/go-prompt"
	"github.com/pubgo/g/xcmd"
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/mubu/mubu"
	"github.com/pubgo/mubu/mubu/models"
	"os"
	"strings"
)

func List() *xcmd.Command {
	var _ps []prompt.Suggest
	var _prompt = "> "

	return &xcmd.Command{
		Use:   "list",
		Short: "list doc and directory",
		RunE: func(cmd *xcmd.Command, args []string) (err error) {
			defer xerror.RespErr(&err)

			xerror.Panic(xenv.LoadFile(".env"))

			mb := mubu.New()
			xerror.PanicErr(mb.API().Login(xenv.GetEnv("phone"), xenv.GetEnv("password")))

			xerror.Panic(mb.API().Walk("", "time", "", "", func(doc *models.ListDoc, e error) error {
				if err != nil {
					return err
				}

				_ps = _ps[:0]
				for _, d := range doc.Data.Documents {
					_ps = append(_ps, prompt.Suggest{Text: d.Name})
				}

				for {
					_d := prompt.Input(_prompt, func(d prompt.Document) []prompt.Suggest {
						return prompt.FilterContains(_ps, d.GetWordBeforeCursor(), true)
					})
					if strings.TrimSpace(_d) == "" {
						break
					}

					if strings.TrimSpace(_d) == "exit" {
						os.Exit(-1)
					}
				}
				return nil
			}))
			return
		},
	}
}
