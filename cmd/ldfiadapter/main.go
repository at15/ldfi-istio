package main

import (
	"fmt"
	"os"

	"github.com/at15/ldfi-istio/ldfiadapter"
)

func main() {
	addr := ""
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}

	s, err := ldfiadapter.NewLDFIAdapter(addr)
	if err != nil {
		fmt.Printf("unable to start server: %v", err)
		os.Exit(-1)
	}

	shutdown := make(chan error, 1)
	go func() {
		s.Run(shutdown)
	}()
	_ = <-shutdown
}
