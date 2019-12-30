package cmds

import (
	"github.com/pubgo/g/xcmd"
)

const Service = "mubu"

// Execute exec
var Execute = xcmd.Init(func(cmd *xcmd.Command) {
	cmd.AddCommand()
})
