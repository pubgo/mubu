package cmds

import (
	"github.com/pubgo/g/xcmd"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/mubu/cmds/mubu"
	"github.com/pubgo/mubu/version"
)

const Service = "mubu"
const EnvPrefix = "MB"

// Execute exec
var Execute = xcmd.Init(EnvPrefix, func(cmd *xcmd.Command) {
	defer xerror.Assert()

	cmd.Use = Service
	cmd.Version = version.Version

	cmd.AddCommand(
		mubu.Version(),
	)

})
