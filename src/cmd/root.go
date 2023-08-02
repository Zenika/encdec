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

var version = "0.200-1 (2023.07.31)"

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
		if !executor.FileEncryptionDecryption {
			// encode a string
			fmt.Printf("Encoded string is: %s\n\n", executor.Encode(args[0]))
			os.Exit(0)
		}
		// encode a file
		if len(args) < 2 {
			fmt.Println("You need to specify the source filename")
			os.Exit(1)
		}
		if err := executor.EncodeFile(args[0]); err != nil {
			fmt.Printf("Error encoding %s : %v", args[0], err)
			os.Exit(2)
		}
	},
}

var decodeCmd = &cobra.Command{
	Use:     "decode",
	Aliases: []string{"dec", "decrypt"},
	Short:   "Decrypts a string or a file",
	Run: func(cmd *cobra.Command, args []string) {
		if !executor.FileEncryptionDecryption {
			// decode a string
			fmt.Printf("Decoded string is: %s\n\n", executor.Decode(args[0]))
			os.Exit(0)
		}
		// decode a file
		if len(args) < 2 {
			fmt.Println("You need to specify the source filename")
			os.Exit(1)
		}
		if err := executor.DecodeFile(args[0]); err != nil {
			fmt.Printf("Error decoding %s : %v", args[0], err)
			os.Exit(2)
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

	rootCmd.PersistentFlags().BoolVarP(&executor.Prompt4K, "prompt", "p", false, "Should we prompt for a secret key")
	rootCmd.PersistentFlags().BoolVarP(&executor.FileEncryptionDecryption, "file", "f", false, "Are we dealing with a file or not")
	rootCmd.PersistentFlags().BoolVarP(&executor.FileEncryptionDecryption, "keep", "k", false, "Should we keep the original file")
}
