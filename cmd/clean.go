package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cleanFile   bool
	cleanFolder bool
	content     string
	dedup       bool
	dryRun      bool
	noLog       bool
	noUndo      bool
	yes         bool
)

var cleanCmd = &cobra.Command{
	Use:   "clean [path]",
	Short: "Remove arquivos ou pastas vazias, por tamanho ou duplicados.",
	Args:  cobra.MaximumNArgs(1),
	Run:   runClean,
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	cleanCmd.Flags().BoolVar(&cleanFile, "file", false, "Alvo são os arquivos")
	cleanCmd.Flags().BoolVar(&cleanFolder, "folder", false, "Alvo são as pastas")
	cleanCmd.Flags().StringVar(&content, "content", "", "Critério de conteúdo (ex: 0b, 3Mb)")
	cleanCmd.Flags().BoolVar(&dedup, "dedup", false, "Remove arquivos duplicados via checksum")
	cleanCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Simula a execução")
	cleanCmd.Flags().BoolVar(&noLog, "no-log", false, "Desativa log de ações")
	cleanCmd.Flags().BoolVar(&noUndo, "no-undo", false, "Desativa arquivo de reversão")
	cleanCmd.Flags().BoolVarP(&yes, "yes", "y", false, "Executa sem pedir confirmação")
}
