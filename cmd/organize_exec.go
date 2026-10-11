package cmd

import (
	"bufio"
	"fmt"
	"housekeeper/internal/logger"
	"housekeeper/internal/scanner"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type FileMovePlan struct {
	SourcePath string
	DestPath   string
	Extension  string
}

func runOrganize(cmd *cobra.Command, args []string) {
	targetPath := "."
	if len(args) > 0 {
		targetPath = args[0]
	}

	if !orgExt && orgName == "" && orgDate == "" {
		fmt.Println("Erro: você deve especificar ao menos um critério de organização (--ext, --name ou --date).")
		return
	}

	if orgName != "" && orgName != "alphabet" {
		fmt.Printf("Erro: critério de nome desconhecido '%s'. Opções válidas: alphabet\n", orgName)
		return
	}

	if orgDate != "" && orgDate != "createat" {
		fmt.Printf("Erro: critério de data desconhecido '%s'. Opções válidas: createat\n", orgDate)
		return
	}

	fmt.Printf("Varrendo diretório para organização: %s...\n\n", targetPath)

	var movePlans []FileMovePlan

	err := scanner.WalkDir(targetPath, func(path string, info os.FileInfo, err error) error {
		if path == targetPath {
			return nil
		}

		if orgExt {
			collectExtMovePlan(targetPath, path, info, &movePlans)
		}

		if orgName == "alphabet" {
			collectAlphabetMovePlan(targetPath, path, info, &movePlans)
		}

		if orgDate == "createat" {
			collectDateMovePlan(targetPath, path, info, &movePlans)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Erro na varredura: %v\n", err)
		return
	}

	if len(movePlans) == 0 {
		fmt.Println("Nenhum arquivo elegível para organização encontrado.")
		return
	}

	fmt.Printf("\nTotal de arquivos mapeados para organização: %d\n", len(movePlans))

	if dryRun {
		fmt.Println("[DRY-RUN] Nenhuma alteração foi realizada.")
		return
	}

	if !yes {
		fmt.Print("\nDeseja realmente executar a movimentação desses arquivos? [s/N]: ")
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))
		if response != "s" && response != "sim" && response != "y" {
			fmt.Println("Operação cancelada pelo usuário.")
			return
		}
	}

	executeMovePlans(movePlans, targetPath)

	fmt.Println("\nOrganização concluída com sucesso!")
}

func executeMovePlans(plans []FileMovePlan, targetPath string) {
	fmt.Println("\nMovendo arquivos...")
	for _, plan := range plans {
		destDir := filepath.Dir(plan.DestPath)
		if err := os.MkdirAll(destDir, 0755); err != nil {
			fmt.Printf("[ERRO] Falha ao criar diretório %s: %v\n", destDir, err)
			continue
		}

		if err := os.Rename(plan.SourcePath, plan.DestPath); err != nil {
			fmt.Printf("[ERRO] Falha ao mover %s: %v\n", filepath.Base(plan.SourcePath), err)
		} else {
			fmt.Printf("[MOVIDO] %s -> %s/\n", filepath.Base(plan.SourcePath), plan.Extension)
			_ = logger.LogAction(targetPath, fmt.Sprintf("MOVED: %s -> %s", plan.SourcePath, plan.DestPath), noLog)
		}
	}
}
