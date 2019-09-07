// Copyright © 2017 Moroen <moroen@gmail.com>

package cmd

import (
	"fmt"

	coap "github.com/moroen/tradfricoap"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if lights, plugs, groups, err := coap.GetDevices(); err == nil {
			fmt.Println("Lights:")
			for i := range lights {
				fmt.Println(lights[i].Describe())
			}

			fmt.Println("\nPlugs:")
			for i := range plugs {
				fmt.Println(plugs[i].Describe())
			}

			fmt.Println("\nGroups:")
			for i := range groups {
				fmt.Println(groups[i].Describe())
			}
		}

		/*
			if lights, err := coap.GetLights(); err == nil {
				fmt.Println("Lights:")
				for i := range lights {
					fmt.Println(lights[i].Describe())
				}
			}

			if plugs, err := coap.GetPlugs(); err == nil {
				fmt.Println("\nPlugs:")
				for i := range plugs {
					fmt.Println(plugs[i].Describe())
				}
			}

			if groups, err := coap.GetGroups(); err == nil {
				fmt.Println("\nGroups:")
				for i := range groups {
					fmt.Println(groups[i].Describe())
				}
			}
		*/
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
