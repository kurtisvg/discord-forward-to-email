package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	opts := parseFlags(os.Args[1:])

	_ = opts // will be used when server is implemented
	fmt.Println("server not yet implemented")
}
