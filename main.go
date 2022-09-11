package main

import (
	"debitinho/file"

	"github.com/spf13/cobra"
)

func main() {
	var (
		rootCmd = &cobra.Command{
			Use: "debitinho",
			CompletionOptions: cobra.CompletionOptions{
				DisableDefaultCmd: true,
			},
		}
		showFileCommand = &cobra.Command{
			Use:   "exibir [arquivo]",
			Short: "exibe um arquivo de remessa/retorno",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				file.ShowFile(file.ParseFile(args[0]))
			},
		}
	)
	rootCmd.AddCommand(showFileCommand)
	rootCmd.Execute()
}
