package main

import (
	b "baa/baml_client"
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	version = "0.1.0"
	rootCmd = &cobra.Command{
		Use:   "baa [category] [category]",
		Short: "baa - categorize input from stdin",
		Long:  "baa - categorize input from stdin",
		Run:   executeDefault,
	}
)

func executeDefault(cmd *cobra.Command, args []string) {
	err := readFromStdin(categorize, args)
	if err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}
}

func readFromStdin(fn func(string, *b.TypeBuilder), args []string) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err == io.EOF && line == "" {
			break
		}
		if err != nil && err != io.EOF {
			return err
		}
		tb, err := buildCategory(args)
		fn(line, tb)
		if err == io.EOF {
			break
		}
	}
	return nil
}

func categorize(input string, tb *b.TypeBuilder) {
	ctx := context.Background()
	category, _ := b.Categorize(ctx, input, b.WithTypeBuilder(tb))
	if category != "" {
		fmt.Printf("%v\n", category)
	}
}

func buildCategory(items []string) (*b.TypeBuilder, error) {
	tb, err := b.NewTypeBuilder()
	if err != nil {
		return nil, err
	}

	category, err := tb.Category()
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		_, err := category.AddValue(item)
		if err != nil {
			return nil, err
		}
	}
	return tb, nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
