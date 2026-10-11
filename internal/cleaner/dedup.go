package cleaner

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"sort"
)

// CalculateSHA256 calcula o hash SHA-256 de um arquivo
func CalculateSHA256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// SortPathsByOldest ordena a lista de arquivos do mais antigo para o mais novo
func SortPathsByOldest(paths []string) {
	sort.Slice(paths, func(i, j int) bool {
		infoI, errI := os.Stat(paths[i])
		infoJ, errJ := os.Stat(paths[j])
		if errI != nil || errJ != nil {
			return false
		}
		return infoI.ModTime().Before(infoJ.ModTime())
	})
}
