package cmds

import (
	"github.com/pubgo/g/xcmd"
	"github.com/pubgo/mubu/cmds/mubu"
)

// Execute exec
var Execute = xcmd.Init(func(cmd *xcmd.Command) {
	cmd.AddCommand(
		mubu.Test(),
	)
})
