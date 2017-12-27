// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"

	coap "github.com/moroen/tradfricoap"

	"github.com/spf13/cobra"
)

var flagGateway string
var flagIdent string
var flagKey string
var flagPort string

// gatewayCmd represents the gateway command
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gateway called")
	},
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := coap.GetConfig()
		isDirty := false

		if flagGateway != "" {
			conf.Gateway = fmt.Sprintf("%s:%s", flagGateway, flagPort)
			isDirty = true
		}

		if flagIdent != "" {
			conf.Identity = flagIdent
			isDirty = true
		}

		if flagKey != "" {
			conf.Passkey = flagKey
			isDirty = true
		}

		if isDirty {
			coap.SetConfig(conf)
			coap.SaveConfig(conf)
		} else {
			log.Println(coap.GetConfig().Describe())
		}
	},
}

// createidCmd represents the createid command
var createidCmd = &cobra.Command{
	Use:   "createid",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		coap.CreateIdent(args[0], args[1], args[2])
	},
}

func init() {

	configCmd.Flags().StringVar(&flagGateway, "gateway", "", "Set gateway address")
	configCmd.Flags().StringVar(&flagIdent, "ident", "", "Set ident")
	configCmd.Flags().StringVar(&flagKey, "key", "", "Set PSK for ident")
	configCmd.Flags().StringVar(&flagPort, "port", coap.DefaultPort, "Gateway port")

	gatewayCmd.AddCommand(configCmd)
	gatewayCmd.AddCommand(createidCmd)
	rootCmd.AddCommand(gatewayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
