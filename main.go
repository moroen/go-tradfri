// Copyright © 2017 Moroen <moroen@gmail.com>

package main

import (
	"tradfri/cmd"

	coap "github.com/moroen/tradfricoap"
	// "github.com/spf13/viper"
	"fmt"
)

func init() {

}

func main() {
	err := coap.LoadConfig()

	cmd.Execute()

	if err != nil {
		fmt.Println("\nNo config found!")
	} else {

	}

}
