// Copyright © 2017 Moroen <moroen@gmail.com>

package cmd

import (
	"fmt"
	"strconv"

	coap "github.com/moroen/tradfricoap"

	"github.com/spf13/cobra"
)

// switchCmd represents the switch command

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Turn on a device",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("switch called")
	},
}

var switchOnCmd = &cobra.Command{
	Use:   "on",
	Short: "Turn on a device",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Wrong number of arguments")
		}

		if err := coap.ValidateDeviceID(args[0]); err != nil {
			return err
		}
		return nil
		// return coap.ValidateOnOff(args[1])
	},
	Run: func(cmd *cobra.Command, args []string) {

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		device, err := coap.SetState(int64(id), 1)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(device.Describe())
		}
	},
}

var switchOffCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn off a device",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Wrong number of arguments")
		}

		if err := coap.ValidateDeviceID(args[0]); err != nil {
			return err
		}
		return nil
		// return coap.ValidateOnOff(args[1])
	},
	Run: func(cmd *cobra.Command, args []string) {

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		device, err := coap.SetState(int64(id), 0)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(device.Describe())
		}
	},
}

var switchListCmd = &cobra.Command{
	Use:   "list",
	Short: "List lights",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		lights, err := coap.GetLights()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for i := range lights {
			fmt.Println(lights[i].Describe())
		}

	},
}

var dimmerCmd = &cobra.Command{
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
			fmt.Println(err.Error())
			return
		}

		level, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		device, err := coap.SetLevel(int64(id), level)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(device.Describe())
		}

	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
	switchCmd.AddCommand(switchListCmd)
	switchCmd.AddCommand(switchOnCmd)
	switchCmd.AddCommand(switchOffCmd)
	switchCmd.AddCommand(dimmerCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// switchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// switchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
