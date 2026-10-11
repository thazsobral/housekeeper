package cmd

import (
	"fmt"
	"housekeeper/internal/organizer"
	"os"
	"path/filepath"
	"strings"
)

func collectDateMovePlan(targetPath string, path string, info os.FileInfo, plans *[]FileMovePlan) {
	if info.IsDir() {
		return
	}

	rel, err := filepath.Rel(targetPath, path)
	if err != nil || strings.Contains(rel, string(filepath.Separator)) {
		return
	}

	folder := organizer.GetDateFolder(info.ModTime())
	destDir := filepath.Join(targetPath, folder)
	destPath := filepath.Join(destDir, info.Name())

	absSource, _ := filepath.Abs(path)
	absDest, _ := filepath.Abs(destPath)
	if absSource == absDest {
		return
	}

	*plans = append(*plans, FileMovePlan{
		SourcePath: path,
		DestPath:   destPath,
		Extension:  folder,
	})

	modTimeStr := info.ModTime().Format("02/01/2006 15:04:05")
	fmt.Printf("[PLANO] %s -> %s/ (%d bytes, Modificado em: %s)\n", info.Name(), folder, info.Size(), modTimeStr)
}
