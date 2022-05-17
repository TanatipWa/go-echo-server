package main

import (
	"fmt"
	"os"

	"github.com/TanatipWa/go-echo-server/pkg/server"
)

func main() {
	if err := server.RunServer(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
