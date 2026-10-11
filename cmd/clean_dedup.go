package cmd

import (
	"fmt"
	"housekeeper/internal/cleaner"
	"housekeeper/internal/logger"
	"os"
)

func processDedupScan(path string, info os.FileInfo, fileHashMap map[string][]string) {
	if info.IsDir() || info.Size() == 0 {
		return
	}
	hash, err := cleaner.CalculateSHA256(path)
	if err != nil {
		fmt.Printf("[AVISO] Não foi possível calcular hash de %s: %v\n", path, err)
		return
	}
	fileHashMap[hash] = append(fileHashMap[hash], path)
}

func processDuplicateGroups(dedup bool, fileHashMap map[string][]string) []string {
	var duplicateFilesToDelete []string
	if !dedup {
		return duplicateFilesToDelete
	}

	for _, paths := range fileHashMap {
		if len(paths) > 1 {
			cleaner.SortPathsByOldest(paths)

			original := paths[0]
			origInfo, _ := os.Stat(original)
			origMod := "N/A"
			if origInfo != nil {
				origMod = origInfo.ModTime().Format("02/01/2006 15:04:05")
			}

			fmt.Printf("\n[GRUPO DUPLICADO] Original (Mantido): %s (Mod: %s)\n", original, origMod)

			for _, dup := range paths[1:] {
				dupInfo, _ := os.Stat(dup)
				dupMod := "N/A"
				if dupInfo != nil {
					dupMod = dupInfo.ModTime().Format("02/01/2006 15:04:05")
				}
				fmt.Printf("  └─ [DUPLICADO ALVO] %s (Mod: %s)\n", dup, dupMod)
				duplicateFilesToDelete = append(duplicateFilesToDelete, dup)
			}
		}
	}
	return duplicateFilesToDelete
}

func executeDedupRemovals(dedup bool, duplicateFilesToDelete []string, targetPath string) {
	if !dedup || len(duplicateFilesToDelete) == 0 {
		return
	}
	fmt.Println("\nRemovendo arquivos duplicados...")
	for _, file := range duplicateFilesToDelete {
		if err := os.Remove(file); err != nil {
			fmt.Printf("[ERRO] Falha ao remover duplicado %s: %v\n", file, err)
		} else {
			fmt.Printf("[REMOVIDO DUPLICADO] %s\n", file)
			// Registra a ação no log (respeitando a flag --no-log via variável global noLog)
			_ = logger.LogAction(targetPath, fmt.Sprintf("REMOVED_DUPLICATE: %s", file), noLog)
		}
	}
}
