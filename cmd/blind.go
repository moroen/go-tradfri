/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	coap "github.com/moroen/go-tradfricoap"
	"github.com/spf13/cobra"
)

// blindCmd represents the blind command
var blindCmd = &cobra.Command{
	Use:   "blind ID level",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {

		if len(args) != 2 {
			return fmt.Errorf("wrong number of arguments")
		}

		if err := coap.ValidateDeviceID(args[0]); err != nil {
			return err
		}

		if args[1] == "open" || args[1] == "close" {
			return nil
		}

		if _, err := strconv.Atoi(args[1]); err != nil {
			return fmt.Errorf("%s doesn't appear to be a valid blind level", args[1])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var level int

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if args[1] == "open" {
			level = 0
		} else if args[1] == "close" {
			level = 100
		} else {
			level, err = strconv.Atoi(args[1])
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}

		blind, err := coap.SetBlind(int64(id), level)
		if err != nil {
			log.Println(err.Error())
			return
		}

		fmt.Println(blind.Describe())
	},
}

func init() {
	rootCmd.AddCommand(blindCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// blindCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// blindCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
