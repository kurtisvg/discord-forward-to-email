package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	opts := parseFlags(os.Args[1:])

	if opts.register {
		opts.require("discord-token", "discord-app-id")
		runRegister(opts)
		return
	}

	_ = opts // will be used when server is implemented
	fmt.Println("server not yet implemented")
}
