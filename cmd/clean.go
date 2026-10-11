package cmd

import (
	"bufio"
	"fmt"
	"housekeeper/internal/cleaner"
	"housekeeper/internal/scanner"
	"os"
	"sort"
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

		if !cleanFile && !cleanFolder && !dedup {
			fmt.Println("Erro: você deve especificar ao menos um escopo de limpeza (--file, --folder ou --dedup).")
			return
		}

		fmt.Printf("Varrendo diretório: %s...\n\n", targetPath)

		var filesToDelete []string
		var foldersToDelete []string

		err := scanner.WalkDir(targetPath, func(path string, info os.FileInfo, err error) error {
			if path == targetPath {
				return nil
			}

			// Valida arquivo apenas se --file foi passado
			if cleanFile && !info.IsDir() {
				matched, sizeInfo, err := cleaner.CheckFileMatch(path, content)
				if err != nil {
					return err
				}
				if matched {
					filesToDelete = append(filesToDelete, path)
					fmt.Printf("[ARQUIVO ALVO] %s (%d bytes)\n", path, sizeInfo)
				}
			}

			// Valida pasta apenas se --folder foi passado
			if cleanFolder && info.IsDir() {
				isEmpty, err := cleaner.IsDirEmpty(path)
				if err == nil && isEmpty {
					foldersToDelete = append(foldersToDelete, path)
					fmt.Printf("[PASTA VAZIA ALVO] %s\n", path)
				}
			}

			return nil
		})

		if err != nil {
			fmt.Printf("Erro na varredura: %v\n", err)
			return
		}

		totalTargets := len(filesToDelete) + len(foldersToDelete)
		if totalTargets == 0 {
			fmt.Println("Nenhum item encontrado com os critérios especificados.")
			return
		}

		fmt.Printf("\nTotal encontrados: %d arquivos, %d pastas\n", len(filesToDelete), len(foldersToDelete))

		if dryRun {
			fmt.Println("[DRY-RUN] Nenhuma alteração foi realizada.")
			return
		}

		if !yes {
			fmt.Print("\nDeseja realmente executar a remoção desses itens? [s/N]: ")
			reader := bufio.NewReader(os.Stdin)
			response, _ := reader.ReadString('\n')
			response = strings.TrimSpace(strings.ToLower(response))
			if response != "s" && response != "sim" && response != "y" {
				fmt.Println("Operação cancelada pelo usuário.")
				return
			}
		}

		// Executa remoção de arquivos somente se --file foi acionado
		if cleanFile && len(filesToDelete) > 0 {
			fmt.Println("\nRemovendo arquivos...")
			for _, file := range filesToDelete {
				if err := os.Remove(file); err != nil {
					fmt.Printf("[ERRO] Falha ao remover arquivo %s: %v\n", file, err)
				} else {
					fmt.Printf("[REMOVIDO] %s\n", file)
				}
			}
		}

		// Executa remoção de pastas somente se --folder foi acionado
		if cleanFolder && len(foldersToDelete) > 0 {
			sort.Slice(foldersToDelete, func(i, j int) bool {
				return len(foldersToDelete[i]) > len(foldersToDelete[j])
			})

			fmt.Println("\nRemovendo pastas vazias...")
			for _, folder := range foldersToDelete {
				if empty, _ := cleaner.IsDirEmpty(folder); empty {
					if err := os.Remove(folder); err != nil {
						fmt.Printf("[ERRO] Falha ao remover pasta %s: %v\n", folder, err)
					} else {
						fmt.Printf("[REMOVIDO] %s\n", folder)
					}
				}
			}
		}

		fmt.Println("\nLimpeza concluída com sucesso!")
	},
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
