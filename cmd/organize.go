package cmd

import (
	"github.com/spf13/cobra"
)

var (
	orgExt  bool
	orgName string
	orgDate string
)

var organizeCmd = &cobra.Command{
	Use:   "organize [path]",
	Short: "Organiza arquivos por extensão, nome ou data.",
	Args:  cobra.MaximumNArgs(1),
	Run:   runOrganize,
}

func init() {
	rootCmd.AddCommand(organizeCmd)

	organizeCmd.Flags().BoolVar(&orgExt, "ext", false, "Organiza os arquivos por extensão")
	organizeCmd.Flags().StringVar(&orgName, "name", "", "Organiza por critério de nome (ex: alphabet)")
	organizeCmd.Flags().StringVar(&orgDate, "date", "", "Organiza por critério de data (ex: createat)")
	organizeCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Simula a execução")
	organizeCmd.Flags().BoolVar(&noLog, "no-log", false, "Desativa log de ações")
	organizeCmd.Flags().BoolVarP(&yes, "yes", "y", false, "Executa sem pedir confirmação")
}
