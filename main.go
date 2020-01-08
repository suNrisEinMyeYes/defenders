package main

import (
	"fmt"
	"os"
	"time"

	"2laba/client"
	"2laba/server"
)

func main() {
	fmt.Println("Please, enter your port and ip")
	var addr string
	fmt.Fscan(os.Stdin, &addr)
	if len(addr) > 5 {
		cln := client.Client {
			Addr: addr,
		}
		err := cln.StartClient()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		var max int
		fmt.Fscan(os.Stdin, &max)
		srv := server.Server {
			Addr:        addr,
			IdleTimeout: 3 * time.Minute,
			MaxInit:     max,
		}
		err := srv.StartServer()
		if err != nil {
			fmt.Println(err)
		}
	}
}
