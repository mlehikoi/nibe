package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mlehikoi/nibe/internal/utils"
)

func main() {
	if len(os.Args) > 1 {
		port, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("Could not parse the port number from ", os.Args[1])
		}
		utils.Authorize(port)
		return
	}
	fmt.Printf("Run:\n"+
		"%s [port]    to authorize the application\n", os.Args[0])
}
