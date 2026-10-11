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

	rel, err := filepath.Rel(targetPath, path)
	if err != nil || strings.Contains(rel, string(filepath.Separator)) {
		return
	}

	ext := organizer.GetCleanExtension(info.Name())
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
