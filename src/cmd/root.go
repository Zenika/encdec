// encdec : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/root.go

package cmd

import (
	"encdec/executor"
	"encdec/helpers"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var version = "0.200-0 (2023.07.29)"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "encdec",
	Short:   "Encode and decode a string or file to-from AES-256",
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

var encodeCmd = &cobra.Command{
	Use:     "encode",
	Aliases: []string{"enc", "encrypt"},
	Short:   "Encrypts a string or a file",
	Run: func(cmd *cobra.Command, args []string) {
		if executor.FileEncryptionDecryption {
			if len(args) < 2 {
				fmt.Println("You need to specify both the source and destination files")
				os.Exit(1)
			}
			if err := executor.EncodeFile(args[0], args[1]); err != nil {
				fmt.Sprintf("Error encoding %s : %v", args[0], err)
			}
		} else {
			fmt.Printf("Encoded string is: %s\n\n", executor.Encode(args[0]))
		}
	},
}

var decodeCmd = &cobra.Command{
	Use:     "decode",
	Aliases: []string{"dec", "decrypt"},
	Short:   "Decrypts a string or a file",
	Run: func(cmd *cobra.Command, args []string) {
		if executor.FileEncryptionDecryption {
			if len(args) < 2 {
				fmt.Println("You need to specify both the source and destination files")
				os.Exit(1)
			}
			if err := executor.DecodeFile(args[0], args[1]); err != nil {
				fmt.Sprintf("Error decoding %s : %v", args[0], err)
			}
		} else {
			fmt.Printf("Decoded string is: %s\n\n", executor.Decode(args[0]))
		}
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
	rootCmd.AddCommand(encodeCmd)
	rootCmd.AddCommand(decodeCmd)

	rootCmd.PersistentFlags().BoolVarP(&executor.Prompt4K, "key", "k", false, "Should we prompt for a secret key")
	rootCmd.PersistentFlags().BoolVarP(&executor.FileEncryptionDecryption, "file", "f", false, "Are we dealing with a file or not")
}
