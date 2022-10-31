package main

import (
	"fmt"

	"github.com/BinMunawir/mashi/src/delivery/infra/web"
)

func main() {
	fmt.Println("Hello from Mashi")
	web.InitializeHTTPServer()
}
