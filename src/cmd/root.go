// encdec : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/root.go

package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"encdec/helpers"
)

var version = "0.100-0 (2023.07.09)"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "encdec",
	Short:   "Encode and decode a string to-from AES-256",
	Version: version,
}

var clCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: []string{"cl"},
	Short:   "Shows changelog",
	Run: func(cmd *cobra.Command, args []string) {
		helpers.Changelog()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(clCmd)

	rootCmd.PersistentFlags().StringVarP(&config.PortainerHostConfigFile, "environment", "e", "push2registry.json", "Environment file.")
}
