// Package cmd is cobra cli
package cmd

import (
	"fmt"
	"os"

	"github.com/jellybro99/sha256_cli/sha256"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sha",
	Short: "A go implementation of the sha256 hashing algorithm",
	Long: `sha is a CLI tool for the sha256 hashing algorithm.
	You can provide text as arguments or pipe it in via stdin.`,
	Args: cobra.ArbitraryArgs,
	Run:  runHasher,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Bool("sha256", false, "Use the sha256 hash function")
}

func runHasher(cmd *cobra.Command, args []string) {
	result, err := cmd.Flags().GetBool("sha256")
	if err != nil {
		os.Exit(1)
	}
	if result {
		fmt.Printf("%s: %X\n", args[0], sha256.Hash(args[0]))
	} else {
		fmt.Println("use -h")
	}
}
