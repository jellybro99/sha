// Package cmd is cobra cli
package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/jellybro99/sha/sha256"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sha",
	Short: "A go implementation of sha hashing algorithms",
	Long:  "sha is a CLI tool for computing sha hashes. You can provide text as arguments or pipe it in via stdin.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("output", "o", "hex", "Output in the given format (hex, dec, bin)")

	rootCmd.AddCommand(newHashCommand("sha256", sha256.Hash))
}

func newHashCommand(hashName string, hashFunction Hasher) *cobra.Command {
	return &cobra.Command{
		Use:   hashName,
		Short: "Compute the " + hashName + " hash of the given input",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runHasher(cmd, args, hashFunction)
		},
	}
}

func runHasher(cmd *cobra.Command, args []string, hashFunction Hasher) error {
	outputFormat, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}
	switch outputFormat {
	case "hex", "dec", "bin":
	default:
		return fmt.Errorf("invalid output format %q (expected hex, dec, or bin)", outputFormat)
	}

	messages, err := getInputs(args)
	if err != nil {
		return err
	}

	return hasher(hashFunction, outputFormat, messages)
}

func getInputs(args []string) ([]string, error) {
	if len(args) > 0 {
		return args, nil
	}

	file, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if file.Size() == 0 {
		return nil, errors.New("no input given")
	}

	message, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}

	return []string{string(message)}, nil
}
