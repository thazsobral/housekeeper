package cleaner

import (
	"fmt"
	"os"
)

// Verifica se o arquivo atende ao critério de tamanho exato
func CheckFileMatch(path string, content string) (bool, int64, error) {
	if content == "" {
		return false, 0, nil
	}

	targetSize, err := ParseSize(content)
	if err != nil {
		return false, 0, fmt.Errorf("formato de tamanho inválido: %s", content)
	}

	info, err := os.Stat(path)
	if err != nil {
		return false, 0, err
	}

	if !info.IsDir() && info.Size() == targetSize {
		return true, info.Size(), nil
	}

	return false, 0, nil
}
