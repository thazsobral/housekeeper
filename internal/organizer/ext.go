package organizer

import (
	"path/filepath"
	"strings"
)

// GetCleanExtension extrai e limpa a extensão do arquivo (ex: ".PDF" -> "pdf")
func GetCleanExtension(filename string) string {
	ext := filepath.Ext(filename)
	if ext == "" {
		return "no-extension"
	}
	return strings.ToLower(strings.TrimPrefix(ext, "."))
}
