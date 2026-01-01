// Package cmd is cobra cli
package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sha",
	Short: "A go implementation of the sha256 hashing algorithm",
	Long:  "sha is a CLI tool for the sha256 hashing algorithm. You can provide text as arguments or pipe it in via stdin.",
	Args:  cobra.ArbitraryArgs,
	RunE:  runHasher,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().Bool("sha256", true, "Use the sha256 hashing algorithm")
	rootCmd.Flags().StringP("output", "o", "hex", "Output in the given format (hex, dec)")
}

func runHasher(cmd *cobra.Command, args []string) error {
	useSha256, err := cmd.Flags().GetBool("sha256")
	if err != nil {
		return err
	}
	outputFormat, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	if err := hasher(args, outputFormat, useSha256); err != nil {
		return err
	}
	return nil
}
