package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jellybro99/sha256_cli/sha256"
	"github.com/spf13/cobra"
)

func runHasher(cmd *cobra.Command, args []string) {
	useSha256, err := cmd.Flags().GetBool("sha256")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	outputFormat, err := cmd.Flags().GetString("output")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	messages, err := getInputs(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(messages) == 0 {
		fmt.Println("no input given")
	}

	if useSha256 {
		if len(messages) == 1 {
			fmt.Print(formatHash(sha256.Hash(messages[0]), outputFormat))
		} else {
			for _, message := range messages {
				// only print message if its not super duper long
				fmt.Printf("%s: %s", message, formatHash(sha256.Hash(message), outputFormat))
			}
		}
	} else {
		fmt.Println("Given hash function is not supported")
	}
}

func formatHash(hash [8]uint32, outputFormat string) string {
	var sb strings.Builder

	for _, word := range hash {
		switch outputFormat {
		case "hex":
			fmt.Fprintf(&sb, "%X ", word)
		case "dec":
			fmt.Fprintf(&sb, "%d ", word)
		}
	}
	sb.WriteString("\n")

	return sb.String()
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
		return nil, errors.New("no input")
	}

	message, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}

	return []string{string(message)}, nil
}
