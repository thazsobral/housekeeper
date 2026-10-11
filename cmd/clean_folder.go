package cmd

import (
	"fmt"
	"housekeeper/internal/cleaner"
	"os"
	"sort"
)

func processFolderScan(path string, info os.FileInfo, foldersToDelete *[]string) {
	if !info.IsDir() {
		return
	}
	isEmpty, err := cleaner.IsDirEmpty(path)
	if err == nil && isEmpty {
		*foldersToDelete = append(*foldersToDelete, path)
		modTimeStr := info.ModTime().Format("02/01/2006 15:04:05")
		fmt.Printf("[PASTA VAZIA ALVO] %s (Modificado em: %s)\n", path, modTimeStr)
	}
}

func executeFolderRemovals(cleanFolder bool, foldersToDelete []string) {
	if !cleanFolder || len(foldersToDelete) == 0 {
		return
	}

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
