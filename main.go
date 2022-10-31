package main

import (
	"fmt"

	"github.com/BinMunawir/mashi/src/delivery/infrastructure/web"
)

func main() {
	fmt.Println("Hello from Mashi")
	web.InitializeHTTPServer()
}
