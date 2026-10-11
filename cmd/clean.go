package cmd

import (
	"bufio"
	"fmt"
	"housekeeper/internal/cleaner"
	"housekeeper/internal/scanner"
	"os"
	"strings"

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
	Run: func(cmd *cobra.Command, args []string) {
		targetPath := "."
		if len(args) > 0 {
			targetPath = args[0]
		}

		fmt.Printf("Varrendo diretório: %s...\n\n", targetPath)

		// Coletamos primeiro os arquivos que atendem ao critério
		var filesToDelete []string

		err := scanner.WalkDir(targetPath, func(path string, info os.FileInfo, err error) error {
			if cleanFile && !info.IsDir() {
				// Verifica se atende ao critério de vazio ou tamanho
				matched, sizeInfo, err := cleaner.CheckMatch(path, cleanFile, cleanFolder, content)
				if err != nil {
					return err
				}
				if matched {
					filesToDelete = append(filesToDelete, path)
					fmt.Printf("[ALVO ENCONTRADO] %s (%d bytes)\n", path, sizeInfo)
				}
			}
			return nil
		})

		if err != nil {
			fmt.Printf("Erro na varredura: %v\n", err)
			return
		}

		if len(filesToDelete) == 0 {
			fmt.Println("Nenhum arquivo encontrado com os critérios especificados.")
			return
		}

		fmt.Printf("\nTotal de arquivos encontrados para limpeza: %d\n", len(filesToDelete))

		// Se for dry-run, apenas encerra aqui
		if dryRun {
			fmt.Println("[DRY-RUN] Nenhuma alteração foi realizada.")
			return
		}

		// Lógica de confirmação (-y ou pergunta interativa)
		if !yes {
			fmt.Print("\nDeseja realmente executar a remoção desses arquivos? [s/N]: ")
			reader := bufio.NewReader(os.Stdin)
			response, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Erro ao ler confirmação.")
				return
			}
			response = strings.TrimSpace(strings.ToLower(response))
			if response != "s" && response != "sim" && response != "y" {
				fmt.Println("Operação cancelada pelo usuário.")
				return
			}
		}

		// Execução real da remoção
		fmt.Println("\nExecutando limpeza...")
		for _, file := range filesToDelete {
			err := os.Remove(file)
			if err != nil {
				fmt.Printf("[ERRO] Falha ao remover %s: %v\n", file, err)
			} else {
				fmt.Printf("[REMOVIDO] %s\n", file)
			}
		}
		fmt.Println("Limpeza concluída com sucesso!")
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	cleanCmd.Flags().BoolVar(&cleanFile, "file", false, "Alvo são os arquivos")
	cleanCmd.Flags().BoolVar(&cleanFolder, "folder", false, "Alvo são as pastas")
	cleanCmd.Flags().StringVar(&content, "content", "", "Critério de conteúdo (ex: empty, 3Mb)")
	cleanCmd.Flags().BoolVar(&dedup, "dedup", false, "Remove arquivos duplicados via checksum")
	cleanCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Simula a execução")
	cleanCmd.Flags().BoolVar(&noLog, "no-log", false, "Desativa log de ações")
	cleanCmd.Flags().BoolVar(&noUndo, "no-undo", false, "Desativa arquivo de reversão")
	cleanCmd.Flags().BoolVarP(&yes, "yes", "y", false, "Executa sem pedir confirmação")
}
