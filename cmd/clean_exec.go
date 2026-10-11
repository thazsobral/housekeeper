package cmd

import (
	"bufio"
	"fmt"
	"housekeeper/internal/scanner"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func runClean(cmd *cobra.Command, args []string) {
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
	fileHashMap := make(map[string][]string)

	err := scanner.WalkDir(targetPath, func(path string, info os.FileInfo, err error) error {
		if path == targetPath {
			return nil
		}

		if cleanFile {
			processFileScan(path, info, content, &filesToDelete)
		}

		if dedup {
			processDedupScan(path, info, fileHashMap)
		}

		if cleanFolder {
			processFolderScan(path, info, &foldersToDelete)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Erro na varredura: %v\n", err)
		return
	}

	duplicateFilesToDelete := processDuplicateGroups(dedup, fileHashMap)

	totalTargets := len(filesToDelete) + len(foldersToDelete) + len(duplicateFilesToDelete)
	if totalTargets == 0 {
		fmt.Println("\nNenhum item encontrado com os critérios especificados.")
		return
	}

	fmt.Printf("\nResumo: %d arquivos por tamanho, %d duplicados, %d pastas vazias\n",
		len(filesToDelete), len(duplicateFilesToDelete), len(foldersToDelete))

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

	executeFileRemovals(cleanFile, filesToDelete)
	executeDedupRemovals(dedup, duplicateFilesToDelete)
	executeFolderRemovals(cleanFolder, foldersToDelete)

	fmt.Println("\nLimpeza concluída com sucesso!")
}
