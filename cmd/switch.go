// Copyright © 2017 Moroen <moroen@gmail.com>

package cmd

import (
	"fmt"
	"strconv"
	"strings"

	coap "github.com/moroen/go-tradfricoap"

	"github.com/spf13/cobra"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
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

		if strings.ToLower(args[1]) == "on" || strings.ToLower(args[1]) == "off" || strings.ToLower(args[1]) == "1" || strings.ToLower(args[1]) == "0" {
			return nil
		} else {
			return fmt.Errorf("%s isn't an allowed setting, use 'on', 'off', '1' or '0'", args[1])
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		state := 0

		id, err := strconv.Atoi(args[0])

		if err != nil {
			panic(err.Error())
		}

		if strings.ToLower(args[1]) == "on" {
			state = 1
		}

		_, err = coap.SetState(int64(id), state)
		if err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// switchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// switchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
