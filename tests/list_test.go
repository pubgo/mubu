package tests

import (
	"fmt"
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/mubu/mubu"
	"github.com/pubgo/mubu/mubu/models"
	"testing"
)

func init() {
	xerror.Panic(xenv.LoadFile("../.env"))
}

func TestLogin(t *testing.T) {
	mb := mubu.New()
	xerror.Debug(mb.API().Login(xenv.GetEnv("phone"), xenv.GetEnv("password")))
	//_, err := mb.API().ListDoc("i-LGHOgX7", "time", "", "")
	//xerror.P(err, "")
	//fmt.Printf("%#v", xerror.PanicErr(mb.API().ListDoc("i-LGHOgX7", "time", "", "")))
	ld := xerror.PanicErr(mb.API().ListDoc("", "time", "", "")).(*models.ListDoc)
	fmt.Println(ld.Data.Dir)
	fmt.Println(ld.Data.Documents)
	fmt.Println(ld.Data.FolderID)
	fmt.Println(ld.Data.Folders)
}

func TestWalk(t *testing.T) {
	mb := mubu.New()
	xerror.Debug(mb.API().Login(xenv.GetEnv("phone"), xenv.GetEnv("password")))
	xerror.Exit(mb.API().Walk("", "time", "", "", func(doc *models.ListDoc, err error) error {
		if err != nil {
			return err
		}

		fmt.Println(doc.Data.FolderID, "test Walk")
		return nil
	}), "Walk error")
}

func TestDoc(t *testing.T) {
	mb := mubu.New()
	xerror.Debug(mb.API().Login(xenv.GetEnv("phone"), xenv.GetEnv("password")))
	//_, err := mb.API().ListDoc("i-LGHOgX7", "time", "", "")
	//xerror.P(err, "")
	//fmt.Printf("%#v", xerror.PanicErr(mb.API().ListDoc("i-LGHOgX7", "time", "", "")))
	ld := xerror.PanicErr(mb.API().GetDoc("9Gx9BDBXD")).(*models.GetDoc)
	fmt.Println(ld.Title())
	fmt.Println(ld.Version())
	fmt.Println(ld.Role())
	fmt.Println(ld.Nodes())
}
