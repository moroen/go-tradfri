// Copyright © 2017 Moroen <moroen@gmail.com>

package cmd

import (
	"fmt"

	coap "github.com/moroen/go-tradfricoap"
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
		lights, err := coap.GetDevices()
		if err != nil {
			if err == coap.ErrorTimeout {
				fmt.Println("Timeout getting lights")
				return
			}

			if err == coap.ErrorBadIdent {
				fmt.Println("Error reading dtls-stream. Bad credentials?")
				return
			}
			panic(err.Error())
		}

		fmt.Println("Lights:")
		for i := range lights {
			fmt.Println(lights[i].Describe())
		}

		groups, err := coap.GetGroups()
		if err != nil {
			if err == coap.ErrorTimeout {
				fmt.Println("Error: Timeout getting groups")
				return
			}

			if err == coap.ErrorBadIdent {
				fmt.Println("Error reading dtls-stream. Bad credentials?")
				return
			}
			panic(err.Error())
		}

		fmt.Println("\nGroups:")
		for i := range groups {
			fmt.Println(groups[i].Describe())
		}
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
