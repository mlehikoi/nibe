package main

import (
	"fmt"
	"os"

	"github.com/mlehikoi/nibe/uplink"
)

func main() {
	uplink := uplink.NewUplink(os.Args[1], os.Args[2])
	uplink.Update()
	for _, sys := range uplink.Systems {
		fmt.Println("System Name:", sys.Name)
		fmt.Println(sys.Status)
		fmt.Println(sys.Compressor)
		fmt.Println(sys.Ventilation)
		fmt.Println(sys.Climate)
		fmt.Println(sys.Addition)
	}
}
