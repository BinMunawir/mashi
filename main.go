package main

import (
	"fmt"

	"github.com/BinMunawir/mashi/src/core/services"
)

func main() {
	fmt.Println("Hello from Mashi")
	// web.InitializeHTTPServer()
	// services.NewKafka().Produce([]byte("hi"), "demo")
	services.NewKafka().Consume("demo")

}
