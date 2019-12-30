package cmds

import (
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/mubu/version"
)

func init() {
	xenv.InitVersion("Mubu", version.Version, version.BuildV, version.CommitV)
}
