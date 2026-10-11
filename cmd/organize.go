package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	orgFile      bool
	orgName      string
	orgExtension bool
	orgDryRun    bool
	orgNoSecure  bool
)

var organizeCmd = &cobra.Command{
	Use:   "organize [path]",
	Short: "Organiza arquivos por extensão, nome ou data de criação.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		targetPath := "."
		if len(args) > 0 {
			targetPath = args[0]
		}

		fmt.Printf("Organizando em: %s\n", targetPath)
		fmt.Printf("Flags -> file: %v | name: %s | extension: %v | dry-run: %v | no-secure: %v\n",
			orgFile, orgName, orgExtension, orgDryRun, orgNoSecure)
	},
}

func init() {
	rootCmd.AddCommand(organizeCmd)

	organizeCmd.Flags().BoolVar(&orgFile, "file", false, "Alvo são os arquivos")
	organizeCmd.Flags().StringVar(&orgName, "name", "", "Critério por nome (ex: alphabet, createat)")
	organizeCmd.Flags().BoolVar(&orgExtension, "extension", false, "Organiza por extensão")
	organizeCmd.Flags().BoolVar(&orgDryRun, "dry-run", false, "Simula a execução")
	organizeCmd.Flags().BoolVar(&orgNoSecure, "no-secure", false, "Desativa o log de undo")
}
