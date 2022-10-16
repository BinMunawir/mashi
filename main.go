package main

import (
	"fmt"

	"github.com/BinMunawir/mashi/src/outer_layers/infrastructure/web"
)

func main() {
	fmt.Println("Hello from Mashi")
	web.InitializeHTTPServer()
}
