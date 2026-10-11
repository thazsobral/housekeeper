package cmd

import (
	"fmt"
	"housekeeper/internal/organizer"
	"os"
	"path/filepath"
	"strings"
)

func collectExtMovePlan(targetPath string, path string, info os.FileInfo, plans *[]FileMovePlan) {
	if info.IsDir() {
		return
	}

	// Processa apenas arquivos na raiz imediata da busca para evitar re-organizar subpastas
	rel, err := filepath.Rel(targetPath, path)
	if err != nil || strings.Contains(rel, string(filepath.Separator)) {
		return
	}

	ext := organizer.GetCleanExtension(info.Name())

	// Define o diretório de destino (ex: targetPath/pdf/)
	destDir := filepath.Join(targetPath, ext)
	destPath := filepath.Join(destDir, info.Name())

	absSource, _ := filepath.Abs(path)
	absDest, _ := filepath.Abs(destPath)
	if absSource == absDest {
		return
	}

	*plans = append(*plans, FileMovePlan{
		SourcePath: path,
		DestPath:   destPath,
		Extension:  ext,
	})

	modTimeStr := info.ModTime().Format("02/01/2006 15:04:05")
	fmt.Printf("[PLANO] %s -> %s/ (%d bytes, Modificado em: %s)\n", info.Name(), ext, info.Size(), modTimeStr)
}

func executeMovePlans(plans []FileMovePlan) {
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
		}
	}
}
