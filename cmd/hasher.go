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
				fmt.Printf("%s: %s", truncateLabel(message, 8), formatHash(hashFunction(message), outputFormat))
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

func truncateLabel(s string, maxRunes int) string {
	runes := []rune(s)
	if len(runes) <= maxRunes {
		return s
	}
	return string(runes[:maxRunes])
}

func formatHash(hash [8]uint32, outputFormat string) string {
	var sb strings.Builder

	for _, word := range hash {
		switch outputFormat {
		case "hex":
			fmt.Fprintf(&sb, "%08X ", word)
		case "dec":
			fmt.Fprintf(&sb, "%d ", word)
		case "bin":
			fmt.Fprintf(&sb, "%032b ", word)
		}
	}
	sb.WriteString("\n")

	return sb.String()
}
