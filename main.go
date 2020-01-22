// Copyright © 2017 Moroen <moroen@gmail.com>

package main

import (
	"fmt"

	"github.com/moroen/go-tradfri/cmd"

	coap "github.com/moroen/go-tradfri/tradfricoap"
	// "github.com/spf13/viper"

	_ "net/http/pprof"
)

func init() {

}

func NoConfigError() {
	fmt.Println("No config found! Please set configuration with 'tradfri gateway config <IP> <KEY>'")
}

func main() {
	_, err := coap.LoadConfig()
	if err != nil {
		panic(err.Error())
	}
	cmd.Execute()

	if err != nil {
		fmt.Println("No config found! Please set configuration with 'tradfri gateway config <IP> <KEY>'")
	}
}
