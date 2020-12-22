package main

import (
	"fmt"
	"os"

	// pb "getitqec.com/server/auth/pkg/api/v1"

	cmd "getitqec.com/server/mailnotification/pkg/cmd"
)

func main() {
	fmt.Printf("Starting mail notification service...\n")
	if err := cmd.RunServer(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
