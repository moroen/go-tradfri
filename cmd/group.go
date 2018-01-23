// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

	coap "github.com/moroen/go-tradfricoap"
	"github.com/spf13/cobra"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Commands for working with groups",
	Long:  `Commands for working with groups.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("group called")
	},
}

var listGroupCmd = &cobra.Command{
	Use:   "list",
	Short: "List groups",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		groups, err := coap.GetGroups()
		if err != nil {
			panic(err.Error())
		}

		for i := range groups {
			fmt.Println(groups[i].Describe())
		}
	},
}

var groupOffCmd = &cobra.Command{
	Use:   "off",
	Short: "Switch a group off",
	Long:  `Commands for working with groups.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Wrong number of arguments")
		}
		err := coap.ValidateDeviceID(args[0])
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err.Error())
		}
		_, err = coap.GroupSetState(int64(id), 0)
		if err != nil {
			panic(err.Error())
		}
	},
}

var groupOnCmd = &cobra.Command{
	Use:   "on",
	Short: "Switch a group on",
	Long:  `Commands for working with groups.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Wrong number of arguments")
		}
		err := coap.ValidateDeviceID(args[0])
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err.Error())
		}
		_, err = coap.GroupSetState(int64(id), 1)
		if err != nil {
			panic(err.Error())
		}
	},
}

var groupLevelCmd = &cobra.Command{
	Use:   "level",
	Short: "Set brighness level (0-254)",
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

		if _, err := strconv.Atoi(args[1]); err != nil {
			return fmt.Errorf("%s doesn't appear to be a valid dimmer level", args[1])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err.Error())
		}

		level, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err.Error())
		}

		_, err = coap.GroupSetLevel(int64(id), level)
		if err != nil {
			panic(err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(groupCmd)
	groupCmd.AddCommand(listGroupCmd)
	groupCmd.AddCommand(groupOnCmd)
	groupCmd.AddCommand(groupOffCmd)
	groupCmd.AddCommand(groupLevelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// groupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// groupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
