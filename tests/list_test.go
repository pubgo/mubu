package tests

import (
	"encoding/json"
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
	//fmt.Println(ld.Title())
	//fmt.Println(ld.Version())
	//fmt.Println(ld.Role())
	//fmt.Println(ld.Nodes())
	fmt.Println(ld.Markdown())
}

func TestMode(t *testing.T) {
	_dt := []byte(`{"code":0,"msg":"Success","data":{"role":31,"baseVersion":101,"name":"测试","definition":"{\"nodes\":[{\"heading\":1,\"id\":\"2f316f6c790c6e0a3\",\"modified\":1578072873383,\"text\":\"\"},{\"children\":[{\"children\":[{\"children\":[{\"color\":\"#dc2d1e\",\"id\":\"31116f31f23c880912\",\"modified\":1577090951704,\"text\":\"三生三世\"}],\"heading\":3,\"id\":\"13e16f31f23c881261\",\"modified\":1577090951704,\"text\":\"heade3\"}],\"heading\":2,\"id\":\"15a16f31f23a4e118\",\"modified\":1577090956425,\"text\":\"heade2\"}],\"collapsed\":false,\"finish\":false,\"heading\":1,\"id\":\"fd16f31ef4491159\",\"modified\":1578072868987,\"text\":\"<span class=\\\"bold\\\">heade1</span>\"},{\"children\":[{\"children\":[{\"id\":\"26916f31f229fb1461\",\"modified\":1577090953772,\"text\":\"<span class=\\\"bold\\\">三生三世</span>\"}],\"collapsed\":false,\"finish\":false,\"heading\":3,\"id\":\"2216f31f04b90192\",\"modified\":1577090951675,\"text\":\"heade3\"}],\"collapsed\":false,\"finish\":false,\"heading\":2,\"id\":\"18016f31ef56ee14a\",\"modified\":1578142929193,\"note\":\"\",\"text\":\"<span class=\\\"bold\\\">heade2</span>\"},{\"children\":[{\"heading\":0,\"id\":\"27e16f31f02b3e146\",\"modified\":1577090821909,\"text\":\"<span class=\\\"bold underline\\\">三生三世</span>\"}],\"heading\":3,\"id\":\"2a216f31ef5f4a124\",\"modified\":1577090777596,\"text\":\"<span class=\\\"bold\\\">heade3</span>\"},{\"children\":[{\"color\":\"#dc2d1e\",\"id\":\"9916f754d531d156\",\"modified\":1578221000217,\"text\":\"hello\"}],\"color\":\"#dc2d1e\",\"heading\":0,\"id\":\"20516f31efd3f2076\",\"modified\":1577090798782,\"text\":\"三生三世\"},{\"children\":[{\"children\":[],\"collapsed\":false,\"color\":\"#3da8f5\",\"finish\":false,\"heading\":2,\"id\":\"3e016f3267062708b\",\"images\":[{\"id\":\"1d916f3267c0b118a-40263\",\"oh\":1004,\"ow\":742,\"uri\":\"document_image/7fabd28a-8c59-4ffe-b9f3-ab2ef4c91549-40263.jpg\",\"w\":87}],\"modified\":1577098747251,\"text\":\"<span class=\\\"bold italic underline\\\">测试图片</span>\"},{\"children\":[{\"color\":\"#3da8f5\",\"heading\":2,\"id\":\"2cc16f328574a3155\",\"modified\":1577100573992,\"text\":\"ssss\"},{\"color\":\"#dc2d1e\",\"heading\":2,\"id\":\"3e016f32856b9017e\",\"modified\":1577100602455,\"text\":\"<a class=\\\"content-link\\\" target=\\\"_blank\\\" rel=\\\"noreferrer\\\" href=\\\"https://mubu.com/doclcoXBPA2D\\\"><span class=\\\"bold italic\\\">https://mubu.com/doclcoXBPA2D</span></a>\"},{\"color\":\"#dc2d1e\",\"heading\":0,\"id\":\"34a16f3289c6a7118\",\"images\":[{\"id\":\"37d16f3289cb16101\",\"oh\":1004,\"ow\":742,\"uri\":\"document_image/7fabd28a-8c59-4ffe-b9f3-ab2ef4c91549-40263.jpg\",\"w\":87}],\"modified\":1577100887822,\"text\":\"<a class=\\\"content-link\\\" target=\\\"_blank\\\" rel=\\\"noreferrer\\\" href=\\\"https://mubu.com/doclcoXBPA2D\\\"><span class=\\\"bold italic\\\">https://mubu.com/doclcoXBPA2D</span></a>\"},{\"color\":\"#dc2d1e\",\"heading\":2,\"id\":\"37116f32861a4f117\",\"modified\":1577100647111,\"text\":\"<span class=\\\"italic\\\">sss</span>\"}],\"color\":\"#3da8f5\",\"heading\":2,\"id\":\"1a816f3269677b199\",\"modified\":1577098773570,\"text\":\"是是是\"}],\"collapsed\":false,\"color\":\"#333333\",\"heading\":2,\"id\":\"2b316f32691d7317\",\"modified\":1577098754204,\"text\":\"测试\"},{\"color\":\"#dc2d1e\",\"heading\":2,\"id\":\"20216f671c9316125\",\"modified\":1577982923660,\"text\":\"\"},{\"color\":\"#dc2d1e\",\"heading\":2,\"id\":\"10716f671c949f054\",\"modified\":1577982924384,\"text\":\"ok\"},{\"color\":\"#dc2d1e\",\"heading\":2,\"id\":\"1a216f6b3c7044178\",\"modified\":1578052121547,\"text\":\"标签\"}]}"}}`)

	n := &models.GetDoc{}
	xerror.Panic(json.Unmarshal(_dt, n))
	//xerror.Debug(n.Nodes())
	fmt.Println(n.Markdown())
}
