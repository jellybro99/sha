package cmd

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

type Hasher func(string) [8]uint32

func hasher(hashFunction Hasher, outputFormat string, messages []string) error {
	if len(messages) == 1 {
		fmt.Print(formatHash(hashFunction(messages[0]), outputFormat))
		return nil
	}

	var wg sync.WaitGroup
	jobs := make(chan string)

	workers := runtime.GOMAXPROCS(0)
	for range workers {
		wg.Go(func() {
			for message := range jobs {
				fmt.Printf("%s: %s", message, formatHash(hashFunction(message), outputFormat))
			}
		})
	}

	for _, message := range messages {
		jobs <- message
	}
	close(jobs)

	wg.Wait()
	return nil
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
