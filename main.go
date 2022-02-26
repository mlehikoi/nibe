package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "authorize" {
			port := 8080
			if len(os.Args) > 2 {
				port, _ = strconv.Atoi(os.Args[2])
			}
			Authorize(port)
			return
		} else if os.Args[1] == "collect" {
			Collect()
			return
		}
	}
	fmt.Printf("Run:\n"+
		"%s authorize [port]    to authorize the application\n"+
		"%s collect             to start collecting data\n",
		os.Args[0], os.Args[0])
}
