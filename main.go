// Copyright © 2017 Moroen <moroen@gmail.com>

package main

import (
	"tradfri/cmd"

	coap "github.com/moroen/go-tradfricoap"
	// "github.com/spf13/viper"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func init() {

}

func main() {
	err := coap.LoadConfig()

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	cmd.Execute()

	if err != nil {
		fmt.Println("\nNo config found!")
	} else {

	}

}
