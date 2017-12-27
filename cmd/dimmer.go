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
	"strconv"

	coap "github.com/moroen/go-tradfricoap"

	"github.com/spf13/cobra"
)

// dimmerCmd represents the dimmer command
var dimmerCmd = &cobra.Command{
	Use:   "dimmer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		level, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err.Error())
		}

		_, err = coap.SetLevel(args[0], level)
		if err != nil {
			panic(err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(dimmerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dimmerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dimmerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
