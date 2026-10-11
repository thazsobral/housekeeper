package cmd

import (
	"fmt"
	"housekeeper/internal/cleaner"
	"housekeeper/internal/logger"
	"os"
)

func processFileScan(path string, info os.FileInfo, contentCrit string, filesToDelete *[]string) {
	if info.IsDir() {
		return
	}
	matched, sizeInfo, err := cleaner.CheckFileMatch(path, contentCrit)
	if err != nil {
		return
	}
	if matched {
		*filesToDelete = append(*filesToDelete, path)
		modTimeStr := info.ModTime().Format("02/01/2006 15:04:05")
		fmt.Printf("[ARQUIVO ALVO] %s (%d bytes, Modificado em: %s)\n", path, sizeInfo, modTimeStr)
	}
}

func executeFileRemovals(cleanFile bool, filesToDelete []string, targetPath string) {
	if !cleanFile || len(filesToDelete) == 0 {
		return
	}
	fmt.Println("\nRemovendo arquivos por tamanho...")
	for _, file := range filesToDelete {
		if err := os.Remove(file); err != nil {
			fmt.Printf("[ERRO] Falha ao remover %s: %v\n", file, err)
		} else {
			fmt.Printf("[REMOVIDO] %s\n", file)
			// Registra a ação no log (respeitando a flag --no-log)
			_ = logger.LogAction(targetPath, fmt.Sprintf("REMOVED_FILE: %s", file), noLog)
		}
	}
}
