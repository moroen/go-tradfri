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
	"strconv"
	"strings"

	coap "github.com/moroen/go-tradfricoap"
	"github.com/spf13/cobra"
)

// colorCmd represents the color command
var colorCmd = &cobra.Command{
	Use:   "color",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {

		if err := coap.ValidateDeviceID(args[0]); err != nil {
			return err
		}

		if strings.ToLower(args[1]) == "list" {
			return nil
		}

		if _, err := strconv.Atoi(args[1]); err != nil {
			return fmt.Errorf("%s doesn't appear to be a valid color level", args[1])
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("color called")

		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err.Error())
		}

		if strings.ToLower(args[1]) == "list" {
			if device, err := coap.GetLight(int64(id)); err == nil {
				coap.ListColorsInMap(device.Colors)
			}
		} else {
			level, err := strconv.Atoi(args[1])
			if err != nil {
				panic(err.Error())
			}
			coap.SetHexForLevel(int64(id), level)
		}
	},
}

func init() {
	rootCmd.AddCommand(colorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// colorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// colorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
